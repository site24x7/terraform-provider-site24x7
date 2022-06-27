terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.17"
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

// Data source to fetch URL monitor starting with the name "REST" and is of the monitor type "RESTAPI"
data "site24x7_monitor" "s247monitor" {
  // (Optional) Regular expression denoting the name of the monitor.
  name_regex = "^REST"
  // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
  monitor_type = "RESTAPI"
}


// Displays the monitor ID
output "s247monitor_monitor_id" {
  description = "Monitor ID : "
  value       = data.site24x7_monitor.s247monitor.id
}
// Displays the name
output "s247monitor_display_name" {
  description = "Monitor Display Name : "
  value       = data.site24x7_monitor.s247monitor.display_name
}
// Displays the user group IDs associated to the monitor
output "monitor_user_group_ids" {
  description = "Monitor User Group IDs : "
  value       = data.site24x7_monitor.s247monitor.user_group_ids
}
// Displays the notification profile ID associated to the monitor
output "s247monitor_notification_profile_id" {
  description = "Monitor Notification Profile ID : "
  value       = data.site24x7_monitor.s247monitor.notification_profile_id
}