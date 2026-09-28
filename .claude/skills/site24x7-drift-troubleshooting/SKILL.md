---
name: site24x7-drift-troubleshooting
description: >
  Diagnose and fix a Site24x7 Terraform plan that never settles — a permanent diff, a
  perpetual "update in-place" after every apply, a list that keeps re-ordering, a password
  that shows as changed every run (or a rotated one that never applies), a value that flips
  back to a server default, a location/notification profile that changes on its own, or a
  resource Terraform wants to recreate every time. Classifies the diff by its shape, tells
  the user the config-side fix they can apply today, and names the provider-side code fix.
  Use whenever someone says "terraform plan always shows changes", "perpetual diff",
  "drift", "plan is never clean", "it keeps wanting to update", "ignore_changes", "after
  import the plan still isn't empty", or pastes a plan with a `~` that comes back after
  apply — even if they don't say "skill".
user-invocable: true
argument-hint: "<the resource and attribute that keeps changing, or paste the plan output>"
allowed-tools:
  - Read
  - Write
  - Edit
  - Glob
  - Grep
  - Bash
  - PowerShell
---

# Troubleshoot Site24x7 drift and perpetual diffs

A diff that comes back after a successful apply has one of two root causes. Decide which
before changing anything:

- **Config-side.** The HCL asks for something the API stores differently: other casing,
  another order, a default the API rewrites. The user can fix this today, in their `.tf`.
- **Provider-side.** The provider's Read maps the API response back into state wrongly: it
  writes a masked secret, uses a list where it should use a set, or declares a `Default` the
  API doesn't honour. The user can only **work around** it (`ignore_changes`). The real fix
  is in this repo's Go code, and the `site24x7-add-resource` skill's Step 8 has the correct
  patterns.

Always give the user **both**: the workaround they can apply now, and the provider fix, with
file and attribute, if they (or you) can change the provider.

Never "fix" drift by running `terraform apply` repeatedly. Each apply may be rewriting a live
monitor, and a secret-blanking diff (Class A) **breaks the monitor's authentication** on
apply.

## Step 1: Reproduce and capture

```sh
terraform plan -out=tfplan               # the diff as reported
terraform show -json tfplan > plan.json  # exact before/after values, including set elements
terraform plan -refresh-only             # what changed remotely, as opposed to what config wants
```

When the diff is on a sensitive attribute the plan just shows `(sensitive value)`. Compare
state instead: `terraform state show -show-sensitive <address>`. Don't paste the output
anywhere public.

To see what the API really sent and returned, run with debug logging. Every request and
response body goes to the log:

```sh
TF_LOG=DEBUG TF_LOG_PATH=tf.log terraform plan
```

(PowerShell: `$env:TF_LOG='DEBUG'; $env:TF_LOG_PATH='tf.log'; terraform plan`.) The log
contains credentials. Delete it when done.

A useful discriminator: **apply once, then plan twice.**

| Second plan | Meaning |
|---|---|
| Same diff every time | Classic perpetual diff. Go to Step 2. |
| Clean, but the diff returns later | Someone or something is changing it remotely: a UI user, another Terraform workspace, or another resource (Class G) |
| Only shows after `terraform refresh` | Server-side defaulting that happens asynchronously after create (Class D) |

## Step 2: Classify the diff by its shape

### Class A: a secret shows as changed on every plan

Shape: `~ password = (sensitive value)`, or for non-sensitive attributes
`~ auth_pass = "****" -> "real"` or `"" -> "real"`, every time.

Cause: the API returns the secret masked, encrypted or empty, and the resource's **Read
writes that value into state**. Config and state never match.

This audit shows how each secret attribute is handled today. Re-run the audit command below
after provider changes.

