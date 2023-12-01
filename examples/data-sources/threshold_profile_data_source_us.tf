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
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP).
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

// Data source to fetch a threshold profile
data "site24x7_threshold_profile" "s247thresholdprofile" {
  // (Required) Regular expression denoting the name of the threshold profile.
  name_regex = "web"
  
}

// Displays the Threshold Profile ID
output "s247_threshold_profile_id" {
  description = "Threshold profile ID : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.id
}

// Displays the Threshold Profile Name
output "s247_threshold_profile_name" {
  description = "Threshold profile name : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.profile_name
}

// Displays the matching profile IDs
output "s247_matching_ids" {
  description = "Matching threshold profile IDs : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.matching_ids
}

// Displays the matching profile IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching threshold profile IDs and names : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.matching_ids_and_names
}

// Displays the Monitor Type
output "s247_montior_type" {
  description = "Monitor Type : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.type
}

// Displays the Profile Type
output "s247_profile_type" {
  description = "Profile Type : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.profile_type
}

// Displays the Down Location Threshold
output "s247_down_location_threshold" {
  description = "Down Location Threshold : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.down_location_threshold
}

// Displays the Website Content Modified
output "s247_website_content_modified" {
  description = "Website Content Modified : "
  value       = data.site24x7_threshold_profile.s247thresholdprofile.website_content_modified
}

// Iterating the threshold profile data source
data "site24x7_threshold_profile" "profilelist" {
  for_each = toset(["s", "a"])
  name_regex = each.key
}

locals {
  threshold_profile_ids = toset([for prof in data.site24x7_threshold_profile.profilelist : prof.id])
  threshold_profile_names = toset([for prof in data.site24x7_threshold_profile.profilelist : prof.profile_name])
}

output "s247_threshold_profile_ids" {
  description = "Matching threshold profile IDs : "
  value       = local.threshold_profile_ids
}

output "s247_threshold_profile_names" {
  description = "Matching threshold profile names : "
  value       = local.threshold_profile_names
}
