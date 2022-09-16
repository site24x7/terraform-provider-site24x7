---
layout: "site24x7"
page_title: "Site24x7: site24x7_monitor_group"
sidebar_current: "docs-site24x7-data-source-monitor-group"
description: |-
  Get information about a monitor group in Site24x7.
---

# Data Source: site24x7\_monitor\_group

Use this data source to retrieve information about an existing monitor group in Site24x7.

## Example Usage

```hcl

// Data source to fetch a monitor group
data "site24x7_monitor_group" "s247monitorgroup" {
  // (Required) Regular expression denoting the name of the monitor group.
  name_regex = "1"
  
}

// Displays the Monitor Group ID
output "s247_monitor_group_id" {
  description = "Monitor Group ID : "
  value       = data.site24x7_monitor_group.s247monitorgroup.id
}

// Displays the Monitor Group Name
output "s247_monitor_group_name" {
  description = "Monitor Group Name : "
  value       = data.site24x7_monitor_group.s247monitorgroup.display_name
}

// Displays the description
output "s247_monitor_group_description" {
  description = "Monitor Group description : "
  value       = data.site24x7_monitor_group.s247monitorgroup.description
}

// Displays the health threshold count
output "s247_monitor_group_health_threshold_count" {
  description = "Monitor Group health threshold count : "
  value       = data.site24x7_monitor_group.s247monitorgroup.health_threshold_count
}

// Displays the monitors associated
output "s247_monitor_group_monitors" {
  description = "Monitors Associated : "
  value       = data.site24x7_monitor_group.s247monitorgroup.monitors
}

// Displays the dependency resource IDs
output "s247_monitor_group_dependency_resource_ids" {
  description = "Dependency resource IDs : "
  value       = data.site24x7_monitor_group.s247monitorgroup.dependency_resource_ids
}

// Displays the suppress alert
output "s247_monitor_group_suppress_alert" {
  description = "Suppress Alert : "
  value       = data.site24x7_monitor_group.s247monitorgroup.suppress_alert
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the monitor_group.

### Read-Only

* `id` (String) The ID of this resource.
* `display_name` (String) Display Name for the Monitor Group.
* `description` (String) Description for the Monitor Group.
* `health_threshold_count` (String) Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status.
* `monitors` (List) List of monitors associated to the group.
* `dependency_resource_ids` (List) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `suppress_alert` (Boolean) Attribute indicating whether alert will be suppressed when the dependent monitor is down.


 