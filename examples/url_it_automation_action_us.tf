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

// Site24x7 IT Automation API doc - https://www.site24x7.com/help/api/#it-automation
resource "site24x7_url_action" "action_us" {
  // (Required) Display name for the action.
  name = "IT Action"

  // (Required) URL to be invoked for action execution.
  url = "https://www.example.com"

  // (Optional) HTTP Method to access the URL. Default: "P". See
  // https://www.site24x7.com/help/api/#http_methods for allowed values.
  method = "G"

  // (Optional) If send_custom_parameters is set as true. Custom parameters to
  // be passed while accessing the URL.
  custom_parameters = "param=value"

  // (Optional) Configuration to send custom parameters while executing the action.
  send_custom_parameters = true

  // (Optional) Configuration to enable json format for post parameters.
  send_in_json_format = true

  // (Optional) Configuration to send incident parameters while executing the action.
  send_incident_parameters = true

  // (Optional) The amount of time a connection waits to time out. Range 1 - 90. Default: 30.
  timeout = 10
}

