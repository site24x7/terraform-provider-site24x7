---
layout: "site24x7"
page_title: "Site24x7: site24x7_dns_server_monitor"
sidebar_current: "docs-site24x7-resource-dns-server-monitor"
description: |-
  Create and manage a DNS server monitor in Site24x7.
---

# Resource: site24x7\_dns\_server\_monitor

Use this resource to create, update and delete a DNS server monitor in Site24x7.

## Example Usage

```hcl

// DNS Server API doc: https://www.site24x7.com/help/api/#dns-server
resource "site24x7_dns_server_monitor" "dns_monitor_basic" {
  // (Required) Name for the monitor.
  display_name              = "Nowatt basic DNS monitor - Terraform"
  
  // (Required) DNS Name Server to be monitored
  dns_host                  = "185.43.51.84"
  
  // (Required) Domain name to be resolved.
  domain_name               = "www.nowatt.com"
}

// DNS Server API doc: https://www.site24x7.com/help/api/#dns-server
resource "site24x7_dns_server_monitor" "dns_server_monitor" {
  // (Required) Display Name for the monitor.
  display_name              = "Nowatt DNS monitor - Terraform"
  
  // (Required) DNS Name Server to be monitored
  dns_host                  = "185.43.51.84"

  // (Required) Domain name to be resolved.
  domain_name               = "www.nowatt.com"
  
  // (Optional) Port for DNS access. Default value: 53
  dns_port                  = "53"

  // (Optional)  Interval at which your DNS server has to be monitored. Default value is 5 minutes.
  check_frequency           = "5"

  // (Optional)  Timeout for connecting to your DNS server. Default value is 10. Range 1 - 45.
  timeout                   = 10

  // (Optional) Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.
  use_ipv6                  = false

  // (Optional) Value to be checked against resolved values. Choose parameters based on your configured lookup type. See https://www.site24x7.com/help/api/#dns_search_config
  search_config {
    lookup_type = "A"
    addr        = "1.2.3.4"
    ttlo        = "2"
    ttl         = "60"
  }

  // (Optional) Lookup Type - See https://www.site24x7.com/help/api/#dns_lookup_type
  // Lookup Types supported: 1 - A, 255 - ALL, 28 - AAAA, 2 - NS, 15 - MX, 5 - CNAME, 6 - SOA, 12 - PTR, 33 - SRV, 16 - TXT, 48 - DNSKEY, 257 - CAA, 43 - DS
  lookup_type               = 1

  // (Optional) Pass dnssec parameter to enable Site24x7 to validate DNS responses. See https://www.site24x7.com/help/internet-service-metrics/dns-monitor.html#dnssec
  dnssec                    = false

  // (Optional) Enable this attribute to auto discover and set up monitoring for all the related resources for the domain_name.
  deep_discovery            = false

  // (Optional) Location Profile to be associated with the monitor. If 
  // location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_id = "123"

  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

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

  // (Optional) Threshold profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/threshold_profiles
  // endpoint for the DNS server monitor (https://www.site24x7.com/help/api/#list-threshold-profiles) will
  // be used.
  threshold_profile_id = "123"

  // (Optional) List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    "123",
    "456"
  ]

  // (Optional) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
  dependency_resource_ids = [
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
    "Network",
  ]

  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]

  // (Optional) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
  actions = {
    "1" = "123"
  }
}

```

## Attributes Reference

### Required

* `display_name` (String) Display Name for the monitor.
* `dns_host` (String) DNS Name Server to be monitored.
* `domain_name` (String) Domain name to be resolved.

### Optional

* `id` (String) The ID of this resource.
* `dns_port` (String) Port for DNS access. Default value: 53.
* `check_frequency` (String) Interval at which your DNS server has to be monitored. Default value is 5 minutes.
* `timeout` (Number) Timeout for connecting to your DNS server. Default value is 10. Range 1 - 45.
* `use_ipv6` (Boolean) Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.
* `deep_discovery` (Boolean) Enable this attribute to auto discover and set up monitoring for all the related resources for the domain_name.
* `lookup_type` (Number) Lookup Types supported: 1 - A, 255 - ALL, 28 - AAAA, 2 - NS, 15 - MX, 5 - CNAME, 6 - SOA, 12 - PTR, 33 - SRV, 16 - TXT, 48 - DNSKEY, 257 - CAA, 43 - DS. DNS Server Lookup Type Constants. See https://www.site24x7.com/help/api/#dns_lookup_type
* `dnssec` (Boolean) Pass dnssec parameter to enable Site24x7 to validate DNS responses.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `actions` (Map of String) Action to be performed on monitor IT Automation templates. 


Refer [API documentation](https://www.site24x7.com/help/api/#dns-server) for more information about attributes.