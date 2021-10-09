---
layout: "site24x7"
page_title: "Site24x7: site24x7_rest_api_monitor"
sidebar_current: "docs-site24x7-resource-rest-api-monitor"
description: |-
  Create and manage a REST API monitor in Site24x7.
---

# Resource: site24x7\_rest\_api\_monitor

Use this resource to create, update, and delete a REST API monitor in Site24x7.

## Example Usage

```hcl
// Site24x7 Rest API Monitor API doc - https://www.site24x7.com/help/api/#rest-api
resource "site24x7_rest_api_monitor" "rest_api_monitor_us" {
  // (Required) Display name for the monitor
  display_name = "rest api - terraform"
  // (Required) Website address to monitor.
  website = "https://dummy.restapiexample.com/"
  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"
  // (Optional) Check for the keyword in the website response.
  matching_keyword = {
 	  severity= 2
 	  value= "aaa"
 	}
  // (Optional) Check for non existence of keyword in the website response.
  unmatching_keyword = {
 	  severity= 2
 	  value= "bbb"
 	}
  // (Optional) Match the regular expression in the website response.
  match_regex = {
 	  severity= 2
 	  value= ".*aaa.*"
 	}
}
```

## Attributes Reference


### Required

* `display_name` (String)
* `website` (String)

### Optional

* `auth_pass` (String)
* `auth_user` (String)
* `check_frequency` (Number)
* `client_certificate_password` (String)
* `custom_headers` (Map of String)
* `http_method` (String)
* `http_protocol` (String)
* `id` (String) The ID of this resource.
* `json_schema_check` (Boolean)
* `jwt_id` (String)
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `match_case` (Boolean)
* `match_regex` (Map of String)
* `matching_keyword` (Map of String)
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `notification_profile_id` (String) Notification profile to be associated with the monitor.
* `oauth2_provider` (String)
* `request_content_type` (String)
* `request_param` (String)
* `response_content_type` (String)
* `ssl_protocol` (String)
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `timeout` (Number)
* `unmatching_keyword` (Map of String)
* `use_alpn` (Boolean)
* `use_ipv6` (Boolean)
* `use_name_server` (Boolean)
* `user_agent` (String)
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down.


Refer [API documentation](https://www.site24x7.com/help/api/#rest-api) for more information about attributes.
