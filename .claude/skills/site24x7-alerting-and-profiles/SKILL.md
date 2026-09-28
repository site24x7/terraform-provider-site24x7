---
name: site24x7-alerting-and-profiles
description: >
  Author and review the Site24x7 resources that are not monitors — location, notification and
  threshold profiles, monitor groups and subgroups, user groups, users, attribute alert groups,
  tags, business hours, the seven alerting integrations (Slack, PagerDuty, ServiceNow, OpsGenie,
  Connectwise, Telegram, webhook), credential profiles, OAuth2 providers, scheduled maintenance
  and reports, SLA settings, milestone markers, IT automation actions and MSP customers. Encodes
  the build order these resources impose on each other, the required attributes of each, which
  credentials never come back from the API, and what `terraform destroy` actually does to each.
  Use whenever someone wants to create, wire up or review anything in a Site24x7 account that
  isn't a monitor — "set up Slack alerts", "add a notification profile", "schedule a maintenance
  window", "put the payments team on call", "SLA report", "add a customer under our MSP" — even
  if they don't say "skill".
user-invocable: true
argument-hint: "<what to set up, e.g. 'PagerDuty for the payments team' or 'business hours plus escalation'>"
allowed-tools:
  - Read
  - Write
  - Edit
  - Glob
  - Grep
  - Bash
---

# Site24x7 alerting, profiles and account resources

This skill covers the 25 registered resources that are **not** monitors and not APM.

- Monitors, and the profile/group associations *on* a monitor → `site24x7-monitor-authoring`.
- Adopting things that already exist in the account → `site24x7-import-existing`.
- APM Insight → `site24x7-monitor-authoring`, Step 5.

Everything here is written the same way as a monitor: `terraform init`, provider block with
`data_center` matching where the OAuth credentials were issued, credentials from
`SITE24X7_OAUTH2_CLIENT_ID` / `_CLIENT_SECRET` / `_REFRESH_TOKEN` and never from a `.tf` file.

## Build order

These are hard requirements — a `Required` attribute that holds another resource's ID. Get the
order wrong and the apply fails on a missing reference, or worse, succeeds against the wrong ID.

1. **`site24x7_user`** and **`site24x7_attribute_alert_group`** — depend on nothing.
2. **`site24x7_user_group`** — requires **both** `users` (a set of user IDs) **and**
   `attribute_group_id`. This is the trap in the whole graph: you cannot create a user alert
   group without first having an attribute alert group, and the error does not say so.
3. **`site24x7_businesshour`** — depends on nothing; feeds notification profiles and SLA
   settings as an *optional* input.
4. **`site24x7_notification_profile`** — only `profile_name` is required. Its escalation blocks
   reference `user_group_id`, and `business_hours_id` defaults to `"-1"`, meaning *All Hours*.
5. **`site24x7_monitor_group`**, then **`site24x7_subgroup`**, which requires
   `parent_group_id` and `top_group_id`.
6. **`site24x7_location_profile`**, **`site24x7_threshold_profile`**, **`site24x7_tag`** —
   independent of each other.
7. Monitors (other skill).
8. Things that point *at* monitors, so they come last: `site24x7_schedule_maintenance`,
   `site24x7_schedule_report`, `site24x7_sla_setting`, `site24x7_milestone_marker`, and any
   integration with `selection_type = 2`.

## Required attributes

Top-level `Required` attributes only. Everything else is optional, and the provider will
default it — read the resource's `docs/resources/` page before assuming a default is harmless.

| Resource | Required |
|---|---|
| `site24x7_attribute_alert_group` | `display_name`, `attribute_list` |
| `site24x7_user` | `display_name`, `email_address`, `user_role`, `notification_medium`, `down_notification_medium`, `critical_notification_medium`, `trouble_notification_medium`, `up_notification_medium` |
| `site24x7_user_group` | `display_name`, `users`, `attribute_group_id` |
| `site24x7_businesshour` | `display_name`, `time_config` |
| `site24x7_notification_profile` | `profile_name` |
| `site24x7_location_profile` | `profile_name`, `primary_location`, `secondary_locations` |
| `site24x7_threshold_profile` | `profile_name`, `type` |
| `site24x7_tag` | `tag_name`, `tag_color` (`tag_value` is optional) |
| `site24x7_monitor_group` | `display_name` |
| `site24x7_subgroup` | `display_name`, `parent_group_id`, `top_group_id` |
| `site24x7_url_action` | `name`, `url` |
| `site24x7_credential_profile` | `credential_type`, `credential_name`, `username`, `password` |
| `site24x7_oauth2_provider` | `provider_name`, `oauth2_flow`, `client_id`, `client_secret`, `access_token_uri`, `send_token_as` |
| `site24x7_schedule_maintenance` | `display_name`, `maintenance_type`, `start_time`, `end_time` |
| `site24x7_schedule_report` | `display_name`, `selection_type`, `report_format`, `scheduled_time`, `user_groups` |
| `site24x7_sla_setting` | `display_name`, `type`, `selection_type`, `sla_targets` |
| `site24x7_milestone_marker` | `marker_time`, `label` |
| `site24x7_customer` | `country_code`, `timezone`, `customer_company`, `display_name`, `email_address`, `portal_name` |

