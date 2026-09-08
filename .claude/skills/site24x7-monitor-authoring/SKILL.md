---
name: site24x7-monitor-authoring
description: >
  Write and review Site24x7 Terraform configuration — choosing the right monitor resource,
  wiring location / notification / threshold profiles, user groups and tags without
  hand-copying 18-digit IDs, and scaling to many monitors with for_each. Encodes the
  provider's name-resolution and defaulting behaviour, which is substring-based and
  silently non-deterministic if used naively. Use whenever someone wants to add, edit or
  review a Site24x7 monitor in Terraform, asks which site24x7_* resource to use, hits an
  "Unable to find ... matching the string" error, sees a monitor land on the wrong profile,
  or wants to bulk-create monitors — even if they don't say "skill".
user-invocable: true
argument-hint: "<what to monitor, e.g. 'the 12 prod API endpoints, alert the payments on-call'>"
allowed-tools:
  - Read
  - Write
  - Edit
  - Glob
  - Grep
  - Bash
---

# Authoring Site24x7 Terraform configuration

The provider is friendlier than it looks: it will resolve profile and group **names** to
IDs for you, and will invent defaults for anything you leave out. Both behaviours are
traps if you don't know their exact rules. This skill is mostly about those rules.

This skill covers **monitors** and APM. For the resources a monitor points at or alerts
through — location / notification / threshold profiles, monitor groups and subgroups, user
groups, users, attribute alert groups, tags, business hours, the seven integrations, credential
profiles, OAuth2 providers, scheduled maintenance and reports, SLA settings, milestone markers,
IT automation actions and MSP customers — use `site24x7-alerting-and-profiles`, which carries
their required attributes and the order they have to be created in.

## Step 0 — Bootstrap, if the directory has no provider configuration yet

Check for a `terraform` block before writing any resource. If there isn't one (an empty
directory, or a project that has never used this provider), write `versions.tf` first —
do not ask the user to create it, just write it and say you did:

```hcl
terraform {
  required_version = ">= 1.5.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      version = "~> 1.0"
    }
  }
}

provider "site24x7" {
  # Credentials come from SITE24X7_OAUTH2_CLIENT_ID / _CLIENT_SECRET / _REFRESH_TOKEN.
  # Never write them into a .tf file.
  data_center = "US" # US | EU | IN | AU | CN | JP | CA — must match where the
                     # credentials were issued, or auth fails confusingly
}
```

Confirm the three environment variables are set (`env | grep SITE24X7` or the PowerShell
equivalent) before running anything, and pick `data_center` from what the user tells you —
ask if they haven't said, because guessing costs a confusing auth failure later.

Then `terraform init`, and carry on to Step 1.

## Step 1 — Pick the resource

Ask what is being checked, not what the customer calls it.

| Intent | Resource |
|---|---|
| An HTTP(S) page is up / contains text | `site24x7_website_monitor` |
| Page load performance, real browser | `site24x7_web_page_speed_monitor` |
| A multi-step browser journey | `site24x7_web_transaction_browser_monitor` |
| A single API endpoint, status/JSON assertions | `site24x7_rest_api_monitor` |
| A chained sequence of API calls | `site24x7_rest_api_transaction_monitor` |
| A SOAP endpoint | `site24x7_soap_monitor` |
| TLS certificate validity/expiry | `site24x7_ssl_monitor` |
| Domain registration expiry | `site24x7_domain_expiry_monitor` |
| A host via the Site24x7 agent | `site24x7_server_monitor` |
| A batch job that must check in | `site24x7_cron_monitor` / `site24x7_heartbeat_monitor` |
| DNS records resolve correctly | `site24x7_dns_server_monitor` |
| A TCP/UDP port answers | `site24x7_port_monitor` |
| ICMP reachability | `site24x7_ping_monitor` |
| ISP / network path quality | `site24x7_isp_monitor` |
| FTP upload/download works | `site24x7_ftp_transfer_monitor` |
| AWS / GCP / Azure account discovery | `site24x7_amazon_monitor` / `_gcp_monitor` / `_azure_monitor` |
| Application internals — transactions, traces, slow SQL, apdex | **not a monitor** — APM Insight, see Step 5 |

`cron` vs `heartbeat`: use `site24x7_cron_monitor` when the job has a schedule Site24x7
should expect it against; `site24x7_heartbeat_monitor` when you only care that a ping
arrives within a window.

## Step 2 — The four associations every monitor has

Every monitor carries `location_profile_id`, `notification_profile_id`,
`threshold_profile_id`, plus `user_group_ids` and `tag_ids`. Prefer never writing a raw ID.

