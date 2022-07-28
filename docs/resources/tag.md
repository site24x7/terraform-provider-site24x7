---
layout: "site24x7"
page_title: "Site24x7: site24x7_tag"
sidebar_current: "docs-site24x7-resource-tag"
description: |-
  Create and manage a tag in Site24x7.
---

# Resource: site24x7\_tag

Use this resource to create, update and delete a tag in Site24x7.

## Example Usage

```hcl
// Site24x7 Tag API doc - https://www.site24x7.com/help/api/#tags
resource "site24x7_tag" "tag_us" {
  // (Required) Display Name for the Tag.
  tag_name = "Website"

  // Value for the Tag.
  tag_value = "Zoho websites"

  // Color code for the Tag. Possible values are '#B7DA9E','#73C7A3','#B5DCDF','#D4ABBB','#4895A8','#DFE897','#FCEA8B','#FFC36D','#F79953','#F16B3C','#E55445','#F2E2B6','#DEC57B','#CBBD80','#AAB3D4','#7085BA','#F6BDAE','#EFAB6D','#CA765C','#999','#4A148C','#009688','#00ACC1','#0091EA','#8BC34A','#558B2F'
  tag_color = "#B7DA9E"

}
```

## Attributes Reference


### Required

* `tag_name` (String) Display Name for the Tag.
* `tag_value` (String) Value for the Tag.
* `tag_color` (String) Color code for the Tag. Possible values are '#B7DA9E','#73C7A3','#B5DCDF','#D4ABBB','#4895A8','#DFE897','#FCEA8B','#FFC36D','#F79953','#F16B3C','#E55445','#F2E2B6','#DEC57B','#CBBD80','#AAB3D4','#7085BA','#F6BDAE','#EFAB6D','#CA765C','#999','#4A148C','#009688','#00ACC1','#0091EA','#8BC34A','#558B2F'.

Refer [API documentation](https://www.site24x7.com/help/api/#tags) for more information about attributes.
 
