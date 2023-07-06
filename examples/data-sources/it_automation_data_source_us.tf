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

// Data source to fetch an IT automation
data "site24x7_it_automation" "s247itautomation" {
  // (Required) Regular expression denoting the name of the IT automation.
  name_regex = "url"
  
}

// Displays the IT automation ID
output "s247_it_automation_id" {
  description = "IT automation ID : "
  value       = data.site24x7_it_automation.s247itautomation.id
}

// Displays the IT automation Name
output "s247_it_automation_name" {
  description = "IT automation name : "
  value       = data.site24x7_it_automation.s247itautomation.action_name
}

// Displays the IT automation Type
output "s247_it_automation_type" {
  description = "IT automation type : "
  value       = data.site24x7_it_automation.s247itautomation.action_type
}
// Displays the matching IT automation object
output "s247_matching_it_automation" {
  description = "Matching IT automation : "
  value       = data.site24x7_it_automation.s247itautomation
}

// Displays the matching IT automation IDs
output "s247_matching_ids" {
  description = "Matching IT automation IDs : "
  value       = data.site24x7_it_automation.s247itautomation.matching_ids
}

// Displays the matching IT automation IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching IT automation IDs and names : "
  value       = data.site24x7_it_automation.s247itautomation.matching_ids_and_names
}