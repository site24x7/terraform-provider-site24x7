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

// DNS Server API doc: https://www.site24x7.com/help/api/#dns-server
resource "site24x7_dns_server_monitor" "dns_monitor_basic" {
  // (Required) Name for the monitor.
  display_name              = "Nowatt basic DNS monitor - Terraform"
  
  // (Required) DNS Name Server to be monitored
  dns_host                  = "185.43.51.84"
  
  // (Required) Domain name to be resolved.
  domain_name               = "www.nowatt.com"
}

// DNS Server API doc: https://www.site24x7.com/help/api/#dns-server
resource "site24x7_dns_server_monitor" "dns_server_monitor" {
  // (Required) Display Name for the monitor.
  display_name              = "Nowatt DNS monitor - Terraform"
  
  // (Required) DNS Name Server to be monitored
  dns_host                  = "185.43.51.84"

  // (Required) Domain name to be resolved.
  domain_name               = "www.nowatt.com"
  
  // (Optional) Port for DNS access. Default value: 53
  dns_port                  = "53"

  // (Optional)  Interval at which your DNS server has to be monitored. Default value is 5 minutes.
  check_frequency           = "5"

  // (Optional)  Timeout for connecting to your DNS server. Default value is 10. Range 1 - 45.
  timeout                   = 10

  // (Optional) Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.
  use_ipv6                  = false

  // (Optional) Value to be checked against resolved values. Choose parameters based on your configured lookup type. See https://www.site24x7.com/help/api/#dns_search_config
  search_config {
    lookup_type = "A"
    addr        = "1.2.3.4"
    ttlo        = "2"
    ttl         = "60"
  }

  // (Optional) Lookup Type - See https://www.site24x7.com/help/api/#dns_lookup_type
  // Lookup Types supported: 1 - A, 255 - ALL, 28 - AAAA, 2 - NS, 15 - MX, 5 - CNAME, 6 - SOA, 12 - PTR, 33 - SRV, 16 - TXT, 48 - DNSKEY, 257 - CAA, 43 - DS
  lookup_type               = 1

  // (Optional) Pass dnssec parameter to enable Site24x7 to validate DNS responses. See https://www.site24x7.com/help/internet-service-metrics/dns-monitor.html#dnssec
  dnssec                    = false

  // (Optional) Enable this attribute to auto discover and set up monitoring for all the related resources for the domain_name.
  deep_discovery            = false

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

  // (Optional) Threshold profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/threshold_profiles
  // endpoint for the DNS server monitor (https://www.site24x7.com/help/api/#list-threshold-profiles) will
  // be used.
  threshold_profile_id = "123"

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

  // (Optional) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
  actions = {
    "1" = "123"
  }
}

