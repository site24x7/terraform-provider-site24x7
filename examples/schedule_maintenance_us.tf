terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest

    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
  data_center = "US"

  // (Optional) ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

// Site24x7 Schedule Maintenance API doc - https://www.site24x7.com/help/api/#schedule-maintenances
//
// Value reference
// ---------------
// maintenance_type      1 = Daily
//                       2 = Weekly (By Time)   - from a start day/time to an end day/time
//                       3 = Once               - between two dates
//                       5 = Monthly (By Date)  - a date of the month, start time + duration
//                       6 = Monthly (By Day)   - a week of the month + day of the week
//                       8 = Weekly (By Day)    - chosen days of the week, start time + duration
//                       (9 = Yearly is not supported by this resource)
// Days of the week      0 = Sunday, 1 = Monday, 2 = Tuesday, 3 = Wednesday,
//                       4 = Thursday, 5 = Friday, 6 = Saturday
// start_week            1 = First, 2 = Second, 3 = Third, 4 = Fourth, 5 = Last week of the month
// maintenance_end_type  0 = Never ends (default)
//                       1 = Ends after maintenance_end_after_times runs
//                       2 = Ends on maintenance_end_on
// selection_type        1 = Monitor Groups, 2 = Monitors (default), 3 = Tags
//
// Formats
// -------
// Time      "hh:mm", 24-hour clock with two-digit hours, e.g. "09:05", "22:30". Not "9:05".
// Date      "yyyy-mm-dd", e.g. "2026-10-15".
// Duration  Whole minutes, e.g. 90 for 1 hour 30 minutes. Must be less than a day.
// IDs       Strings of digits, e.g. "123456000007534005".
//
// Each maintenance_type uses only its own attributes. terraform plan fails if
// an attribute the type needs is missing, or if one it does not use is set.

// -----------------------------------------------------------------------------
// 3 - Once: a single window between two dates.
// -----------------------------------------------------------------------------
resource "site24x7_schedule_maintenance" "once" {
  // (Required) Display name for the maintenance.
  display_name = "Switch upgrade - Terraform"

  // (Optional) Description for the maintenance.
  description = "Core switch firmware upgrade"

  // (Required) Type of maintenance. 3 = Once.
  maintenance_type = 3

  // (Required for Once) Date the window starts. Format: "yyyy-mm-dd".
  start_date = "2026-10-15"

  // (Required) Time the window starts. Format: "hh:mm" (24-hour).
  start_time = "19:30"

  // (Required for Once) Date the window ends. Format: "yyyy-mm-dd".
  // Same as start_date, or later for a window that runs past midnight.
  end_date = "2026-10-15"

  // (Required for Once, Daily and Weekly (By Time)) Time the window ends.
  // Format: "hh:mm" (24-hour).
  end_time = "21:00"

  // (Optional) Time zone of start_time and end_time, e.g. "GMT", "IST", "PST".
  // Default: your account time zone.
  time_zone = "GMT"

  // (Optional) Keep checking uptime during the window, so outages are still
  // recorded while alerts are suppressed. true or false. Default: true.
  perform_monitoring = true

  // (Optional) Resources the maintenance applies to.
  // 1 = Monitor Groups, 2 = Monitors, 3 = Tags. Default: 2.
  selection_type = 2

  // (Required when selection_type = 2) List of monitor IDs.
  monitors = ["123456000007534005"]
}

// -----------------------------------------------------------------------------
// 1 - Daily: every day at the same time.
// -----------------------------------------------------------------------------
resource "site24x7_schedule_maintenance" "daily" {
  // (Required) Display name for the maintenance.
  display_name = "Nightly backup - Terraform"

  // (Optional) Description for the maintenance.
  description = "Backup job window"

  // (Required) Type of maintenance. 1 = Daily.
  maintenance_type = 1

  // (Required) Time the window starts every day. Format: "hh:mm" (24-hour).
  start_time = "01:00"

  // (Required for Daily) Time the window ends every day. Format: "hh:mm" (24-hour).
  end_time = "02:00"

  // (Optional) First day the maintenance runs. Format: "yyyy-mm-dd".
  // Default: today, in the maintenance time zone.
  maintenance_start_on = "2026-10-10"

  // (Optional) When the recurring maintenance stops.
  // 0 = Never, 1 = after maintenance_end_after_times runs, 2 = on maintenance_end_on. Default: 0.
  maintenance_end_type = 1

  // (Required when maintenance_end_type = 1) Number of runs before the
  // maintenance stops. Whole number, 1 or more.
  maintenance_end_after_times = 30

  // (Required when selection_type = 2, the default) List of monitor IDs.
  monitors = ["123456000007534005"]
}

// -----------------------------------------------------------------------------
// 2 - Weekly (By Time): from a start day and time to an end day and time.
// This example runs from Saturday 22:00 to Sunday 04:00, every other week.
// -----------------------------------------------------------------------------
resource "site24x7_schedule_maintenance" "weekly_by_time" {
  // (Required) Display name for the maintenance.
  display_name = "Weekend patching - Terraform"

  // (Required) Type of maintenance. 2 = Weekly (By Time).
  maintenance_type = 2

  // (Required for Weekly (By Time)) Day of the week the window starts.
  // 0 = Sunday ... 6 = Saturday.
  start_day = 6

  // (Required) Time the window starts on start_day. Format: "hh:mm" (24-hour).
  start_time = "22:00"

  // (Required for Weekly (By Time)) Day of the week the window ends.
  // 0 = Sunday ... 6 = Saturday.
  end_day = 0

  // (Required for Weekly (By Time)) Time the window ends on end_day.
  // Format: "hh:mm" (24-hour).
  end_time = "04:00"

  // (Optional) Run every N weeks, 1 to 4. 1 = every week, 2 = every other week.
  // Default: 1.
  execute_every = 2

  // (Optional) First day the maintenance runs. Format: "yyyy-mm-dd".
  // Default: today, in the maintenance time zone.
  maintenance_start_on = "2026-10-10"

  // (Optional) When the recurring maintenance stops.
  // 0 = Never, 1 = after maintenance_end_after_times runs, 2 = on maintenance_end_on. Default: 0.
  maintenance_end_type = 2

  // (Required when maintenance_end_type = 2) Last date the maintenance runs.
  // Format: "yyyy-mm-dd".
  maintenance_end_on = "2027-03-31"

  // (Required when selection_type = 2, the default) List of monitor IDs.
  monitors = ["123456000007534005"]
}

