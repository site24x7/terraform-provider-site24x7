# Site24x7 Terraform provider — Claude Code skills

Three skills ship with this repository. Each one encodes provider behaviour that is easy to get
wrong from the documentation alone: name resolution that is substring-based, defaults that pick
element `[0]`, resources that cannot be imported, deletes that do nothing, and credentials that
never come back from the API.

| Skill | Covers |
|---|---|
| [`site24x7-monitor-authoring`](site24x7-monitor-authoring/SKILL.md) | Writing and reviewing monitors, and APM Insight |
| [`site24x7-alerting-and-profiles`](site24x7-alerting-and-profiles/SKILL.md) | The 25 resources that are not monitors — profiles, groups, users, integrations, schedules, SLAs, credentials, MSP customers |
| [`site24x7-import-existing`](site24x7-import-existing/SKILL.md) | Bringing an existing (ClickOps) account under Terraform management |

## Which one do I want?

| I want to… | Skill |
|---|---|
| Add a monitor for a URL, host, API, port, DNS record, certificate | `site24x7-monitor-authoring` |
| Work out which `site24x7_*` monitor resource fits | `site24x7-monitor-authoring` |
| Fix "Unable to find … matching the string" | `site24x7-monitor-authoring` |
| Understand why a monitor landed on the wrong profile | `site24x7-monitor-authoring` |
| Bulk-create monitors with `for_each` | `site24x7-monitor-authoring` |
| Tune an APM Insight agent, or adopt an APM application | `site24x7-monitor-authoring` (Step 5) |
| Create a notification / location / threshold profile | `site24x7-alerting-and-profiles` |
| Put a team on call — users, user groups, attribute alert groups | `site24x7-alerting-and-profiles` |
| Send alerts to Slack, PagerDuty, ServiceNow, OpsGenie, Connectwise, Telegram, a webhook | `site24x7-alerting-and-profiles` |
| Schedule a maintenance window or a report, define an SLA | `site24x7-alerting-and-profiles` |
| Store a credential once and reuse it across monitors | `site24x7-alerting-and-profiles` |
| Add a customer under an MSP or BU | `site24x7-alerting-and-profiles` |
| Know what `terraform destroy` really does to a resource | `site24x7-alerting-and-profiles` |
| Import monitors that already exist | `site24x7-import-existing` |
| Adopt a whole hand-built account into Terraform | `site24x7-import-existing` |
| Drive `terraform plan` to "No changes" after an import | `site24x7-import-existing` |
| Decide between `import` blocks, `terraform import`, and `site24x7_importer.py` | `site24x7-import-existing` |

## How to invoke them

