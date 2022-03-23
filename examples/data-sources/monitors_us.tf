terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.8.1"
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

// Data source to fetch all URL monitors
data "site24x7_monitors" "s247monitors" {
  // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
  monitor_type = "URL"
}
// Displays the monitor IDs
output "s247monitors_ids" {
  description = "Monitor IDs : "
  value       = data.site24x7_monitors.s247monitors.ids
}
// Displays the monitor IDs and names of the monitors
output "s247monitors_ids_and_names" {
  description = "Monitor IDs and Names : "
  value       = data.site24x7_monitors.s247monitors.ids_and_names
}

// Data source to fetch URL monitors starting with the name "zylker"
data "site24x7_monitors" "zylkerMonitorIDs" {
  // (Optional) Regular expression denoting the name of the monitor.
  name_regex = "^zylker"
  // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
  monitor_type = "URL"
}

// Displays the monitor IDs
output "zylkerMonitorIDs_monitor_id" {
  description = "Zylker Monitor IDs : "
  value       = data.site24x7_monitors.zylkerMonitorIDs.ids
}

// Displays the monitor IDs and names of the monitors
output "zylkerMonitorIDs_monitor_id_and_names" {
  description = "Zylker Monitor IDs and Names : "
  value       = data.site24x7_monitors.zylkerMonitorIDs.ids_and_names
}