**None of this applies to the APM Insight resources.** They are not monitors and expose no
location, notification, threshold, user group or tag attributes at all — writing
`location_profile_name` into one is a schema error, not a resolution failure. Skip to
Step 5 for those.

### Name-based attributes (resolved at apply time)

| Write this | Instead of | Available on |
|---|---|---|
| `location_profile_name` | `location_profile_id` | 14 resources |
| `notification_profile_name` | `notification_profile_id` | 19 resources |
| `user_group_names` (list) | `user_group_ids` | 19 resources |
| `tag_names` (list) | `tag_ids` | 19 resources |

Not every monitor type exposes every one — `grep '"location_profile_name"'` in the
resource's Go file, or check its page under `docs/resources/`, before assuming.

**There is no `threshold_profile_name`.** Threshold profiles are ID-only — resolve them
with the `site24x7_threshold_profile` data source, or leave the attribute out entirely and
let the provider pick by monitor type (Step 3).

### The matching rules — read these before writing a name

All name matching is `strings.Contains`, not equality (`site24x7/monitor_defaults.go`):

- **`location_profile_name` and `notification_profile_name`: substring, last match wins.**
  The loop has no early exit, so if `"US"` matches both "US - East" and "US - West", the
  monitor silently gets whichever the API returned *last*. Always pass a string that is
  unique across the account, or use the data source and pass an ID.
- **`user_group_names` and `tag_names`: substring, union of all matches.** `"prod"` attaches
  *every* group or tag containing "prod" — `prod-db`, `prod-web`, `non-prod`. This is
  occasionally what people want and usually not.
- **No match is a hard error**, e.g. `Unable to find location profile matching the string:
  "..." in Site24x7`. That error means the substring matched nothing — check for typos and
  case (matching is case-sensitive) before assuming the profile is missing.

### What happens when you omit them

Omission is not "no association" — the provider fills it in and writes the result to state:

- `location_profile_id` omitted → **`locationProfiles[0]`**, the first profile the API
  happens to return. Not deterministic across accounts, and it will differ between the
  customer's staging and production tenants.
- `notification_profile_id` omitted → **`notificationProfiles[0]`**, same caveat.
- `user_group_ids` omitted → **`userGroups[0]`**. A monitor with no explicit group is a
  monitor whose alerts go somewhere arbitrary.
- `threshold_profile_id` omitted → the first profile **whose type matches the monitor type**
  (`URL`, `RESTAPI`, `SERVER`, …). If the account has no threshold profile of that type the
  apply fails with `Please configure threshold profile for the monitor type: <TYPE>`. This
  is the one default that is actually sensible; leaving it out is fine and usually right.

So: always set location, notification and user group explicitly. Threshold may be omitted.

## Step 3 — Write it

Resolve the pieces once with data sources, then reference them. This is the shape to reach
for:

```hcl
data "site24x7_location_profile" "eu" {
  name_regex = "^EU - All$"
}

resource "site24x7_rest_api_monitor" "payments_api" {
  display_name = "payments-api / health"
  website      = "https://api.example.com/health"

  check_frequency = "5"          # minutes, as a string
  timeout         = 10

  location_profile_id       = data.site24x7_location_profile.eu.id
  notification_profile_name = "Payments — Sev1"   # unique substring, verified
  user_group_names          = ["payments-oncall"]
  tag_names                 = ["prod", "team-payments"]

  # threshold_profile_id omitted: provider selects the RESTAPI-type profile
}
```

Note `check_frequency` is a **string** on monitor resources, not a number — a very common
plan-time type error.

### Many monitors

Use `for_each` over a map, never `count` over a list. The repo's own
`examples/bulk_resource_addition_us.tf` uses `count` with `display_name = count.index`;
don't copy it. Removing one entry from a `count` list re-indexes and destroys/recreates
every monitor after it, and integer display names are unusable in alerts.

```hcl
variable "endpoints" {
  type = map(object({ url = string, team = string }))
  default = {
    payments = { url = "https://api.example.com/payments/health", team = "payments" }
    search   = { url = "https://api.example.com/search/health",   team = "search" }
  }
}

resource "site24x7_website_monitor" "api" {
  for_each = var.endpoints

  display_name              = "api / ${each.key}"
  website                   = each.value.url
  check_frequency           = "5"
  location_profile_id       = data.site24x7_location_profile.eu.id
  notification_profile_name = "Payments — Sev1"
  user_group_names          = ["${each.value.team}-oncall"]
  tag_names                 = ["prod", "team-${each.value.team}"]
}
```

## Step 4 — Validate

```sh
terraform fmt -recursive
terraform validate
terraform plan
```