`site24x7_threshold_profile` looks small in that table and is not: nearly everything is a
nested block (`*_trouble_threshold`, `*_critical_threshold`, each with `severity`,
`comparison_operator`, `value`, `strategy`, `polls_check`). Pick the blocks that match the
`type` you set — a server-type profile ignores website blocks.

`site24x7_milestone_marker`'s `monitor_id` is **optional**, `ForceNew`, and defaults to `"-1"`,
which creates a *global* marker applying to every monitor. Leaving it out is not "no monitor",
it is "all monitors".

## The seven integrations

| Resource | Required beyond `name` |
|---|---|
| `site24x7_webhook_integration` | `url` |
| `site24x7_opsgenie_integration` | `url` |
| `site24x7_slack_integration` | `url`, `sender_name`, `title` |
| `site24x7_telegram_integration` | `channel_url`, `token`, `title` |
| `site24x7_pagerduty_integration` | `service_key`, `sender_name`, `title` |
| `site24x7_servicenow_integration` | `instance_url`, `title`, `sender_name`, `user_name`, `password` |
| `site24x7_connectwise_integration` | `url`, `company`, `public_key`, `private_key`, `company_id`, `close_status` |

All seven share the same targeting shape: `selection_type`, `monitors`, `tags`.

- `selection_type = 0` (**the default**) — All Monitors.
- `selection_type = 2` — only the monitors listed in `monitors`.
- `selection_type = 3` — only the monitors carrying the tags listed in `tags`.

**Nothing cross-checks these.** The provider passes `selection_type` straight through and never
validates that `monitors` is populated when you set `2`, or `tags` when you set `3`. Set `2` and
forget `monitors`, and the integration is created targeting nothing — no plan error, no apply
error, and no alerts. Equally, populating `monitors` while leaving `selection_type` at its
default `0` silently routes *every* monitor's alerts to that integration. Always write both, or
neither.

## Credentials that never come back

Seven credential-shaped attributes appear on fourteen resources, 19 attribute/resource pairs
in all, and the API does not return them in usable form. Generated or refreshed config shows
them empty, masked or encrypted, and applying that value overwrites the live credential.
Source them from variables backed by a vault before the first apply.

**15 of the 19 pairs are marked `Sensitive`.** The remaining four are all `auth_pass` on HTTP
monitors, and they are **printed in cleartext** in plan output, `terraform show` and CI logs:

| Attribute | On | `Sensitive` |
|---|---|---|
| `auth_pass` | `site24x7_website_monitor`, `site24x7_rest_api_monitor`, `site24x7_web_page_speed_monitor`, `site24x7_rest_api_transaction_monitor` (inside each `steps` → `step_details` block) | **no** |
| `auth_pass` | `site24x7_oauth2_provider` | yes |
| `client_certificate_password` | `site24x7_website_monitor`, `site24x7_rest_api_monitor`, `site24x7_rest_api_transaction_monitor` (inside `steps`) | yes |
| `password` | `site24x7_credential_profile`, `site24x7_ftp_transfer_monitor`, `site24x7_servicenow_integration`, `site24x7_webhook_integration` | yes |
| `password` (in the `auth_details` map) | `site24x7_web_transaction_browser_monitor` | yes, via the map. The SDK ignores `Sensitive` on a map's nested schema, so the whole `auth_details` map is marked instead |
| `client_secret` | `site24x7_oauth2_provider`, `site24x7_azure_monitor` | yes |
| `private_key` | `site24x7_gcp_monitor`, `site24x7_connectwise_integration` | yes |
| `token` | `site24x7_telegram_integration` | yes |
| `access_token` | `site24x7_oauth2_provider` (computed, server-issued) | yes |

