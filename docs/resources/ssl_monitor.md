---
layout: "site24x7"
page_title: "Site24x7: site24x7_ssl_monitor"
sidebar_current: "docs-site24x7-resource-ssl-monitor"
description: |-
  Create and manage a SSL monitor in Site24x7.
---

# Resource: site24x7\_ssl\_monitor

Use this resource to create, update and delete a SSL monitor in Site24x7.

## Example Usage

```hcl
// Site24x7 SSL Certificate Monitor API doc - https://www.site24x7.com/help/api/#ssl-certificate
resource "site24x7_ssl_monitor" "ssl_monitor_us" {
  // (Required) Display name for the monitor
  display_name = "Example SSL Monitor"
  // (Required) Domain name to be verified for SSL Certificate.
  domain_name = "www.example.com"
  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"
  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
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
* `domain_name` (String) Domain name to be verified for SSL Certificate.

### Optional

* `id` (String) The ID of this resource.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `timeout` (Number) Timeout for connecting to the host. Range 1 - 45.
* `expire_days` (Number) Day threshold for certificate expiry notification. Range 1 - 999.
* `http_protocol_version` (String) Version of the HTTP protocol.
* `ignore_domain_mismatch` (Boolean) Boolean to ignore domain name mismatch errors.
* `ignore_trust` (Boolean) To ignore the validation of SSL/TLS certificate chain.
* `port` (Number) Server Port.
* `protocol` (String) Supported protocols are HTTPS, SMTPS, POPS, IMAPS, FTPS or CUSTOM



Refer [API documentation](https://www.site24x7.com/help/api/#ssl-certificate) for more information about attributes.