**Automatically.** Just describe the task. Each skill's `description` field decides whether it
loads, and all three are written to trigger on the way people actually phrase things ("how do I
get my existing monitors into Terraform", "set up Slack alerts") without anyone saying "skill".

**Explicitly**, by name, with a free-text argument:

```
/site24x7-import-existing all SERVER monitors
/site24x7-monitor-authoring the 12 prod API endpoints, alert the payments on-call
/site24x7-alerting-and-profiles PagerDuty for the payments team
```

The argument is optional but worth giving — it is what lets the skill skip straight to its fast
path instead of asking. All three are `user-invocable`.

## What each will do

### `site24x7-monitor-authoring`

1. Writes `versions.tf` and a provider block if the directory has none — it does this without
   asking.
2. Picks the resource from what is being checked, not from what you call it.
3. Wires location / notification / threshold profiles, user groups and tags — preferring names
   over 18-digit IDs, while warning where name matching is ambiguous.
4. Writes the HCL, using `for_each` over a map for many monitors.
5. Runs `terraform fmt`, `validate`, `plan`.

It will not run `apply` for you on a new naming convention without telling you to check the
first monitor's resolved profile, because name resolution is server-side and only proves itself
on apply.

### `site24x7-alerting-and-profiles`

1. Establishes the build order — several of these resources have hard references to each other,
   and `site24x7_user_group` needs both a user list and an attribute alert group.
2. Gives the required attributes for each resource, so the first apply is not a guessing game.
3. Flags the traps: integration targeting that silently alerts nothing, credentials that print
   in cleartext, `site24x7_customer` destroys that delete nothing.
4. Reviews existing HCL against a checklist.

### `site24x7-import-existing`

1. Establishes a working provider block **in the customer's own project**, and proves auth and
   data centre with one cheap read before touching imports.
2. Inventories the account — monitors by type, their dependencies, and APM separately.
3. Picks the import mechanism (`import` blocks by default).
4. Generates configuration and reconciles the diff.
5. Repeats until `terraform plan` reports "No changes", then hands over.

It has a fast path for a single named monitor, so asking it to adopt one SSL monitor will not
trigger a full account sweep. It will not run `apply` during an adoption unless the plan has
been read line by line — an apply against a half-written config overwrites live monitors.

## Prerequisites for all three

- **Credentials in the environment**, never in a `.tf` file:
  `SITE24X7_OAUTH2_CLIENT_ID`, `SITE24X7_OAUTH2_CLIENT_SECRET`, `SITE24X7_OAUTH2_REFRESH_TOKEN`.
- **`data_center` matching where those credentials were issued** — `US`, `EU`, `IN`, `AU`, `CN`,
  `JP` or `CA`. A mismatch fails as a confusing auth error, not a clear one.
- **`zaaid`** when the target is a customer under an MSP or BU. Without it you query the parent
  account and see none of the customer's monitors.
- **Terraform ≥ 1.5** to use `import` blocks and `terraform plan -generate-config-out`. The
  import skill has a path for older versions, but it is much more work.
- The provider from the registry: `source = "site24x7/site24x7"`, `version = "~> 1.0"`.

## Two things to know before you start

**These are project skills.** They live in `.claude/skills/`, so they load when Claude Code is
working inside this repository. The import skill correctly tells you *not* to run an adoption
inside a clone of the provider repo — which means the skill will not be loaded in the directory
where you are supposed to do the work. To use them from a customer's Terraform project, copy the
skill directories to `~/.claude/skills/` (user scope, available everywhere):

```powershell
New-Item -ItemType Directory -Force $env:USERPROFILE\.claude\skills | Out-Null
Copy-Item -Recurse -Force .claude\skills\site24x7-* $env:USERPROFILE\.claude\skills\
```

```sh
mkdir -p ~/.claude/skills
cp -r .claude/skills/site24x7-* ~/.claude/skills/
```

Copied that way they are a snapshot: re-copy after pulling changes to this repo.

**The shell commands assume a POSIX shell.** If you are on PowerShell, substitute:

| In the skills | PowerShell |
|---|---|
| `terraform state list \| wc -l` | `(terraform state list \| Measure-Object -Line).Lines` |
| `env \| grep SITE24X7` | `Get-ChildItem Env:SITE24X7*` |
| `TF_LOG=DEBUG terraform apply` | `$env:TF_LOG='DEBUG'; terraform apply` |
| `grep -rE 'oauth2_\|auth_pass\|secret' *.tf` | `Select-String -Path *.tf -Pattern 'oauth2_\|auth_pass\|secret'` |

## Maintaining these skills

Each `SKILL.md` has YAML frontmatter with `name` (**must equal its directory name**),
`description` (what drives automatic loading — keep the trigger phrases in it),
`user-invocable`, `argument-hint` and `allowed-tools` (all three are granted `Read`, `Write`,
`Edit`, `Glob`, `Grep`, `Bash`, so they can write `.tf` files and run `terraform`).

The skills state facts about the provider that change when the provider changes: how many
resources are importable, which resources expose `location_profile_name`, the monitor type
codes, which attributes are `Sensitive`. After adding or changing a resource, re-check them.

Every `site24x7_*` name mentioned across the skills can be validated against the provider's
registrations in one command:

```powershell
$lines = Get-Content provider/provider.go | Where-Object { $_ -notmatch '^\s*//' }
$known = @{}
($lines | Select-String -Pattern '"(site24x7_[a-z0-9_]+)":').Matches |
  ForEach-Object { $known[$_.Groups[1].Value] = $true }
Get-ChildItem -Recurse .claude/skills -Filter SKILL.md | ForEach-Object {
  ([regex]::Matches((Get-Content $_.FullName -Raw), 'site24x7_[a-z0-9_]+')).Value |
    Sort-Object -Unique | Where-Object { -not $known.ContainsKey($_) } |
    ForEach-Object { "$($_): not registered in provider.go" }
}
```

Two known false positives: `site24x7_user_constants` (an anchor in a Site24x7 docs URL) and
`site24x7_importer` (the Python script's filename).

When adding a resource to the provider, the places to update are the resource/data-source map
and the destroy-semantics list in `site24x7-alerting-and-profiles`, the importable count and
dependency order in `site24x7-import-existing`, and — if it is a monitor — the resource table in
`site24x7-monitor-authoring`.
