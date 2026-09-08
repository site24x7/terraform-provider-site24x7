---
name: site24x7-import-existing
description: >
  Bring monitors, profiles, tags and integrations that already exist in a Site24x7
  account under Terraform management, producing HCL that plans clean. Inventories the
  account with the site24x7_monitors data source, picks the right import mechanism per
  resource type, generates configuration, and iterates until `terraform plan` reports no
  changes. Use whenever someone wants to import existing Site24x7 monitors, adopt a
  hand-built (ClickOps) Site24x7 account into Terraform, migrate monitors into state,
  run site24x7_importer.py, or asks "how do I get my existing monitors into Terraform"
  — even if they don't say "skill".
user-invocable: true
argument-hint: "<what to import, e.g. 'all SERVER monitors' or 'the 40 website monitors tagged prod'>"
allowed-tools:
  - Read
  - Write
  - Edit
  - Glob
  - Grep
  - Bash
---

# Import existing Site24x7 resources into Terraform

Adopting an existing account is a **one-way loop**: inventory → import into state →
generate config → reconcile the diff → repeat until `terraform plan` is empty. Nothing is
adopted until the plan is empty. Do not stop at "the import command succeeded."

Never run `terraform apply` during an adoption unless the plan has been read line by line.
An apply against a half-written config **overwrites live monitors**.

## Phase 0 — Establish a working provider block

Create (or confirm) a working directory that is the customer's own Terraform project.
Do **not** work inside a clone of `terraform-provider-site24x7` — the README's import
walkthrough tells you to, and it leads to state and config living in the wrong repo.

```hcl
terraform {
  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      version = "~> 1.0"   # pin; the provider is on the 1.x line
    }
  }
}

provider "site24x7" {
  # Read from SITE24X7_OAUTH2_CLIENT_ID / _CLIENT_SECRET / _REFRESH_TOKEN
  data_center = "US"       # US | EU | IN | AU | CN | JP | CA
  # zaaid     = "1234"     # only for a customer under an MSP or BU
}
```

Two failure modes to rule out **before** touching imports, because both surface as
confusing auth or empty-list errors:

- **Data-center mismatch.** OAuth credentials are scoped to the DC they were generated in.
  Each DC has its own API base URL and token URL (`site24x7/data_center.go`). Credentials
  from the US DC against `data_center = "EU"` will not work.
- **Missing `zaaid`.** Under an MSP/BU, without `zaaid` you are querying the parent account
  and will see none of the customer's monitors.

Verify with a cheap read before going further:

```hcl
data "site24x7_device_key" "check" {}
output "device_key" { value = data.site24x7_device_key.check.id }
```

`terraform init && terraform apply` on just that. If it returns a key, auth and DC are right.

## Phase 1 — Inventory the account

`site24x7_monitors` returns IDs and `id__name` pairs, filtered by `monitor_type` **or**
`name_regex` (the data source handles one filter at a time — passing both takes the
`monitor_type` branch only).

```hcl
data "site24x7_monitors" "all" {
  monitor_type = "SERVER"
}
output "inventory" { value = data.site24x7_monitors.all.ids_and_names }
```

Monitor type codes (`api/constants.go` — these are the exact accepted strings):

| Code | Resource | Code | Resource |
|---|---|---|---|
| `URL` | `site24x7_website_monitor` | `PORT` | `site24x7_port_monitor` |
| `HOMEPAGE` | `site24x7_web_page_speed_monitor` | `PING` | `site24x7_ping_monitor` |
| `SSL_CERT` | `site24x7_ssl_monitor` | `SOAP` | `site24x7_soap_monitor` |
| `RESTAPI` | `site24x7_rest_api_monitor` | `ISP` | `site24x7_isp_monitor` |
| `RESTAPISEQ` | `site24x7_rest_api_transaction_monitor` | `FTP` | `site24x7_ftp_transfer_monitor` |
| `SERVER` | `site24x7_server_monitor` | `DNS` | `site24x7_dns_server_monitor` |
| `CRON` | `site24x7_cron_monitor` | `DOMAINEXPIRY` | `site24x7_domain_expiry_monitor` |
| `HEARTBEAT` | `site24x7_heartbeat_monitor` | `REALBROWSER` | `site24x7_web_transaction_browser_monitor` |
| `AMAZON` | `site24x7_amazon_monitor` | `GCP` | `site24x7_gcp_monitor` |
| `AZURE` | `site24x7_azure_monitor` | | |

Also inventory the things monitors *depend on*, and import those **first** — a monitor
referencing a location/notification/threshold profile that isn't in state will churn
forever: `site24x7_location_profile`, `site24x7_notification_profile`,
`site24x7_threshold_profile`, `site24x7_user_group`, `site24x7_tag`,
`site24x7_monitor_group`. Each has a matching data source for discovery, as do
`site24x7_user`, `site24x7_attribute_alert_group`, `site24x7_credential_profile`,
`site24x7_oauth2_provider`, `site24x7_sla_setting`, `site24x7_it_automation` (the data
source for the `site24x7_url_action` resource — the names differ) and `site24x7_customer`.

