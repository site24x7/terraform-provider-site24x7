terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.3"
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

// Site24x7 Monitor Group API doc - https://www.site24x7.com/help/api/#monitor-groups
resource "site24x7_monitor_group" "monitor_group_us" {
  // (Required) Display Name for the Monitor Group.
  display_name = "Website Group"

  // (Optional) Description for the Monitor Group.
  description = "This is the description of the monitor group from terraform"

  // Number of monitors' health that decide the group status. ‘0’ implies that all the monitors 
  // are considered for determining the group status. Default value is 1
  health_threshold_count = 1
  // (Optional) List of dependent resource ids.
  dependency_resource_id = ["100000000005938013"]
  // (Optional) Boolean value indicating whether to suppress alert when the dependent monitor is down
  // Setting suppress_alert = true with an empty dependency_resource_id is meaningless.
  suppress_alert = true
}