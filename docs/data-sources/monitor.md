---
layout: "site24x7"
page_title: "Site24x7: site24x7_monitor"
sidebar_current: "docs-site24x7-data-source-monitor"
description: |-
  Get information about a monitor in Site24x7.
---

# Data Source: site24x7\_monitor

Use this data source to retrieve information about an existing monitor in Site24x7.

## Example Usage

```hcl
// Data source to fetch URL monitor starting with the name "REST" and is of the monitor type "RESTAPI"
data "site24x7_monitor" "s247monitor" {
  // (Optional) Regular expression denoting the name of the monitor.
  name_regex = "^REST"
  // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
  monitor_type = "RESTAPI"
}


// Displays the monitor ID
output "s247monitor_monitor_id" {
  description = "Monitor ID ============================ "
  value       = data.site24x7_monitor.s247monitor.id
}
// Displays the name
output "s247monitor_display_name" {
  description = "Monitor Display Name ============================ "
  value       = data.site24x7_monitor.s247monitor.display_name
}
// Displays the user group IDs associated to the monitor
output "monitor_user_group_ids" {
  description = "Monitor User Group IDs ============================ "
  value       = data.site24x7_monitor.s247monitor.user_group_ids
}
// Displays the notification profile ID associated to the monitor
output "s247monitor_notification_profile_id" {
  description = "Monitor Notification Profile ID ============================ "
  value       = data.site24x7_monitor.s247monitor.notification_profile_id
}
```

## Attributes Reference


### Optional

* `name_regex` (String) Regular expression denoting the name of the monitor.
* `monitor_type` (String) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.

### Read-Only

* `id` (String) The ID of this resource.
* `display_name` (String) Display Name for the monitor.
* `notification_profile_id` (String) Notification profile associated with the monitor.
* `threshold_profile_id` (String) Threshold profile associated with the monitor.
* `location_profile_id` (String) Location profile associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups associated with the monitor.
* `user_group_ids` (List of String) List of user groups associated with the monitor.
* `tag_ids` (List of String) List of tags IDs associated with the monitor.
* `third_party_service_ids` (List of String) List of Third Party Service IDs associated with the monitor.


 