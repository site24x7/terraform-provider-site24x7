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


  // Site24x7 Web-Transaction-Browser Monitor API doc - https://www.site24x7.com/help/api/#web-transaction-(browser)
  resource "site24x7_web_transaction_browser_monitor" 
  "web_transaction_browser_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "RBM-www.demoqa.com"
   
  //(Optional) Type of the monitor
  type="REALBROWSER"

  // (Required) Unique name to be used in the ping URL.
 base_url: "https://www.demoqa.com"

  //(Optional)Recorded transaction script type.(txt , side)
  selenium_script="{\"id\":\"4bab335d-6abe-4450-9ca7-ab160de8fcb6\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://demoqa.com\",\"tests\":[{\"id\":\"742a0bbc-e523-461a-afeb-0636271a4361\",\"name\":\"\",\"commands\":[{\"id\":\"0621950a-7705-4aed-a9b7-74fcb968cac3\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://demoqa.com/\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"af22dd1f-886c-4251-8214-98e774f19579\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"b58182de-36a4-4575-9839-a4b0793891b2\",\"comment\":\"\",\"command\":\"click\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:position\"],[\"css=.card:nth-child(2) h5\",\"css:finder\"],[\"xpath=//h5[contains(text(),'Forms')]\",\"xpath:link\"],[\"xpath=//div[@id='app']/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div[2]/div/div[3]/h5\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"3c4eac2b-536a-4de1-9ce1-18fa72eab52c\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"click Practice Form\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/forms\",\"stepCount\":2,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"7673bc49-3288-495b-93bc-39e524edf2e7\",\"comment\":\"\",\"command\":\"clickAndWait\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"xpath:position\"],[\"css=.show .text\",\"css:finder\"],[\"xpath=//span[contains(text(),'Practice Form')]\",\"xpath:link\"],[\"xpath=(//li[@id='item-0']/span)[2]\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div/div/div[2]/div/ul/li/span\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"c026cc61-1301-428c-b7ca-3e4a76ee1b1d\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"742a0bbc-e523-461a-afeb-0636271a4361\"]}],\"urls\":[\"https://demoqa.com/\"],\"plugins\":[]}"

  //(Optional)Recorded transaction script type.(txt , side)
  script_type="txt"

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

## Attributes Reference

### Required
* `base_url`     (String) BaseURL of the transaction.
* `selenium_script`(String)Recorded Trasanction script to create a monitor.
* `script_type`   (String)Recorded transaction script type.(txt , side)
* `proxy_details` (List of String) Proxy checking in the web response
* `cookies` (Map of String) Cookies for Advanced settings
### Optional
* `id` (String) The ID of this resource.
* `display_name` (String) Display Name for the monitor.
* `type`         (String) REALBROWSER.
* `check_frequency` (String)Check interval for monitoring.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `page_load_time` (Number)Timeout for page load.
* `async_dc_enabled` (Boolean) When asynchronous data collection is enabled, polling will be carried out from all the locations at the same time. If it is disabled, polling will be done consecutively from the selected locations.
* `think_time`  (Number)Think time between each steps.
<!-- * `parallel_polling` Enable parallel polling to initiate data collection from all the configured monitoring locations simultaneously during hourly polls -->
* `perform_automation` (Boolean) Check box to do automation or not
* `resolution` (String) Screen resolution for running the script.
* `browser_type` (Number) Type of the browser. 1 - Firefox, 2 - Chrome. Default value is 1.
* `browser_version` (Number) Version of the browser. 83 - Firefox, 88 - Chrome. Default value is 83.
* `custom_headers` (Map of String) A Map of Header name and value.
* `user_agent` (String) User Agent to be used while monitoring the website.
* `auth_pass` (String) Authentication password to access the website.
* `auth_user` (String) Authentication user name to access the website.
* `credential_profile_id` (String)Credential Profile to associate the website with. Notes: If you're using Auth user and Auth password, you can't configure Credential Profile
* `ip_type` (Boolean)Whether ipv6 or ipv4
* `ignore_cert_err` (Boolean) ssl certificate configuration for the monitor
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.

Refer [API documentation](https://www.site24x7.com/help/api/#web-transaction-(browser)) for more information about attributes.
