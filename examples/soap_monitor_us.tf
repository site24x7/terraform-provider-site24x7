terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.13.0"

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

// Site24x7 PING monitor API doc - https://www.site24x7.com/help/api/#PING
resource "site24x7_soap_monitor" "soap_monitor_basic" {
    // (Required) Display name for the monitor
    display_name = "Example ping Monitor"

    // (Required) host name of the monitor
    website = "www.example.com"
    //(Required)Request param for the monitor
    request_param= "param=value"

}

// Site24x7 PING monitor API doc - https://www.site24x7.com/help/api/#PING
resource "site24x7_ping_monitor" "ping_monitor_us" {
     // (Required) Display name for the monitor
    display_name = "Example ping Monitor"

    // (Required) host name of the monitor
    website = "www.example.com"

//(Required)Request param for the monitor
    request_param= "param=value"

// (Optional) Map of HTTP response headers to check.
      response_headers_severity = 0 // Can take values 0 or 2. '0' denotes Down and '2' denotes Trouble.
      response_headers = {
        "Content-Encoding" = "gzip"
        "Connection"       = "Keep-Alive"
      }

  //Soap attribute severity to be set
soap_attributes_severity=0
soap_attributes={
    "name" = "soap attribute name"
    "value" = "soap attribute value"
  }

  //Check frequency of the monitor
   check_frequency= "15"

   // (Optional) Id of the Location Profile to be associated to the monitor.
  location_profile_id= "456418000007860019"

  // (Optional) Credential Profile IDs to be associated to the monitor.
  credential_profile_id="456418000008650007"

   timeout= 45

  response_type="A"

  //Enable ALPN
  use_alpn = true

  // (Optional) Resolve the IP address using Domain Name Server. Default: true.
  use_name_server = false

    // HTTP Configuration
    // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response.
    // You can specify individual status codes, as well as ranges separated with a colon.
   up_status_codes = "400:500"
  
  ssl_protocol="TLSv1.2"
  
  http_protocol="H2"

  use_ipv6=true

 // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  // (Optional) List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    "123",
    "456"
  ]

  // (Optional) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
  dependency_resource_ids = [
    "123",
    "456"
  ]

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

    // (Optional) List of Threshold Profile IDs to be associated to the monitor.
threshold_profile_id="456418000008650009"

 perform_automation=true


  // (Optional) List of User Group IDs to be associated to the monitor.
 user_group_ids=["456418000000025007"]

   // (Optional) List of On call schedule ids IDs to be associated to the monitor.
 on_call_schedule_id="456418000001258016"


}