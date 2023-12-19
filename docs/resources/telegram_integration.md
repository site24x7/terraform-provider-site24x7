---
layout: "site24x7"
page_title: "Site24x7: site24x7_telegram_integration"
sidebar_current: "docs-site24x7-telegram-integration"
description: |-
Create and manage a telegram integration in Site24x7.
---

# Resource: site24x7\_telegram\_integration

Use this resource to create, update, and delete telegram integration in Site24x7.

## Example Usage

```hcl

// Telegram Integration API doc: https://www.site24x7.com/help/api/#create-telegram
resource "site24x7_telegram_integration" "telegram_integration" {
  // (Required) Display name for the integration
  name           = "Telegram Integration With Site24x7"
  // (Required) Web URL of your telegram channel to which the message will be posted.
  channel_url    = "https://web.telegram.org/z/#-86576685"
  // (Required) Bot token created by botfather.
  token    = "3468:zFiWfmzab-u-rom-TI8"
  // (Required) Title of the incident.
  title = "Website is Down"
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) Monitors to be associated with the integration when the selection_type = 2.
  monitors      = ["756"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3.
  tags          = ["345"]
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id = ["123"]
}

```

## Attributes Reference


### Required

* `name` (String) Display name for the integration.
* `channel_url` (String) Web URL of your telegram channel to which the message will be posted.
* `token` (String) Bot token created by botfather.
* `title` (String) Title of the incident.


### Optional

* `id` (String) The ID of this resource.
* `selection_type` (Number) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `trouble_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.  Default value is 'true'.
* `critical_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `down_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `monitors` (List of String) Monitors to be associated with the integration when the selection_type = 2.
* `tags` (List of String) Tags to be associated with the integration when the selection_type = 3.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-telegram) for more information about attributes.