| Resource | Attribute | `Sensitive` | Diff suppressed | Read writes API value |
|---|---|---|---|---|
| `site24x7_credential_profile` | `password` | yes | no | **yes** |
| `site24x7_oauth2_provider` | `client_secret`, `auth_pass` | yes | no | **yes** |
| `site24x7_connectwise_integration` | `private_key` | yes | no | **yes** |
| `site24x7_telegram_integration` | `token` | yes | no | **yes** |
| `site24x7_ftp_transfer_monitor` | `password` | yes | no | **yes** |
| `site24x7_gcp_monitor` | `private_key` | yes | no | **yes** |
| `site24x7_web_page_speed_monitor` | `auth_pass` | **no** | no | **yes** |
| `site24x7_servicenow_integration` | `password` | yes | partial (see Class B) | yes |
| `site24x7_website_monitor` | `auth_pass` | **no** | always | yes |
| `site24x7_website_monitor` | `client_certificate_password` | yes | always | yes |
| `site24x7_rest_api_monitor` | `auth_pass` | **no** | always | yes |
| `site24x7_rest_api_monitor` | `client_certificate_password` | yes | always | yes |
| `site24x7_webhook_integration` | `password` | yes | always | yes |
| `site24x7_rest_api_transaction_monitor` | `steps` → `auth_pass` | **no** | always | yes (via `d.Set("steps", …)`) |
| `site24x7_rest_api_transaction_monitor` | `steps` → `client_certificate_password` | yes | always | yes (via `d.Set("steps", …)`) |
| `site24x7_web_transaction_browser_monitor` | `auth_details` → `password` | yes (on the map) | no | **yes** (via `d.Set("auth_details", …)`) |
| `site24x7_azure_monitor` | `client_secret` | yes | no | no ✓ |

`site24x7_soap_monitor` has no `client_certificate_password` attribute; it's commented out
in the schema.

Rows with "Read writes API value = yes" and no suppression are Class A candidates. Confirm
with `terraform state show -show-sensitive`: if state holds `****`, `ENC(…)` or `""` while
config holds the real value, it's Class A.

- **Workaround now:** `lifecycle { ignore_changes = [password] }`. Tell the user the cost
  honestly: Terraform will then **never** push a rotated secret either. Rotation becomes
  "remove the `ignore_changes`, apply, put it back", or a change in the UI.
- **Provider fix:** stop `d.Set`-ing that attribute in the resource's Read (and in the
  update-state helper Read calls), so state keeps the last applied value. Keep
  `Sensitive: true`. Don't use an always-true `DiffSuppressFunc` instead, because that
  causes Class B.

Rows marked **`Sensitive` = no** also print the secret in plan output and CI logs. Flag this
whether or not it drifts.

### Class B: a rotated secret never reaches the API

Shape: the user changed the password in HCL, `terraform plan` says **"No changes"**, and
the monitor keeps failing auth with the old password.

Cause: `DiffSuppressFunc: func(...) bool { return true }` (the "always" rows above). The
diff is suppressed, so Update is never called. The new value only goes out if some other
attribute of the same resource changes in the same apply, because Update sends the whole
object from config. `site24x7_servicenow_integration` suppresses only when the old state
value isn't empty, `****` or `ENC(`-prefixed, so whether it drifts or sticks depends on what
the API returned last.

- **Workaround now, in order of preference:**
  1. Change the secret in the Site24x7 UI as well. State doesn't care.
  2. Make a harmless change to another attribute in the same apply (e.g. append to
     `display_name`, then revert). Both applies send the new secret.
  3. **Not** `terraform apply -replace`. It recreates the monitor with a **new ID and no
     history**, and anything referencing the old ID breaks.
- **Provider fix:** remove the `DiffSuppressFunc` and stop setting the attribute in Read.
  Both are needed: removing only the suppression turns Class B into Class A.

### Class C: a list re-orders on every plan

Shape: the same elements on both sides with the order swapped, or `- "123"` and `+ "123"`
on one list:

```
~ user_group_ids = [
    - "123000000001",
      "123000000002",
    + "123000000001",
  ]
```

Cause: the attribute is `TypeList`, and the API returns IDs in its own order (usually
ascending ID, sometimes creation order). Unordered ID collections should be `TypeSet`. The
provider is inconsistent here:

| Attribute | `TypeSet` in | `TypeList` (drift-prone) in |
|---|---|---|
| `monitor_groups` | website, cron, heartbeat, user | ssl, rest_api, rest_api_transaction, server, port, ping, dns_server, domain_expiry, isp, soap, ftp_transfer, web_page_speed, web_transaction_browser, schedule_maintenance, sla_setting |
| `user_group_ids` | cron, heartbeat, user | website, ssl, rest_api, rest_api_transaction, server, port, ping, dns_server, domain_expiry, isp, soap, ftp_transfer, web_page_speed, web_transaction_browser, amazon, azure, gcp, monitor_group, subgroup, oauth2_provider |
| `third_party_service_ids` | cron, heartbeat | most other monitors, monitor_group, subgroup |
| `monitors` | subgroup | all seven integrations, schedule_maintenance, sla_setting |
| `tags` | none | all seven integrations, schedule_maintenance |
| `tag_ids` | every monitor, monitor_group, subgroup | none |

