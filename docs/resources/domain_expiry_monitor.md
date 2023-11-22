---
layout: "site24x7"
page_title: "Site24x7: site24x7_domain_expiry_monitor"
sidebar_current: "docs-site24x7-resource-domain-expiry-monitor"
description: |-
  Create and manage a Domain Expiry monitor in Site24x7.
---

# Resource: site24x7\_domain\_expiry\_monitor

Use this resource to create, update and delete a Domain-Expiry monitor in Site24x7.

## Example Usage


// Site24x7 Domain-Expiry Monitor API doc - https://www.site24x7.com/help/api/#domain-expiry
resource "site24x7_domain_expiry_monitor" 
"domain_expiry_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Domain Expiry Monitor"

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

## Attributes Reference

### Required
* `display_name` (String) Display Name for the monitor.
* `type` DOMAINEXPIRY
### Optional
* `id` (String) The ID of this resource.
* `host_name`(String)Registered domain name.
* `domain_name`(String)Who is server.
* `port`(int)  Whois Server Port
* `timeout`(int) Timeout for connecting to the host.
* `use_ipv6`(bool) Prefer IPV6
* `expire_days`(int) Day threshold for domain expiry notification.
* `ignore_registry_date`(bool)Ignores the registry expiry date and prefer registrar expiry date when notifying for domain expiry.
* `matching_keyword` (Map of String) Check for the keyword in the website response.
* `unmatching_keyword` (Map of String) Check for non existence of keyword in the website response.
* `match_case` (Boolean) Perform case sensitive keyword search or not.
* `match_regex` (Map of String) Match the regular expression in the website response.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.

Refer [API documentation](https://www.site24x7.com/help/api/#domain-expiry) for more information about attributes.
