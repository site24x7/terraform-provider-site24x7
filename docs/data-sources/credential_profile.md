---
layout: "site24x7"
page_title: "Site24x7: site24x7_credential_profile"
sidebar_current: "docs-site24x7-data-source-credential-profile"
description: |-
  Get information about a Credential Profile in Site24x7.
---

# Data Source: site24x7\_credential\_profile

Use this data source to retrieve information about an existing Credential Profiles in Site24x7.

## Example Usage

```hcl

// Data source to fetch an Credential Profile
data "site24x7_credential_profile" "s247credentialprofile" {
  // (Required) Regular expression denoting the name of the Credential Profile.
  name_regex = "url"
  
}

// Displays the Credential Profile ID
output "s247_credential_profile_id" {
  description = "Credential Profile ID : "
  value       = data.site24x7_credential_profile.s247credentialprofile.id
}

// Displays the Credential Profile Name
output "s247_credential_profile_name" {
  description = "Credential Profile name : "
  value       = data.site24x7_credential_profile.s247credentialprofile.credential_name
}

// Displays the Credential Profile Type
output "s247_credential_profile_type" {
  description = "Credential Profile type : "
  value       = data.site24x7_credential_profile.s247credentialprofile.credential_type
}
// Displays the Credential Profile username
output "s247_credential_profile_username" {
  description = "Credential Profile username: "
  value       = data.site24x7_credential_profile.s247credentialprofile.username
}

```

## Attributes Reference

### Required

* `name_regex` (String) Regular expression denoting the name of the Credential Profile.

### Read-Only

* `id` (String) The ID of the matching Credential Profile.
* `credential_type` (Integer) Type for the Credential Profile.
* `credential_name` (String) Name for the Credential Profile.
* `username` (String) Username for the Credential Profile.
 