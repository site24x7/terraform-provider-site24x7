terraform {
  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      
    }
  }
}

provider "site24x7" {
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
	// environment variable if the attribute is empty or omitted.
	oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
	// environment variable if the attribute is empty or omitted.
	oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
    
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
	// environment variable if the attribute is empty or omitted.
	oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"
    
	// (Required) Specify the data center from which you have obtained your
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
	data_center = "US"
  
	// (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
	// requests. This is the upper limit for the wait duration with exponential
	// backoff.
	retry_max_wait = 30
  
	// (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
	max_retries = 4
  
  
}

resource "site24x7_azure_monitor" "example" {
  // (Required) Display name for the Azure monitor
  display_name            = "My Azure Monitor Terraform"

  // (Required) Azure Entra ID (Tenant ID)
  tenant_id               = "<AZURE_ENTRA_TENANT_ID>"

  // (Required) Application (client) ID from Azure App Registration
  client_id               = "<AZURE_APP_REGISTRATION_CLIENT_ID>"

  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) Application client secret from Azure App Registration
  client_secret           = "<AZURE_APP_REGISTRATION_CLIENT_SECRET>"

  // (Required) Monitor type should be "AZURE"
  type                    = "AZURE"

  // (Required) Azure services to discover. Format: ["<ResourceProvider>/<ResourceType>"]
  // Example: ["Microsoft.Compute/virtualMachines"]
  services                = ["Microsoft.Compute/virtualMachines"] # Modify with actual service type IDs you want to discover

  // (Required) Set to 0 for Azure Account-based discovery, 1 for Management Group-based discovery
  management_group_reg    = 0        # Use 0 for Azure Account, 1 for Management Group

  // (Optional) AZURE discover interval in minutes
  discovery_interval      = "30"     # Optional: Discovery interval in minutes

  // (Optional) Automatically add newly discovered Azure subscriptions
  auto_add_subscription   = 1        # Optionally set to 1 to auto-add subscriptions

  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_id = "Terraform Profile"

  // (Optional) List if user group IDs to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_ids          = ["123"]

  // (Optional) Threshold profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/threshold_profiles
  // endpoint for the AZURE monitor type (https://www.site24x7.com/help/api/#list-threshold-profiles) will
  // be used.
  threshold_profile_id    = "azure threshold profile"

  // (Optional) Include Azure resources with matching tags.
  // Format: "key:value"
  include_tags = [
    "Environment:Production",
    "Team:DevOps"
  ]

  // (Optional) Exclude Azure resources with matching tags.
  // Format: "key:value"
  exclude_tags = [
    "Environment:Test",
    "Owner:External"
  ]


}