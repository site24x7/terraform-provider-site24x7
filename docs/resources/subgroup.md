---
layout: "site24x7"
page_title: "Site24x7: site24x7_subgroup"
sidebar_current: "docs-site24x7-resource-subgroup"
description: |-
  Create and manage a subgroup in Site24x7.
---

# Resource: site24x7\_subgroup

Use this resource to create, update and delete a subgroup in Site24x7.

## Example Usage

```hcl

// Subgroup API doc: https://www.site24x7.com/help/api/#subgroups
resource "site24x7_subgroup" "subgroup_default" {
  // (Required) Display Name for the Subgroup.
  display_name = "Default subgroup - Terraform"
  // (Required) Unique ID of the top monitor group for which business view has been configured.
  top_group_id = "123456000033743001"
  // (Required) Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).
  parent_group_id = "123456000033743001"
}

// Subgroup API doc: https://www.site24x7.com/help/api/#subgroups
resource "site24x7_subgroup" "subgroup_zoho" {
  // (Required) Display Name for the Subgroup.
  display_name = "Zoho Subgroup - Terraform"
  // (Required) Unique ID of the top monitor group for which business view has been configured.
  top_group_id = "123456000033743001"
  // (Required) Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).
  parent_group_id = "123456000033743001"
  // (Optional) Description for the Subgroup.
  description = "This is the description of the subgroup"
  // (Optional) Denotes the type of monitors that can be associated. ‘1’ implies that all type of monitors can be associated with this subgroup. Default value is 1. '2' - Web, '3' - Port/Ping, '4' - Server, '5' - Database, '6' - Synthetic Transaction, '7' - Web API, '8' - APM Insight,'9' - Network Devices, '10' - RUM, '11' - AppLogs Monitor
  group_type = 1
  // (Optional) Monitors to be associated with the Subgroup.
  monitors = [
    "123456000024411005",
  ]
  // (Optional) Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.
  health_threshold_count = 1
  // (Optional) Threshold profile to be associated with the Subgroup.
  healthcheck_profile_id = "123456000000025005"
  // (Optional) Notification profile to be associated with the Subgroup.
  notification_profile_id = "123456000000029001"
  // (Optional) The user groups to be notified during an outage. Mandatory, if on_call_schedule_id is not chosen.
  user_group_ids = [
    "123456000000025009",
  ]
  // (Optional) On-Call Schedule to be associated with the Subgroup. Mandatory, if user_group_ids is not chosen.
  on_call_schedule_id = "123456000000025007"
  // (Optional) List of Third Party Service IDs to be associated to the Subgroup.
  third_party_service_ids = [
    "123456000000025011",
  ]
  // (Optional) List of tag IDs to be associated to the Subgroup.
  tag_ids = [
    "123456000000025013",
  ]
}

```

## Attributes Reference


### Required

* `display_name` (String) Display Name for the Subgroup.
* `parent_group_id` (String) Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).
* `top_group_id` (String) Unique ID of the top monitor group for which business view has been configured.

### Optional

* `id` (String) The ID of this resource.
* `description` (String) Description for the Subgroup.
* `group_type` (Number) Denotes the type of monitors that can be associated. Default value is 1. '1' implies that all type of monitors can be associated with this subgroup. '2' - Web, '3' - Port/Ping, '4' - Server, '5' - Database, '6' - Synthetic Transaction, '7' - Web API, '8' - APM Insight,'9' - Network Devices, '10' - RUM, '11' - AppLogs Monitor.
* `monitors` (List of String) List of monitors to be associated with the Subgroup.
* `health_threshold_count` (Number) Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.
* `healthcheck_profile_id` (String) Threshold profile to be associated with the Subgroup.
* `notification_profile_id` (String) Notification profile to be associated with the Subgroup.
* `user_group_ids` (List of String) The user groups to be notified during an outage. Mandatory, if `on_call_schedule_id` is not chosen.
* `on_call_schedule_id` (String) On-Call Schedule to be associated with the Subgroup. Mandatory, if `user_group_ids` is not chosen.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the Subgroup.
* `tag_ids` (Set of String) List of tag IDs to be associated to the Subgroup.

### Read-Only

* `check_frequency` (Number) Health Check polling interval.


Refer [API documentation](https://www.site24x7.com/help/api/#subgroups) for more information about attributes.
 