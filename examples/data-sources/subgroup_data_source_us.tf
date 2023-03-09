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

// Data source to fetch a Subgroup
data "site24x7_subgroup" "s247subgroup" {
  // (Required) Regular expression denoting the name of the monitor group.
  name_regex = "a"
  
}

// Displays the Subgroup ID
output "s247_subgroup_id" {
  description = "Subgroup ID : "
  value       = data.site24x7_subgroup.s247subgroup.id
}

// Displays the Subgroup Name
output "s247_subgroup_name" {
  description = "Subgroup Name : "
  value       = data.site24x7_subgroup.s247subgroup.display_name
}

// Displays the description
output "s247_subgroup_description" {
  description = "Subgroup description : "
  value       = data.site24x7_subgroup.s247subgroup.description
}

// Displays the parent group ID
output "s247_subgroup_parent_group_id" {
  description = "Parent Group IDs : "
  value       = data.site24x7_subgroup.s247subgroup.parent_group_id
}

// Displays the Top Group ID
output "s247_subgroup_top_group_id" {
  description = "Top Group ID : "
  value       = data.site24x7_subgroup.s247subgroup.top_group_id
}

// Displays the health threshold count
output "s247_subgroup_health_threshold_count" {
  description = "Subgroup health threshold count : "
  value       = data.site24x7_subgroup.s247subgroup.health_threshold_count
}

// Displays the monitors associated
output "s247_subgroup_monitors" {
  description = "Monitors Associated : "
  value       = data.site24x7_subgroup.s247subgroup.monitors
}

// Displays the type of the subgroup
output "s247_subgroup_type" {
  description = "Subgroup Type : "
  value       = data.site24x7_subgroup.s247subgroup.group_type
}
