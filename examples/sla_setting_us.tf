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

// Site24x7 SLA Settings API doc - https://www.site24x7.com/help/api/#sla-settings

// ============================
// Composite SLA with monitors
// ============================
resource "site24x7_sla_setting" "availability_sla" {
  // (Required) Display name for the SLA.
  display_name = "Monitors Availability"

  // (Required) Type of SLA: 1 - Availability, 2 - Response time, 3 - Composite.
  type = 1

  // (Optional) Description for the SLA.
  description = "99.9% Availability SLA."

  // (Required) Resource type for the SLA Report.
  // Associate multiple monitors or all monitors in a group.
  selection_type = 2

  // (Required if selection_type is Monitor IDs) Monitor IDs to associate.
  monitors = [
    "113770000039133011",
    "113770000008080001"
  ]

  // (Optional) Business hours during which outage details reports are generated.
  business_hours_id = "113770000010999123"

  // (Required) SLA targets to be achieved.
  sla_targets {
    target_name      = "Success"
    target_color     = "#33CC00"
    target_condition = 0
    target_value     = 99
  }

  sla_targets {
    target_name      = "Poor"
    target_color     = "#0045cc"
    target_condition = 3
    target_value     = 90
  }

  sla_targets {
    target_name      = "Medium"
    target_color     = "#33CC00"
    target_condition = 4
    target_value     = 90
  }

  // (Optional) SLA target for monitor availability (for composite SLA type).
  slo_availability {
    availability = 99
    condition    = 0
    weightage    = 50
  }

  // (Optional) SLA target for monitor response time (for composite SLA type).
  slo_responsetime {
    responsetime   = 100
    time_available = 99
    condition      = 3
    weightage      = 50
  }
}

// ====================================
// SLA with Monitor Groups
// ====================================
resource "site24x7_sla_setting" "group_sla" {
  display_name   = "Monitor Group SLA"
  type           = 1
  description    = "SLA for all monitors in a group."
  selection_type = 3

  // (Required if selection_type is Monitor Groups) Monitor Group IDs.
  monitor_groups = [
    "113770000012345001"
  ]

  sla_targets {
    target_name      = "Good"
    target_color     = "#33CC00"
    target_condition = 0
    target_value     = 99.5
  }

  sla_targets {
    target_name      = "Acceptable"
    target_color     = "#FF9900"
    target_condition = 3
    target_value     = 95
  }
}

// ========================
// Data Source - Read Only
// ========================
data "site24x7_sla_setting" "existing" {
  sla_id = site24x7_sla_setting.availability_sla.id
}

// Output the SLA details
output "sla_display_name" {
  value = data.site24x7_sla_setting.existing.display_name
}

output "sla_type" {
  value = data.site24x7_sla_setting.existing.type
}
