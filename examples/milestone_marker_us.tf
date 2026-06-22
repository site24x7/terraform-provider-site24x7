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

// Site24x7 Milestone Marker API doc - https://www.site24x7.com/help/api/#milestone-markers
//
// Milestone markers help you record significant events like build deployments,
// product updates, feature enhancements and infrastructure upgrades.
//
// There are three types of milestone markers:
//   1. Global         - Applies to all monitors. Omit monitor_id or set it to "-1".
//   2. Monitor-level  - Associated with a specific monitor. Set monitor_id to a valid Monitor ID.
//   3. Group-level    - Associated with a monitor group. Set monitor_id to a valid Monitor Group ID.

// ============================
// Global Milestone Marker
// ============================
// A global milestone marker applies to ALL monitors in your account.
// Simply omit the monitor_id attribute (defaults to "-1").
resource "site24x7_milestone_marker" "global_release" {
  // (Required) Timestamp of milestone creation.
  marker_time = "2026-06-04T14:45:00+0530"

  // (Required) Milestone label.
  label = "v2.0 Global Release"

  // (Optional) Milestone description.
  message = "Global milestone marker for v2.0 release across all monitors"

  // monitor_id is omitted - this creates a GLOBAL milestone marker.
}

// ====================================
// Monitor-specific Milestone Marker
// ====================================
// Associate the milestone with a specific monitor by providing its Monitor ID.
resource "site24x7_milestone_marker" "monitor_deployment" {
  // (Optional) Monitor ID to associate this milestone with.
  // Provide a valid Monitor ID for a monitor-specific milestone.
  monitor_id = "15698000017614001"

  // (Required) Timestamp of milestone creation.
  marker_time = "2026-06-03T11:00:00+0530"

  // (Required) Milestone label.
  label = "Build version 9.35"

  // (Optional) Milestone description.
  message = "Deployed build version 9.35 to production"
}

// ====================================
// Monitor Group Milestone Marker
// ====================================
// Associate the milestone with a monitor group by providing its Group ID.
resource "site24x7_milestone_marker" "group_upgrade" {
  // (Optional) Monitor Group ID to associate this milestone with.
  // Provide a valid Monitor Group ID for a group-level milestone.
  monitor_id = "15698000067539089"

  // (Required) Timestamp of milestone creation.
  marker_time = "2026-06-03T14:30:00+0530"

  // (Required) Milestone label.
  label = "Infrastructure Upgrade"

  // (Optional) Milestone description.
  message = "Upgraded all servers in the group to latest OS version"
}
