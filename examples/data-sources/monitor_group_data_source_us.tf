terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.26"
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

// Data source to fetch a monitor group
data "site24x7_monitor_group" "s247monitorgroup" {
  // (Required) Regular expression denoting the name of the monitor group.
  name_regex = "1"
  
}

// Displays the Monitor Group ID
output "s247_monitor_group_id" {
  description = "Monitor Group ID : "
  value       = data.site24x7_monitor_group.s247monitorgroup.id
}

// Displays the Monitor Group Name
output "s247_monitor_group_name" {
  description = "Monitor Group Name : "
  value       = data.site24x7_monitor_group.s247monitorgroup.display_name
}

// Displays the description
output "s247_monitor_group_description" {
  description = "Monitor Group description : "
  value       = data.site24x7_monitor_group.s247monitorgroup.description
}

// Displays the health threshold count
output "s247_monitor_group_health_threshold_count" {
  description = "Monitor Group health threshold count : "
  value       = data.site24x7_monitor_group.s247monitorgroup.health_threshold_count
}

// Displays the monitors associated
output "s247_monitor_group_monitors" {
  description = "Monitors Associated : "
  value       = data.site24x7_monitor_group.s247monitorgroup.monitors
}

// Displays the dependency resource IDs
output "s247_monitor_group_dependency_resource_ids" {
  description = "Dependency resource IDs : "
  value       = data.site24x7_monitor_group.s247monitorgroup.dependency_resource_ids
}

// Displays the suppress alert
output "s247_monitor_group_suppress_alert" {
  description = "Suppress Alert : "
  value       = data.site24x7_monitor_group.s247monitorgroup.suppress_alert
}
