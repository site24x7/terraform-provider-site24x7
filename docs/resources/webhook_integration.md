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
// Webhook Integration API doc: https://www.site24x7.com/help/api/#create-webhook
resource "site24x7_webhook_integration" "webhook_integration" {
  // (Required) Display name for the WebHook
  name                            = "Test WebHook"
  // (Required) URL to be invoked for action execution
  url                             = "http://example.com"
  // (Required) The amount of time a connection waits to time out.Range 1 - 45
  timeout                         = 30
  // (Required) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  method                          = "P"
  // (Required) Resource Type associated with this integration
  // https://www.site24x7.com/help/api/#resource_type_constants
  // Monitor Group not supported
  selection_type                  = 2
  // (Required) URL to be invoked from an On-Premise Poller agent
  is_poller_webhook               = false
  // (Optional) Denotes On-Premise Poller ID
  poller                          = "113770000023231022"
  // (Optional) Configuration to send incident parameters while executing the action
  send_incident_parameters        = false
  // (Optional) Configuration to send custom parameters while executing the action
  send_custom_parameters          = true
  // (Optional) Custom parameters to be passed while accessing the URL
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
  // (Optional) Monitors associated with the integration
  monitors                        = ["756"]
  // (Required) Configuration to handle ticketing based integration
  manage_tickets                  = false
  // (Optional) URL to be invoked to update the request
  update_url                      = "http://test.tld"
  // (Optional) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  update_method                   = "P"
  // (Optional) Configuration to send incident parameters while executing the action
  update_send_incident_parameters = false
  // (Optional) Configuration to send custom parameters while executing the action
  update_send_custom_parameters   = false
  // (Optional) When update_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL
  update_custom_parameters        = "param=value"
  // (Optional) URL to be invoked to close the request
  close_url                       = "http://test.tld"
  // (Optional) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  close_method                    = "P"
  // (Optional) Configuration to send incident parameters while executing the action
  close_send_incident_parameters  = false
  // (Optional) Configuration to send custom parameters while executing the action
  close_send_custom_parameters    = false
  // (Optional) When close_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL
  close_custom_parameters         = "param=value"
  // (Optional) List of tag IDs to be associated with the integration
  alert_tags_id                   = ["123"]
}
```

## Attributes Reference


### Required

* `name` (String) Display name for the integration.
* `url` (String) Hook URL to which the message will be posted.
* `timeout` (Number) The amount of time a connection waits to time out.Range 1 - 45.
* `method` (String) Resource Type associated with this integration.Please refer [API documentation](https://www.site24x7.com/help/api/#http_methods).
* `selection_type` (Number) Resource Type associated with this integration.Monitor Group not supported.Please refer [API documentation](https://www.site24x7.com/help/api/#resource_type_constants).
* `is_poller_webhook` (Bool) URL to be invoked from an On-Premise Poller agent.
* `manage_tickets` (Bool) Configuration to handle ticketing based integration.


### Optional

* `id` (String) The ID of this resource.
* `poller` (String) Denotes On-Premise Poller ID.
* `send_incident_parameters` (Bool) Configuration to send incident parameters while executing the action.
* `send_custom_parameters` (Bool) Configuration to send custom parameters while executing the action.
* `custom_parameters` (String) Custom parameters to be passed while accessing the URL.
* `send_in_json_format` (Bool) Configuration to enable json format for post parameters.
* `auth_method` (String) Authentication method to access the action url.Please refer [API documentation](https://www.site24x7.com/help/api/#auth_method).
* `username` (String) Username for Authentication.
* `password` (String) Password for Authentication.
* `oauth2_provider` (String) Provider ID of the OAuth Provider to be associated with the action.Please refer [API documentation](https://www.site24x7.com/help/api/#list-oauth-providers).
* `user_agent` (String) User Agent to be used while monitoring the website.
* `monitors` (List of String) Monitors associated with the integration.
* `manage_tickets` (Bool) Configuration to handle ticketing based integration.
* `update_url` (String) URL to be invoked to update the request.
* `update_method` (String) HTTP Method to access the URL.Please refer [API documentation](https://www.site24x7.com/help/api/#http_methods).
* `update_send_incident_parameters` (Bool) Configuration to send incident parameters while executing the action. 
* `update_send_custom_parameters` (Bool) Configuration to send custom parameters while executing the action.
* `update_custom_parameters` (String) When update_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL.
* `close_url` (String) URL to be invoked to close the request.
* `close_method` (String) HTTP Method to access the URL.Please refer [API documentation](https://www.site24x7.com/help/api/#http_methods).
* `close_send_incident_parameters` (Bool) Configuration to send incident parameters while executing the action.
* `close_send_custom_parameters` (Bool) Configuration to send custom parameters while executing the action.
* `close_custom_parameters` (String) When close_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL
* `alert_tags_id` (List of String) List of tags to be associated with the integration.

Refer [API documentation](https://www.site24x7.com/help/api/#create-webhook) for more information about attributes.