The `site24x7_credential_profile` **data source** also marks its `password` output `Sensitive`.

`Sensitive` only hides the value from CLI output. **It is still stored in plaintext in the
state file**, so the state backend needs encryption and access control whatever this table
says.

`site24x7_soap_monitor` has no `client_certificate_password` attribute. The schema entry is
commented out, so a SOAP monitor that needs a client certificate password can't be fully
managed from Terraform.

`site24x7_credential_profile` is the one to reach for rather than repeating a password across
monitors: five monitor resources accept `credential_profile_id` (`website`, `rest_api`,
`soap`, `web_page_speed`, `ftp_transfer`). Its own `password` is required. It's `Sensitive`,
so it no longer prints, but it still needs vault-sourced input.

Whether each secret also causes a perpetual diff, or never applies after rotation, depends on
how the resource reads it back. The `site24x7-drift-troubleshooting` skill has that table
(Classes A and B).

## What `terraform destroy` actually does

Every resource here passes the delete straight through to the API with no confirmation and no
guard, and most treat an already-deleted resource as success. Two exceptions, both verified in
the provider source:

- **`site24x7_apm_application`** is guarded by `delete_on_destroy`, default `false`: destroy
  releases it from state and leaves the application running. At `true` it permanently deletes
  the application *and its monitoring history*.
- **`site24x7_customer` does not delete anything.** `customerDelete` in
  `site24x7/msp/customer.go` is an empty function returning `nil`, even though the endpoint
  client implements `Delete`. A destroy drops the customer from Terraform state while the
  customer account continues to exist — and continues to bill. Removing a customer is a manual
  action in the MSP portal. Tell the user this explicitly; the plan output will say "destroy"
  and mean nothing of the sort.

What the provider does *not* protect you from, because the behaviour is server-side: deleting a
monitor group that still contains monitors, or a location / notification / threshold profile
still referenced by one. Check the cascade in a sandbox account before running it in
production — the provider will forward the call either way.

## Complete resource and data source map

46 registered resources, 24 registered data sources. `import` means the resource carries a
`ResourceImporter`, so `import` blocks and config generation work on it.

| Resource | Import | Matching data source |
|---|---|---|
| 16 standard monitors (`website`, `rest_api`, `rest_api_transaction`, `ssl`, `web_page_speed`, `web_transaction_browser`, `server`, `cron`, `heartbeat`, `dns_server`, `domain_expiry`, `isp`, `ftp_transfer`, `port`, `ping`, `soap`) | yes | `site24x7_monitor`, `site24x7_monitors` |
| `site24x7_amazon_monitor`, `_gcp_monitor`, `_azure_monitor` | **no** | `site24x7_monitor`, `site24x7_monitors` |
| `site24x7_monitor_group` | yes | `site24x7_monitor_group` |
| `site24x7_subgroup` | yes | **none registered** — commented out at `provider/provider.go:137` |
| `site24x7_location_profile` | yes | `site24x7_location_profile` |
| `site24x7_notification_profile` | yes | `site24x7_notification_profile` |
| `site24x7_threshold_profile` | yes | `site24x7_threshold_profile` |
| `site24x7_user_group` | yes | `site24x7_user_group` |
| `site24x7_user` | yes | `site24x7_user` |
| `site24x7_tag` | yes | `site24x7_tag` |
| `site24x7_attribute_alert_group` | yes | `site24x7_attribute_alert_group` |
| `site24x7_url_action` | yes | `site24x7_it_automation` — **note the name change** |
| `site24x7_credential_profile` | yes | `site24x7_credential_profile` |
| `site24x7_oauth2_provider` | yes | `site24x7_oauth2_provider` |
| `site24x7_sla_setting` | yes | `site24x7_sla_setting` |
| `site24x7_customer` | yes | `site24x7_customer`, `site24x7_msp` |
| `site24x7_businesshour` | yes | none |
| `site24x7_schedule_maintenance` | yes | `site24x7_schedule_maintenances` — list, filterable by `monitor_id`, `monitor_group_id`, `tag_id`, `maintenance_type`, `maintenance_status`, `name_regex` |
| `site24x7_schedule_report` | yes | none |
| all 7 `site24x7_*_integration` | yes | none |
| `site24x7_milestone_marker` | **no** | none |
| `site24x7_apm_application` | yes | `site24x7_apm_application`, `site24x7_apm_applications` |
| `site24x7_apm_agent_config_profile` | yes | `site24x7_apm_agent_config_profile`, `site24x7_apm_agent_config_profiles` |

