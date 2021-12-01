---
layout: "site24x7"
page_title: "Site24x7: site24x7_website_monitor"
sidebar_current: "docs-site24x7-resource-website-monitor"
description: |-
  Create and manage a Website monitor in Site24x7.
---

# Resource: site24x7\_website\_monitor

Use this resource to create, update, and delete a website monitor in Site24x7.

## Example Usage

```hcl
// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Example Monitor"

  // (Required) Website address to monitor.
  website = "https://www.example.com"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = 1

  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

}
```

## Attributes Reference

### Required

* `display_name` (String) Display Name for the monitor.
* `website` (String) Website address to monitor.

### Optional

* `id` (String) The ID of this resource.
* `notification_profile_id` (String) Notification profile to be associated with the monitor.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down.
* `tag_ids` (List of String) List of tags to be associated to the monitor.
* `actions` (Map of String) Action to be performed on monitor status changes.
* `auth_pass` (String) Authentication password to access the website.
* `auth_user` (String) Authentication user name to access the website.
* `check_frequency` (Number) Interval at which your website has to be monitored. Default value is 1 minute.
* `http_method` (String) HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported.
* `match_case` (Boolean) Perform case sensitive keyword search or not.
* `match_regex_severity` (Number) Severity with which alert has to raised when the matching regex is found in the website response.
* `match_regex_value` (String) Match the regular expression in the website response.
* `matching_keyword_severity` (Number) Severity with which alert has to raised when the matching keyword is found in the website response.
* `matching_keyword_value` (String)
* `timeout` (Number) Timeout for connecting to website. Default value is 10. Range 1 - 45.
* `unmatching_keyword_severity` (Number) Severity with which alert has to raised when the keyword is not present in the website response.
* `unmatching_keyword_value` (String)
* `up_status_codes` (String) Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.
* `use_name_server` (Boolean) Resolve the IP address using Domain Name Server.
* `user_agent` (String) User Agent to be used while monitoring the website.
* `custom_headers` (Map of String) A Map of Header name and value.
* `response_headers` (Map of String) A Map of Header name and value.
* `response_headers_severity` (Number) Severity with which alert has to raised when the header is found in the website response. Default value is 2. '0' denotes Down and '2' denotes Trouble.


Refer [API documentation](https://www.site24x7.com/help/api/#website) for more information about attributes.