`validate` will not catch any of the name-resolution problems above — they are apply-time,
server-side lookups. `plan` catches types and required attributes. Only `apply` proves a
name resolved, so on the first apply of a new naming convention, apply **one** monitor and
read back which profile it landed on before rolling out the rest.

Reading resolution decisions:

```sh
TF_LOG=DEBUG terraform apply 2>&1 | grep -i "Finding match\|Match found\|Associating"
```

The provider logs each name→ID resolution at DEBUG. This is the fastest way to prove a
substring matched what you intended.

## Step 5 — APM Insight (not a monitor)

APM Insight watches the inside of an application — transactions, traces, slow SQL, apdex —
via a language agent, not a synthetic check. It has its own API (`/apminsight/*`) and
therefore its own shape. Two resources, six data sources:

| Purpose | Resource / data source |
|---|---|
| Configure agents of one type | `site24x7_apm_agent_config_profile` (resource) |
| Adopt an application's lifecycle | `site24x7_apm_application` (resource) |
| Read a profile, or the default for an agent type | `site24x7_apm_agent_config_profile` (data source) |
| Read all profiles, optionally by agent type | `site24x7_apm_agent_config_profiles` |
| Read applications | `site24x7_apm_application` / `site24x7_apm_applications` |
| Read agent instances | `site24x7_apm_instance` / `site24x7_apm_instances` |

### Agent configuration profiles

Attribute names are the API's dotted `agent_config` keys with the dots replaced by
underscores — `transaction.trace.enabled` becomes `transaction_trace_enabled`. That rule
holds for every setting, so you can read the Site24x7 API docs directly.

Every setting is optional and defaults to the value Site24x7 ships in its own default
profiles, so declare only what differs:

```hcl
resource "site24x7_apm_agent_config_profile" "java_verbose" {
  profile_name = "Java — verbose tracing"
  agent_type   = "JAVA" # JAVA | DOTNET | PHP | RUBY | NODEJS | PYTHON

  transaction_trace_threshold           = 1   # default 2, seconds
  transaction_tracking_request_interval = 1   # 1 = every request, 5 = one in five
  apdex_threshold                       = 0.5 # a float, not an int
}
```

Three things that bite:

- **`is_default = true` demotes the previous default silently.** Only one profile per agent
  type can be the default, and Terraform cannot see the other profile change. Keep it
  `true` on exactly one profile per `agent_type`, or successive applies will fight.
- **`agent_type` is not validated locally**, deliberately, so new Site24x7 agent types keep
  working. A typo therefore surfaces as an API error at apply, not a plan error.
- **`cloud_instance_cleanup_threshold` must be 1–15**; `apdex_threshold` is a float
  (`0.5`), despite the API docs calling it an int.

### Applications

`site24x7_apm_application` **adopts**, it does not create — an application exists because
an agent reported in, and the API has no create endpoint. `managed` is the only attribute
with a remote effect (it calls `manage`/`unmanage`); `time_window` only affects the request
path. Resolve the ID by name rather than pasting it:

```hcl
data "site24x7_apm_application" "checkout" {
  name_regex = "^checkout-api$"
}

resource "site24x7_apm_application" "checkout" {
  application_id    = data.site24x7_apm_application.checkout.application_id
  managed           = true
  delete_on_destroy = false # true makes destroy delete the app and all its history
}
```

Instances are read-only by design: their IDs change on every autoscale, so a resource
pinned to one would churn.

Validate exactly as in Step 4.

## Review checklist

When reviewing existing Site24x7 HCL, flag:

- [ ] `location_profile_name` / `notification_profile_name` whose value is a substring of
      more than one profile in the account (last-match-wins hazard).
- [ ] Both `*_name` and `*_id` set for the same association — the name wins and silently
      overwrites the ID every apply. Pick one.
- [ ] Missing `location_profile_*`, `notification_profile_*` or `user_group_*` — the
      provider will pick element `[0]` and the customer will not know.
- [ ] `count` used where `for_each` belongs.
- [ ] `check_frequency` written as a number.
- [ ] Credentials (`oauth2_*`, `auth_pass`, `client_certificate_password`) as literals in
      `.tf` rather than variables from a vault.
- [ ] No `tag_names` at all — untagged monitors are invisible to every later bulk operation.
- [ ] `data_center` not matching the DC the OAuth credentials were issued in.
- [ ] `is_default = true` on more than one `site24x7_apm_agent_config_profile` sharing an
      `agent_type` — successive applies will undo each other.
- [ ] `delete_on_destroy = true` on a `site24x7_apm_application` — destroy then deletes the
      application and its monitoring history irreversibly.
- [ ] Profile/group/tag attributes written onto an APM resource — they do not exist there.
