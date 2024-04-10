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

// Data source to fetch a MSP customer
data "site24x7_msp" "s247mspcustomer" {
  // (Required) Regular expression denoting the name of the MSP customer.
  customer_name_regex = "a"
  
}

// Displays MSP customer ID
output "s247_msp_id" {
  description = "MSP customer ID : "
  value       = data.site24x7_msp.s247mspcustomer.id
}

// Displays MSP customer name
output "s247_msp_name" {
  description = "MSP customer name : "
  value       = data.site24x7_msp.s247mspcustomer.customer_name
}

// Displays the matching ZAAIDs
output "s247_matching_ids" {
  description = "Matching MSP customer IDs : "
  value       = data.site24x7_msp.s247mspcustomer.matching_zaaids
}

// Displays the matching ZAAIDs and names
output "s247_matching_ids_and_names" {
  description = "Matching MSP customer IDs and names : "
  value       = data.site24x7_msp.s247mspcustomer.matching_zaaids_and_names
}

// Displays ZAAID of the customer
output "s247_msp_customer_zaaid" {
  description = "ZAAID : "
  value       = data.site24x7_msp.s247mspcustomer.zaaid
}

// Displays user ID of the customer
output "s247_msp_customer_userid" {
  description = "User ID : "
  value       = data.site24x7_msp.s247mspcustomer.user_id
}