- **Workaround now:** write the list **in the order the API returns it**. Copy the order
  from `terraform state show <address>` after a refresh, not from the UI. If the list comes
  from a `for` expression or data source, wrap it in `sort()`, which usually matches the
  API's ascending-ID order. Confirm with a second plan. `ignore_changes` also works but
  hides real membership changes, so avoid it here.
- **Provider fix:** change the attribute to `TypeSet`, and read it in the expand function
  with `d.Get(k).(*schema.Set).List()`. This changes the state format. Users see a one-time
  diff after upgrading, so call it out in the changelog.

### Class D: a value flips back to a default, `""` or `0`

Shape: `~ some_attr = "B" -> "O"`, `"" -> null`, `0 -> 5`, or `true -> false`, returning
after every apply.

Causes, most likely first:

1. **Schema `Default` ≠ API default.** The provider declares `Default: X` but the API stores
   or returns `Y` when the field is omitted or irrelevant (e.g. `auth_method` defaults to
   `"B"` on `site24x7_website_monitor`). The attribute doesn't apply in the user's setup,
   but the provider keeps sending the default.
2. **The API ignores the field in this configuration.** The value only applies with another
   setting (e.g. `request_content_type` only matters when `http_method = "P"`). The API
   drops it and Read writes back `""`.
3. **The user set a value the API normalises**: `"200, 201"` → `"200,201"`,
   `"https://x.com"` → `"https://x.com/"`, colour `#fff` → `#FFF`, a JSON request body
   re-serialised, a timezone name mapped to another alias. The same shape as causes 1 and
   2, but both sides are non-empty and look "the same".
4. **The API fills a value asynchronously after create.** It only shows after `terraform
   refresh`.

- **Workaround now:**
  - For 1 and 2, set the attribute explicitly to exactly the value state shows after
    refresh, or remove it from config if the provider has no `Default`.
  - For 3, write the canonical form the API returns.
  - For 4, `ignore_changes` on that attribute, or set it to the server value.
- **Provider fix:**
  - For 1 and 4, make it `Optional + Computed` and drop the `Default`.
  - For 2, don't send the field unless its precondition holds.
  - For 3, add a normalising `DiffSuppressFunc` that compares **canonicalised** values (trim,
    case-fold, JSON-equivalent). Never add an always-true one.

After an import, many Class D diffs are just `""` / `0` written by `-generate-config-out`.
Delete those lines. `site24x7-import-existing` Phase 4 covers that loop.

### Class E: a profile or user-group ID changes on its own

Shape: `~ location_profile_id = "…001" -> "…002"`, or `notification_profile_id` or
`threshold_profile_id` changing, especially when `*_profile_name` is set, or when nothing
profile-related is set at all.

Cause: the provider's **name resolution and defaulting** in
[site24x7/monitor_defaults.go](../../../site24x7/monitor_defaults.go). `*_name` is matched as
a substring and the matching profile is written over `*_id`. With neither set, the provider
picks a default from the list the API returns. Setting both `_name` and `_id` means the name
wins on every apply.

- **Workaround now:** reference profiles by ID from a resource or data source
  (`site24x7_location_profile.x.id`), and don't set the `_name` variant as well.
- The full explanation and the naming rules that avoid ambiguous matches are in the
  **`site24x7-monitor-authoring`** skill. Load it rather than re-deriving them here.

### Class F: Terraform wants to create the resource on every plan, or plan fails

Shape: `+ create` for a resource that exists, or `Error: … not found` during refresh.

Causes:

- **It was deleted in the UI.** Resources that handle a 404 in Read drop out of state, and
  the plan correctly proposes recreating them. That's not a bug.
- **Plan fails with a 404 error instead.** These resources don't handle 404 in Read, so one
  UI deletion breaks every plan: `site24x7_credential_profile`, `site24x7_oauth2_provider`,
  `site24x7_sla_setting`, `site24x7_attribute_alert_group`, `site24x7_milestone_marker`.
  Workaround: `terraform state rm <address>`, then re-apply to recreate or re-import.
  Provider fix: `if apierrors.IsNotFound(err) { d.SetId(""); return nil }` in Read.
- **Wrong account or data centre.** A different `zaaid`, `data_center` or credential set
  means every ID 404s. Everything in the workspace plans as create at once. Check the
  provider block before anything else.
