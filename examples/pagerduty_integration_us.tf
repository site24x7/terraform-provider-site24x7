terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "0.0.1-beta.14"
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

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

// PagerDuty API doc: https://www.site24x7.com/help/api/#create-pagerduty
resource "site24x7_pagerduty_integration" "pagerduty_integration_basic" {
  // (Required) Display name for the PagerDuty Integration.
  name = "PageDuty Integration - Terraform"
  // (Required) Unique integration key provided by PagerDuty to facilitate incident creation in PagerDuty.
  service_key = "service_key"
  // (Required) Name of the service who posted the incident.
  sender_name = "Site24x7 - Terraform"
  // (Required) Title of the incident.
  title = "$MONITORNAME is $STATUS"
}

// PagerDuty API doc: https://www.site24x7.com/help/api/#create-pagerduty
resource "site24x7_pagerduty_integration" "pagerduty_integration" {
  // (Required) Display name for the PagerDuty Integration.
  name = "PageDuty Integration - Terraform"
  // (Required) Unique integration key provided by PagerDuty to facilitate incident creation in PagerDuty.
  service_key = "service_key"
  // (Required) Name of the service who posted the incident.
  sender_name = "Site24x7 - Terraform"
  // (Required) Title of the incident.
  title = "$MONITORNAME is $STATUS"
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) Configuration to resolve the incidents manually when the monitor changes to UP status.
  manual_resolve = false
  // (Optional) Configuration to send custom parameters while executing the action.
  send_custom_parameters = true
  // (Optional) Mandatory, if send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL.
  custom_parameters               = "{\"test\":\"abcd\"}"
  // (Optional) Monitors to be associated with the integration when the selection_type = 2.
  monitors                        = ["756"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3.
  tags                        = ["345"]
  // (Optional) List of tag IDs to be associated with the integration.
  alert_tags_id                   = ["123"]
}