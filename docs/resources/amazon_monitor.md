---
layout: "site24x7"
page_title: "Site24x7: site24x7_amazon_monitor"
sidebar_current: "docs-site24x7-resource-amazon-monitor"
description: |-
  Create and manage amazon monitors in Site24x7.
---

# Resource: site24x7\_amazon\_monitor

Use this resource to create, update, and delete amazon monitors in Site24x7.

## Example Usage

```hcl
// Site24x7 Amazon Monitor API doc - https://www.site24x7.com/help/api/#amazon-webservice-monitor
resource "site24x7_amazon_monitor" "aws_monitor_site24x7" {
  // (Required) Display name for the monitor
  display_name = "aws_added_via_terraform"
  // (Required) AWS access key
  aws_access_key = ""
  // (Required) AWS secret key
  aws_secret_key = ""
  // (Optional) AWS discover frequency
  aws_discovery_frequency = 5
  // (Optional) AWS services to discover. See https://www.site24x7.com/help/api/#aws_discover_services 
  // for knowing service ID.
  aws_discover_services = ["1"]
  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]
  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]
}

```

## Attributes Reference

### Required

* `display_name` (String) Display name for the AWS monitor.
* `aws_access_key` (String) Access Key ID for the AWS account.
* `aws_secret_key` (String) Secret Access key for the AWS account.

### Optional

* `id` (String) The ID of this resource.
* `aws_discover_services` (List of String) List of AWS services that needs to be discovered. Please refer [API documentation](https://www.site24x7.com/help/api/#aws_discover_services) for knowing AWS service ID's.
* `aws_discovery_frequency` (Number) Rediscovery polling interval for the AWS account. Please refer [API documentation](https://www.site24x7.com/help/api/#aws_discover_frequency) for knowing values that can be configured.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.

Refer [API documentation](https://www.site24x7.com/help/api/#amazon-webservice-monitor) for more information about attributes.