- **Not importable.** `site24x7_amazon_monitor`, `site24x7_gcp_monitor`,
  `site24x7_azure_monitor` and `site24x7_milestone_marker` have no `Importer`, so an
  "adopted" one was never really in state.

### Class G: two things own the same setting

Shape: the plan is clean right after apply, then the diff comes back on its own. Or two
resources alternate: applying one creates a diff on the other.

Causes seen in this provider:

- **APM default profile.** Two `site24x7_apm_agent_config_profile` resources of the same
  `agent_type` both set `is_default = true`. Promoting one demotes the other remotely. Set
  it on exactly one.
- **Monitor ↔ integration association set from both ends.** An integration's `monitors`
  (with `selection_type = 2`) and a monitor's `third_party_service_ids` describe the same
  link. Manage the association from **one** side only.
- **Monitor ↔ group membership from both ends.** A monitor's `monitor_groups` versus
  membership managed on the group or subgroup side. Pick one owner.
- **Multiple workspaces or ClickOps.** Another Terraform workspace, or people editing in the
  UI. `terraform plan -refresh-only` shows what changed remotely. The fix is organisational,
  not technical.

### Class H: "update in-place" with no visible attribute change

Shape: `~ resource "…" { … }` with only `# (N unchanged attributes hidden)`, or a nested
block showing the same values.

Causes: a change inside a `Sensitive` attribute (Class A hiding itself), a nested
`TypeList` + `MaxItems: 1` block where the API returns an extra or empty element, or a
`TypeMap` (`request_headers`) whose values the API normalises. Use
`terraform show -json tfplan` and diff `before` against `after` for that resource to find
the real path, then reclassify.

## Step 3: Apply the fix and prove it

For a **config-side** fix:

1. Edit the HCL.
2. `terraform plan`. The diff for that attribute is gone.
3. `terraform apply` (only if other intended changes remain), then `terraform plan` twice.
   Both must say **"No changes."**

For a **provider-side** fix in this repo:

1. Make the change using the patterns in `site24x7-add-resource` Step 8.
2. Add a unit test that fails before the fix. For Class A: build ResourceData with the
   secret, call Read with a fake returning `****`, and assert the state value is unchanged.
3. Run `go build ./...`, `go vet ./...` and `go test ./site24x7/...`.
4. Build and install locally. Against a real account: apply, then plan twice, which must be
   clean. Then change the secret or list order and confirm the plan shows **exactly** that
   change.
5. Update the offender tables in this skill, and the secrets table in
   `site24x7-alerting-and-profiles`, so they stay accurate.

Only use `ignore_changes` as a deliberate, documented trade-off, never as the first move. If
the user uses it, add a comment in their HCL saying why and which provider issue it works
around.

## Re-auditing the offender tables

The tables above come from the source. Re-derive them after provider changes. Secrets
(PowerShell):

```powershell
$secrets = 'auth_pass','client_certificate_password','password','private_key','token','client_secret','access_token'
Get-ChildItem -Recurse site24x7 -Filter *.go | ? { $_.Name -notlike '*_test.go' -and $_.Name -notlike '*data_source*' } | % {
  $txt = Get-Content $_.FullName -Raw
  foreach ($s in $secrets) { if ($txt -match "`"$s`":\s*\{") {
    $block = [regex]::Match($txt, "`"$s`":\s*\{[\s\S]*?\n\t\},").Value
    '{0,-30} {1,-28} Sensitive={2,-5} DSF={3,-5} ReadSets={4}' -f $_.Name, $s,
      ($block -match 'Sensitive:\s*true'), ($block -match 'DiffSuppressFunc'), ($txt -match "d\.Set\(`"$s`"")
  } }
}
```

Lists vs sets:

```sh
grep -rnE '"(monitor_groups|user_group_ids|third_party_service_ids|monitors|tags|tag_ids)": \{' site24x7 -A1 | grep -E 'TypeList|TypeSet'
```

Read-without-404 handling: resources whose file has `Create:` but no `IsNotFound`.

## Checklist before closing

- [ ] Diff reproduced, with a second plan after apply
- [ ] Classified (A–H), with the evidence: state value vs config value vs API response
- [ ] User given a workaround they can apply now, including its trade-off
- [ ] Provider-side cause named with file and attribute, and fixed with a test if in scope
- [ ] Two consecutive `terraform plan` runs show "No changes."
- [ ] No secret, debug log or `-show-sensitive` output left in the repo or pasted publicly
