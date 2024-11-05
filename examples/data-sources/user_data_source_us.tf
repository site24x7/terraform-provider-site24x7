terraform {
  # Require Terraform version 0.15.x (recommended)

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      //version = "1.0.92"
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


// Data source to fetch a user
data "site24x7_user" "s247user" {
  // (Required) Regular expression denoting the name of the user.
  name_regex = "vinoth"  
}

// Displays the User ID
output "s247_user_id" {
  description = "User ID : "
  value       = data.site24x7_user.s247user.id
}

// Displays the User ID
output "s247_user_email" {
  description = "Email : "
  value       = data.site24x7_user.s247user.email
}

// Displays the User Role
output "s247_user_displayname" {
  description = "Display_Name : "
  value       = data.site24x7_user.s247user.display_name
}

// Displays the matched userIDS
output "s247_user_matchingUserIDs" {
  description = "Role : "
  value       = data.site24x7_user.s247user.matching_ids
}

// Displays the matched IDs and Names 

output "s247_user_matchingUserIDsAndNames" {
  description = "Role : "
  value       = data.site24x7_user.s247user.matching_ids_and_names
}
//
