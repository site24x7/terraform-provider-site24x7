terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      // Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.4"
      # source  = "registry.terraform.io/site24x7/site24x7"
      # version = "1.0.0"
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  # // The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  # // environment variable if the attribute is empty or omitted.
  # oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  # // The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  # // environment variable if the attribute is empty or omitted.
  # oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
  data_center = "US"

  // The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

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
  # location_profile_name = "North America"
  // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response. 
  // You can specify individual status codes, as well as ranges separated with a colon.
  # up_status_codes = "400:500"

  // ================ JSON ASSERTION ATTRIBUTES
  // (Optional) Response content type. Default value is 'T'
  // 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML
  // https://www.site24x7.com/help/api/#res_content_type
  response_content_type = "J"
  // (Optional) Provide multiple JSON Path expressions to enable evaluation of JSON Path expression assertions. 
  // The assertions must successfully parse the JSON Path in the JSON. JSON expression assertions fails if the expressions does not match.
  match_json_path = [
    "$.store.book[*].author",
    "$..author",
    "$.store.*"
  ]
  // (Optional) Trigger an alert when the JSON path assertion fails during a test. 
  // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
  match_json_path_severity = 0
  // (Optional) JSON schema to be validated against the JSON response.
  json_schema = "{\"test\":\"abcd\"}"
  // (Optional) Trigger an alert when the JSON schema assertion fails during a test. 
  // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
  json_schema_severity = 2
  // (Optional) JSON Schema check allows you to annotate and validate all JSON endpoints for your web service.
  json_schema_check = true
  // JSON ASSERTION ATTRIBUTES ================
}

