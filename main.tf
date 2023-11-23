terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      # source = "site24x7/site24x7"
      // Uncomment for local build
      source  = "registry.terraform.io/site24x7/site24x7"
      version = "1.0.0"
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_CLIENT_SECRET>"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_REFRESH_TOKEN>"

  // (Optional) The access token will be looked up in the SITE24X7_OAUTH2_ACCESS_TOKEN
  // environment variable if the attribute is empty or omitted. You need not configure oauth2_access_token
  // when oauth2_refresh_token is set.
  # oauth2_access_token = "<SITE24X7_OAUTH2_ACCESS_TOKEN>"

  // (Optional) oauth2_access_token expiry in seconds. Specify access_token_expiry when oauth2_access_token is configured.
  # access_token_expiry = "0"

  // (Optional) ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
  data_center = "US"

  // (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4
}



// DomainExpiry Server API doc: https://www.site24x7.com/help/api/#domain-expiry
resource "site24x7_domain_expiry_monitor" "domain_expiry_monitor_basic3" {
  // (Required) Display name for the monitor
  display_name = "ignore registry 1"
  host_name = "file.com"
  domain_name="whios.iana.org"
  port=443
  timeout = 50
  expire_days=40
  use_ipv6=true
  //matching_keyword={2="aaa"}
  location_profile_name = "North America"
  ignore_registry_date = false
  type = "DOMAINEXPIRY"
  actions={0=456418000001238121}
   on_call_schedule_id = "456418000001258016"
  
}
// DomainExpiry Server API doc: https://www.site24x7.com/help/api/#domain-expiry
resource "site24x7_domain_expiry_monitor" "domain_expiry_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "ignore registry 2"
  host_name = "file.com"
  domain_name="whios.iana.org"
  port=443
  timeout = 50
  expire_days=40
  use_ipv6=false
  match_case=false 
  matching_keyword = {
 	  severity= 2
 	  value= "aaa"
 	}
  unmatching_keyword = {
 	  severity= 0
 	  value= "bbb"
 	}
  match_regex = {
 	  severity= 0
 	  value= "test(.*)\\d"
 	}
  perform_automation=true
  location_profile_name = "North America"
  ignore_registry_date = false
  type = "DOMAINEXPIRY"
  actions={1=456418000001238121}
  notification_profile_name="System Generated"
  
}


