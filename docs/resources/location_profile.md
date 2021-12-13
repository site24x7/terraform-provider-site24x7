---
layout: "site24x7"
page_title: "Site24x7: site24x7_location_profile"
sidebar_current: "docs-site24x7-resource-location-profile"
description: |-
  Create and manage a location profile in Site24x7.
---

# Resource: site24x7\_location\_profile

Use this resource to create, update, and delete a location profile in Site24x7.

## Example Usage

```hcl
// Site24x7 Location Profile API doc - https://www.site24x7.com/help/api/#location-profiles
resource "site24x7_location_profile" "location_profile_us" {
  // (Required) Display name for the location profile.
  profile_name = "Location Profile - Terraform"

  // (Required) Primary location for monitoring.
  primary_location = "20"

  // (Optional) List of secondary locations for monitoring.
  secondary_locations = [
    "106",
	  "8",
	  "113",
	  # "94"
  ]

  // (Optional) Restricts polling of the resource from the selected locations alone in the Location Profile, overrides the alternate location poll logic.
  restrict_alternate_location_polling = true
}
```

## Attributes Reference

### Required

* `profile_name` (String) Display name for the location profile.
* `primary_location` (String) Primary location for monitoring.

### Optional

* `secondary_locations` (List of String) List of secondary locations for monitoring.
* `restrict_alternate_location_polling` (Boolean) Restricts polling of the resource from the selected locations alone in the Location Profile, overrides the alternate location poll logic.


Refer [API documentation](https://www.site24x7.com/help/api/#location-profiles) for more information about attributes.


