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

// Site24x7 Schedule Maintenance API doc - https://www.site24x7.com/help/api/#schedule-maintenances
resource "site24x7_schedule_maintenance" "schedule_maintenance_basic" {
  // (Required) Display name for the maintenance.
  display_name = "Schedule Maintenance - Terraform"
  // (Optional) Description for the maintenance.
  description = "Switch upgrade"
  // (Required) Mandatory, if the maintenance_type chosen is Once. Maintenance start date. Format - yyyy-mm-dd.
  start_date = "2022-06-15"
  // (Required) Mandatory, if the maintenance_type chosen is Once. Maintenance end date. Format - yyyy-mm-dd.
  end_date = "2022-06-15"
  // (Required) Maintenance start time. Format - hh:mm
  start_time = "19:41"
  // (Required) Maintenance end time. Format - hh:mm
  end_time = "20:44"
  // (Optional) Resource Type associated with this integration. Default value is '2'. Can take values 1|2|3. '1' denotes 'Monitor Group', '2' denotes 'Monitors', '3' denotes 'Tags'.
  selection_type = 2
  // (Optional) Monitors that need to be associated with the maintenance window when the selection_type = 2.
  monitors = ["123456000007534005"]
  // (Optional) Monitor Groups that need to be associated with the maintenance window when the selection_type = 1.
  # monitor_groups = ["756"]
  # // (Optional) Tags that need to be associated with the maintenance window when the selection_type = 3.
  # tags = ["345"]
  // (Optional) Enable this to perform uptime monitoring of the resource during the maintenance window.
  perform_monitoring = true
}