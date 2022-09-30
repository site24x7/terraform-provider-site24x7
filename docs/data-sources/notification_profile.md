---
layout: "site24x7"
page_title: "Site24x7: site24x7_notification_profile"
sidebar_current: "docs-site24x7-data-source-notification-profile"
description: |-
  Get information about a notification profile in Site24x7.
---

# Data Source: site24x7\_notification\_profile

Use this data source to retrieve information about an existing notification profile in Site24x7.

## Example Usage

```hcl

// Data source to fetch a notification profile
data "site24x7_notification_profile" "s247notificationprofile" {
  // (Required) Regular expression denoting the name of the notification profile.
  name_regex = "t"
  
}

// Displays the Notification Profile ID
output "s247_notification_profile_id" {
  description = "Notification profile ID : "
  value       = data.site24x7_notification_profile.s247notificationprofile.id
}

// Displays the Notification Profile Name
output "s247_notification_profile_name" {
  description = "Notification profile name : "
  value       = data.site24x7_notification_profile.s247notificationprofile.profile_name
}

// Displays the matching profile IDs
output "s247_matching_ids" {
  description = "Matching notification profile IDs : "
  value       = data.site24x7_notification_profile.s247notificationprofile.matching_ids
}

// Displays the matching profile IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching notification profile IDs and names : "
  value       = data.site24x7_notification_profile.s247notificationprofile.matching_ids_and_names
}

// Displays the RCA Needed
output "s247_rca_needed" {
  description = "RCA Needed : "
  value       = data.site24x7_notification_profile.s247notificationprofile.rca_needed
}


// Displays the Notfiy after executing actions
output "s247_notify_after_executing_actions" {
  description = "Notfiy after executing actions : "
  value       = data.site24x7_notification_profile.s247notificationprofile.notify_after_executing_actions
}

// Displays the Suppress Automation
output "s247_suppress_automation" {
  description = "Suppress automation : "
  value       = data.site24x7_notification_profile.s247notificationprofile.suppress_automation
}


// Iterating the notification profile data source
data "site24x7_notification_profile" "profilelist" {
  for_each = toset(["terra", "a"])
  name_regex = each.key
}

locals {
  notification_profile_ids = toset([for prof in data.site24x7_notification_profile.profilelist : prof.id])
  notification_profile_names = toset([for prof in data.site24x7_notification_profile.profilelist : prof.profile_name])
}

output "s247_notification_profile_ids" {
  description = "Matching notification profile IDs : "
  value       = local.notification_profile_ids
}

output "s247_notification_profile_names" {
  description = "Matching notification profile names : "
  value       = local.notification_profile_names
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the notification profile.

### Read-Only

* `id` (String) The ID of this resource.
* `matching_ids` (List) List of notification profile IDs matching the `name_regex`.
* `matching_ids_and_names` (List) List of notification profile IDs and names matching the `name_regex`.
* `profile_name` (String) Display name for the notification profile.
* `rca_needed` (Boolean) Configuration denoting whether send root cause analysis when the monitor is down is enabled for this profile.
* `notify_after_executing_actions` (Boolean) Configuration denoting whether to raise alerts for downtime only after executing the pre-configured monitor actions.
* `template_id` (String) Email template ID for notification.
* `suppress_automation` (Boolean) Configuration denoting whether to stop automation from being executed on the dependent monitors.



 