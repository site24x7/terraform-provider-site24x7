---
layout: "site24x7"
page_title: "Site24x7: site24x7_ftp_transfer_monitor"
sidebar_current: "docs-site24x7-resource-ftp-transfer-monitor"
description: |-
  Create and manage a FTP Transfer monitor in Site24x7.
---

# Resource: site24x7\_ftp\_transfer\_monitor

Use this resource to create, update and delete a FTP Transfer monitor in Site24x7.

## Example Usage


```hcl

// Site24x7 FTP Monitor API doc - https://www.site24x7.com/help/api/#ftp-transfer

resource "site24x7_ftp_transfer_monitor" "ftp_transfer_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "FTP Transfer Monitor"

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
* `hostname`  (String)Host or domain name for the monitor
* `username`    (String)username to access the file
* `destination` (String)Destination file path of the monitor
### Optional
* `id` (String) The ID of this resource.
* `type` (String) FTP
* `password` (String) password to access the file
* `check_upload` (Boolean) To check upload or not
* `check_download` (Boolean) To check download or not
* `protocol`    (String) Protocol of the monitor
* `port`  (Number) Port of the monitor
* `timeout` (Number) timeout period of the monitor
* `dependency_resource_ids` (List) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
*  `check_frequency` (String)The Endpoints are mentioned at this interval
* `perform_automation` (Boolean) To perform automation or not
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `threshold_profile_id` (String) Threshold profile associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.

Refer [API documentation](https://www.site24x7.com/help/api/) for more information about attributes.
