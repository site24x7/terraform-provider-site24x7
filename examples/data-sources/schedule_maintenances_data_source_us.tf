terraform {
  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest

    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  # oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // (Optional) ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
  data_center = "US"
}

// Site24x7 Schedule Maintenances API doc - https://www.site24x7.com/help/api/#schedule-maintenances
// Every filter is optional; the filters that are set are combined.

// All maintenances in the account.
data "site24x7_schedule_maintenances" "all" {}

// Maintenances that list this monitor directly (selection_type = 2).
// Maintenances that reach the monitor through a group or tag are not included.
data "site24x7_schedule_maintenances" "for_monitor" {
  // (Optional) Monitor ID.
  monitor_id = "113770000039133011"
}

// Maintenances for a monitor group (selection_type = 1), looked up by name.
data "site24x7_monitor_group" "payments" {
  name_regex = "^Payments$"
}

data "site24x7_schedule_maintenances" "for_group" {
  // (Optional) Monitor group ID.
  monitor_group_id = data.site24x7_monitor_group.payments.id
}

// Maintenances for a tag (selection_type = 3).
data "site24x7_schedule_maintenances" "for_tag" {
  // (Optional) Tag ID.
  tag_id = "113770000039135099"
}

// One-time (maintenance_type = 3) maintenances whose name starts with "Release".
data "site24x7_schedule_maintenances" "one_time_releases" {
  // (Optional) Regular expression matched against the display name.
  name_regex = "^Release"
  // (Optional) Maintenance type code.
  maintenance_type = 3
}

output "all_maintenances" {
  description = "Every maintenance as <id>__<display_name>"
  value       = data.site24x7_schedule_maintenances.all.ids_and_names
}

output "monitor_maintenance_ids" {
  description = "Maintenances that list the monitor directly"
  value       = data.site24x7_schedule_maintenances.for_monitor.ids
}

output "group_maintenance_windows" {
  description = "Schedule of each maintenance for the Payments group"
  value = [
    for m in data.site24x7_schedule_maintenances.for_group.maintenances :
    "${m.display_name}: ${m.start_date} ${m.start_time} - ${m.end_date} ${m.end_time} (${m.time_zone})"
  ]
}
