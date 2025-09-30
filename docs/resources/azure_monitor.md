---
layout: "site24x7"
page_title: "Site24x7: site24x7_azure_monitor"
sidebar_current: "docs-site24x7-resource-azure-monitor"
description: |-
  Create and manage Azure monitors in Site24x7.
---

# Resource: site24x7_azure_monitor

Use this resource to create, update and delete Azure monitors in Site24x7.

## Example Usage

```hcl
terraform {
  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest
    }
  }
}

provider "site24x7" {
  oauth2_client_id     = "<SITE24X7_OAUTH2_CLIENT_ID>"
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"
  data_center          = "US"
  retry_max_wait       = 30
  max_retries          = 4
}

resource "site24x7_azure_monitor" "example" {
  // (Required) Display name for the Azure monitor
  display_name = "My Azure Monitor Terraform"

  // (Required) Azure Entra ID (Tenant ID)
  tenant_id = "<AZURE_ENTRA_TENANT_ID>"

  // (Required) Application (client) ID from Azure App Registration
  client_id = "<AZURE_APP_REGISTRATION_CLIENT_ID>"

  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) Application client secret from Azure App Registration
  client_secret = "<AZURE_APP_REGISTRATION_CLIENT_SECRET>"

  // (Required) Monitor type should be "AZURE"
  type = "AZURE"

  // (Required) Azure services to discover. Format: ["<ResourceProvider>/<ResourceType>"]
  services = ["Microsoft.Compute/virtualMachines"]

  // (Required) Set to 0 for Azure Account-based discovery, 1 for Management Group-based discovery
  management_group_reg = 0

  // (Optional) Discovery interval in minutes
  discovery_interval = "30"

  // (Optional) Automatically add newly discovered Azure subscriptions
  auto_add_subscription = 1

  // (Optional) Name of the notification profile to be associated with the monitor.
  // Either specify notification_profile_id or notification_profile_name.
  notification_profile_id = "Terraform Profile"

  // (Optional) List of user group IDs to be notified on down.
  user_group_ids = ["123"]

  // (Optional) Threshold profile to be associated with the monitor.
  threshold_profile_id = "azure threshold profile"

  // (Optional) Include Azure resources with matching tags. Format: "key:value"
  include_tags = [
    "Environment:Production",
    "Team:DevOps"
  ]

  // (Optional) Exclude Azure resources with matching tags. Format: "key:value"
  exclude_tags = [
    "Environment:Test",
    "Owner:External"
  ]
}
Attributes Reference
Required
display_name (String) Display name for the Azure monitor.

tenant_id (String) Azure Entra ID (Tenant ID).

client_id (String) Application (client) ID from Azure App Registration.

client_secret (String) Client secret from Azure App Registration.

type (String) Monitor type. Must be "AZURE".

services (List of String) Azure services to discover. Format: ["<ResourceProvider>/<ResourceType>"]

management_group_reg (Number) Set to 0 for Azure Account-based discovery, 1 for Management Group-based discovery.

Optional
discovery_interval (String) Rediscovery polling interval in minutes.

auto_add_subscription (Number) Automatically add newly discovered subscriptions. Set to 1 to enable.

notification_profile_id (String) Notification profile ID to associate with the monitor. If omitted, the first profile from the /api/notification_profiles endpoint will be used.

notification_profile_name (String) Notification profile name to associate with the monitor. Supports partial or full name match.

user_group_ids (List of String) List of user group IDs to be notified on down. If omitted, the first group from the /api/user_groups endpoint will be used.

user_group_names (List of String) List of user group names to be notified on down. Supports partial or full name match.

threshold_profile_id (String) Threshold profile ID to associate with the monitor. If omitted, the first profile for the AZURE monitor type will be used.

tag_ids (List of String) List of tag IDs to be associated to the monitor. Either use tag_ids or tag_names.

tag_names (List of String) List of tag names to be associated. Supports partial or full name match. Either use tag_names or tag_ids.

third_party_service_ids (List of String) List of Third Party Service IDs to associate to the monitor.

include_tags (List of String) Include resources with these tags. Format: "key:value".

exclude_tags (List of String) Exclude resources with these tags. Format: "key:value".

Output
id (String) The ID of this resource.

Refer API documentation https://www.site24x7.com/help/api/#azure-monitor for more information about attributes.