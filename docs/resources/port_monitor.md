---
layout: "site24x7"
page_title: "Site24x7: site24x7_port_monitor"
sidebar_current: "docs-site24x7-resource-port-monitor"
description: |-
  Create and manage a Port monitor in Site24x7.
---

# Resource: site24x7\_port\_monitor

Use this resource to create, update and delete a Port monitor in Site24x7.

## Example Usage

```hcl

// Site24x7 Port Monitor API doc - https://www.site24x7.com/help/api/#port-(custom-protocol)
resource "site24x7_port_monitor" "port_monitor_example" {

  // (Required) Display name for the monitor
  display_name = "Port Monitor"

  //(Required)Host name of the monitor
  host_name="www.example.com"

  //(Optional)Whois Server Port.Default value is 43
  port=443

  //Select IPv6 for monitoring the websites hosted with IPv6 address. If 
  //you choose non IPv6 supported locations, monitoring will happen through
  //IPv4.
  use_ipv6=true

  //Use secure connection (SSL).
  use_ssl=true

  //Invert the default behaviour of PORT check.
  invert_port_check=false

  //(Optional)Timeout for connecting to the host.Range 1 - 45.
  timeout=10 

  //(Optional)Toggle button to perform automation or not
  perform_automation=true

  //(Optional)if user_group_ids is not choosen
  //On-Call Schedule of your choice.
  //Create new On-Call Schedule or find your preferred On-Call Schedule ID.
  on_call_schedule_id="456418000001258016"

  // (Optional) Check for the keyword in the response.
  matching_keyword = {
     severity= 2
     value= "sample"
    }
  
  // (Optional) Check for non existence of keyword in the response.
  unmatching_keyword = {
  	  severity= 2
  	  value= "smile"
    }
  
  // (Optional) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
  actions = {1=465545643755}

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
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be 
  //used.
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

}
```
## Attributes Reference

### Required
* `display_name` (String) Display Name for the monitor.
* `host_name`(String)Registered domain name.
### Optional
* `id` (String) The ID of this resource.
* `type` DOMAINEXPIRY
* `domain_name`(String)Who is server.
* `port`(int)  Whois Server Port
* `timeout`(int) Timeout for connecting to the host.
* `use_ipv6`(bool) Prefer IPV6
* `use_ssl` (bool) Prefer SSL
* `perform_automation` (bool) Automating the scheduled maintenance
* `invert_port_check` (bool) Invert the default behaviour of PORT check.
* `command` (string) Use this to send plain text commands to the host.
* `application_type` (string) Provide the application type that runs on the configured port.
* `check_frequency` Check interval for monitoring.
* `matching_keyword` (Map of String) Check for the keyword in the website response.
* `unmatching_keyword` (Map of String) Check for non existence of keyword in the website response.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.

Refer [API documentation](https://www.site24x7.com/help/api/#port-(custom-protocol)) for more information about attributes.
