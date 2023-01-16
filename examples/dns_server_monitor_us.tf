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

// DNS Server API doc: https://www.site24x7.com/help/api/#dns-server
resource "site24x7_dns_server_monitor" "dns_server_monitor" {
  // (Required) Name for the monitor.
  display_name              = "mydnsservermonitor"
  
  // (Required) DNS Name Server to be monitored
  dns_host                  = "8.8.8.8"
  
  // Port for DNS access. Default value: 53
  dns_port                  = 53
  
  // Select IPv6 for monitoring the websites hosted with IPv6 address. If you choose non IPv6 supported locations, monitoring will happen through IPv4.
  use_ipv6                  = false

  // (Required) Domain name to be resolved.
  domain_name               = "www.example.com"

  // (Required) Check interval for monitoring.
  check_frequency           = "5"

  // (Required) Timeout for connecting to DNS Server. Range 1 - 45.
  timeout                   = 10

  // (Required) Location Profile to be associated with the monitor.
  location_profile_name     = "Global"

  // (Required) Notification profile to be associated with the monitor.
  notification_profile_name = "Default Notification"

  // (Required) Threshold profile to be associated with the monitor.
  threshold_profile_id      = "12345"

  // Required if on_call_schedule_id is not choosen User group to be notified on down.
  user_group_names          = ["Admin"]

  // Group IDs to associate monitor.
  monitor_groups            = ["12345"]

  // Lookup Type - See https://www.site24x7.com/help/api/#dns_lookup_type
  lookup_type               = 1

  // Pass dnssec parameter to enable Site24x7 to validate DNS responses. See https://www.site24x7.com/help/internet-service-metrics/dns-monitor.html#dnssec
  dnssec                    = false


  // (Optional) List if tag names to be associated to the monitor.
  tag_names = [
    "example-tag",
  ]

  // Value to be checked against resolved values. Choose parameters based on your configured lookup type. See https://www.site24x7.com/help/api/#dns_search_config
  search_config {
    lookup_type = "A"
    addr        = "1.2.3.4"
    ttlo        = "2"
    ttl         = "60"
  }

  // (Optional) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
  actions_ids = {
    "1" = "123"
  }

  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_services = [
    "4567"
  ]

    // 	Enable this attribute to auto discover and set up monitoring for all the related resources for the domain_name.
  deep_discovery            = false
}
