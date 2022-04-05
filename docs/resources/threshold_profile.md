---
layout: "site24x7"
page_title: "Site24x7: site24x7_threshold_profile"
sidebar_current: "docs-site24x7-resource-threshold-profile"
description: |-
  Create and manage a threshold profile in Site24x7.
---

# Resource: site24x7\_threshold\_profile

Use this resource to create, update and delete a threshold profile in Site24x7.

## Example Usage

```hcl
// Site24x7 Threshold Profile API doc](https://www.site24x7.com/help/api/#threshold-website))
resource "site24x7_threshold_profile" "website_threshold_profile_us" {
  // (Required) Name of the profile
  profile_name = "URL Threshold Profile - Terraform"
  // (Required) Type of the profile - Denotes monitor type (eg) RESTAPI, SSL_CERT
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
  // (Optional) Response time threshold configuration
  primary_response_time_trouble_threshold = {
    // https://www.site24x7.com/help/api/#threshold_severity
    // 2 - Trouble
    severity = 2
    // https://www.site24x7.com/help/api/#constants
    // 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=), 6 - Not Equal to (≠)
    comparison_operator = 2
    // Attribute Threshold Value
    value               = 1000
    // https://www.site24x7.com/help/api/#threshold_strategy
    // 1 - Poll Count, 2 - Poll Average, 3 - Time Range, 4 - Average Time
    strategy            = 2
    // Poll Check Value
    polls_check         = 5
  }

  primary_response_time_critical_threshold = {
    // https://www.site24x7.com/help/api/#threshold_severity
    // 3 - Critical
    severity = 3
    // https://www.site24x7.com/help/api/#constants
    // 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=), 6 - Not Equal to (≠)
    comparison_operator = 1
    // Attribute Threshold Value
    value               = 2000
    // https://www.site24x7.com/help/api/#threshold_strategy
    // 1 - Poll Count, 2 - Poll Average, 3 - Time Range, 4 - Average Time
    strategy            = 2
    // Poll Check Value
    polls_check         = 5
  }

  secondary_response_time_trouble_threshold = {
    // https://www.site24x7.com/help/api/#threshold_severity
    // 2 - Trouble
    severity = 2
    // https://www.site24x7.com/help/api/#constants
    // 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=), 6 - Not Equal to (≠)
    comparison_operator = 1
    // Attribute Threshold Value
    value               = 3000
    // https://www.site24x7.com/help/api/#threshold_strategy
    // 1 - Poll Count, 2 - Poll Average, 3 - Time Range, 4 - Average Time
    strategy            = 2
    // Poll Check Value
    polls_check         = 5
  }

  secondary_response_time_critical_threshold = {
    // https://www.site24x7.com/help/api/#threshold_severity
    // 3 - Critical
    severity = 3
    // https://www.site24x7.com/help/api/#constants
    // 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=), 6 - Not Equal to (≠)
    comparison_operator = 1
    // Attribute Threshold Value
    value               = 4000
    // https://www.site24x7.com/help/api/#threshold_strategy
    // 1 - Poll Count, 2 - Poll Average, 3 - Time Range, 4 - Average Time
    strategy            = 2
    // Poll Check Value
    polls_check         = 5
  }
}
```

## Attributes Reference

### Required

* `profile_name` (String) Display Name for the threshold profile
* `type` (String) Type of the monitor for which the threshold profile is being created. Refer [API documentation](https://www.site24x7.com/help/api/#threshold-parameters) for more information about type.

### Optional

* `id` (String) The ID of this resource.
* `down_location_threshold` (Number) Triggers alert when the monitor is down from configured number of locations. Default value is '3'
* `profile_type` (Number) Static Threshold(1) or AI-based Threshold(2)
* `website_content_changes` (Block List) Triggers alert when Website content changes by configured percentage. (see [below for nested schema](#nestedblock--website_content_changes))
* `website_content_modified` (Boolean) Triggers alert when the website content is modified.
* `primary_response_time_trouble_threshold` (Map of Number) Response time trouble threshold for the primary monitoring location. (see [below for map schema](#nestedblock--response_time_threshold))
* `primary_response_time_critical_threshold` (Map of Number) Response time critical threshold for the primary monitoring location. (see [below for map schema](#nestedblock--response_time_threshold))
* `secondary_response_time_trouble_threshold` (Map of Number) Response time trouble threshold for the secondary monitoring location. (see [below for map schema](#nestedblock--response_time_threshold))
* `secondary_response_time_critical_threshold` (Map of Number) Response time critical threshold for the secondary monitoring location. (see [below for map schema](#nestedblock--response_time_threshold))

<a id="nestedblock--website_content_changes"></a>
### Nested Schema for `website_content_changes`

### Required

* `severity` (Number) Trouble(2), Critical(3). Refer [API documentation](https://www.site24x7.com/help/api/#threshold_severity) for more information about threshold severity.
* `value` (Number) Attribute Threshold Value

### Optional

* `comparison_operator` (Number) 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=), 6 - Not Equal to (≠). Refer [API documentation](https://www.site24x7.com/help/api/#constants) for more information about comparison operator.

<a id="nestedblock--response_time_threshold"></a>
### Map Schema for `primary_response_time_trouble_threshold`, `primary_response_time_critical_threshold`, `secondary_response_time_trouble_threshold`, `secondary_response_time_critical_threshold`

### Required

* `severity` (Number) Trouble(2), Critical(3). Refer [API documentation](https://www.site24x7.com/help/api/#threshold_severity) for more information about threshold severity.
* `comparison_operator` (Number) 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=), 6 - Not Equal to (≠). Refer [API documentation](https://www.site24x7.com/help/api/#constants) for more information about comparison operator.
* `value` (Number) Attribute Threshold Value
* `strategy` (Number) Poll Count(1), Poll Average(2), Time Range(3), Average Time(4). Refer [API documentation](https://www.site24x7.com/help/api/#threshold_strategy) for more information about threshold strategy.
* `polls_check` (Number) Poll Check Value


Refer [API documentation](https://www.site24x7.com/help/api/#threshold-website) for more information about attributes.

