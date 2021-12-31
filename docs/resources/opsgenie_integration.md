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
  // Monitor Group not supported
  selection_type = 2
  // (Optional) Monitors associated with the integration
  monitors       = ["4567"]
  // (Optional) Configuration to create an incident during a TROUBLE alert
  trouble_alert  = false
  // (Optional) Configuration to resolve the incidents manually when the monitor changes to UP status
  manual_resolve = true
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id  = ["123"]
}
```

## Attributes Reference


### Required

* `name` (String) Display name for the integration.
* `url` (String) URL to be invoked for action execution.
* `selection_type` (Number) Resource Type associated with this integration.Monitor Group not supported.Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).



### Optional

* `id` (String) The ID of this resource.
* `trouble_alert` (Boolean) Configuration to create an incident during a TROUBLE alert.
* `manual_resolve` (Boolean) Configuration to resolve the incidents manually when the monitor changes to UP status.
* `monitors` (List of String) Monitors associated with the integration.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-opsgenie) for more information about attributes.