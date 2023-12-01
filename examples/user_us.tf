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

// User API doc: https://www.site24x7.com/help/api/#users
resource "site24x7_user" "user_basic" {

  // (Required) Name of the User.
  display_name = "User - Terraform"

  // (Required) Email address of the user.  Email verification has to be done manually.
  email_address = "jim@example.com"

  // (Required) Phone number configurations to receive alerts.
  mobile_settings = {
    "country_code" = "93"
    "mobile_number"= "434388234"
  }

  // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  notification_medium = [
    1,
  ]

  // (Required) Role assigned to the user for accessing Site24x7. Role will be updated only after the user accepts the invitation. Refer https://www.site24x7.com/help/api/#site24x7_user_constants
  user_role = 10
  
  // (Required) Medium through which you’d wish to receive the down alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  down_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the critical alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  critical_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the trouble alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  trouble_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the up alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  up_notification_medium = [
    1,
  ]
}

// User API doc: https://www.site24x7.com/help/api/#users
resource "site24x7_user" "user_example" {

  // (Required) Name of the User.
  display_name = "User - Terraform"

  // (Required) Email address of the user. Email verification has to be done manually.
  email_address = "jim@example.com"

  // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  notification_medium = [
    1,
  ]

  // (Required) Role assigned to the user for accessing Site24x7. Role will be updated only after the user accepts the invitation. Refer https://www.site24x7.com/help/api/#site24x7_user_constants
  user_role = 10
  
  // (Required) Medium through which you’d wish to receive the down alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  down_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the critical alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  critical_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the trouble alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  trouble_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the up alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  up_notification_medium = [
    1,
  ]

  // (Optional) Phone number configurations to receive alerts.
  mobile_settings = {
    "country_code" = "93"
    "mobile_number"= "434388234"
  }

  // (Optional) Provide your job title to be added in Site24x7. Refer https://www.site24x7.com/help/api/#job_title
  job_title = 1

  // (Optional) Resource type associated to this user. Default value is '0'. Can take values 0|1. '0' denotes 'All Monitors', '1' denotes 'Monitor Group'.
  selection_type = 1

  // (Optional) List of monitor groups to which the user has access to. 'monitor_groups' attribute is mandatory when the 'selection_type' is '1'.
  monitor_groups = [
    "306947000021059031",
    "306947000033224882"
  ]

  // (Optional) Groups to be associated for the user for receiving alerts.
  user_group_ids = [
    "306947000000025005",
    "306947000000025009",
    "306947000000025007"
  ]

}


