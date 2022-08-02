terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.24"
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // ZAAID of the customer under a MSP or BU
  zaaid = "1234"

  // Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
  data_center = "US"

  // The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

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