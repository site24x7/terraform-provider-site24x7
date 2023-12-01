---
layout: "site24x7"
page_title: "Site24x7: site24x7_credential_profile"
sidebar_current: "docs-site24x7-resource-credential-profile"
description: |-
  Create and manage a Credential Profile monitor in Site24x7.
---

# Resource: site24x7\_Credential\_Profile

Use this resource to create, update and delete a Credential Profile in Site24x7.

## Example Usage

```hcl

// Site24x7 Credential Profile API doc - https://www.site24x7.com/help/api/#credential-profile
resource "site24x7_credential_profile" "credential_profile_us" {
  // (Required) Credential Profile Name.
  credential_name = "Credential profile - terraform"
  // (Required) Credential Profile Type.
  credential_type = 3
  // (Required) Username for the Credential Profile.
  username = "Testing"
  // (Required) Password for the Credential Profile.
  password = "Test"
}

```

## Attributes Reference

### Required

* `credential_name` (String) Credential Profile Name.
* `credential_type` (Integer) Credential Profile Type.
* `username` (String) Username for the Credential Profile.
* `password` (String) Password for the Credential Profile.


Refer [API documentation](https://www.site24x7.com/help/api/#credential-profile) for more information about attributes.
