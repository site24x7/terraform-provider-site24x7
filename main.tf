terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      // Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.15"
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

  // ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

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

// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Example Monitor"

  // (Required) Website address to monitor.
  website = "https://www.example.com"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = "1"

  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"
}

// Site24x7 Rest API Monitor API doc - https://www.site24x7.com/help/api/#rest-api
resource "site24x7_rest_api_monitor" "rest_api_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "REST API Monitor - terraform"
  // (Required) Website address to monitor.
  website = "https://swapi-graphql.netlify.app/.netlify/functions/index"
  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

  // (Optional) Provide content type for request params. "G" denotes GraphQL.
  request_content_type = "G"
  // (Optional) Provide the GraphQL query to get specific response from GraphQL based API service. request_content_type = "G"
  graphql_query = "query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}"
  // (Optional) Provide the GraphQL variables to get specific response from GraphQL based API service. request_content_type = "G"
  graphql_variables = "{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}"

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
