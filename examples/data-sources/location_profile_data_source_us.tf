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
	// (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
	// environment variable if the attribute is empty or omitted.
	oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"
  
	// (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
	// environment variable if the attribute is empty or omitted.
	oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
  
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

// Data source to fetch a location profile
data "site24x7_location_profile" "s247locationprofile" {
  // (Required) Regular expression denoting the name of the location profile.
  name_regex = "a"
  
}

// Displays the Location Profile ID
output "s247_location_profile_id" {
  description = "Location Profile ID : "
  value       = data.site24x7_location_profile.s247locationprofile.id
}

// Displays the Location Profile Name
output "s247_location_profile_name" {
  description = "Location Profile Name : "
  value       = data.site24x7_location_profile.s247locationprofile.profile_name
}

// Displays the matching profile IDs and names
output "s247_matching_ids_and_names" {
  description = "Location Profile IDs and names : "
  value       = data.site24x7_location_profile.s247locationprofile.matching_ids_and_names
}

// Displays the Primary Location
output "s247_primary_location" {
  description = "Primary Location : "
  value       = data.site24x7_location_profile.s247locationprofile.primary_location
}


// Displays the Secondary Locations
output "s247_secondary_locations" {
  description = "Secondary Locations : "
  value       = data.site24x7_location_profile.s247locationprofile.secondary_locations
}

// Displays the Location consent for outer regions
output "s247_outer_regions_location_consent" {
  description = "LocationConsentForOuterRegions : "
  value       = data.site24x7_location_profile.s247locationprofile.outer_regions_location_consent
}


// Iterating the location profile data source
data "site24x7_location_profile" "profilelist" {
  for_each = toset(["America", "Asia", "Europe"])
  name_regex = each.key
}

locals {
  location_profile_ids = toset([for prof in data.site24x7_location_profile.profilelist : prof.id])
  location_profile_names = toset([for prof in data.site24x7_location_profile.profilelist : prof.profile_name])
}

output "s247_location_profile_ids" {
  description = "Location Profile IDs : "
  value       = local.location_profile_ids
}

output "s247_location_profile_names" {
  description = "Location Profile Names : "
  value       = local.location_profile_names
}