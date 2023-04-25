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
  oauth2_client_id = "1000.D5HV4RBPJ7L0IM7KG3JS84EO9N1T6W"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "dcaecd9e2036c076f783a3ad3515be41b05ad8480a"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "1000.d270d3da512b1ce9587798e04b142c93.9985bc636390932c7b7fc7ab2f7efd82"

  // (Optional) The access token will be looked up in the SITE24X7_OAUTH2_ACCESS_TOKEN
  // environment variable if the attribute is empty or omitted. You need not configure oauth2_access_token
  // when oauth2_refresh_token is set.
  #  oauth2_access_token = "1000.0d54698c58bb0d2854e12a9663f0f545.41e3b975a264c02823823aea85adb69a"

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

// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Website Monitor - Terraform"

  // (Required) Website address to monitor.
  website = "https://www.example2.com"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = "5"

  // (Optional) Name of the Location Profile that has to be associated with the monitor.
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
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

  // (Optional) Value to be checked against resolved values. Choose parameters based on your configured lookup type. See https://www.site24x7.com/help/api/#dns_search_config
  search_config {
    lookup_type = "AAAA"
    addr        = "1.2.3.4"
    ttlo        = "3"
    ttl         = "60"
  }

  // (Optional) Value to be checked against resolved values. Choose parameters based on your configured lookup type. See https://www.site24x7.com/help/api/#dns_search_config
  search_config {
    lookup_type = "AAAA"
    addr        = "1.2.3.4"
    ttlo        = "1"
    ttl         = "60"
  }

  // (Optional) Lookup Type - See https://www.site24x7.com/help/api/#dns_lookup_type
  // Lookup Types supported: 1 - A, 255 - ALL, 28 - AAAA, 2 - NS, 15 - MX, 5 - CNAME, 6 - SOA, 12 - PTR, 33 - SRV, 16 - TXT, 48 - DNSKEY, 257 - CAA, 43 - DS
  lookup_type               = 1

  // (Optional) Pass dnssec parameter to enable Site24x7 to validate DNS responses. See https://www.site24x7.com/help/internet-service-metrics/dns-monitor.html#dnssec
  dnssec                    = false

  // (Optional) Enable this attribute to auto discover and set up monitoring for all the related resources for the domain_name.
  deep_discovery            = false
}



// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_rest_api_transaction_monitor" "rest_api_transaction_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "RestAPI Transaction Monitor"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = "5"

  // (Optional) Name of the Location Profile that has to be associated with the monitor.
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  steps {
    display_name = "Location changes done"
    step_details {
      step_url = "https://www.example1.com"
      // (Optional) Map of HTTP response headers to check.
      response_headers_severity = 0 // Can take values 0 or 2. '0' denotes Down and '2' denotes Trouble.
      response_headers = {
        "Content-Encoding" = "gzip"
        "Connection" = "Keep-Alive"
      }
    }
  }

  steps {
    display_name = "RestAPI Transaction Monitor"
    step_details {
      step_url = "https://www.example1.com"
      // (Optional) Map of HTTP response headers to check.
      response_headers_severity = 0 // Can take values 0 or 2. '0' denotes Down and '2' denotes Trouble.
      response_headers = {
        "Content-Encoding" = "gzip"
        "Connection" = "Keep-Alive"
      }
    }
  }
}

// Site24x7 Rest API Monitor API doc - https://www.site24x7.com/help/api/#rest-api
resource "site24x7_rest_api_monitor" "rest_api_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "REST API Monitor - terraform"
  // (Required) Website address to monitor.
  website = "https://dummy.restapiexample.com/"
  // (Optional) Name of the Location Profile that has to be associated with the monitor.
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = "5"

}
