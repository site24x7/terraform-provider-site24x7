terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.28"
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
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
	data_center = "US"
	
	// (Optional) ZAAID of the customer under a MSP or BU
	zaaid = "1234"
  
	// (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
	retry_min_wait = 1
  
	// (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
	// requests. This is the upper limit for the wait duration with exponential
	// backoff.
	retry_max_wait = 30
  
	// (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
	max_retries = 4
  
  }

// Subgroup API doc: https://www.site24x7.com/help/api/#subgroups
resource "site24x7_subgroup" "subgroup_default" {
  // (Required) Display Name for the Subgroup.
  display_name = "Default subgroup - Terraform"
  // (Required) Unique ID of the top monitor group for which business view has been configured.
  top_group_id = "123456000033743001"
  // (Required) Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).
  parent_group_id = "123456000033743001"
}

// Subgroup API doc: https://www.site24x7.com/help/api/#subgroups
resource "site24x7_subgroup" "subgroup_zoho" {
  // (Required) Display Name for the Subgroup.
  display_name = "Zoho Subgroup - Terraform"
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