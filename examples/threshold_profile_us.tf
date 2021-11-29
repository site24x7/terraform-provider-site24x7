terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "0.0.1-beta.4"
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  # // The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  # // environment variable if the attribute is empty or omitted.
  # oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  # // The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  # // environment variable if the attribute is empty or omitted.
  # oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

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