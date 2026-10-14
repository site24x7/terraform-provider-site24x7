---
layout: "site24x7"
page_title: "Site24x7: site24x7_schedule_maintenance"
sidebar_current: "docs-site24x7-schedule-maintenance"
description: |-
Create and manage a Schedule Maintenance in Site24x7.
---

# Resource: site24x7\_schedule\_maintenance

Use this resource to create, update, and delete Schedule Maintenance in Site24x7. It supports one-time, daily, weekly and monthly maintenance windows.

## Example Usage

```hcl

// Site24x7 Schedule Maintenance API doc - https://www.site24x7.com/help/api/#schedule-maintenances

// Once: a single window between two dates.
resource "site24x7_schedule_maintenance" "once" {
  display_name     = "Switch upgrade"
  description      = "Core switch firmware upgrade"
  maintenance_type = 3
  start_date       = "2026-10-15"
  end_date         = "2026-10-15"
  start_time       = "19:30"
  end_time         = "21:00"
  time_zone        = "GMT"
  selection_type   = 2
  monitors         = ["123456000007534005"]
}

// Daily: every day from 01:00 to 02:00, ending after 30 runs.
resource "site24x7_schedule_maintenance" "daily" {
  display_name                = "Nightly backup"
  maintenance_type            = 1
  start_time                  = "01:00"
  end_time                    = "02:00"
  maintenance_start_on        = "2026-10-01"
  maintenance_end_type        = 1
  maintenance_end_after_times = 30
  monitors                    = ["123456000007534005"]
}

// Weekly (By Time): from Saturday 22:00 to Sunday 04:00, every other week.
resource "site24x7_schedule_maintenance" "weekly_by_time" {
  display_name     = "Weekend patching"
  maintenance_type = 2
  start_day        = 6 // Saturday
  start_time       = "22:00"
  end_day          = 0 // Sunday
  end_time         = "04:00"
  execute_every    = 2
  monitors         = ["123456000007534005"]
}

// Weekly (By Day): Wednesdays and Fridays at 14:30 for 30 minutes, bi-weekly.
resource "site24x7_schedule_maintenance" "weekly_by_day" {
  display_name     = "Bi-weekly deploys"
  maintenance_type = 8
  week_days        = [3, 5]
  start_time       = "14:30"
  duration         = 30
  execute_every    = 2
  monitors         = ["123456000007534005"]
}

// Monthly (By Date): the 15th of every month at 03:00 for 2 hours, until a date.
resource "site24x7_schedule_maintenance" "monthly_by_date" {
  display_name         = "Monthly DB maintenance"
  maintenance_type     = 5
  monthly_start_date   = 15
  start_time           = "03:00"
  duration             = 120
  maintenance_end_type = 2
  maintenance_end_on   = "2027-12-31"
  monitors             = ["123456000007534005"]
}

// Monthly (By Day): Patch Tuesday, the second Tuesday of the month, for a monitor group.
resource "site24x7_schedule_maintenance" "patch_tuesday" {
  display_name      = "Patch Tuesday"
  maintenance_type  = 6
  start_week        = 2
  start_day         = 2 // Tuesday
  start_time        = "23:00"
  duration          = 180
  selection_type    = 1
  monitor_groups    = ["123456000007534010"]
  subgroup_monitors = false
}

```

## Maintenance types

`maintenance_type` sets which schedule attributes apply. Setting an attribute that the chosen type does not use is rejected at plan time: the API would ignore it and every later plan would show a diff.

| `maintenance_type` | Required | Optional |
|---|---|---|
| `3` Once | `start_date`, `end_date`, `start_time`, `end_time` | |
| `1` Daily | `start_time`, `end_time` | recurrence attributes |
| `2` Weekly (By Time) | `start_day`, `start_time`, `end_day`, `end_time` | `execute_every`, recurrence attributes |
| `8` Weekly (By Day) | `start_time`, `duration` | `week_days`, `execute_every`, recurrence attributes |
| `5` Monthly (By Date) | `monthly_start_date`, `start_time`, `duration` | recurrence attributes |
| `6` Monthly (By Day) | `start_week`, `start_day`, `start_time`, `duration` | `start_after`, recurrence attributes |

