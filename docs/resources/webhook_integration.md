---
layout: "site24x7"
page_title: "Site24x7: site24x7_webhook_integration"
sidebar_current: "docs-site24x7-webhook-integration"
description: |-
Create and manage a webhook integration in Site24x7.
---

# Resource: site24x7\_webhook\_integration

Use this resource to create, update, and delete webhook integration in Site24x7.

## Example Usage

```hcl

// Webhook API doc: https://www.site24x7.com/help/api/#create-webhook
resource "site24x7_webhook_integration" "webhook_integration" {
  // (Required) Display name for the Webhook
  name                            = "Test Webhook"
  // (Required) URL to be invoked for action execution
  url                             = "http://example.com"
  // (Required) The amount of time a connection waits to time out.Range 1 - 45
  timeout                         = 30
  // (Optional) HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported. Default value is 'G'.
  // https://www.site24x7.com/help/api/#http_methods
  method                          = "P"
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) URL to be invoked from an On-Premise Poller agent
  is_poller_webhook               = false
  // (Optional) Denotes On-Premise Poller ID
  poller                          = "113770000023231022"
  // (Optional) Configuration to send incident parameters while executing the action
  send_incident_parameters        = false
  // (Optional) Configuration to send custom parameters while executing the action
  send_custom_parameters          = true
  // (Optional) Mandatory when send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL
  custom_parameters               = "{\"test\":\"abcd\"}"
  // (Optional) Configuration to enable json format for post parameters
  send_in_json_format             = true
  // (Optional) Authentication method to access the action url.
  // https://www.site24x7.com/help/api/#auth_method
  auth_method                     = "B"
  // (Optional) Username for Authentication
  username                        = "username"
  // (Optional) Password for Authentication
  password                        = "password"
  // (Optional) Provider ID of the OAuth Provider to be associated with the action
  // https://www.site24x7.com/help/api/#list-oauth-providers
  oauth2_provider                 = "113770000023231001"
  // (Optional) User Agent to be used while monitoring the website
  user_agent                      = "Mozilla"
  // (Optional) Map of custom HTTP headers to send.
  custom_headers = {
    "Accept" = "application/json"
  }
  // (Optional) Monitors to be associated with the integration when the selection_type = 2.
  monitors                        = ["756"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3.
  tags                        = ["345"]
  // (Optional) List of tag IDs to be associated with the integration.
  alert_tags_id                   = ["123"]
  // (Optional) Configuration to handle ticketing based integration
  manage_tickets                  = false
  // (Optional) URL to be invoked to update the request
  update_url                      = "http://test.tld"
  // (Optional) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  update_method                   = "P"
  // (Optional) Configuration to send incident parameters while updating the ticket.
  update_send_incident_parameters = false
  // (Optional) Configuration to send custom parameters while updating the ticket.
  update_send_custom_parameters   = false
  // (Optional) Mandatory when update_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL
  update_custom_parameters        = "param=value"
  // (Optional) Configuration to post in JSON format while updating the ticket.
  update_send_in_json_format = true
  // (Optional) URL to be invoked to close the request
  close_url                       = "http://test.tld"
  // (Optional) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  close_method                    = "P"
  // (Optional) Configuration to send incident parameters while closing the ticket.
  close_send_incident_parameters  = false
  // (Optional) Configuration to send custom parameters while closing the ticket.
  close_send_custom_parameters    = false
  // (Optional) Mandatory when close_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL
  close_custom_parameters         = "param=value"
  // (Optional) Configuration to post in JSON format while closing the ticket.
  close_send_in_json_format = true
}

```

## Attributes Reference

### Required

* `name` (String) Display name for the Webhook integration.
* `url` (String) Hook URL to which the message will be posted.

### Optional

* `id` (String) The ID of this resource.
* `timeout` (Number) The amount of time a connection waits to time out. Default value is 30. Range 1 - 45.
* `method` (String) HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported. Default value is 'G'. Please refer [API documentation](https://www.site24x7.com/help/api/#http_methods).
* `selection_type` (Number) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'. Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `trouble_alert` (Boolean) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.  Default value is 'true'.
* `critical_alert` (Boolean) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `down_alert` (Boolean) Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
* `is_poller_webhook` (Boolean) Boolean indicating whether it is an On-Premise Poller based Webhook.
* `poller` (String) Denotes On-Premise Poller ID.
* `send_incident_parameters` (Boolean) Configuration to send incident parameters while executing the action.
* `send_custom_parameters` (Boolean) Configuration to send custom parameters while executing the action.
* `custom_parameters` (String) Mandatory when `send_custom_parameters` is set as true. Custom parameters to be passed while accessing the URL.
* `send_in_json_format` (Boolean) Configuration to enable json format for post parameters.
* `auth_method` (String) Authentication method to access the action url. Please refer [API documentation](https://www.site24x7.com/help/api/#auth_method).
* `username` (String) Username for Authentication.
* `password` (String) Password for Authentication.
* `oauth2_provider` (String) Provider ID of the OAuth Provider to be associated with the action. Please refer [API documentation](https://www.site24x7.com/help/api/#list-oauth-providers).
* `user_agent` (String) User Agent to be used while monitoring the website.
* `custom_headers` (Map of String) A Map of Header name and value.
* `monitors` (List of String) Monitors to be associated with the integration when the selection_type = 2.
* `tags` (List of String) Tags to be associated with the integration when the selection_type = 3.
* `alert_tags_id` (List of String) List of tags to be associated with the integration.
* `manage_tickets` (Boolean) Configuration to handle ticketing based integration.
* `update_url` (String) URL to be invoked to update the request.
* `update_method` (String) HTTP Method to access the URL.Please refer [API documentation](https://www.site24x7.com/help/api/#http_methods).
* `update_send_incident_parameters` (Boolean) Configuration to send incident parameters while updating the ticket.
* `update_send_custom_parameters` (Boolean) Configuration to send custom parameters while updating the ticket.
* `update_custom_parameters` (String) Mandatory when `update_send_custom_parameters` is set as true.Custom parameters to be passed while accessing the URL.
* `update_send_in_json_format` (Boolean) Configuration to post in JSON format while updating the ticket.
* `close_url` (String) URL to be invoked to close the request.
* `close_method` (String) HTTP Method to access the URL.Please refer [API documentation](https://www.site24x7.com/help/api/#http_methods).
* `close_send_incident_parameters` (Boolean) Configuration to send incident parameters while closing the ticket.
* `close_send_custom_parameters` (Boolean) Configuration to send custom parameters while closing the ticket.
* `close_custom_parameters` (String) Mandatory when `close_send_custom_parameters` is set as true.Custom parameters to be passed while accessing the URL.
* `close_send_in_json_format` (Boolean) Configuration to post in JSON format while closing the ticket.


Refer [API documentation](https://www.site24x7.com/help/api/#create-webhook) for more information about attributes.