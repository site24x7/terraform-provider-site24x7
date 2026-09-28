---
layout: "site24x7"
page_title: "Site24x7: site24x7_schedule_maintenances"
sidebar_current: "docs-site24x7-data-source-schedule-maintenances"
description: |-
  Get information about scheduled maintenance windows in Site24x7, filtered by monitor, monitor group, tag, type, status or name.
---

# Data Source: site24x7\_schedule\_maintenances

Use this data source to list scheduled maintenance windows, optionally filtered to the ones that apply to a particular monitor, monitor group or tag.

Every filter is optional, and the filters that are set are combined: a maintenance must pass all of them. With no filters, every maintenance in the account is returned. An empty result is not an error, because a monitor with no maintenance window is a valid answer.

~> **Filtering happens in the provider.** The Site24x7 List Maintenance API (`GET /maintenance`) takes no query parameters, so this data source fetches every maintenance and filters the list itself. In an account with many maintenances, each read transfers all of them.

~> **`monitor_id` matches direct membership only.** A maintenance matches `monitor_id` only when it lists that monitor itself (`selection_type = 2`). A maintenance that covers the monitor through one of its monitor groups (`selection_type = 1`) or tags (`selection_type = 3`) is **not** returned. To find those, query the monitor's groups with `monitor_group_id` and its tags with `tag_id`.

## Example Usage

```hcl
// Every maintenance that lists this monitor directly.
data "site24x7_schedule_maintenances" "for_monitor" {
  monitor_id = "113770000039133011"
}

// Every maintenance for a monitor group, looked up by name.
data "site24x7_monitor_group" "payments" {
  name_regex = "^Payments$"
}

data "site24x7_schedule_maintenances" "for_group" {
  monitor_group_id = data.site24x7_monitor_group.payments.id
}

// Every maintenance targeting a tag.
data "site24x7_schedule_maintenances" "for_tag" {
  tag_id = "113770000039135099"
}

// Filters combine: one-time (type 3) maintenances whose name starts with "Release".
data "site24x7_schedule_maintenances" "one_time_releases" {
  name_regex       = "^Release"
  maintenance_type = 3
}

output "monitor_maintenance_ids" {
  value = data.site24x7_schedule_maintenances.for_monitor.ids
}

// The `maintenances` attribute carries the full configuration, so you can
// read schedules without a second lookup per maintenance.
output "group_maintenance_windows" {
  value = [
    for m in data.site24x7_schedule_maintenances.for_group.maintenances :
    "${m.display_name}: ${m.start_date} ${m.start_time} - ${m.end_date} ${m.end_time} (${m.time_zone})"
  ]
}
```

## Attributes Reference

### Optional

* `name_regex` (String) Regular expression matched against the maintenance display name.
* `monitor_id` (String) Return only maintenances that list this monitor directly (`selection_type = 2`). Maintenances that reach the monitor through a monitor group or tag are not matched.
* `monitor_group_id` (String) Return only maintenances that list this monitor group (`selection_type = 1`).
* `tag_id` (String) Return only maintenances that list this tag (`selection_type = 3`).
* `maintenance_type` (Number) Return only maintenances of this type, for example `3` for Once. Omit it, or set `0`, to match every type.
* `maintenance_status` (String) Return only maintenances with this status code, exactly as the API returns it, for example `V`.

### Read-Only

* `id` (String) The ID of this data source, derived from the filters.
* `ids` (List of String) IDs of the matching maintenances, sorted.
* `ids_and_names` (List of String) Matching maintenances as `<id>__<display_name>`, sorted by ID.
* `maintenances` (List of Object) Matching maintenances, sorted by maintenance ID. Each entry has:
  * `maintenance_id` (String) Unique ID of the maintenance.
  * `display_name` (String) Display name.
  * `description` (String) Description.
  * `maintenance_type` (Number) Once / Daily / Weekly / Monthly / Yearly configuration code.
  * `maintenance_status` (String) Status of the maintenance.
  * `selection_type` (Number) `1` = Monitor Groups, `2` = Monitors, `3` = Tags.
  * `monitors` (List of String) Monitor IDs, when `selection_type = 2`.
  * `monitor_groups` (List of String) Monitor group IDs, when `selection_type = 1`.
  * `tags` (List of String) Tag IDs, when `selection_type = 3`.
  * `subgroup_monitors` (Boolean) Whether subgroups of the selected monitor groups are included. `true` when the API omits it, which is the API's documented default.
  * `perform_monitoring` (Boolean) Whether uptime monitoring continues during the maintenance.
  * `time_zone` (String) Time zone of the maintenance.
  * `start_date`, `end_date` (String) Start and end dates, `yyyy-mm-dd`. Set for one-time maintenances.
  * `start_time`, `end_time` (String) Start and end times, `hh:mm`.
  * `duration` (Number) Duration in minutes. `0` when the API returns none.
  * `start_day`, `end_day` (Number) Weekly start and end days.
  * `week_days` (List of Number) Days of the week on which the maintenance recurs.
  * `execute_every` (Number) Interval, in weeks, at which a weekly maintenance recurs.
  * `start_week` (Number) Week of the month on which a monthly (by day) maintenance recurs.
  * `start_after` (Number) Days after which a monthly maintenance begins.
  * `monthly_start_date` (Number) Date on which a monthly (by date) maintenance recurs.
  * `maintenance_start_on` (String) Date from which a recurring maintenance starts, `yyyy-mm-dd`.
  * `maintenance_end_type` (Number) How a recurring maintenance ends.
  * `maintenance_end_after_times` (Number) Number of executions after which it ends, when `maintenance_end_type = 1`.
  * `maintenance_end_on` (String) Date on which it ends, when `maintenance_end_type = 2`.
  * `maintenance_start_time` (String) When the maintenance was created or last updated, ISO 8601.
  * `zuid` (String) ID of the user who created the maintenance.

Refer [API documentation](https://www.site24x7.com/help/api/#schedule-maintenances) for more information about attributes.
