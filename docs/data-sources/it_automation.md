---
layout: "site24x7"
page_title: "Site24x7: site24x7_it_automation"
sidebar_current: "docs-site24x7-data-source-it-automation"
description: |-
  Get information about a IT automation in Site24x7.
---

# Data Source: site24x7\_it\_automation

Use this data source to retrieve information about an existing IT automation in Site24x7.

## Example Usage

```hcl

// Data source to fetch an IT automation
data "site24x7_it_automation" "s247itautomation" {
  // (Required) Regular expression denoting the name of the IT automation.
  name_regex = "url"
  
}

// Displays the IT automation ID
output "s247_it_automation_id" {
  description = "IT automation ID : "
  value       = data.site24x7_it_automation.s247itautomation.id
}

// Displays the IT automation Name
output "s247_it_automation_name" {
  description = "IT automation name : "
  value       = data.site24x7_it_automation.s247itautomation.action_name
}

// Displays the IT automation Type
output "s247_it_automation_type" {
  description = "IT automation type : "
  value       = data.site24x7_it_automation.s247itautomation.action_type
}
// Displays the matching IT automation object
output "s247_matching_it_automation" {
  description = "Matching IT automation : "
  value       = data.site24x7_it_automation.s247itautomation
}

// Displays the matching IT automation IDs
output "s247_matching_ids" {
  description = "Matching IT automation IDs : "
  value       = data.site24x7_it_automation.s247itautomation.matching_ids
}

// Displays the matching IT automation IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching IT automation IDs and names : "
  value       = data.site24x7_it_automation.s247itautomation.matching_ids_and_names
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the IT automation.

### Read-Only

* `id` (String) The ID of the matching IT automation.
* `matching_ids` (List) List of IT automation IDs matching the `name_regex`.
* `matching_ids_and_names` (List) List of IT automation IDs and names matching the `name_regex`.
* `action_name` (String) Display name for the IT automation.
* `action_type` (Number) Type of the IT automation.
* `url` (String) URL invoked for action execution.
* `method` (String) HTTP Method to access the action url.
* `timeout` (Number) "Timeout for connecting the website.
* `send_custom_parameters` (Boolean) Custom parameters sent while executing the action.
* `custom_parameters` (String) Custom parameters passed while accessing the action url.
* `send_in_json_format` (Boolean) Configuration to enable json format for post parameters.
* `send_email` (Boolean) Boolean indicating whether to send email or not.
* `send_incident_parameters` (Boolean) Configuration to send incident parameters while executing the action.
* `auth_method` (String) Authentication method to access the action url.
* `user_agent` (String) User Agent used while monitoring the website.
 