terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.18"
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

  // ZAAID of the customer under a MSP or BU
  zaaid = "1234"

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

// Site24x7 notification profile doc - https://www.site24x7.com/help/api/#notification-profiles
resource "site24x7_notification_profile" "notification_profile_us" {
  // (Required) Display name for the notification profile.
  profile_name = "Notification Profile - Terraform"
}

// Site24x7 notification profile API doc - https://www.site24x7.com/help/api/#notification-profiles
resource "site24x7_notification_profile" "notification_profile_all_attributes_us" {
  // (Required) Display name for the notification profile.
  profile_name = "Notification Profile All Attributes - Terraform"

  // (Optional) Settings to send root cause analysis when monitor goes down. Default is true.
  rca_needed= true

  // (Optional) Settings to downtime only after executing configured monitor actions.
  notify_after_executing_actions = true

  // (Optional) Configuration for delayed notification. Default value is 1. Can take values 1, 2, 3, 4, 5.
  downtime_notification_delay = 2

  // (Optional) Settings to receive persistent notification after number of errors. Can take values 1, 2, 3, 4, 5.
  persistent_notification = 1

  // (Optional) User group ID for downtime escalation.
  escalation_user_group_id = "123456000000025005"

  // (Optional) Duration of Downtime before Escalation. Mandatory if any user group is added for escalation.
  escalation_wait_time = 30

  // (Optional) Email template ID for notification
  template_id = 123456000024578001

  // (Optional) Settings to stop an automation being executed on the dependent monitors.
  suppress_automation = true

  // (Optional) Execute configured IT automations during an escalation.
  escalation_automations = [
    "123456000000047001"
  ]

  // (Optional) Invoke and manage escalations in your preferred third party services.
  escalation_services = [
    "123456000008777001"
  ]

}