---
layout: "site24x7"
page_title: "Site24x7: site24x7_gcp_monitor"
sidebar_current: "docs-site24x7-resource-gcp-monitor"
description: |-
  Create and manage Google Cloud Platform monitors in Site24x7.
---
# Resource: site24x7\_gcp\_monitor

Use this resource to create, update and delete gcp monitors in Site24x7.

## Example Usage

```hcl

terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
    
  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // (Required) Specify the data center (US/EU/IN/AU/CN/JP/CA).
  data_center = "US"
}

# Require Google provider
provider "google" {
  project = "GOOGLE PROJECT ID"
  region  = "GOOGLE DEFUALT REGION"
}

# Resource to define a GCP IAM Service Account
resource "google_service_account" "site24x7_sa" {
  account_id   = "site24x7-monitor"
  display_name = "Site24x7 Monitoring Service Account"
}

# IAM role binding for the service account
resource "google_project_iam_member" "site24x7_monitor_role" {
  project = "GOOGLE PROJECT ID"
  role    = "roles/viewer"
  member  = "serviceAccount:${google_service_account.site24x7_sa.email}"
}

resource "google_service_account_key" "site24x7_sa" {
  service_account_id = google_service_account.site24x7_sa.name
}


// Site24x7 GCP Monitor API doc - https://www.site24x7.com/help/api/#google-cloud-platform-monitor
resource "site24x7_gcp_monitor" "gcp_monitor_site24x7" {
  // (Required) Display name for the monitor
  display_name = "gcp_monitor_via_terraform"
  
  // (Required) GCP Project ID
  project_id = jsondecode(base64decode(google_service_account_key.site24x7_sa.private_key))["project_id"]
  
  // (Required) GCP Service Account Email and Private JSON Key
  client_email = jsondecode(base64decode(google_service_account_key.site24x7_sa.private_key))["client_email"]
  private_key = jsondecode(base64decode(google_service_account_key.site24x7_sa.private_key))["private_key"] 


  
  // (Optional) GCP services to monitor. See API documentation for service IDs.
  gcp_discover_services = [1,2,3,4,5]
  
  // (Optional) GCP Discovery Frequency in Mins
  gcp_discovery_frequency   = "30"

 // (Optional) Auto discover new resources 1 for Enable and 0 for disable
  stop_rediscover_option    = 1

  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  // (Optional) List if user group IDs to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_ids = [
    "123",
  ]

  // (Optional) List if user group names to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]
  
  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]
}

## Attributes Reference

### Required

* `display_name` (String) Display name for the GCP monitor.

* `project_id` (String) The GCP Project ID.

* `client_email` (String) The client email for authentication.

* `private_key` (String) The private key for authentication.

Optional

* `gcp_discover_services` (List of Integer) List of GCP services to be discovered.

* `gcp_discovery_frequency` (Number) Rediscovery polling interval in minutes.

* `stop_rediscover_option (Number) Option to auto-discover new resources (1 to enable, 0 to disable).

* `notification_profile_name` (String) Name of the notification profile to associate with the monitor.

* `user_group_ids` (List of String) List of user group IDs to be notified when the monitor is down.

* `user_group_names` (List of String) List of user group names to be notified when the monitor is down.

* `tag_ids` (List of String) List of tag IDs to be associated with the monitor.

Output

* `id` (String) The ID of this resource.