---
name: site24x7-import-existing
description: >
  Bring monitors, profiles, tags and integrations that already exist in a Site24x7
  account under Terraform management, producing HCL that plans clean. Inventories the
  account with the site24x7_monitors data source, picks the right import mechanism per
  resource type, generates configuration, and iterates until `terraform plan` reports no
  changes. Use whenever someone wants to import existing Site24x7 monitors, import every
  monitor of one type (all website / SERVER / SSL monitors), import a
  monitor group together with the monitors in it, adopt a hand-built (ClickOps) Site24x7
  account into Terraform, migrate monitors into state, run site24x7_importer.py, or asks
  "how do I get my existing monitors into Terraform" — even if they don't say "skill".
user-invocable: true
argument-hint: "<what to import, e.g. 'all SERVER monitors', 'monitor group PROD_WEB and its monitors' or 'the 40 website monitors tagged prod'>"
allowed-tools:
  - Read
  - Write
  - Edit
  - Glob
  - Grep
  - Bash
  - PowerShell
---

# Import existing Site24x7 resources into Terraform

Adopting an existing account is a **one-way loop**: inventory → import into state →
generate config → reconcile the diff → repeat until `terraform plan` is empty. Nothing is
adopted until the plan is empty. Do not stop at "the import command succeeded."

Never run `terraform apply` during an adoption unless the plan has been read line by line.
An apply against a half-written config **overwrites live monitors**.

## Running Terraform during an import

These rules apply to every phase and every fast path. Breaking them is how an import ends
up with two processes fighting over one state file and a session polling with `sleep`.

**One Terraform process per directory, in the foreground.**

- Never put `terraform import`, `plan` or `apply` in a shell loop, a script run in the
  background, or `run_in_background`. One `terraform plan -generate-config-out` imports every
  `import` block in the directory, so a loop is never needed.
- Never start Terraform while another Terraform process is running in the same directory.
  Two processes write the same `terraform.tfstate`; one fails on the state lock and the other
  can lose imports.
- A large plan can take minutes. Give the command a long timeout (up to 10 minutes) and wait.
  If it is still running, report that to the user. Do not start a second run and do not poll
  with `sleep`.
- `Error acquiring the state lock` means another process holds it. Find that process and tell
  the user. Run `terraform force-unlock` only after the user confirms no Terraform process is
  still running.
- If a process has to be stopped on Windows, note that Git Bash `kill` takes Git Bash process
  IDs, while `ps -W` and Task Manager show Windows process IDs. Stop a Windows process ID with
  PowerShell `Stop-Process -Id <id>`. If that returns "Access is denied", stop and ask the user
  to end it. Do not retry.

**Credentials from a `.env` file.** Terraform and the provider never read `.env`. The provider
reads only the `SITE24X7_OAUTH2_CLIENT_ID`, `SITE24X7_OAUTH2_CLIENT_SECRET` and
`SITE24X7_OAUTH2_REFRESH_TOKEN` environment variables. Each shell command starts a new shell,
so load the file **in the same command** that runs Terraform:

```sh
set -a; . <(tr -d '\r' < .env); set +a; terraform plan
```

```powershell
Get-Content .env | Where-Object { $_ -match '^\s*(?:export\s+)?([A-Za-z_][A-Za-z0-9_]*)\s*=\s*(.*)$' } |
  ForEach-Object { Set-Item "Env:$($Matches[1])" $Matches[2].Trim().Trim('"').Trim("'") }
terraform plan
```

`tr -d '\r'` matters: a `.env` saved with Windows line endings puts an invisible carriage
return at the end of each value, and OAuth then fails. Never print the values, and never copy
them into a `.tf` file.

**A run that looks hung is usually retrying.** Each failed API call is retried up to
`max_retries` times (default 4) with waits of up to `retry_max_wait` seconds (default 30), so
bad credentials or a wrong `data_center` can make a plan sit for minutes before it fails.
Prove authentication with the cheap read in Phase 0 before any import.

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
  data_center = "US"       # US | EU | IN | AU | CN | JP | CA | UK | AE | SA
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

`site24x7_monitors` returns IDs and `id__name` pairs, filtered by `monitor_type`,
`name_regex`, or both. It **fails the plan** with `Unable to find monitors matching …` when
nothing matches, so declare it only for types the account is known to have. It cannot filter
by monitor group; for a group, use the monitor group fast path below. To import every
monitor of one type, use the monitor type fast path below.

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

**These resource families have no data source at all**, so nothing in Terraform can discover
them by name: subgroups, business hours, scheduled reports, all seven integrations, and
milestone markers. (Scheduled maintenances can be inventoried with the
`site24x7_schedule_maintenances` data source.) Their IDs have to come from the Site24x7 UI or a direct
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
   zero-value attributes it emitted, and re-plan until it reports
   `1 to import, 0 to add, 0 to change, 0 to destroy`.
