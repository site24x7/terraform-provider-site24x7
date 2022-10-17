---
layout: "site24x7"
page_title: "Site24x7: site24x7_connectwise_integration"
sidebar_current: "docs-site24x7-connectwise-integration"
description: |-
Create and manage a connectwise integration in Site24x7.
---

# Resource: site24x7\_connectwise\_integration

Use this resource to create, update, and delete connectwise integration in Site24x7.

## Example Usage

```hcl

// Connectwise Integration API doc: https://www.site24x7.com/help/api/#create-connectwise
resource "site24x7_connectwise_integration" "connectwise_integration" {
  // (Required) Display name for the integration
  name           = "Connectwise Integration With Site24x7"
  // (Required) Hook URL to which the message will be posted
  url            = "https://stawcdwging.connectwisedev.com/"
  // (Required) Name of the comapny for Authentication.
  company        = "zylker_c"
  // (Required) Public Key for Authentication.
  public_key 	   = "K3xKPKiP88i6rmAb"
  // (Required) Private Key for Authentication.
  private_key 	 = "Fkb45lqwhQGIxcc5"
  // (Required) Company ID for which the message will be posted.
  company_id 	   = "GreenInc"
  // (Required) Provide the configuration settings to resolve or close incidents automatically in Connectwise, when the monitor status changes to UP.
  close_status   = "Closed (resolved)"
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) Monitors to be associated with the integration when the selection_type = 2.
  monitors                        = ["756"]
  // (Optional) User groups to be associated with the integration which will be notified when there is an error in ConnectWise Manage Integration.
  user_groups                        = ["72356"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3.
  tags                        = ["345"]
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id  = ["123"]
}

```

## Attributes Reference


### Required

* `name` (String) Display name for the integration.
* `url` (String) Hook URL to which the message will be posted.
* `company` (String) Name of the service who posted the message.
* `public_key` (String) Title of the incident.
* `private_key` (String) Title of the incident.
* `company_id` (String) Title of the incident.
* `close_status` (String) Title of the incident.


### Optional

* `id` (String) The ID of this resource.
* `selection_type` (Number) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `trouble_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.  Default value is 'true'.
* `critical_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `down_alert` (Boolean) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `monitors` (List of String) Monitors to be associated with the integration when the selection_type = 2.
* `tags` (List of String) Tags to be associated with the integration when the selection_type = 3.
* `user_groups` (List of String) User groups to be associated with the integration that will be notified when there is an error in ConnectWise Manage Integration.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-connectwise) for more information about attributes.


