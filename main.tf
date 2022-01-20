terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      // Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "0.0.1-beta.14"
      // Uncomment for local setup
      # source  = "registry.zoho.io/zoho/site24x7"
      # version = "1.0.0"
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

// Site24x7 Tag API doc - https://www.site24x7.com/help/api/#tags
# resource "site24x7_tag" "tag_us" {
#   // (Required) Display Name for the Tag.
#   tag_name = "Website Tag - Terraform"

#   // (Required) Value for the Tag.
#   tag_value = "Zoho Domains - Terraform"

#   // Color code for the Tag. Possible values are '#B7DA9E','#73C7A3','#B5DCDF','#D4ABBB','#4895A8','#DFE897','#FCEA8B','#FFC36D','#F79953','#F16B3C','#E55445','#F2E2B6','#DEC57B','#CBBD80','#AAB3D4','#7085BA','#F6BDAE','#EFAB6D','#CA765C','#999','#4A148C','#009688','#00ACC1','#0091EA','#8BC34A','#558B2F'
#   tag_color = "#B7DA9E"
# }

// Site24x7 notification profile doc - https://www.site24x7.com/help/api/#notification-profiles
# resource "site24x7_notification_profile" "notification_profile_us" {
#   // (Required) Display name for the notification profile.
#   profile_name = "Notification Profile - Terraform"
# }

// Destroy command --> terraform destroy -target site24x7_website_monitor.website_monitor_example

// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Example Monitor"

  // (Required) Website address to monitor.
  website = "https://www.example.com"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = 1

  // (Optional) Name of the location profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  # notification_profile_name = "Terraform"
  # notification_profile_id="123456000000029001" // Default Notification
  # notification_profile_id="123456000024606003" // Terraform Profile

  // (Optional) List if user group names to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]

  // (Optional) List if user group IDs to be notified on down. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  # user_group_ids = [
  #   "123456000000025005", // Admin
  #   "123456000000025007", // Network
  # ]

  // (Optional) List of tag names to be associated to the monitor. Tag name matching works for both exact and 
  //  partial match. Either specify tag_ids or tag_names.
  # tag_names = [
  #   "Terraform",
  #   "Network",
  # ]

  // (Optional) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
  # tag_ids = [
  #   # "123456000024829001", // Terraform
  #   # "123456000024829005", // Network
  #   "123456000024829007", // Server
  #   "123456000024829003", // Website
  # ]
}
