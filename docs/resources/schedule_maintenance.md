---
layout: "site24x7"
page_title: "Site24x7: site24x7_schedule_maintenance"
sidebar_current: "docs-site24x7-schedule-maintenance"
description: |-
Create and manage a Schedule Maintenance in Site24x7.
---

# Resource: site24x7\_schedule\_maintenance

Use this resource to create, update, and delete Schedule Maintenance in Site24x7.

## Example Usage

```hcl

// Site24x7 Schedule Maintenance API doc - https://www.site24x7.com/help/api/#schedule-maintenances
resource "site24x7_schedule_maintenance" "schedule_maintenance_basic" {
  // (Required) Display name for the maintenance.
  display_name = "Schedule Maintenance - Terraform"
  // (Optional) Description for the maintenance.
  description = "Switch upgrade"
  // (Required) Mandatory, if the maintenance_type chosen is Once. Maintenance start date. Format - yyyy-mm-dd.
  start_date = "2022-06-15"
  // (Required) Mandatory, if the maintenance_type chosen is Once. Maintenance end date. Format - yyyy-mm-dd.
  end_date = "2022-06-15"
  // (Required) Maintenance start time. Format - hh:mm
  start_time = "19:41"
  // (Required) Maintenance end time. Format - hh:mm
  end_time = "20:44"
  // (Optional) Resource Type associated with this integration. Default value is '2'. Can take values 1|2|3. '1' denotes 'Monitor Group', '2' denotes 'Monitors', '3' denotes 'Tags'.
  selection_type = 2
  // (Optional) Monitors that need to be associated with the maintenance window when the selection_type = 2.
  monitors = ["123456000007534005"]
  // (Optional) Monitor Groups that need to be associated with the maintenance window when the selection_type = 1.
  # monitor_groups = ["756"]
  # // (Optional) Tags that need to be associated with the maintenance window when the selection_type = 3.
  # tags = ["345"]
  // (Optional) Enable this to perform uptime monitoring of the resource during the maintenance window.
  perform_monitoring = true
}

```

## Attributes Reference


### Required

* `display_name` (String) Display name for the maintenance.
* `start_date` (String) Mandatory, if the maintenance_type chosen is Once. Maintenance start date. Format - yyyy-mm-dd.
* `end_date` (String) Mandatory, if the maintenance_type chosen is Once. Maintenance end date. Format - yyyy-mm-dd.
* `start_time` (String) Maintenance start time. Format - hh:mm
* `end_time` (String) Maintenance end time. Format - hh:mm


### Optional

* `description` (String) Description for the maintenance.
* `selection_type` (Number) Resource Type associated with this integration. Default value is '2'. Can take values 1|2|3. '1' denotes 'Monitor Group', '2' denotes 'Monitors', '3' denotes 'Tags'. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `monitors` (List of String) Monitors that need to be associated with the maintenance window when the selection_type = 2.
* `monitor_groups` (List of String) Tags that need to be associated with the maintenance window when the selection_type = 3.
* `tags` (List of String) Tags that need to be associated with the maintenance window when the selection_type = 3.
* `perform_monitoring` (Boolean) Enable this to perform uptime monitoring of the resource during the maintenance window.

Refer [API documentation](https://www.site24x7.com/help/api/#schedule-maintenances) for more information about attributes.


