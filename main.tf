terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      # source  = "site24x7/site24x7"
      # // Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      # version = "1.0.25"
      // Uncomment for local build
      source  = "registry.terraform.io/site24x7/site24x7"
      version = "1.0.0"
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

  # // (Optional) The access token will be looked up in the SITE24X7_OAUTH2_ACCESS_TOKEN
  # // environment variable if the attribute is empty or omitted. You need not configure oauth2_access_token
  # // when oauth2_refresh_token is set.
  # oauth2_access_token = "<SITE24X7_OAUTH2_ACCESS_TOKEN>"

	// (Optional) oauth2_access_token expiry in seconds. Specify access_token_expiry when oauth2_access_token is configured.
  # access_token_expiry = "0"

  // (Optional) ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
  data_center = "US"

  // (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Example Monitor - Terraform 1"

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

  // (Optional) List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    site24x7_monitor_group.Site24x7.id 
  ]

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

  # tag_ids = [site24x7_tag.tag_us.id]

  tag_names = [ "Server Tag" ]
}

// Subgroup API doc: https://www.site24x7.com/help/api/#subgroups
resource "site24x7_subgroup" "subgroup_default" {
  // (Required) Display Name for the Subgroup.
  display_name = "Default subgroup"
  // (Required) Unique ID of the top monitor group for which business view has been configured.
  top_group_id = "123456000033743001"
  // (Required) Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).
  parent_group_id = "123456000033743001"
}

// Subgroup API doc: https://www.site24x7.com/help/api/#subgroups
resource "site24x7_subgroup" "subgroup_zoho" {
  // (Required) Display Name for the Subgroup.
  display_name = "Zoho Subgroup"
  // (Required) Unique ID of the top monitor group for which business view has been configured.
  top_group_id = "123456000033743001"
  // (Required) Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).
  parent_group_id = "123456000033743001"
  // (Optional) Description for the Subgroup.
  description = "This is the description of the subgroup"
  // (Optional) Denotes the type of monitors that can be associated. ‘1’ implies that all type of monitors can be associated with this subgroup. Default value is 1. '2' - Web, '3' - Port/Ping, '4' - Server, '5' - Database, '6' - Synthetic Transaction, '7' - Web API, '8' - APM Insight,'9' - Network Devices, '10' - RUM, '11' - AppLogs Monitor
  group_type = 1
  // (Optional) Monitors to be associated with the Subgroup.
  monitors = [
    "123456000024411005",
  ]
  // (Optional) Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.
  health_threshold_count = 1
}

// Monitor Group API doc: https://www.site24x7.com/help/api/#monitor-groups
resource "site24x7_monitor_group" "Site24x7" {
#// (Required) Display Name for the Monitor Group.
display_name = "Site24x7 Group - Terraform"
#// (Optional) Description for the Monitor Group.
description = "This is the description of the group"
}

// Data source to fetch a monitor group
data "site24x7_monitor_group" "s247monitorgroup" {
  // (Required) Regular expression denoting the name of the monitor group.
  name_regex = "1"
  
}

// Displays the Monitor Group ID
output "s247_monitor_group_id" {
  description = "Monitor Group ID : "
  value       = data.site24x7_monitor_group.s247monitorgroup.id
}

// Displays the Monitor Group Name
output "s247_monitor_group_name" {
  description = "Monitor Group Name : "
  value       = data.site24x7_monitor_group.s247monitorgroup.display_name
}

// Displays the description
output "s247_monitor_group_description" {
  description = "Monitor Group description : "
  value       = data.site24x7_monitor_group.s247monitorgroup.description
}

// Displays the health threshold count
output "s247_monitor_group_health_threshold_count" {
  description = "Monitor Group health threshold count : "
  value       = data.site24x7_monitor_group.s247monitorgroup.health_threshold_count
}

// Displays the monitors associated
output "s247_monitor_group_monitors" {
  description = "Monitors Associated : "
  value       = data.site24x7_monitor_group.s247monitorgroup.monitors
}

// Displays the dependency resource IDs
output "s247_monitor_group_dependency_resource_ids" {
  description = "Dependency resource IDs : "
  value       = data.site24x7_monitor_group.s247monitorgroup.dependency_resource_ids
}

// Displays the suppress alert
output "s247_monitor_group_suppress_alert" {
  description = "Suppress Alert : "
  value       = data.site24x7_monitor_group.s247monitorgroup.suppress_alert
}




