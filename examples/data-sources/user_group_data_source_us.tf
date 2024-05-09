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
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
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

// Data source to fetch a user group
data "site24x7_user_group" "s247usergroup" {
  // (Required) Regular expression denoting the name of the user group.
  name_regex = "a"
  
}

// Displays the User Group ID
output "s247_user_group_id" {
  description = "User group ID : "
  value       = data.site24x7_user_group.s247usergroup.id
}

// Displays the User Group Name
output "s247_user_group_name" {
  description = "User group name : "
  value       = data.site24x7_user_group.s247usergroup.display_name
}

// Displays the matching usergroup IDs
output "s247_matching_ids" {
  description = "Matching user group IDs : "
  value       = data.site24x7_user_group.s247usergroup.matching_ids
}

// Displays the matching usergroup IDs and names
output "s247_matching_ids_and_names" {
  description = "Matching user group IDs and names : "
  value       = data.site24x7_user_group.s247usergroup.matching_ids_and_names
}

// Displays the Users
output "s247_usergroup_users" {
  description = "Users : "
  value       = data.site24x7_user_group.s247usergroup.users
}

// Displays the attribute group ID
output "s247_attribute_group_id" {
  description = "Attribute group ID : "
  value       = data.site24x7_user_group.s247usergroup.attribute_group_id
}

// Displays the product ID
output "s247_product_id" {
  description = "Product ID : "
  value       = data.site24x7_user_group.s247usergroup.product_id
}

// Iterating the user group data source
data "site24x7_user_group" "usergrouplist" {
  for_each = toset(["e", "a"])
  name_regex = each.key
}

locals {
  user_group_ids = toset([for prof in data.site24x7_user_group.usergrouplist : prof.id])
  user_group_names = toset([for prof in data.site24x7_user_group.usergrouplist : prof.display_name])
}

output "s247_user_group_ids" {
  description = "Matching user group IDs : "
  value       = local.user_group_ids
}

output "s247_user_group_names" {
  description = "Matching user group names : "
  value       = local.user_group_names
}
