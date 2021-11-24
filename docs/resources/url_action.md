---
layout: "site24x7"
page_title: "Site24x7: site24x7_url_action"
sidebar_current: "docs-site24x7-resource-url-action"
description: |-
  Create and manage a URL action in Site24x7.
---

# Resource: site24x7\_url\_action

Use this resource to create, update, and delete a URL IT automation action in Site24x7.

## Example Usage

```hcl
// Site24x7 IT Automation API doc - https://www.site24x7.com/help/api/#it-automation
resource "site24x7_url_action" "action_us" {
  // (Required) Display name for the action.
  name = "IT Action"

  // (Required) URL to be invoked for action execution.
  url = "https://www.example.com"

  // (Optional) HTTP Method to access the URL. Default: "P". See
  // https://www.site24x7.com/help/api/#http_methods for allowed values.
  method = "G"

  // (Optional) If send_custom_parameters is set as true. Custom parameters to
  // be passed while accessing the URL.
  custom_parameters = "param=value"

  // (Optional) Configuration to send custom parameters while executing the action.
  send_custom_parameters = true

  // (Optional) Configuration to enable json format for post parameters.
  send_in_json_format = true

  // (Optional) Configuration to send incident parameters while executing the action.
  send_incident_parameters = true

  // (Optional) The amount of time a connection waits to time out. Range 1 - 90. Default: 30.
  timeout = 10
}
```

## Attributes Reference

### Required

* `name` (String) Display name for the Action.
* `url` (String) URL to be invoked for action execution.

### Optional

* `auth_method` (String) Authentication method to access the action url.
* `custom_parameters` (String) Mandatory, if send_custom_parameters is set as true. Custom parameters to be passed while accessing the action url.
* `id` (String) The ID of this resource.
* `method` (String) HTTP Method to access the action url.
* `timeout` (Number) Timeout for connecting to URL. Default value is 10. Range 1 - 90.
* `send_custom_parameters` (Boolean) Configuration to send custom parameters while executing the action.
* `send_email` (Boolean) Boolean indicating whether to send email or not.
* `send_in_json_format` (Boolean) Optional, use only if HTTP Method chosen is GET. Configuration to enable json format for post parameters.
* `send_incident_parameters` (Boolean) Configuration to send incident parameters while executing the action.
* `user_agent` (String) User Agent to be used while monitoring the website.

Refer [API documentation](https://www.site24x7.com/help/api/#it-automation) for more information about attributes.

