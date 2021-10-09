---
layout: "site24x7"
page_title: "Site24x7: site24x7_monitor_group"
sidebar_current: "docs-site24x7-resource-monitor-group"
description: |-
  Create and manage a monitor group in Site24x7.
---

# Resource: site24x7\_monitor\_group

Use this resource to create, update, and delete a monitor group in Site24x7.

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
  // (Optional) List of dependent resource ids.
  dependency_resource_id = ["306947000005938013"]
  // (Optional) Boolean value indicating whether to suppress alert when the dependent monitor is down
  // Setting suppress_alert = true with an empty dependency_resource_id is meaningless.
  suppress_alert = true
}
```

## Attributes Reference


### Required

* `description` (String) Description for the Monitor Group.
* `display_name` (String) Display Name for the Monitor Group.

### Optional

* `dependency_resource_id` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `health_threshold_count` (Number) Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.
* `id` (String) The ID of this resource.
* `suppress_alert` (Boolean) Boolean value indicating whether to suppress alert when the dependent monitor is down. Setting suppress_alert = true with an empty dependency_resource_id is meaningless.


Refer [API documentation](https://www.site24x7.com/help/api/#monitor-groups) for more information about attributes.
