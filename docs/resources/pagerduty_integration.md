---
layout: "site24x7"
page_title: "Site24x7: site24x7_pagerduty_integration"
sidebar_current: "docs-site24x7-pagerduty-integration"
description: |-
Create and manage a PagerDuty integration in Site24x7.
---

# Resource: site24x7\_pagerduty\_integration

Use this resource to create, update, and delete PagerDuty integration in Site24x7.

## Example Usage

```hcl
// PagerDuty API doc: https://www.site24x7.com/help/api/#create-pagerduty
resource "site24x7_pagerduty_integration" "pagerduty_integration_example" {
  // (Required) Display name for the PagerDuty Integration.
  name = "PageDuty Integration - Terraform"
  // (Required) Unique integration key provided by PagerDuty to facilitate incident creation in PagerDuty.
  service_key = "service_key"
  // (Required) Name of the service who posted the incident.
  sender_name = "Site24x7 - Terraform"
  // (Required) Title of the incident.
  title = "$MONITORNAME is $STATUS"
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) Configuration to resolve the incidents manually when the monitor changes to UP status.
  manual_resolve = false
  // (Optional) Configuration to send custom parameters while executing the action.
  send_custom_parameters = true
  // (Optional) Mandatory, if send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL.
  custom_parameters               = "{\"test\":\"abcd\"}"
  // (Optional) Monitors to be associated with the integration when the selection_type = 2
  monitors                        = ["756"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3
  tags                        = ["345"]
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id                   = ["123"]
}
```

## Attributes Reference


### Required

* `name` (String) Display name for the PagerDuty Integration.
* `service_key` (String) Unique integration key provided by PagerDuty to facilitate incident creation in PagerDuty.
* `sender_name` (String) Name of the service who posted the message.
* `title` (String) Title of the incident.


### Optional

* `id` (String) The ID of this resource.
* `selection_type` (Number) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `trouble_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.  Default value is 'true'.
* `critical_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `down_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `manual_resolve` (Boolean) Configuration to resolve the incidents manually when the monitor changes to UP status.
* `send_custom_parameters` (Boolean) Configuration to send custom parameters while executing the action.
* `custom_parameters` (String) Mandatory when `send_custom_parameters` is set as true. Custom parameters to be passed while accessing the URL.
* `monitors` (List of String) Monitors to be associated with the integration when the selection_type = 2.
* `tags` (List of String) Tags to be associated with the integration when the selection_type = 3.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-pagerduty) for more information about attributes.