Three data sources have no resource behind them and exist purely to read account state:
`site24x7_device_key` (the account's device key — also the cheapest auth smoke test),
`site24x7_aws_external_id`, and `site24x7_apm_instance` / `site24x7_apm_instances`.

Resources with **no** data source cannot be discovered by name from Terraform. To adopt or
reference one you need its ID from the Site24x7 UI or a direct API call — that applies to
business hours, scheduled reports, every integration, milestone markers and subgroups.

`site24x7_schedule_maintenances` filters in the provider, because `GET /maintenance` takes no
query parameters. Its `monitor_id` filter matches only maintenances that list the monitor
directly (`selection_type = 2`). A maintenance that covers the monitor through a monitor group
or tag is not returned, so query those with `monitor_group_id` and `tag_id` as well.

## Worked example — a complete alerting path

```hcl
resource "site24x7_attribute_alert_group" "payments" {
  display_name   = "payments-attributes"
  attribute_list = [1, 2] # attribute IDs; see docs/resources/attribute_alert_group.md
}

resource "site24x7_user" "oncall" {
  display_name  = "Payments on-call"
  email_address = "payments-oncall@example.com"

  # Numbers, not strings: user_role is TypeInt and every notification medium is
  # a set of ints. Codes are at
  # https://www.site24x7.com/help/api/#site24x7_user_constants
  user_role           = 1
  notification_medium = [1] # 1 = email

  down_notification_medium     = [1]
  critical_notification_medium = [1]
  trouble_notification_medium  = [1]
  up_notification_medium       = [1]
}

resource "site24x7_user_group" "payments_oncall" {
  display_name = "payments-oncall"
  users        = [site24x7_user.oncall.id]

  # Required. A user group cannot exist without one.
  attribute_group_id = site24x7_attribute_alert_group.payments.id
}

resource "site24x7_businesshour" "office" {
  display_name = "Office hours"

  time_config {
    day        = 1
    start_time = "09:00"
    end_time   = "18:00"
  }
}

resource "site24x7_notification_profile" "payments" {
  profile_name = "Payments — Sev1"
  # business_hours_id defaults to "-1" (All Hours); set it to restrict alerting
  # to the window above.
  business_hours_id = site24x7_businesshour.office.id
}

resource "site24x7_slack_integration" "payments" {
  name        = "payments-slack"
  url         = var.slack_webhook_url
  sender_name = "Site24x7"
  title       = "$MONITORNAME is $STATUS"

  # Target explicitly: selection_type and monitors must agree.
  selection_type = 3
  tags           = [site24x7_tag.payments.id]
}

resource "site24x7_tag" "payments" {
  tag_name  = "team-payments"
  tag_color = "#B7DA9E"
}
```

Verify attribute IDs, user roles and notification medium codes against the resource's page under
`docs/resources/` — they are numeric enums, and a wrong-but-valid number applies cleanly and
routes alerts somewhere unintended.

## Review checklist

- [ ] `site24x7_user_group` without `attribute_group_id` wired from a real
      `site24x7_attribute_alert_group` — it is required, and a hardcoded ID from another tenant
      applies cleanly and misroutes.
- [ ] An integration with `selection_type = 2` and no `monitors`, or `= 3` and no `tags` —
      created, silent, alerts nothing.
- [ ] An integration with `monitors` or `tags` populated but `selection_type` left at `0` —
      every monitor in the account routes to it.
- [ ] Any credential written as a literal. `Sensitive` ones still land in state in plaintext,
      and `auth_pass` on the website, REST API, web page speed and REST API transaction
      monitors is not `Sensitive`, so it also appears in plan output and CI logs in cleartext.
- [ ] A destroy plan containing `site24x7_customer` — it deletes nothing; say so.
- [ ] `delete_on_destroy = true` on `site24x7_apm_application`.
- [ ] `site24x7_milestone_marker` with no `monitor_id`, when a specific monitor was meant — it
      defaults to a global marker across all monitors.
- [ ] `site24x7_subgroup` referenced by name — there is no data source for it; the ID has to
      come from the UI or the API.
- [ ] Repeated passwords across monitors where one `site24x7_credential_profile` would do.
