---
layout: "site24x7"
page_title: "Site24x7: site24x7_web_page_speed_monitor"
sidebar_current: "docs-site24x7-resource-web-page-speed-monitor"
description: |-
  Create and manage a Web Page Speed monitor in Site24x7.
---

# Resource: site24x7\_web\_page\_speed\_monitor

Use this resource to create, update and delete a web page speed monitor in Site24x7.

## Example Usage

```hcl

// Website Monitor API doc: https://www.site24x7.com/help/api/#web-page-speed-(browser)
resource "site24x7_web_page_speed_monitor" "web_page_speed_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "mymonitor"

  // (Required) Website address to monitor.
  website = "https://foo.bar"

  // (Required) Type of the browser. 1 - Firefox, 2 - Chrome. Default value is 1
  browser_type = 2

  // (Required) Version of the browser. 83 - Firefox, 88 - Chrome. Default value is 83
  browser_version = 88

  // (Optional) Check interval for monitoring. Default: 1. See
  // https://www.site24x7.com/help/api/#check-interval for all supported
  // values.
  check_frequency = "15"

  // (Optional) Timeout for connecting to the website. Range 1 - 45. Default: 30
  timeout = 40

  // (Optional) Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure. 
  use_ipv6 = false

  // (Optional) Type of content the website page has. 1 - Static Website,	2 - Dynamic Website, 3 - Flash-Based Website. Default value is 1.
  website_type = 2

  // (Optional) Type of the device used for running the speed test. 1 - Desktop, 2 - Mobile, 3 - Tab. Default value is "1".
  device_type = "1"

  // (Optional) Set a resolution based on your preferred device type.
  wpa_resolution = "1024,768" 

  // HTTP configuration

  // (Optional) HTTP Method to be used for accessing the website. Default: "G".
  // See https://www.site24x7.com/help/api/#http_methods for allowed values.
  http_method = "G"

  // (Optional) User Agent to be used while monitoring the website.
  user_agent = "some user agent string"

  // (Optional) Authentication user name to access the website.
  auth_user = "theuser"

  // (Optional) Authentication password to access the website.
  auth_pass = "thepasswd"

  // (Optional) Map of custom HTTP headers to send.
  custom_headers = {
    "Accept" = "application/json"
    "Accept-Encoding" = "gzip"
  }

  // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon. Default: ""
  up_status_codes = "200,404"


  // Content Check Configuration

  // (Optional) Check for the keyword in the website response.
  matching_keyword_value = "foo"

  // (Optional) Alert type to match on. See
  // https://www.site24x7.com/help/api/#alert-type-constants for available
  // values.
  matching_keyword_severity = 2

  // (Optional) Check for non existence of keyword in the website response.
  unmatching_keyword_value = "error"

  // (Optional) Alert type to match on. See
  // https://www.site24x7.com/help/api/#alert-type-constants for available
  // values.
  unmatching_keyword_severity = 2

  // (Optional) Perform case sensitive keyword search or not. Default: false.
  match_case = true

  // (Optional) Match the regular expression in the website response.
  match_regex_value = ".*imprint.*"

  // (Optional) Alert type to match on. See
  // https://www.site24x7.com/help/api/#alert-type-constants for available
  // values.
  match_regex_severity = 2

  // Configuration Profiles

  // (Optional) Location Profile to be associated with the monitor. If 
  // location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_id = "123"

  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

  // (Optional) Notification profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/notification_profiles
  // endpoint (https://www.site24x7.com/help/api/#list-notification-profiles)
  // will be used.
  notification_profile_id = "123"

  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  // (Optional) Threshold profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/threshold_profiles
  // endpoint for the website monitor (https://www.site24x7.com/help/api/#list-threshold-profiles) will
  // be used.
  threshold_profile_id = "123"

  // (Optional) List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    "123",
    "456"
  ]

  // (Optional) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
  dependency_resource_ids = [
    "123",
    "456"
  ]

  // (Optional) List if user group IDs to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_ids = [
    "123",
  ]

  // (Optional) List if user group names to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]

  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]

  // (Optional) List of tag names to be associated to the monitor. Tag name matching works for both exact and 
  //  partial match. Either specify tag_ids or tag_names.
  tag_names = [
    "Terraform",
    "Network",
  ]

  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]

  // (Optional) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
  actions = {
    "1" = "123"
  }
}

```

## Attributes Reference

### Required

* `display_name` (String) Display Name for the monitor.
* `website` (String) Website address to monitor.

### Optional

* `id` (String) The ID of this resource.
* `check_frequency` (String) Interval at which your website has to be monitored. Default value is 1 minute.
* `timeout` (Number) Timeout for connecting to website. Default value is 10. Range 1 - 45.
* `use_ipv6` (Number) Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.
* `website_type` (Number) Type of content the website page has. 1 - Static Website,	2 - Dynamic Website, 3 - Flash-Based Website. Default value is 1.
* `device_type` (String) Type of the device used for running the speed test. 1 - Desktop, 2 - Mobile, 3 - Tab. Default value is "1".
* `browser_type` (Number) Type of the browser. 1 - Firefox, 2 - Chrome. Default value is 1.
* `browser_version` (Number) Version of the browser. 83 - Firefox, 88 - Chrome. Default value is 83.
* `wpa_resolution` (String) Set a resolution based on your preferred device type.
* `http_method` (String) HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported.
* `custom_headers` (Map of String) A Map of Header name and value.
* `user_agent` (String) User Agent to be used while monitoring the website.
* `auth_pass` (String) Authentication password to access the website.
* `auth_user` (String) Authentication user name to access the website.
* `up_status_codes` (String) Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.
* `matching_keyword_severity` (Number) Severity with which alert has to raised when the matching keyword is found in the website response.
* `matching_keyword_value` (String) Check for the keyword in the website response.
* `unmatching_keyword_severity` (Number) Severity with which alert has to raised when the keyword is not present in the website response.
* `unmatching_keyword_value` (String) Check for the absence of the keyword in the website response.
* `match_case` (Boolean) Perform case sensitive keyword search or not.
* `match_regex_severity` (Number) Severity with which alert has to raised when the matching regex is found in the website response.
* `match_regex_value` (String) Match the regular expression in the website response.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `actions` (Map of String) Action to be performed on monitor IT Automation templates. 

Refer [API documentation](https://www.site24x7.com/help/api/#web-page-speed-(browser)) for more information about attributes.