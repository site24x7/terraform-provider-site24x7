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

// Site24x7 Domain Expiry monitor API doc - https://www.site24x7.com/help/api/#web-transaction-(browser)
resource "site24x7_web_transaction_browser_monitor" "web_transaction_browser_monitor_basic" {

    // (Optional)Display name for the monitor
     display_name="RBM-Terraform"

    //(Optional) Toggle to enable asynchronous data collection or not
     async_dc_enabled = false
     
    //(Required)Base url for the monitor
      base_url= "https://www.demoqa.com/"


    //(Required)Selenium script to add to the monitor
      selenium_script="{\"id\":\"b500a6da-dbb7-4d0c-968e-ae3fd6ab411f\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://www.example.com\",\"tests\":[{\"id\":\"f30156f4-3a70-4031-9a41-dbe8eac7e494\",\"name\":\"\",\"commands\":[{\"id\":\"59b9fca1-28f3-46eb-afc5-0be35a1f582f\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://www.example.com\",\"targets\":[],\"value\":\"\",\"URL\":\"https://www.example.com\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"14bbdd0e-f78f-4591-916e-4b6cf94ce576\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"38e2db3f-b835-47e8-8786-b87792d6fe4f\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"f30156f4-3a70-4031-9a41-dbe8eac7e494\"]}],\"urls\":[\"https://www.example.com/\"],\"plugins\":[]}"

    //(Required)Script type of the selenium script added
      script_type="txt"

    // (Optional) Check interval for monitoring. Default: 1. See
    // https://www.site24x7.com/help/api/#check-interval for all supported
    // values.
      check_frequency = "1"

    //(Optional) browser version of the monitor 
     browser_version=10101

    //(Optional) Think time of the monitor
     think_time=3


     //(Optional)Page loading time
     page_load_time=60

    //(Optional)Resolution of the monitor
     resolution="1600,900"

    //(Optional)Ip type of the monitor
      ip_type = 2

    //(Optional) Browser type of the monitor
      browser_type = 1

    // (Optional) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
      dependency_resource_ids = [
        "123",
        "456"
      ]

    // (Optional) Provide the Custom Header for parameter forwarding in Map format.
      custom_headers ={
        "Accept" = "application/json"
      }

    // (Required) Provide the Cookies for parameter forwarding in Map format.
      cookies ={
        "Accept" = "application/json"
      }

    //(Required) Map of proxy-details 
      proxy_details={webProxyUrl: "qwerty", webProxyUname: "sdcsdcsdc", webProxyPass: "sadasds"}

    //(Optional)Map of auth-details
      auth_details={userName: "12345",  password: "12345"}

    // (Optional) User Agent to be used while monitoring the website.
      user_agent = "some user agent string"

    // (Optional) Threshold profile to be associated with the monitor. If
    // omitted, the first profile returned by the /api/threshold_profiles
    // endpoint for the website monitor (https://www.site24x7.com/help/api/#list-threshold-profiles) will
    // be used.
      threshold_profile_id = "123"

    //(Optional)Toggle button to perform automation or not
      perform_automation=true

    //(Optional)if user_group_ids is not choosen
    //On-Call Schedule of your choice.
    //Create new On-Call Schedule or find your preferred On-Call Schedule ID.
    on_call_schedule_id="456418000001258016"


  // (Optional) Location Profile to be associated with the monitor. If 
  // location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
    location_profile_id = "123"

  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
    location_profile_name = "North America"


  // (Optional) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
    actions = {1=465545643755}


  // (Optional) Notification profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/notification_profiles
  // endpoint (https://www.site24x7.com/help/api/#list-notification-profiles)
  // will be used.
    notification_profile_id = "123"

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

// Site24x7 Domain Expiry monitor API doc - https://www.site24x7.com/help/api/#web-transaction-(browser)
resource "site24x7_web_transaction_browser_monitor" "web_transaction_browser_monitor_basic2" {
    //(Required)Base url for the monitor
      base_url= "https://www.example.com/"


    //(Required)Selenium script to add to the monitor
      selenium_script="{\"id\":\"b500a6da-dbb7-4d0c-968e-ae3fd6ab411f\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://www.example.com\",\"tests\":[{\"id\":\"f30156f4-3a70-4031-9a41-dbe8eac7e494\",\"name\":\"\",\"commands\":[{\"id\":\"59b9fca1-28f3-46eb-afc5-0be35a1f582f\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://www.example.com\",\"targets\":[],\"value\":\"\",\"URL\":\"https://www.example.com\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"14bbdd0e-f78f-4591-916e-4b6cf94ce576\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"38e2db3f-b835-47e8-8786-b87792d6fe4f\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"f30156f4-3a70-4031-9a41-dbe8eac7e494\"]}],\"urls\":[\"https://www.example.com/\"],\"plugins\":[]}"

    //(Required)Script type of the selenium script added
      script_type="txt"


    // (Required) Provide the Cookies for parameter forwarding in Map format.
      cookies ={
        "Accept" = "application/json"
      }

    //(Required) Map of proxy-details 
      proxy_details={webProxyUrl: "qwerty", webProxyUname: "sdcsdcsdc", webProxyPass: "sadasds"}

}