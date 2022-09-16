---
layout: "site24x7"
page_title: "Site24x7: site24x7_location_profile"
sidebar_current: "docs-site24x7-data-source-location-profile"
description: |-
  Get information about a location profile in Site24x7.
---

# Data Source: site24x7\_location\_profile

Use this data source to retrieve information about an existing location profile in Site24x7.

## Example Usage

```hcl

// Data source to fetch a location profile
data "site24x7_location_profile" "s247locationprofile" {
  // (Required) Regular expression denoting the name of the location profile.
  name_regex = "a"
  
}

// Displays the Location Profile ID
output "s247_location_profile_id" {
  description = "Location Profile ID : "
  value       = data.site24x7_location_profile.s247locationprofile.id
}

// Displays the Location Profile Name
output "s247_location_profile_name" {
  description = "Location Profile Name : "
  value       = data.site24x7_location_profile.s247locationprofile.profile_name
}

// Displays the matching profile IDs and names
output "s247_matching_ids_and_names" {
  description = "Location Profile IDs and names : "
  value       = data.site24x7_location_profile.s247locationprofile.matching_ids_and_names
}

// Displays the Primary Location
output "s247_primary_location" {
  description = "Primary Location : "
  value       = data.site24x7_location_profile.s247locationprofile.primary_location
}


// Displays the Secondary Locations
output "s247_secondary_locations" {
  description = "Secondary Locations : "
  value       = data.site24x7_location_profile.s247locationprofile.secondary_locations
}

// Displays the Location consent for outer regions
output "s247_outer_regions_location_consent" {
  description = "LocationConsentForOuterRegions : "
  value       = data.site24x7_location_profile.s247locationprofile.outer_regions_location_consent
}


// Iterating the location profile data source
data "site24x7_location_profile" "profilelist" {
  for_each = toset(["America", "Asia", "Europe"])
  name_regex = each.key
}

locals {
  location_profile_ids = toset([for prof in data.site24x7_location_profile.profilelist : prof.id])
  location_profile_names = toset([for prof in data.site24x7_location_profile.profilelist : prof.profile_name])
}

output "s247_location_profile_ids" {
  description = "Location Profile IDs : "
  value       = local.location_profile_ids
}

output "s247_location_profile_names" {
  description = "Location Profile Names : "
  value       = local.location_profile_names
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the location profile.

### Read-Only

* `id` (String) The ID of this resource.
* `matching_ids_and_names` (List) List of location profile IDs and names matching the `name_regex`.
* `profile_name` (String) Display name for the location profile.
* `primary_location` (String) Primary location for monitoring.
* `secondary_locations` (List) List of secondary locations for monitoring.
* `outer_regions_location_consent` (Boolean) Attribute denoting whether consent is mandatory for monitoring from countries outside the European Economic Area (EEA) and the Adequate countries.
* `restrict_alt_loc` (Boolean) Restricts polling of the resource from the selected locations alone in the Location Profile, overrides the alternate location poll logic.



 