**Seven resource families have no data source at all**, so nothing in Terraform can discover
them by name: subgroups, business hours, scheduled maintenance, scheduled reports, all seven
integrations, and milestone markers. Their IDs have to come from the Site24x7 UI or a direct
API call. The full resource-to-data-source map is in the `site24x7-alerting-and-profiles`
skill; don't promise the customer name-based discovery for anything on that list.

Write the inventory to a scratch file (not into the customer's `main.tf`) so the ID list
survives across steps.

### APM Insight needs its own inventory

APM Insight resources are **not monitors**. They come from a different API
(`/apminsight/*`), have no `monitor_type` code, and never appear in `site24x7_monitors` —
so an adoption driven only by the table above silently leaves the entire APM estate
unmanaged. Inventory it separately:

```hcl
data "site24x7_apm_applications" "all" {}
data "site24x7_apm_agent_config_profiles" "all" {}

output "apm_applications" { value = data.site24x7_apm_applications.all.ids_and_names }
output "apm_profiles"     { value = data.site24x7_apm_agent_config_profiles.all.ids_and_names }
```

Two importable resources come out of this:

- `site24x7_apm_agent_config_profile` — import by profile ID. Read repopulates every
  setting, so an imported profile needs no hand-fixing.
- `site24x7_apm_application` — import by application ID. Adopt-only; see Phase 4.

APM **instances are deliberately read-only** (`site24x7_apm_instance` /
`site24x7_apm_instances` are data sources only): instance IDs change on every autoscale,
so a resource pinned to one would churn. Do not promise the customer instance-level
Terraform management.

## Fast path — one named resource

When the user names a single monitor ("import the SSL monitor for example.com", "adopt
monitor 123456000025786003"), do not run the full adoption loop. Do this instead:

1. Bootstrap the provider block if the directory has none (Phase 0), and `terraform init`.
2. Resolve the ID if you were given a name, using `site24x7_monitors` with `name_regex`.
3. Write one `import` block, generate, plan:

```hcl
import {
  to = site24x7_ssl_monitor.example_com
  id = "123456000025786003"
}
```

```sh
terraform plan -generate-config-out=generated.tf
```

4. Move the generated resource into a sensibly named file, delete the empty-string and
   zero-value attributes it emitted, and re-plan until it reports no changes.
5. Leave the `import` block in place until the plan is clean, then remove it — the
   resource is in state and the block has done its job.

Then stop. Phases 1 and 3 below are for adopting an account in bulk; running them for one
monitor wastes the user's time. The **plan must still reach "No changes"** — that part is
not optional at any scale.

## Phase 2 — Choose the import mechanism

Decide once, per project, in this order:

**A. Terraform ≥ 1.5 `import` blocks + config generation — the default. Use this.**

```hcl
import {
  to = site24x7_server_monitor.web_01
  id = "123456000025786003"
}
```

```sh
terraform plan -generate-config-out=generated.tf
```

Terraform writes real HCL for every `import` block and tells you exactly what still
differs. It works for **42 of the 46 resources** — every resource whose Go definition
carries a `ResourceImporter`.

**B. `terraform import` CLI — only for Terraform < 1.5.** Requires an empty
`resource "..." "..." {}` stub to exist first, and generates no configuration; you write
every attribute by hand from `terraform show -json`.

**C. `utilities/importer/site24x7_importer.py` — legacy; avoid unless A and B are
impossible.** Be honest with the customer about its limits rather than sending them down it:

- Its `Site24x7TerraformResourceVsMonitorType` map has **6 entries**. `--resource` anything
  else raises a `KeyError` on startup. Supported: `site24x7_website_monitor`,
  `site24x7_ssl_monitor`, `site24x7_rest_api_monitor`,
  `site24x7_rest_api_transaction_monitor`, `site24x7_server_monitor`, `site24x7_tag`.
- `conf/resource_vs_attribute_types.json` holds 10 keys, two of which
  (`Site24x7_web_transaction_browser_monitor`, `Site24x7_ftp_transfer_monitor`) are
  capitalised and can never match the lowercase resource name from `--resource`.
- It shells out to `terraform import` from `utilities/importer/`, so state lands in the
  provider repo, not the customer's project.
- It reads `monitors_to_import.json` from its own directory and exits 1 if absent.

If the customer is already mid-flight on C, the recovery is: take the IDs out of
`monitors_to_import.json`, discard the rest, and restart at A in their own project.

### Resources with no import support at all

Four registered resources have **no `ResourceImporter`**, so no mechanism can import them.
They must be recreated in Terraform and the originals deleted in the UI (say this
explicitly — it means a monitoring gap the customer has to schedule):

- `site24x7_amazon_monitor`
- `site24x7_gcp_monitor`
- `site24x7_azure_monitor`
- `site24x7_milestone_marker`

## Phase 3 — Generate configuration

Import in dependency order, in batches of ~20 so a failure is cheap to diagnose:

1. `site24x7_tag`, `site24x7_user`, `site24x7_attribute_alert_group`
2. `site24x7_user_group` — it requires **both** a `users` list and an `attribute_group_id`, so
   it cannot come before step 1
3. `site24x7_businesshour`, `site24x7_location_profile`, `site24x7_notification_profile`,
   `site24x7_threshold_profile`, `site24x7_credential_profile`, `site24x7_oauth2_provider`
4. `site24x7_monitor_group`, then `site24x7_subgroup`
5. Monitors
6. Everything that points *at* monitors: integrations (`site24x7_*_integration`),
   `site24x7_schedule_maintenance`, `site24x7_schedule_report`, `site24x7_sla_setting`,
   `site24x7_url_action`
7. APM: `site24x7_apm_agent_config_profile` (depends on nothing), then
   `site24x7_apm_application`

`site24x7_milestone_marker` cannot be imported at any point — see the list above.

The required attributes and the reasoning behind this order are in the
`site24x7-alerting-and-profiles` skill, which covers all 25 non-monitor resources.

Give resources meaningful Terraform names. The Python importer produces
`SERVER_123456000025786003`; derive names from `ids_and_names` (the `id__name` pairs)
instead — `web_01`, not the monitor ID. Renaming later costs a `moved` block or a state
surgery.

## Phase 4 — Reconcile until the plan is empty

Run `terraform plan` and drive the diff to zero. The recurring causes, in the order they
show up:

- **Computed attributes written into config.** `location_profile_id`,
  `notification_profile_id` and `threshold_profile_id` are `Optional + Computed`. Generated
  config includes them; that is correct and should stay. But do **not** also set the
  `_name` variants — see the next point.
- **`*_name` and `*_id` both present.** If `location_profile_name` is set, the provider
  resolves it and **overwrites** `location_profile_id` on every apply. During adoption keep
  the IDs and delete the `_name` attributes; switch to names later, deliberately, using the
  `site24x7-monitor-authoring` skill.
- **Empty strings and zero values.** The Site24x7 API omits unset fields; Terraform writes
  `""` / `0`. Delete those lines rather than fighting them.
- **List ordering.** `monitor_groups`, `user_group_ids`, `tag_ids` come back in API order.
  Match the order the API returns, or accept a one-time re-order apply.
- **Secrets.** `auth_pass`, `client_certificate_password`, `password`, `private_key`, `token`,
  `client_secret` and `access_token` never come back from the API. Generated config will show
  them empty. Move them to variables sourced from a vault before the first apply, or the apply
  will blank them out on the live monitor or integration. Only 6 of the 11 attribute/resource
  pairs are marked `Sensitive`, so the rest also land in your plan output and CI logs in
  cleartext — the exhaustive table is in the `site24x7-alerting-and-profiles` skill.
- **APM applications are adopt-only.** `site24x7_apm_application` has no create endpoint —
  an application exists because an agent reported in. The resource takes over its
  managed/suspended lifecycle, and `managed` is the only attribute with a remote effect.
  Keep the generated `delete_on_destroy = false`: at `true`, destroy permanently deletes
  the application *and its monitoring history*.
- **Two APM profiles both claiming the default.** Only one profile per agent type can be
  the default, and promoting one demotes the other without Terraform seeing it. If more
  than one imported profile of the same `agent_type` has `is_default = true` in config, the
  plan never settles — each apply undoes the last. Set it on exactly one per agent type.

Stop only when `terraform plan` says **"No changes."**

## Phase 5 — Hand over

- Move credentials out of HCL entirely (env vars or a vault) — the repo's own examples say
  "always best practice to store your credentials in a Vault of your choice."
- Replace the flat generated file with `for_each` over a map where monitors are homogeneous.
- Commit state backend config; local state after an adoption of hundreds of monitors is a
  liability.
- Record which resources could not be imported (the four above) and what the customer
  decided to do about them.
- Warn about the two destroy surprises before handing over state: `site24x7_customer`'s delete
  is a no-op (the customer account survives a `terraform destroy`), and
  `site24x7_apm_application` with `delete_on_destroy = true` deletes the application and its
  history. Both are detailed in the `site24x7-alerting-and-profiles` skill.

## Verification checklist

Before declaring the import done, confirm all of:

- [ ] `terraform plan` → "No changes."
- [ ] Resource count in state matches the inventory count from Phase 1
      (`terraform state list | wc -l`).
- [ ] APM inventoried separately from monitors — `site24x7_monitors` does not return it.
- [ ] `terraform plan` run a second time, after `terraform refresh`, is still empty
      (catches server-side defaulting that only shows on re-read).
- [ ] No credential literal anywhere in `*.tf` (`grep -rE 'oauth2_|auth_pass|secret' *.tf`).
- [ ] Every non-importable resource is listed in the handover notes.
