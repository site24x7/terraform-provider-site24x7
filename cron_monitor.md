---
layout: "site24x7"
page_title: "Site24x7: site24x7_cron_monitor"
sidebar_current: "docs-site24x7-resource-cron-monitor"
description: |-
  Create and manage a Cron monitor in Site24x7.
---

# Resource: site24x7\_cron\_monitor

Use this resource to create, update and delete a Cron monitor in Site24x7.

## Example Usage

```hcl

// Site24x7 Cron Monitor API doc - https://www.site24x7.com/help/api/#cron
resource "site24x7_cron_monitor" "cron_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "Sample test"
  
  // (Required) Cron expression to denote the job schedule.
  cron_expression = "* * * * *"
  
  // (Required) Timezone of the server where job runs.
  cron_tz = "IST"
  
  // (Required) Provide an extended period of time (seconds) to define when your alerts should be triggered. This is basically to avoid false alerts.
  wait_time = 30
}

// Site24x7 Cron Monitor API doc - https://www.site24x7.com/help/api/#cron
resource "site24x7_cron_monitor" "cron_monitor_all_attributes" {
 // (Required) Display name for the monitor
  display_name = "Sample test"
  
  // (Required) Cron expression to denote the job schedule.
  cron_expression = "* * * * *"
  
  // (Required) Timezone of the server where job runs.
  cron_tz = "IST"
  
  // (Required) Provide an extended period of time (seconds) to define when your alerts should be triggered. This is basically to avoid false alerts.
  wait_time = 30

  // (Optional) Threshold profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/threshold_profiles
  // endpoint for the CRON monitor type (https://www.site24x7.com/help/api/#list-threshold-profiles) will
  // be used.
  threshold_profile_id = "123"

  // (Optional) Notification profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/notification_profiles
  // endpoint (https://www.site24x7.com/help/api/#list-notification-profiles)
  // will be used.
  notification_profile_id = "123"

  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  // (Optional) List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    "123",
    "456"
  ]

  // (Optional) List if user group IDs to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_ids = [
    "123",
  ]

  // (Optional) List if user group names to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]

  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]

  // (Optional) List of tag names to be associated to the monitor. Tag name matching works for both exact and 
  //  partial match. Either specify tag_ids or tag_names.
  tag_names = [
    "Terraform",
    "Server",
  ]

  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]

  // (Optional) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.
  on_call_schedule_id = "3455"
}

```

## Attributes Reference

### Required

* `display_name` (String) Display Name for the monitor.
* `name_in_ping_url` (String) Unique name to be used in the ping URL.

### Optional

* `id` (String) The ID of this resource.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.

Refer [API documentation](https://www.site24x7.com/help/api/#cron) for more information about attributes.
