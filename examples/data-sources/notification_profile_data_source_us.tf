terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
	// environment variable if the attribute is empty or omitted.
	oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
	// environment variable if the attribute is empty or omitted.
	oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
    
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
	// environment variable if the attribute is empty or omitted.
	oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"
  
	// (Required) Specify the data center from which you have obtained your
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
	data_center = "US"
	
	// (Optional) ZAAID of the customer under a MSP or BU
	zaaid = "1234"
  
	// (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
	retry_min_wait = 1
  
	// (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
	// requests. This is the upper limit for the wait duration with exponential
	// backoff.
	retry_max_wait = 30
  
	// (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
	max_retries = 4
  
  }

// Data source to fetch a notification profile
data "site24x7_notification_profile" "s247notificationprofile" {
  // (Required) Regular expression denoting the name of the notification profile.
  name_regex = "t"
  
}

// Displays the Notification Profile ID
output "s247_notification_profile_id" {
  description = "Notification profile ID : "
  value       = data.site24x7_notification_profile.s247notificationprofile.id
}

// Displays the Notification Profile Name
output "s247_notification_profile_name" {
  description = "Notification profile name : "
  value       = data.site24x7_notification_profile.s247notificationprofile.profile_name
}

// Displays the matching profile IDs
output "s247_matching_ids" {
  description = "Matching notification profile IDs : "
  value       = data.site24x7_notification_profile.s247notificationprofile.matching_ids
}

// Displays the matching profile IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching notification profile IDs and names : "
  value       = data.site24x7_notification_profile.s247notificationprofile.matching_ids_and_names
}

// Displays the RCA Needed
output "s247_rca_needed" {
  description = "RCA Needed : "
  value       = data.site24x7_notification_profile.s247notificationprofile.rca_needed
}


// Displays the Notfiy after executing actions
output "s247_notify_after_executing_actions" {
  description = "Notfiy after executing actions : "
  value       = data.site24x7_notification_profile.s247notificationprofile.notify_after_executing_actions
}

// Displays the Suppress Automation
output "s247_suppress_automation" {
  description = "Suppress automation : "
  value       = data.site24x7_notification_profile.s247notificationprofile.suppress_automation
}


// Iterating the notification profile data source
data "site24x7_notification_profile" "profilelist" {
  for_each = toset(["terra", "a"])
  name_regex = each.key
}

locals {
  notification_profile_ids = toset([for prof in data.site24x7_notification_profile.profilelist : prof.id])
  notification_profile_names = toset([for prof in data.site24x7_notification_profile.profilelist : prof.profile_name])
}

output "s247_notification_profile_ids" {
  description = "Matching notification profile IDs : "
  value       = local.notification_profile_ids
}

output "s247_notification_profile_names" {
  description = "Matching notification profile names : "
  value       = local.notification_profile_names
}