The recurrence attributes are `maintenance_start_on`, `maintenance_end_type`, `maintenance_end_after_times` and `maintenance_end_on`. Yearly maintenance (`9`) is not supported by this resource.

Days of the week run from `0` (Sunday) to `6` (Saturday).

## Attributes Reference

### Required

* `display_name` (String) Display name for the maintenance.
* `maintenance_type` (Number) Type of maintenance: `1` Daily, `2` Weekly (By Time), `3` Once, `5` Monthly (By Date), `6` Monthly (By Day), `8` Weekly (By Day).
* `start_time` (String) Maintenance start time. Format: hh:mm.

### Optional

* `description` (String) Description for the maintenance.
* `end_time` (String) Maintenance end time. Format: hh:mm. Required for Once, Daily and Weekly (By Time).
* `time_zone` (String) Time zone for the maintenance. Defaults to the account time zone.
* `perform_monitoring` (Boolean) Perform uptime monitoring of the resources during the maintenance window. Default is `true`.
* `start_date` (String) Maintenance start date. Format: yyyy-mm-dd. Required for Once.
* `end_date` (String) Maintenance end date. Format: yyyy-mm-dd. Required for Once.
* `start_day` (Number) Day of the week the maintenance starts on, `0` (Sunday) to `6` (Saturday). Required for Weekly (By Time) and Monthly (By Day).
* `end_day` (Number) Day of the week the maintenance ends on, `0` (Sunday) to `6` (Saturday). Required for Weekly (By Time).
* `week_days` (List of Number) Days of the week the maintenance recurs on, `0` (Sunday) to `6` (Saturday). Used by Weekly (By Day).
* `execute_every` (Number) Interval in weeks at which a weekly maintenance recurs, `1` to `4`. Use `2` for bi-weekly. Default is `1`. Used by Weekly (By Time) and Weekly (By Day).
* `duration` (Number) Maintenance duration in minutes, under a day. Required for Weekly (By Day), Monthly (By Date) and Monthly (By Day).
* `monthly_start_date` (Number) Date of the month the maintenance recurs on, `1` to `31`. Required for Monthly (By Date).
* `start_week` (Number) Week of the month the maintenance recurs in: `1` First, `2` Second, `3` Third, `4` Fourth, `5` Last. Required for Monthly (By Day).
* `start_after` (Number) Number of days after which the maintenance begins every month, `0` to `25`. Default is `0`. Used by Monthly (By Day).
* `maintenance_start_on` (String) Date a recurring maintenance starts on. Format: yyyy-mm-dd. Defaults to the current date in the maintenance time zone. Not used by Once.
* `maintenance_end_type` (Number) When a recurring maintenance ends: `0` Never, `1` after `maintenance_end_after_times` runs, `2` on `maintenance_end_on`. Default is `0`. Not used by Once.
* `maintenance_end_after_times` (Number) Number of runs after which the maintenance ends. Required when `maintenance_end_type` is `1`.
* `maintenance_end_on` (String) Date the maintenance ends on. Format: yyyy-mm-dd. Required when `maintenance_end_type` is `2`.
* `selection_type` (Number) Resources the maintenance applies to: `1` Monitor Groups, `2` Monitors, `3` Tags. Default is `2`. All Monitors is not supported. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `monitors` (List of String) Monitor IDs. Required when `selection_type` is `2`.
* `monitor_groups` (List of String) Monitor group IDs. Required when `selection_type` is `1`.
* `tags` (List of String) Tag IDs. Required when `selection_type` is `3`.
* `subgroup_monitors` (Boolean) Include the subgroups of the selected monitor groups. Default is `true`. Used when `selection_type` is `1`.

## Import

Schedule maintenances can be imported using the maintenance ID, for example:

```
$ terraform import site24x7_schedule_maintenance.once 123456000007534020
```

Refer [API documentation](https://www.site24x7.com/help/api/#schedule-maintenances) for more information about attributes.
