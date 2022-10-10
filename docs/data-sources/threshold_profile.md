---
layout: "site24x7"
page_title: "Site24x7: site24x7_threshold_profile"
sidebar_current: "docs-site24x7-data-source-threshold-profile"
description: |-
  Get information about a threshold profile in Site24x7.
---

# Data Source: site24x7\_threshold\_profile

Use this data source to retrieve information about an existing threshold profile in Site24x7.

## Example Usage

```hcl

// Data source to fetch a threshold profile
data "site24x7_threshold_profile" "s247thresholdprofile" {
  // (Required) Regular expression denoting the name of the threshold profile.
  name_regex = "web"
  
}

// Displays the Threshold Profile ID
output "s247_threshold_profile_id" {
  description = "Threshold profile ID : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.id
}

// Displays the Threshold Profile Name
output "s247_threshold_profile_name" {
  description = "Threshold profile name : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.profile_name
}

// Displays the matching profile IDs
output "s247_matching_ids" {
  description = "Matching threshold profile IDs : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.matching_ids
}

// Displays the matching profile IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching threshold profile IDs and names : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.matching_ids_and_names
}

// Displays the Monitor Type
output "s247_montior_type" {
  description = "Monitor Type : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.type
}

// Displays the Profile Type
output "s247_profile_type" {
  description = "Profile Type : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.profile_type
}

// Displays the Down Location Threshold
output "s247_down_location_threshold" {
  description = "Down Location Threshold : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.down_location_threshold
}

// Displays the Website Content Modified
output "s247_website_content_modified" {
  description = "Website Content Modified : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.website_content_modified
}

// Iterating the threshold profile data source
data "site24x7_threshold_profile" "profilelist" {
  for_each = toset(["s", "a"])
  name_regex = each.key
}

locals {
  threshold_profile_ids = toset([for prof in data.site24x7_threshold_profile.profilelist : prof.id])
  threshold_profile_names = toset([for prof in data.site24x7_threshold_profile.profilelist : prof.profile_name])
}

output "s247_threshold_profile_ids" {
  description = "Matching threshold profile IDs : "
  value       = local.threshold_profile_ids
}

output "s247_threshold_profile_names" {
  description = "Matching threshold profile names : "
  value       = local.threshold_profile_names
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the threshold profile.

### Read-Only

* `id` (String) The ID of this resource.
* `matching_ids` (List) List of threshold profile IDs matching the `name_regex`.
* `matching_ids_and_names` (List) List of threshold profile IDs and names matching the `name_regex`.
* `profile_name` (String) Display name for the threshold profile.
* `type` (String) Type of the monitor for which the threshold profile is being created.
* `profile_type` (Number) Static Threshold(1) or AI-based Threshold(2).
* `down_location_threshold` (Number) Triggers alert when the monitor is down from configured number of locations. Default value is '3'.
* `website_content_modified` (Boolean) Boolean indicating whether alert will be triggered when the website content is modified.



 