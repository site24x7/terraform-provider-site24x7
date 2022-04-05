---
layout: "site24x7"
page_title: "Site24x7: site24x7_user_group"
sidebar_current: "docs-site24x7-resource-user-group"
description: |-
  Create and manage a user group in Site24x7.
---

# Resource: site24x7\_user\_group

Use this resource to create, update and delete a user group in Site24x7.

## Example Usage

```hcl
// User Group API doc: https://www.site24x7.com/help/api/#user-groups
resource "site24x7_user_group" "user_group_us" {
  // (Required) Display name for the user group.
  display_name = "User Group - Terraform"
  // (Required) User IDs of the users to be associated to the group.
  users = ["123"]
  // (Required) Attribute Alert Group to be associated with the User Alert group.
  attribute_group_id = "456"
}
```

## Attributes Reference


### Required

* `attribute_group_id` (String) Attribute Alert Group to be associated with the User Alert group.
* `display_name` (String) Display name for the user group.
* `users` (List of String) User IDs of the users to be associated to the group.

### Optional

* `id` (String) The ID of this resource.
* `product_id` (Number) Product for which the user group is being created. Default value is 0.

Refer [API documentation](https://www.site24x7.com/help/api/#user-groups) for more information about attributes.