// -----------------------------------------------------------------------------
// 8 - Weekly (By Day): on chosen days of the week, for a set duration.
// This example runs Monday, Wednesday and Friday at 14:30 for 60 minutes.
// -----------------------------------------------------------------------------
resource "site24x7_schedule_maintenance" "weekly_by_day" {
  // (Required) Display name for the maintenance.
  display_name = "Deploy window - Terraform"

  // (Required) Type of maintenance. 8 = Weekly (By Day).
  maintenance_type = 8

  // (Optional) Days of the week the window runs on. List of 0 = Sunday ... 6 = Saturday.
  week_days = [1, 3, 5]

  // (Required) Time the window starts on each day. Format: "hh:mm" (24-hour).
  start_time = "14:30"

  // (Required for Weekly (By Day)) Length of each window in minutes, less than a day.
  // Weekly (By Day) uses duration instead of end_time.
  duration = 60

  // (Optional) Run every N weeks, 1 to 4. 1 = every week, 2 = every other week.
  // Default: 1.
  execute_every = 1

  // (Optional) First day the maintenance runs. Format: "yyyy-mm-dd".
  // Default: today, in the maintenance time zone.
  maintenance_start_on = "2026-10-12"

  // (Optional) When the recurring maintenance stops. 0 = Never (default).
  maintenance_end_type = 0

  // (Required when selection_type = 2, the default) List of monitor IDs.
  monitors = ["123456000007534005"]
}

// -----------------------------------------------------------------------------
// 5 - Monthly (By Date): on a date of the month, for a set duration.
// This example runs on the 15th of every month at 03:00 for 2 hours, 12 times.
// -----------------------------------------------------------------------------
resource "site24x7_schedule_maintenance" "monthly_by_date" {
  // (Required) Display name for the maintenance.
  display_name = "Monthly DB maintenance - Terraform"

  // (Required) Type of maintenance. 5 = Monthly (By Date).
  maintenance_type = 5

  // (Required for Monthly (By Date)) Date of the month the window runs on, 1 to 31.
  monthly_start_date = 15

  // (Required) Time the window starts. Format: "hh:mm" (24-hour).
  start_time = "03:00"

  // (Required for Monthly (By Date)) Length of the window in minutes, less than a day.
  duration = 120

  // (Optional) First day the maintenance runs. Format: "yyyy-mm-dd".
  // Default: today, in the maintenance time zone.
  maintenance_start_on = "2026-11-01"

  // (Optional) When the recurring maintenance stops.
  // 0 = Never, 1 = after maintenance_end_after_times runs, 2 = on maintenance_end_on. Default: 0.
  maintenance_end_type = 1

  // (Required when maintenance_end_type = 1) Number of runs before the
  // maintenance stops. Whole number, 1 or more.
  maintenance_end_after_times = 12

  // (Required when selection_type = 2, the default) List of monitor IDs.
  monitors = ["123456000007534005"]
}

// -----------------------------------------------------------------------------
// 6 - Monthly (By Day): on a week of the month and day of the week.
// This example is Patch Tuesday: the second Tuesday of every month, for a
// monitor group without its subgroups.
// -----------------------------------------------------------------------------
resource "site24x7_schedule_maintenance" "monthly_by_day" {
  // (Required) Display name for the maintenance.
  display_name = "Patch Tuesday - Terraform"

  // (Optional) Description for the maintenance.
  description = "Monthly OS patching"

  // (Required) Type of maintenance. 6 = Monthly (By Day).
  maintenance_type = 6

  // (Required for Monthly (By Day)) Week of the month.
  // 1 = First, 2 = Second, 3 = Third, 4 = Fourth, 5 = Last.
  start_week = 2

  // (Required for Monthly (By Day)) Day of the week. 0 = Sunday ... 6 = Saturday.
  start_day = 2

  // (Optional) Start the window this many days after that day, 0 to 25.
  // Default: 0. For example, 2 runs the window on the Thursday after Patch Tuesday.
  start_after = 0

  // (Required) Time the window starts. Format: "hh:mm" (24-hour).
  start_time = "23:00"

  // (Required for Monthly (By Day)) Length of the window in minutes, less than a day.
  duration = 180

  // (Optional) Time zone of start_time, e.g. "GMT", "IST", "PST".
  // Default: your account time zone.
  time_zone = "GMT"

  // (Optional) Keep checking uptime during the window. true or false. Default: true.
  perform_monitoring = false

  // (Optional) Resources the maintenance applies to.
  // 1 = Monitor Groups, 2 = Monitors, 3 = Tags. Default: 2.
  selection_type = 1

  // (Required when selection_type = 1) List of monitor group IDs.
  monitor_groups = ["123456000007534010"]

  // (Optional, only when selection_type = 1) Include the monitors in the
  // subgroups of the selected monitor groups. true or false. Default: true.
  subgroup_monitors = false

  // To target tags instead, set selection_type = 3 and replace monitor_groups with:
  // (Required when selection_type = 3) List of tag IDs.
  # tags = ["123456000007534020"]
}
