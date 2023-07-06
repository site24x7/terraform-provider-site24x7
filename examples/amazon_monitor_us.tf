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

resource "site24x7_amazon_monitor" "aws_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "aws_added_via_terraform"
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) External ID for the AWS account
  external_id = "asdf"
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) AWS Role ARN
  role_arn = "arn:aws:iam::23243:role/TerraformAdminRole"
  // (Optional) AWS discover frequency
  aws_discovery_frequency = 5
  // (Optional) AWS services to discover. See https://www.site24x7.com/help/api/#aws_discover_services 
  // for knowing service ID.
  aws_discover_services = [1,2,3,4,5,6,8,11,14,15,16,17,18,19,20,21,22,23,25,27,29,30,31,32,33,34,35,38,39,40,41,42,43,45,46,49,48,47,53,59,56,57,58,60,61,62,63,65,66,69,70,68,75,76,79,82,83,85,87,92,95,88]
}

// Site24x7 Amazon Monitor API doc - https://www.site24x7.com/help/api/#amazon-webservice-monitor
resource "site24x7_amazon_monitor" "aws_monitor_site24x7" {
  // (Required) Display name for the monitor
  display_name = "aws_added_via_terraform"
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) External ID for the AWS account
  external_id = "asdf"
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) AWS Role ARN
  role_arn = ""
  // (Required) AWS External ID
  aws_external_id = ""
  // (Optional) AWS discover frequency
  aws_discovery_frequency = 5
  // (Optional) AWS services to discover. See https://www.site24x7.com/help/api/#aws_discover_services 
  // for knowing service ID.
  aws_discover_services = [1]
  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  // (Optional) List if user group IDs to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_ids = [
    "123",
  ]

  // (Optional) List if user group names to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]
  
  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]

  // (Optional) List of tag names to be associated to the monitor. Tag name matching works for both exact and 
  //  partial match. Either specify tag_ids or tag_names.
  tag_names = [
    "Terraform",
    "Network",
  ]

  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]
}

