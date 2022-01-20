---
layout: "site24x7"
page_title: "Site24x7: site24x7_opsgenie_integration"
sidebar_current: "docs-site24x7-opsgenie-integration"
description: |-
Create and manage an opsgenie integration in Site24x7.
---

# Resource: site24x7\_opsgenie\_integration

Use this resource to create, update, and delete opsgenie integration in Site24x7.

## Example Usage

```hcl

// Opsgenie Integration API doc: https://www.site24x7.com/help/api/#create-opsgenie
resource "site24x7_opsgenie_integration" "opsgenie_integration" {
  // (Required) Display name for the integration
  name           = "OpsGenie Integration With Site24x7"
  // (Required) URL to be invoked for action execution
  url            = "https://api.opsgenie.com/v1/json/site24x7?apiKey=a19y1cdd-bz7a-455a-z4b1-c1528323502s"
  // (Required) Resource Type associated with this integration
  // https://www.site24x7.com/help/api/#resource_type_constants
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) Configuration to resolve the incidents manually when the monitor changes to UP status.
  manual_resolve = false
  // (Optional) Configuration to send custom parameters while executing the action.
  send_custom_parameters = true
  // (Optional) Mandatory, if send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL.
  custom_parameters               = "{\"test\":\"abcd\"}"
  // (Optional) Monitors to be associated with the integration when the selection_type = 2.
  monitors                        = ["756"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3.
  tags                        = ["345"]
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id  = ["123"]
}

```

## Attributes Reference


### Required

* `name` (String) Display name for the integration.
* `url` (String) URL to be invoked for action execution.

### Optional

* `id` (String) The ID of this resource.
* `selection_type` (Number) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `trouble_alert` (Boolean) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.  Default value is 'true'.
* `critical_alert` (Boolean) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `down_alert` (Boolean) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `manual_resolve` (Boolean) Configuration to resolve the incidents manually when the monitor changes to UP status.
* `send_custom_parameters` (Boolean) Configuration to send custom parameters while executing the action.
* `custom_parameters` (String) Mandatory when `send_custom_parameters` is set as true. Custom parameters to be passed while accessing the URL.
* `monitors` (List of String) Monitors to be associated with the integration when the selection_type = 2.
* `tags` (List of String) Tags to be associated with the integration when the selection_type = 3.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-opsgenie) for more information about attributes.