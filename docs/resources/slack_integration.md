---
layout: "site24x7"
page_title: "Site24x7: site24x7_slack_integration"
sidebar_current: "docs-site24x7-slack-integration"
description: |-
Create and manage a slack integration in Site24x7.
---

# Resource: site24x7\_slack\_integration

Use this resource to create, update, and delete slack integration in Site24x7.

## Example Usage

```hcl
// Slack Integration API doc: https://www.site24x7.com/help/api/#create-slack
resource "site24x7_slack_integration" "slack_integration" {
  // (Required) Display name for the integration
  name           = "Slack Integration With Site24x7"
  // (Required) Hook URL to which the message will be posted
  url            = "https://hooks.slack.com/services/T03NSM5L0/B9XER11N0/7Vk7I5n3C3ac5JnT3J4euf6"
  // (Optional) Monitors associated with the integration
  monitors       = ["756"]
  // (Required) Resource Type associated with this integration
  // https://www.site24x7.com/help/api/#resource_type_constants
  // Monitor Group not supported
  selection_type = 2
  // (Required) Name of the service who posted the message
  sender_name    = "Site24x7"
  // (Required) Title of the incident
  title          = "$MONITORNAME is $STATUS"
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id  = ["123"]
}
```

## Attributes Reference


### Required

* `name` (String) Display name for the integration.
* `url` (String) Hook URL to which the message will be posted.
* `selection_type` (Number) Resource Type associated with this integration.Monitor Group not supported.Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `sender_name` (String) Name of the service who posted the message.
* `title` (String) Title of the incident.


### Optional

* `id` (String) The ID of this resource.
* `monitors` (List of String) Monitors associated with the integration.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-slack) for more information about attributes.


