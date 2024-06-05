---
layout: "site24x7"
page_title: "Site24x7: site24x7_monitor_group"
sidebar_current: "docs-site24x7-resource-monitor-group"
description: |-
  Create and manage a monitor group in Site24x7.
---

# Resource: site24x7\_monitor\_group

Use this resource to create, update and delete a monitor group in Site24x7.

## Example Usage

```hcl

// Site24x7 Monitor Group API doc - https://www.site24x7.com/help/api/#monitor-groups
resource "site24x7_monitor_group" "monitor_group_us" {
  // (Required) Display Name for the Monitor Group.
  display_name = "Website Group"

  // (Optional) Description for the Monitor Group.
  description = "This is the description of the monitor group from terraform"

  // Number of monitors' health that decide the group status. ‘0’ implies that all the monitors 
  // are considered for determining the group status. Default value is 1
  health_threshold_count = 1
  // (Optional) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
  dependency_resource_ids = ["100000000005938013"]
  // (Optional) Boolean value indicating whether to suppress alert when the dependent monitor is down
  // Setting suppress_alert = true with an empty dependency_resource_id is meaningless.
  suppress_alert = true

  // (Optional) Health check profile to be associated with the monitor group.
  healthcheck_profile_id = "100000000000029001"

  // (Optional) Notification profile to be associated with the monitor group. 
  notification_profile_id = "100000000000029001"

  // (Optional) List of user groups to be notified when the monitor group is down.
  user_group_ids = [
    "100000000000025005",
    "100000000000025007"
  ]

  // (Optional) List if tag IDs to be associated to the monitor group.
  tag_ids = [
    "100000000048172001"
  ]

  // (Optional) List of Third Party Service IDs to be associated to the monitor group.
  third_party_service_ids = [
    "100000000048172001"
  ]

  // (Optional) Enable incident management. Default value is false.
  enable_incident_management = true 

  // (Optional) Healing period for the incident.
  healing_period = 10

  // (Optional) Alert frequency for the incident.
  alert_frequency = 10 

  // (Optional) Enable periodic alerting.
  alert_periodically = true
}

```

## Attributes Reference


### Required

* `display_name` (String) Display Name for the Monitor Group.

### Optional

* `description` (String) Description for the Monitor Group.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `health_threshold_count` (Number) Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.
* `id` (String) The ID of this resource.
* `suppress_alert` (Boolean) Boolean value indicating whether to suppress alert when the dependent monitor is down. Setting suppress_alert = true with an empty dependency_resource_id is meaningless.
* `healthcheck_profile_id` (String) Health check profile to be associated with the monitor group.
* `notification_profile_id` (String) Notification profile to be associated with the monitor group.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor group is down.
* `tag_ids` (List of String) List of tag IDs to be associated to the monitor group.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor group.
* `enable_incident_management` (Boolean) Enable incident management. Default value is false.
* `healing_period` (Number) Healing period for the incident.
* `alert_frequency` (Number) Alert frequency for the incident.
* `alert_periodically` (Boolean) Enable periodic alerting.

Refer [API documentation](https://www.site24x7.com/help/api/#monitor-groups) for more information about attributes.
 