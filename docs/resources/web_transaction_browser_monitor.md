---
layout: "site24x7"
page_title: "Site24x7: site24x7_web_transaction_browser_monitor"
sidebar_current: "docs-site24x7-resource-web-transaction-browser-monitor"
description: |-
  Create and manage a Web Transaction Browser monitor in Site24x7.
---

# Resource: site24x7\_web\_transaction\_browser\_monitor

Use this resource to create, upda te and delete a Web-Transaction-Browser monitor in Site24x7.

## Example Usage

```hcl

  // Site24x7 Web-Transaction-Browser Monitor API doc - https://www.site24x7.com/help/api/#web-transaction-(browser)
  resource "site24x7_web_transaction_browser_monitor" "web_transaction_browser_monitor_example" {
    // (Required) Display name for the monitor
    display_name = "RBM-www.demoqa.com"
    
    //(Optional) Type of the monitor
    type="REALBROWSER"

    // (Required) Unique name to be used in the ping URL.
    base_url: "https://www.example.com"

    //(Required while creation)Recorded transaction script type.(txt , side)
    selenium_script="{\"id\":\"b500a6da-dbb7-4d0c-968e-ae3fd6ab411f\",\"version\":\"1.1\",\"name\":\"\",       \"url\":\"https://www.example.com\",\"tests\":[{\"id\":\"f30156f4-3a70-4031-9a41-dbe8eac7e494\",\"name\":\"\",\"commands\":[{\"id\":\"59b9fca1-28f3-46eb-afc5-0be35a1f582f\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://www.example.com\",\"targets\":[],\"value\":\"\",\"URL\":\"https://www.example.com\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"14bbdd0e-f78f-4591-916e-4b6cf94ce576\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"38e2db3f-b835-47e8-8786-b87792d6fe4f\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"f30156f4-3a70-4031-9a41-dbe8eac7e494\"]}],\"urls\":[\"https://www.example.com/\"],\"plugins\":[]}"

    //(Required while creation)Recorded transaction script type.(txt , side)
    script_type="txt"
    // (Required) Provide the Cookies for parameter forwarding in   Map format.
    cookies ={
          "Accept" = "application/json"
        }
    //(Optional)Map of String
    proxy_details={webProxyUrl: "qwerty", webProxyUname: "sdcsdcsdc", webProxyPass: "sadasds"}

    //(Optional)Map of auth-details
    auth_details={userName: "12345",  password: "12345"}

    //(Optional)Check interval for monitoring.
    check_frequency="15"

    //(Optional)
    async_dc_enabled=false

    //(Optional)Choose the browser type. Default is value is 1.
    browser_type=1

    //(Optional)
    browser_version =10101

    //(Optional)Think time between each steps
    think_time=1

    //(Optional)Timeout for page load.
    page_load_time=60

    //(Optional) Screen resolution for running the script.
    resolution="1600,900"

    // (Optional) Threshold profile to be associated with the monitor. If
    // omitted, the first profile returned by the /api/threshold_profiles
    // endpoint for the HEARTBEAT monitor type (https://www.site24x7.com/help/api/#list-threshold-profiles) will
    // be used.
    threshold_profile_id = "456418000342341"

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
      "Server",
    ]

    // (Optional) List of Third Party Service IDs to be associated to the monitor.
    third_party_service_ids = [
      "4567"
    ]

    // (Optional) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.
    on_call_schedule_id = "3455"
}
```
## Attributes Reference

### Required
* `base_url`     (String) BaseURL of the transaction.
* `selenium_script`(String)Recorded Trasanction script to create a monitor.(Required for creation)
* `script_type`   (String)Recorded transaction script type.(txt , side)(Required for creation)
* `cookies` (Map of String) Cookies for Advanced settings(Required)
### Optional
* `id` (String) The ID of this resource.
* `display_name` (String) Display Name for the monitor.
* `type`         (String) REALBROWSER.
* `check_frequency` (String)Check interval for monitoring.
* `resolution` (String) Screen resolution for running the script.
* `browser_type` (Number) Type of the browser. 1 - Firefox, 2 - Chrome. Default value is 1.
* `browser_version` (Number) Version of the browser. 83 - Firefox, 88 - Chrome. Default value is 83.
* `proxy_details` (Map of String) Proxy checking in the web response
* `page_load_time` (Number)Timeout for page load.
* `async_dc_enabled` (Boolean) When asynchronous data collection is enabled, polling will be carried out from all the locations at the same time. If it is disabled, polling will be done consecutively from the selected locations.
* `think_time`  (Number)Think time between each steps.
* `auth_details` (Map of String) auth details to access the monitor details
* `custom_headers` (Map of String) A Map of Header name and value.
* `user_agent` (String) User Agent to be used while monitoring the website.
* `ip_type` (Number)Whether ipv6 or ipv4
* `ignore_cert_err` (Boolean) ssl certificate configuration for the monitor
* `perform_automation` (Boolean) Check box to do automation or not
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.
* `actions` (Map of String) Action to be performed on monitor IT Automation templates. 
Refer [API documentation](https://www.site24x7.com/help/api/#web-transaction-(browser)) for more information about attributes.
