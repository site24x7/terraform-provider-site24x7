---
layout: "site24x7"
page_title: "Site24x7: site24x7_user"
sidebar_current: "docs-site24x7-data-source-user"
description: |-
  Get information about a user  in Site24x7.
---

# Data Source: site24x7\_user

Use this data source to retrieve information about an existing user in Site24x7.

## Example Usage

```hcl

// Data source to fetch a user
data "site24x7_user" "s247user" {
  // (Required) Regular expression denoting the name of the user.
  name_regex = "vinoth"  
}

// Displays the User ID
output "s247_user_id" {
  description = "User ID : "
  value       = data.site24x7_user.s247user.id
}

// Displays the User ID
output "s247_user_email" {
  description = "Email : "
  value       = data.site24x7_user.s247user.email
}

// Displays the User Role
output "s247_user_displayname" {
  description = "Display_Name : "
  value       = data.site24x7_user.s247user.display_name
}

// Displays the matched userIDS
output "s247_user_matchingUserIDs" {
  description = "Role : "
  value       = data.site24x7_user.s247user.matching_ids
}

// Displays the matched IDs and Names 

output "s247_user_matchingUserIDsAndNames" {
  description = "Role : "
  value       = data.site24x7_user.s247user.matching_ids_and_names
}
//

// Iterating the user group data source
data "site24x7_user" "userlist" {
  for_each = toset(["e", "a"])
  name_regex = each.key
}

locals {
  user_group_ids = toset([for prof in data.site24x7_user.userlist : prof.id])
  user_group_names = toset([for prof in data.site24x7_user.userlist : prof.display_name])
}

output "s247_user_ids" {
  description = "Matching user IDs : "
  value       = local.user_group_ids
}

output "s247_user_names" {
  description = "Matching user names : "
  value       = local.user_names
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the user.

### Read-Only

* `id` (String) The ID of this resource.
* `matching_ids` (List) List of user  IDs matching the `name_regex`.
* `matching_ids_and_names` (List) List of user  IDs and names matching the `name_regex`.
* `display_name` (String) Display name for the user.
* `email` (String) Email address of the user.