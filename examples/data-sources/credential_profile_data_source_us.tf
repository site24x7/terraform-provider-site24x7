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

// Data source to fetch an Credential Profile
data "site24x7_credential_profile" "s247credentialprofile" {
  // (Required) Regular expression denoting the name of the Credential Profile.
  name_regex = "url"
  
}

// Displays the Credential Profile ID
output "s247_credential_profile_id" {
  description = "Credential Profile ID : "
  value       = data.site24x7_credential_profile.s247credentialprofile.id
}

// Displays the Credential Profile Name
output "s247_credential_profile_name" {
  description = "Credential Profile name : "
  value       = data.site24x7_credential_profile.s247credentialprofile.credential_name
}

// Displays the Credential Profile Type
output "s247_credential_profile_type" {
  description = "Credential Profile type : "
  value       = data.site24x7_credential_profile.s247credentialprofile.credential_type
}
// Displays the Credential Profile username
output "s247_credential_profile_username" {
  description = "Credential Profile username: "
  value       = data.site24x7_credential_profile.s247credentialprofile.username
}
