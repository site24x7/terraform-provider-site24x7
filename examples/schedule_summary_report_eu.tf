terraform {
  required_providers {
    site24x7 = {
      source  = "site24x7.com/site24x7/site24x7"
      version = "1.0.0"
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
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
  data_center = "EU"

  // (Optional) ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

resource "site24x7_schedule_report" "weekly_summary" {
  # Display name for the scheduled report
  display_name = "Schedule Summary Report"

  # Report type constant
  # 17 = Summary Report
  report_type = 17

  # Resource selection type
  # 0 = All Monitors
  # 2 = Specific Monitors
  # 3 = Tags
  # 4 = Monitor Type
  selection_type = 0

  # Report output format
  # 1 = PDF
  # 2 = CSV
  report_format = 2

  # Report frequency
  # 1 = Daily
  # 2 = Weekly
  # 3 = Monthly
  report_frequency = 2

  # Day of the week the report is generated
  # 1 = Sunday, 2 = Monday, ... 7 = Saturday
  scheduled_day = 3

  # Hour of the day the report is generated (0–23)
  scheduled_time = 10

  # List of user group IDs who will receive the report
  user_groups = [4111500000000]
}