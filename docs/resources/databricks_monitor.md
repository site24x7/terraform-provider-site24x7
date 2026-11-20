---
layout: "site24x7"
page_title: "Site24x7: site24x7_databricks_monitor"
sidebar_current: "docs-site24x7-resource-databricks-monitor"
description: |-
  Create and manage Databricks workspace monitors in Site24x7.
---

# Resource: site24x7_databricks_monitor

Use this resource to create, update and delete Databricks workspace monitors (monitor type `DBRICKS_WS`) in Site24x7.
Site24x7 connects to the workspace with a Databricks service principal over OAuth machine-to-machine authentication.

## Example Usage

```hcl
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

  // (Optional) Notification profile name or ID. Defaults to the first notification profile.
  notification_profile_name = "System Generated"

  // (Optional) User group names or IDs to alert. Defaults to the first user group.
  user_group_names = ["Admin"]
}
```

## Attributes Reference

### Required

* `display_name` (String) Display name for the monitor.
* `workspace_uri` (String) URL of the Databricks workspace, e.g. `https://adb-1234567890123456.7.azuredatabricks.net/`. Must start with `https://`.
* `access_key` (String) Client ID of the Databricks service principal.
* `secret_key` (String, Sensitive) OAuth secret of the Databricks service principal. Site24x7 never returns it, so it is not read back. Changing it in configuration sends the new value with `update_auth_details` set.

### Optional

* `auth_type` (Number) Authentication type. Default `4` (OAuth machine-to-machine).
* `token_profile` (String) Token profile used to obtain the access token. Default `databricks_oidc`.
* `check_frequency` (String) Polling interval in minutes. Default `"30"`.
* `monitored_services` (Set of Number) Databricks service codes to discover. Default `[5002]`.
* `disable_discovery` (Boolean) Stop auto-discovering new Databricks resources. Default `false`.
* `threshold_profile_id` (String) Threshold profile to associate with the monitor. If omitted, the account's `DBRICKS_WS` threshold profile is used, when one exists.
* `notification_profile_id` (String) Notification profile to associate with the monitor.
* `notification_profile_name` (String) Name of the notification profile to associate with the monitor. Matched as a substring. Takes precedence over `notification_profile_id`.
* `user_group_ids` (Set of String) User group IDs to notify on down.
* `user_group_names` (List of String) Names of the user groups to notify on down. Matched as a substring. Takes precedence over `user_group_ids`.
* `tag_ids` (Set of String) Tag IDs to associate with the monitor.
* `tag_names` (List of String) Names of the tags to associate with the monitor. Matched as a substring. Takes precedence over `tag_ids`.

### Read-Only

* `id` (String) The monitor ID.

## Import

Databricks monitors can be imported using the monitor ID:

```hcl
import {
  to = site24x7_databricks_monitor.example
  id = "123456000028010003"
}
```

After import, set `secret_key` in configuration. Site24x7 does not return it, so the first plan after import shows it as an update.

Refer to the [Site24x7 API documentation](https://www.site24x7.com/help/api/#monitors) for more information.
