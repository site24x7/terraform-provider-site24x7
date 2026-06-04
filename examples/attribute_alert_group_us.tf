terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
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
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
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

// Site24x7 Attribute Alert Group API doc - https://www.site24x7.com/help/api/#attribute-alert-group

// ====================================
// DevOps Alert Attributes Group
// ====================================
resource "site24x7_attribute_alert_group" "devops_alerts" {
  // (Required) Display name for the attribute alert group.
  display_name = "devOps Alert Attributes"

  // (Required) List of attribute IDs to be associated to the group.
  // Refer attribute details API for valid attribute IDs.
  // Common attributes:
  //   1 - Availability
  //   2 - Website content change percentage
  //   3 - (refer attribute_details API)
  //  26 - CPU Utilization threshold
  //  27 - Memory Utilization threshold
  attribute_list = [1, 3]
}

// ====================================
// Server Monitoring Attributes Group
// ====================================
resource "site24x7_attribute_alert_group" "server_alerts" {
  display_name   = "Server Alert Attributes"
  attribute_list = [26, 27]
}

// ========================
// Data Source - Read Only
// ========================
data "site24x7_attribute_alert_group" "existing" {
  group_id = site24x7_attribute_alert_group.devops_alerts.id
}

// Output the group details
output "alert_group_name" {
  value = data.site24x7_attribute_alert_group.existing.display_name
}

output "alert_group_attributes" {
  value = data.site24x7_attribute_alert_group.existing.attribute_list
}
