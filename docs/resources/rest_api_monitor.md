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
  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]
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

* `display_name` (String) Display Name for the monitor.
* `website` (String) Website address to monitor.

### Optional

* `id` (String) The ID of this resource.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `notification_profile_id` (String) Notification profile to be associated with the monitor.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down.
* `tag_ids` (List of String) List of tags to be associated to the monitor.
* `check_frequency` (Number) Interval at which your website has to be monitored. Default value is 1 minute.
* `timeout` (Number) Timeout for connecting to website. Default value is 10. Range 1 - 45.
* `client_certificate_password` (String) Password of the uploaded client certificate.
* `http_method` (String) HTTP Method used for accessing the website. Default value is G.
* `http_protocol` (String) Specify the version of the HTTP protocol. Default value is H1.1.
* `auth_pass` (String) Authentication user name to access the website.
* `auth_user` (String) Authentication password to access the website.
* `oauth2_provider` (String) Provider ID of the OAuth Provider to be associated with the monitor.
* `json_schema_check` (Boolean) Enable this option to perform the JSON schema check.
* `jwt_id` (String) Token ID of the Web Token to be associated with the monitor.
* `match_case` (Boolean) Perform case sensitive keyword search or not.
* `match_regex` (Map of String) Match the regular expression in the website response.
* `matching_keyword` (Map of String) Check for the keyword in the website response.
* `unmatching_keyword` (Map of String) Check for non existence of keyword in the website response.
* `request_content_type` (String) Provide content type for request params.
* `request_param` (String) Provide parameters to be passed while accessing the website.
* `response_content_type` (String) Response content type.
* `ssl_protocol` (String) Specify the version of the SSL protocol. If you are not sure about the version, use Auto.
* `use_alpn` (Boolean) Enable ALPN to send supported protocols as part of the TLS handshake.
* `use_ipv6` (Boolean) Select IPv6 for monitoring the websites hosted with IPv6 address. If you choose non IPv6 supported locations, monitoring will happen through IPv4.
* `use_name_server` (Boolean) Resolve the IP address using Domain Name Server.
* `user_agent` (String) User Agent to be used while monitoring the website.
* `custom_headers` (Map of String) A Map of Header name and value.
* `response_headers` (Map of String) A Map of Header name and value.
* `response_headers_severity` (Number) Severity with which alert has to raised when the header is found in the website response. Default value is 2. '0' denotes Down and '2' denotes Trouble.


Refer [API documentation](https://www.site24x7.com/help/api/#rest-api) for more information about attributes.