5. Run `terraform apply` and read the plan it shows: it must be the same import-only plan.
   An import-only apply writes state and changes nothing in Site24x7. **Generating config does
   not import anything.** The monitor is in state only after this apply.
6. Remove the `import` block. `terraform plan` must now say "No changes".

Then stop. Phases 1 and 3 below are for adopting an account in bulk; running them for one
monitor wastes the user's time. The **plan must still reach "No changes"** — that part is
not optional at any scale.

## Fast path — one monitor group and its monitors

When the user names a monitor group ("import the monitors in group PROD_WEB, including the
group"), import the group and its monitors in one pass. Do not sweep the whole account.

**How membership works.** The `site24x7_monitor_group` resource does not manage which
monitors are in the group; it has no `monitors` attribute, and an update keeps the members
the group already has. Membership is set on each monitor through its `monitor_groups`
attribute. So importing the group adds or removes no members, and the monitors' generated
config is what keeps them in the group.

1. **Bootstrap and prove auth** (Phase 0), and `terraform init`.

2. **Resolve the group** with the `site24x7_monitor_group` data source, in a scratch file:

   ```hcl
   data "site24x7_monitor_group" "target" {
     name_regex = "^PROD_WEB$"
   }

   output "group" {
     value = {
       id           = data.site24x7_monitor_group.target.id
       display_name = data.site24x7_monitor_group.target.display_name
       monitors     = data.site24x7_monitor_group.target.monitors
     }
   }
   ```

   `name_regex` is case-insensitive and **unanchored, and the first match wins**. `PROD_WEB`
   alone would also match `PROD_WEB_DR`, and you would import the wrong group without an
   error. Always anchor it with `^…$`, and escape regex characters in the name with a double
   backslash in HCL (`"^App\\.Prod \\(EU\\)$"`). Read the output back to the user: the
   `display_name` must be exactly the group they named. The data source's `id` is the group
   ID, and `monitors` is the set of monitor IDs in the group.

   `monitors` holds direct members only. Subgroups are separate `site24x7_subgroup` resources
   with no data source. Ask whether the user wants them; their IDs come from the Site24x7 UI.

3. **Find each monitor's type.** `monitors` returns IDs only, but an `import` block needs the
   resource type. Make one read-only API call that lists every monitor with its type, and
   keep the members of the group:

   ```powershell
   # Credentials already in the environment (see "Credentials from a .env file").
   # Use the URLs for the account's data_center: site24x7/data_center.go has the table.
   $api   = "https://www.site24x7.com/api"            # EU www.site24x7.eu, IN www.site24x7.in, ...
   $token = (Invoke-RestMethod -Method Post "https://accounts.zoho.com/oauth/v2/token" -Body @{
       client_id     = $env:SITE24X7_OAUTH2_CLIENT_ID
       client_secret = $env:SITE24X7_OAUTH2_CLIENT_SECRET
       refresh_token = $env:SITE24X7_OAUTH2_REFRESH_TOKEN
       grant_type    = "refresh_token" }).access_token
   $headers = @{ Authorization = "Zoho-oauthtoken $token"; Accept = "application/json; version=2.0" }
   $session = New-Object Microsoft.PowerShell.Commands.WebRequestSession
   # Under an MSP or BU only: send the customer's zaaid as a cookie, as the provider does.
   # $session.Cookies.Add((New-Object System.Net.Cookie("zaaid", "<zaaid>", "/", ([uri]$api).Host)))
   $members = @("<monitor id>", "<monitor id>")       # from the group output in step 2
   (Invoke-RestMethod "$api/monitors" -Headers $headers -WebSession $session).data |
     Where-Object { $members -contains $_.monitor_id } |
     Select-Object monitor_id, display_name, type
   ```

   ```sh
   # POSIX shell with curl and python3. Same URLs as above.
   # Under an MSP or BU, add -b "zaaid=<zaaid>" to the /monitors call.
   token=$(curl -s -X POST "https://accounts.zoho.com/oauth/v2/token" \
     -d client_id="$SITE24X7_OAUTH2_CLIENT_ID" -d client_secret="$SITE24X7_OAUTH2_CLIENT_SECRET" \
     -d refresh_token="$SITE24X7_OAUTH2_REFRESH_TOKEN" -d grant_type=refresh_token |
     python3 -c 'import json,sys; print(json.load(sys.stdin)["access_token"])')
   curl -s "https://www.site24x7.com/api/monitors" \
     -H "Authorization: Zoho-oauthtoken $token" -H "Accept: application/json; version=2.0" |
     python3 -c 'import json,sys
   members = set(sys.argv[1:])
   for m in json.load(sys.stdin)["data"]:
       if m["monitor_id"] in members: print(m["monitor_id"], m["type"], m["display_name"], sep="\t")' \
     <monitor id> <monitor id>
   ```

   Keep the token in a variable: never print it or write it to a file.

   Map each `type` to its resource with the type table in Phase 1, and write the result to a
   scratch inventory file. Then report two lists to the user before going on:
   - Monitors of type `AMAZON`, `GCP` or `AZURE`: they cannot be imported (see Phase 2).
   - Group member IDs that `/monitors` did not return: they are not monitors this resource
     family can import.

4. **Write the `import` blocks**, the group first, then one per importable monitor, named from
   the display names (`web_01`, not the ID; make names unique):

   ```hcl
   import {
     to = site24x7_monitor_group.prod_web
     id = "<group id>"
   }

   import {
     to = site24x7_website_monitor.checkout
     id = "<monitor id>"
   }
   ```

5. **Generate once**, in the foreground (see "Running Terraform during an import"):

   ```sh
   terraform plan -generate-config-out=generated.tf
   ```

   For a group of more than about 20 monitors, write the blocks in batches of ~20 and take each
   batch through steps 5–7 before starting the next. `-generate-config-out` refuses to write
   to a file that already exists, so give each batch its own file (`generated_01.tf`,
   `generated_02.tf`, …).

6. **Reconcile** with Phase 4 until the plan shows only imports
   (`N to import, 0 to add, 0 to change, 0 to destroy`), plus two group-specific steps:
   - In each monitor's `monitor_groups`, replace the literal group ID with
     `site24x7_monitor_group.prod_web.id`. Keep any other group IDs the monitor belongs to.
   - Leave the group's `dependency_resource_ids` as literal IDs. The monitors already depend
     on the group, so referencing a monitor from the group creates a cycle.

   The monitors' profiles, user groups and tags are usually shared with the rest of the
   account. Keep them as literal IDs unless the user asks to adopt those too; importing a
   shared profile makes this project its owner for every team that uses it.

7. **Apply the imports.** Run `terraform apply` and check that the plan it shows is the same
   import-only plan. It writes state and changes nothing in Site24x7. Generating config does
   not import anything; nothing is in state until this apply.

8. **Verify**: `terraform state list` shows the group plus every importable monitor. Remove
   the `import` blocks, and `terraform plan` must say "No changes".

## Fast path — every monitor of one type

When the user names a monitor type ("import all our website monitors", "adopt the SERVER
monitors", "all SSL monitors whose name starts with prod-"), import that type only. Do not
sweep the whole account.

1. **Bootstrap and prove auth** (Phase 0), and `terraform init`.

2. **Map the request to one type code and one resource** with the type table in Phase 1:
   "website" is `URL` → `site24x7_website_monitor`, "web page speed" is `HOMEPAGE`, "REST API"
   is `RESTAPI` and "REST API transaction" is `RESTAPISEQ`. If the user names a type the
   table doesn't have, ask; do not guess. If the type is `AMAZON`, `GCP` or `AZURE`, stop
   here and explain: those resources have no importer (see Phase 2).

3. **Inventory that type in a separate directory**, never in the customer's project. The
   inventory needs a `terraform apply`, and an apply in the customer's project would also
   apply any change already pending there. Create `inventory/` next to the project, copy the
   `terraform` and `provider` blocks into it, and add:

   ```hcl
   data "site24x7_monitors" "selected" {
     monitor_type = "URL"
     # name_regex = "^prod-"   # optional, to narrow it down
   }

   output "inventory" { value = data.site24x7_monitors.selected.ids_and_names }
   ```

   ```sh
   terraform -chdir=inventory init
   terraform -chdir=inventory apply     # declares no resources, so it only reads
   terraform -chdir=inventory output -json inventory
   ```

   How the data source matches (`site24x7/monitors/monitors_data_source.go`):
   - `monitor_type` must match the code exactly, in upper case. `url` matches nothing.
   - `name_regex` is **case-sensitive** and unanchored. That is the opposite of the monitor
     group data source. Prefix it with `(?i)` for a case-insensitive match.
   - When nothing matches, the plan fails with `Unable to find monitors matching …`. Tell the
     user the account has no monitors of that type. That is not an auth failure.

   Report the count and the display names to the user and confirm before writing anything.
   For more than ~20 monitors, agree on a batch size.

4. **Leave out monitors the project already manages.** An `import` block that points a new
   address at a monitor that is already in state does not fail. It produces a second resource
   that manages the same monitor, and the two then overwrite each other. Collect the IDs
   already in state:

   ```sh
   terraform show -json > state.json        # in the customer's project; {} if no state yet
   ```

5. **Write the `import` blocks** from the inventory. Name each block after the display name
   (`checkout_api`, not the ID), and make the names unique. For more than a handful of
   monitors, generate the blocks with a script instead of typing them:

   ```powershell
   $resource = "site24x7_website_monitor"
   $state    = Get-Content state.json -Raw | ConvertFrom-Json
   $managed  = @($state.values.root_module.resources | ForEach-Object { $_.values.id })
   $pairs    = terraform -chdir=inventory output -json inventory | ConvertFrom-Json
   $seen = @{}
   $blocks = foreach ($p in $pairs) {
     $id, $name = $p -split '__', 2
     if ($managed -contains $id) { Write-Host "already managed: $id $name"; continue }
     $tf = ($name.ToLower() -replace '[^a-z0-9]+', '_').Trim('_')
     if ($tf -notmatch '^[a-z]') { $tf = "m_$tf" }
     $base = $tf; $n = 2
     while ($seen.ContainsKey($tf)) { $tf = "${base}_$n"; $n++ }
     $seen[$tf] = $true
     "import {`n  to = $resource.$tf`n  id = `"$id`"`n}`n"
   }
   # WriteAllText writes UTF-8 without a BOM; Set-Content in Windows PowerShell does not.
   [IO.File]::WriteAllText("$PWD\imports.tf", ($blocks -join "`n"))
   ```

   ```sh
   terraform -chdir=inventory output -json inventory | python3 -c '
   import json, re, sys
   resource = sys.argv[1]
   state = json.load(open("state.json", encoding="utf-8-sig"))
   managed = {r["values"].get("id") for r in
              state.get("values", {}).get("root_module", {}).get("resources", [])}
   seen = set()
   for pair in json.loads(sys.stdin.buffer.read().decode("utf-8-sig")):
       mid, name = pair.split("__", 1)
       if mid in managed:
           print("already managed:", mid, name, file=sys.stderr); continue
       tf = re.sub(r"[^a-z0-9]+", "_", name.lower()).strip("_")
       if not re.match(r"[a-z]", tf): tf = "m_" + tf
       base, n = tf, 2
       while tf in seen: tf, n = f"{base}_{n}", n + 1
       seen.add(tf)
       print(f"import {{\n  to = {resource}.{tf}\n  id = \"{mid}\"\n}}\n")
   ' site24x7_website_monitor > imports.tf
   ```

   The ID is everything before the first `__`; display names may contain `__` themselves.
   Report the "already managed" lines to the user. For more than ~20 monitors, keep one
   batch in `imports.tf`, park the rest in a scratch file, and take each batch through steps
   6–8 before starting the next.

6. **Generate**, in the foreground (see "Running Terraform during an import"), with a new file
   per batch, because `-generate-config-out` refuses to overwrite one:

   ```sh
   terraform plan -generate-config-out=generated_01.tf
   ```

7. **Reconcile** with Phase 4 until the plan shows only imports
   (`N to import, 0 to add, 0 to change, 0 to destroy`). One type means one schema, so the
   same fix usually applies to every monitor in the batch: fix one, then apply the same edit
   to the rest. Keep profile, user group, tag and `monitor_groups` IDs as literals unless the
   user asks to adopt those too. They are shared with the rest of the account.
   Pay attention to secrets: for website, REST API, web page speed and REST API transaction
   monitors, `auth_pass` comes back empty. Move it to a variable before the apply, or the apply
   blanks it on the live monitor.

8. **Apply the imports.** Run `terraform apply` and check that the plan it shows is the same
   import-only plan. It writes state and changes nothing in Site24x7. Then remove that batch's
   `import` blocks; `terraform plan` must say "No changes".

9. **Verify**: `terraform state list | grep <resource type>` count equals the inventory count
   minus the already-managed ones, and `terraform plan` is still "No changes". Delete
   `inventory/` and `state.json`. Then suggest Phase 5's `for_each` refactor: monitors of a
   single type are where it pays off most.

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
  will blank them out on the live monitor or integration. 15 of the 19 attribute/resource
  pairs are marked `Sensitive`. The exceptions are `auth_pass` on the website, REST API,
  web page speed and REST API transaction monitors, which still land in your plan output and
  CI logs in cleartext — the exhaustive table is in the `site24x7-alerting-and-profiles` skill.
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
- [ ] No Terraform process left running, and no `.terraform.tfstate.lock.info` left behind.
- [ ] For a monitor group import: the group's members in Site24x7 are unchanged, and every
      imported monitor's `monitor_groups` references the group resource.
- [ ] Every non-importable resource is listed in the handover notes.
