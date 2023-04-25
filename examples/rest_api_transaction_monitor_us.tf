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
  # oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  # oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // ZAAID of the customer under a MSP or BU
  zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
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

// Site24x7 Rest API Monitor API doc - https://www.site24x7.com/help/api/#rest-api-transaction
resource "site24x7_rest_api_transaction_monitor" "rest_api_transaction_monitor_basic" {
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
  location_profile_name = "North America"
}

// Site24x7 Rest API Transaction Monitor API doc - https://www.site24x7.com/help/api/#rest-api-transaction
resource "site24x7_rest_api_transaction_monitor" "rest_api_transaction_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "RestAPI Transaction Monitor"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = "5"

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

  // (Required) List of Steps details to be associated to the steps.

  steps {

      // (Required) Display name for the step
      display_name = "RestAPI Transaction Monitor"

      // (Required)  API request details related to this step.
      step_details {

          // (Required) Domain address for the step.
          step_url = "https://www.example1.com"

          // (Optional)  Timeout for connecting to REST API Default value is 10. Range 1 - 45.
          timeout = 10

          // (Optional) Map of custom HTTP headers to send.
          custom_headers = {
            "Accept" = "application/json"
          }

          // (Optional) Check for the keyword in the website response.
          matching_keyword = {
            severity= 2
            value= "aaa"
          }

          // (Optional) Check for non existence of keyword in the website response.
          unmatching_keyword = {
            severity= 2
            value= "bbb"
          }

          // (Optional) Match the regular expression in the website response.
          match_regex = {
            severity= 2
            value= ".*aaa.*"
          }

          // (Optional) Map of HTTP response headers to check.
          response_headers_severity = 0 // Can take values 0 or 2. '0' denotes Down and '2' denotes Trouble.
          response_headers = {
            "Content-Encoding" = "gzip"
            "Connection" = "Keep-Alive"
          }

          // HTTP Configuration
          // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response.
          // You can specify individual status codes, as well as ranges separated with a colon.
          up_status_codes = "400:500"

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

          // ================ HTTP POST with request body
          // (Optional) HTTP Method to be used for accessing the website. Default value is 'G'. 'G' denotes GET, 'P' denotes POST and 'H' denotes HEAD. PUT, PATCH and DELETE are not supported.
          http_method = "P"
          // (Optional) Provide content type for request params when http_method is 'P'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML and 'F' denotes FORM
          request_content_type = "J"
          // (Optional) Provide the content to be passed in the request body while accessing the website.
          request_body = "{\"user_name\":\"joe\"}"
          // (Optional) Map of custom HTTP headers to send.
          request_headers = {
            "Accept" = "application/json"
          }
          // HTTP POST with request body ================

          // ================ GRAPHQL ATTRIBUTES
          // (Optional) Provide content type for request params.
          // request_content_type = "G"
          // (Optional) Provide the GraphQL query to get specific response from GraphQL based API service. request_content_type = "G"
          graphql_query = "query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}"
          // (Optional) Provide the GraphQL variables to get specific response from GraphQL based API service. request_content_type = "G"
          graphql_variables = "{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}"
          // GRAPHQL ATTRIBUTES ================
      }
  }
}