terraform {
  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest
    }
  }
}

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

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token.
  data_center = "US"

  // (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_max_wait = 30

  // (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4
}

variable "databricks_client_secret" {
  type      = string
  sensitive = true
}

resource "site24x7_databricks_monitor" "example" {
  // (Required) Display name for the monitor
  display_name = "Databricks Workspace - Production"

  // (Required) URL of the Databricks workspace
  workspace_uri = "https://adb-1234567890123456.7.azuredatabricks.net/"

  // (Required) Client ID of the Databricks service principal
  access_key = "<DATABRICKS_SERVICE_PRINCIPAL_CLIENT_ID>"

  // (Required) OAuth secret of the service principal. Write-only.
  secret_key = var.databricks_client_secret

  // (Optional) Polling interval in minutes. Default "30".
  check_frequency = "30"

  // (Optional) Databricks service codes to discover. Default [5002].
  // monitored_services = [5002]

  // (Optional) Stop auto-discovering new resources. Default false.
  // disable_discovery = false

  // (Optional) Threshold profile ID. Defaults to the account's DBRICKS_WS threshold profile.
  // threshold_profile_id = "123456000023902001"

  // (Optional) Notification profile name or ID. Defaults to the first notification profile.
  notification_profile_name = "System Generated"

  // (Optional) User group names or IDs to alert. Defaults to the first user group.
  user_group_names = ["Admin"]

  // (Optional) Tag names or IDs to associate with the monitor.
  // tag_names = ["prod"]
}
