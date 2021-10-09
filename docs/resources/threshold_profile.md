---
layout: "site24x7"
page_title: "Site24x7: site24x7_threshold_profile"
sidebar_current: "docs-site24x7-resource-threshold-profile"
description: |-
  Create and manage a threshold profile in Site24x7.
---

# Resource: site24x7\_threshold\_profile

Use this resource to create, update, and delete a threshold profile in Site24x7.

## Example Usage

```hcl
// Site24x7 Threshold Profile API doc](https://www.site24x7.com/help/api/#threshold-website))
resource "site24x7_threshold_profile" "website_threshold_profile_us" {
  // (Required) Name of the profile
  profile_name = "URL Threshold Profile - Terraform"
  // (Required) Type of the profile - Denotes monitor type
  type = "URL"
  // (Optional) Threshold profile types - https://www.site24x7.com/help/api/#threshold_profile_types
  profile_type = 1
  // (Optional) Triggers alert when the monitor is down from configured number of locations.
  down_location_threshold = 1
  // (Optional) Triggers alert when Website content is modified.
  website_content_modified = false
  // (Optional) Triggers alert when Website content changes by configured percentage.
  website_content_changes {
    severity     = 2
    value = 80
  }
  website_content_changes {
    severity     = 3
    value = 95
  }
}
```

## Attributes Reference

### Required

* `profile_name` (String) Display Name for the threshold profile
* `type` (String) Type of the monitor for which the threshold profile is being created.

### Optional

* `down_location_threshold` (Number) Triggers alert when the monitor is down from configured number of locations.
* `id` (String) The ID of this resource.
* `profile_type` (Number) Static Threshold(1) or AI-based Threshold(2)
* `website_content_changes` (Block List) (see [below for nested schema](#nestedblock--website_content_changes))
* `website_content_modified` (Boolean) Triggers alert when the website content is modified.

<a id="nestedblock--website_content_changes"></a>
### Nested Schema for `website_content_changes`

Required:

* `severity` (Number)
* `value` (Number)

Optional:

* `comparison_operator` (Number)


Refer [API documentation](https://www.site24x7.com/help/api/#threshold-website) for more information about attributes.

