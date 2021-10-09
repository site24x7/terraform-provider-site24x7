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
}

```

## Attributes Reference

### Required

* `aws_access_key` (String)
* `aws_secret_key` (String)
* `display_name` (String)

### Optional

* `aws_discover_services` (List of String)
* `aws_discovery_frequency` (Number)
* `id` (String) The ID of this resource.
* `notification_profile_id` (String)
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down.

Refer [API documentation](https://www.site24x7.com/help/api/#amazon-webservice-monitor) for more information about attributes.

