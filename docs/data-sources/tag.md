---
layout: "site24x7"
page_title: "Site24x7: site24x7_tag"
sidebar_current: "docs-site24x7-data-source-tag"
description: |-
  Get information about a tag in Site24x7.
---

# Data Source: site24x7\_tag

Use this data source to retrieve information about an existing tag in Site24x7.

## Example Usage

```hcl

// Data source to fetch a tag
data "site24x7_tag" "s247tag" {
  // (Required) Regular expression denoting the name of the tag.
  tag_name_regex = "tagname"
  // (Optional) Regular expression denoting the value of the tag.
  tag_value_regex = "tagvalue"
}


// Displays the Tag ID
output "s247_tag_id" {
  description = "Tag ID : "
  value       = data.site24x7_tag.s247tag.id
}
// Displays the Tag Name
output "s247_tag_name" {
  description = "Tag Name : "
  value       = data.site24x7_tag.s247tag.tag_name
}
// Displays the Tag Value
output "s247_tag_value" {
  description = "Tag Value : "
  value       = data.site24x7_tag.s247tag.tag_value
}
// Displays the Tag Type
output "s247_tag_type" {
  description = "Tag Type : "
  value       = data.site24x7_tag.s247tag.tag_type
}
// Displays the Tag Color
output "s247_tag_color" {
  description = "Tag Color : "
  value       = data.site24x7_tag.s247tag.tag_color
}

```

## Attributes Reference

### Required

* `tag_name_regex` (String) Regular expression denoting the name of the tag.

### Optional

* `tag_value_regex` (String) Regular expression denoting the value of the tag.

### Read-Only

* `id` (String) The ID of this resource.
* `tag_name` (String) Display Name for the Tag.
* `tag_value` (String) Value for the Tag.
* `tag_type` (String) Type of the Tag.
* `tag_color` (String) Tag color code.


 