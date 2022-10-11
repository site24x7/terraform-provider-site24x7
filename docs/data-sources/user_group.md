---
layout: "site24x7"
page_title: "Site24x7: site24x7_user_group"
sidebar_current: "docs-site24x7-data-source-user-group"
description: |-
  Get information about a user group in Site24x7.
---

# Data Source: site24x7\_user\_group

Use this data source to retrieve information about an existing user group in Site24x7.

## Example Usage

```hcl

// Data source to fetch a user group
data "site24x7_user_group" "s247usergroup" {
  // (Required) Regular expression denoting the name of the user group.
  name_regex = "a"
  
}

// Displays the User Group ID
output "s247_user_group_id" {
  description = "User group ID : "
  value       = data.site24x7_user_group.s247usergroup.id
}

// Displays the User Group Name
output "s247_user_group_name" {
  description = "User group name : "
  value       = data.site24x7_user_group.s247usergroup.display_name
}

// Displays the matching usergroup IDs
output "s247_matching_ids" {
  description = "Matching user group IDs : "
  value       = data.site24x7_user_group.s247usergroup.matching_ids
}

// Displays the matching usergroup IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching user group IDs and names : "
  value       = data.site24x7_user_group.s247usergroup.matching_ids_and_names
}

// Displays the Users
output "s247_usergroup_users" {
  description = "Users : "
  value       = data.site24x7_user_group.s247usergroup.users
}

// Displays the attribute group ID
output "s247_attribute_group_id" {
  description = "Attribute group ID : "
  value       = data.site24x7_user_group.s247usergroup.attribute_group_id
}

// Displays the product ID
output "s247_product_id" {
  description = "Product ID : "
  value       = data.site24x7_user_group.s247usergroup.product_id
}

// Iterating the user group data source
data "site24x7_user_group" "usergrouplist" {
  for_each = toset(["e", "a"])
  name_regex = each.key
}

locals {
  user_group_ids = toset([for prof in data.site24x7_user_group.usergrouplist : prof.id])
  user_group_names = toset([for prof in data.site24x7_user_group.usergrouplist : prof.display_name])
}

output "s247_user_group_ids" {
  description = "Matching user group IDs : "
  value       = local.user_group_ids
}

output "s247_user_group_names" {
  description = "Matching user group names : "
  value       = local.user_group_names
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the user group.

### Read-Only

* `id` (String) The ID of this resource.
* `matching_ids` (List) List of user group IDs matching the `name_regex`.
* `matching_ids_and_names` (List) List of user group IDs and names matching the `name_regex`.
* `display_name` (String) Display name for the user group.
* `users` (List of String) User IDs of the users associated to the group.
* `attribute_group_id` (String) Attribute alert group associated with the user alert group.
* `product_id` (Number) Product for which the user group was created.




 