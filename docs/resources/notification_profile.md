---
layout: "site24x7"
page_title: "Site24x7: site24x7_notification_profile"
sidebar_current: "docs-site24x7-resource-notification-profile"
description: |-
  Create and manage a notification profile in Site24x7.
---

# Resource: site24x7\_notification\_profile

Use this resource to create, update and delete a notification profile in Site24x7.

## Example Usage

```hcl

// Site24x7 notification profile API doc - https://www.site24x7.com/help/api/#enhanced-notification-profiles
resource "site24x7_notification_profile" "notification_profile_default" {
  // (Required) Display name for the notification profile.
  profile_name = "Notification Profile(Terraform) - Default"

}

// Site24x7 notification profile API doc - https://www.site24x7.com/help/api/#enhanced-notification-profiles
resource "site24x7_notification_profile" "notification_profile_us" {
  // (Required) Display name for the notification profile.
  profile_name = "Notification Profile(Terraform) - All Attributes"

  // (Optional) Configuration to send root cause analysis when the monitor is down. Default is true.
  rca_needed= false

  // (Optional) Configuration to raise alerts for downtime only after executing the pre-configured monitor actions. Default is false.
  notify_after_executing_actions = true

  // (Optional) Email template ID for notification
  template_id = 123456000024578001

  // (Optional) Configuration to stop automation from being executed on the dependent monitors. Default is true.
  suppress_automation = false

  // (Optional) Configuration to alert the user. All alerts will be sent through the notification mode of your preference. You can also configure the business hours and the status for which you would like to receive an alert. If you do not set any specific business hours or status preferences, you'll receive alerts for all the status changes throughout the day.
  alert_configuration {     
    // (Optional) Status for which alerts should be raised. '-1' denotes 'Any', '0' denotes 'Down', '1' denotes 'Up', '2' denotes 'Trouble' and '3' denotes 'Critical'.                     
    status = 2    
    // (Optional) Predefined business hours during which alerts should be sent. Default value is -1.
    business_hours_id = "123456000036869001"     
    // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.    
    notification_medium = [
      1,
		  2,
		  6
    ]   
    // (Optional) To specify whether the user would receive alerts within or beyond business hours. Default value is '0'.
    outside_business_hours = "0"    
  }

  // (Optional) Configuration to alert the user. All alerts will be sent through the notification mode of your preference. You can also configure the business hours and the status for which you would like to receive an alert. If you do not set any specific business hours or status preferences, you'll receive alerts for all the status changes throughout the day.
  alert_configuration {     
    // (Optional) Status for which alerts should be raised. '-1' denotes 'Any', '0' denotes 'Down', '1' denotes 'Up', '2' denotes 'Trouble' and '3' denotes 'Critical'.                     
    status = -1 
    // (Optional) Predefined business hours during which alerts should be sent. Default value is -1.
    business_hours_id = "-1"     
    // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.    
    notification_medium = [
      1,
		  2,
		  6
    ]   
    // (Optional) To specify whether the user would receive alerts within or beyond business hours. Default value is '0'.
    outside_business_hours = "0"    
  }

  // (Optional) You can choose to delay and receive Down, Trouble, or Critical notifications if the monitor remains in the same state for a specific number of polls. If you haven't configured any Notification Delay for a specific period, you'll receive alerts immediately.
  notification_delay_configuration {    
    // (Optional) Status for which alerts should be raised. '0' denotes 'Down', '2' denotes 'Trouble' and '3' denotes 'Critical'.
    status = 0    
    // (Optional) Alerting Period - Predefined business hours during which alerts should be sent. Default value is '-1' and it denotes 'All Hours'.
    business_hours_id = "123456000036869001"            
    // (Optional) Notify based on the downtime delay constants define here - https://www.site24x7.com/help/api/#notification-profile-constants. Default value is '1' and it denotes 'Notify immediately after failure'.
    notification_delay = 1  
    // (Optional) To specify whether the user would receive alerts within or beyond business hours. Default value is '0' and it denotes 'Within the business_hours_id configured', '1' denotes 'Outside the business_hours_id configured'.
    outside_business_hours = "1"     
  }

  // (Optional) You can choose to delay and receive Down, Trouble, or Critical notifications if the monitor remains in the same state for a specific number of polls. If you haven't configured any Notification Delay for a specific period, you'll receive alerts immediately.
  notification_delay_configuration {    
    // (Optional) Status for which alerts should be raised. '0' denotes 'Down', '2' denotes 'Trouble' and '3' denotes 'Critical'.
    status = 0    
    // (Optional) Alerting Period - Predefined business hours during which alerts should be sent. Default value is '-1' and it denotes 'All Hours'.
    business_hours_id = "-1"            
    // (Optional) Notify based on the downtime delay constants define here - https://www.site24x7.com/help/api/#notification-profile-constants. Default value is '1' and it denotes 'Notify immediately after failure'.
    notification_delay = 1  
    // (Optional) To specify whether the user would receive alerts within or beyond business hours. Default value is '0' and it denotes 'Within the business_hours_id configured', '1' denotes 'Outside the business_hours_id configured'.
    outside_business_hours = "0"     
  }

  // (Optional) Persistent alerts provide continuous notifications until you acknowledge the Down/Critical/Trouble alarm. You will be receiving alerts until you acknowledge the alarms, at the frequency you've configured in the Notify Every Field.
  persistent_alert_configuration {    
    // (Optional) Denotes the number of times the error has to be ignored before sending a notification. Value ranges from 0-60.
    notify_every = 3   
    // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.    
    notification_medium = [
      1,
		  2,
		  6
    ]
    // (Optional) List of third-party services through which you’d wish to receive the notification.
    third_party_services = [
        "123456000025007005"
    ]
  }

  // (Optional) Persistent alerts provide continuous notifications until you acknowledge the Down/Critical/Trouble alarm. You will be receiving alerts until you acknowledge the alarms, at the frequency you've configured in the Notify Every Field.
  persistent_alert_configuration {    
    // (Optional) Denotes the number of times the error has to be ignored before sending a notification. Value ranges from 0-60.
    notify_every = 5    
    // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.    
    notification_medium = [
      1,
		  2,
		  6
    ]
    // (Optional) List of third-party services through which you’d wish to receive the notification.
    third_party_services = [
        "123456000024411001",
        "123456000024899001",
    ]
  }

  // (Optional) Configuration to receive persistent notifications after a specific number of errors.
  escalation_levels { 
    // (Required) User group ID for downtime escalation.
    user_group_id = "123456000024800001"            
    // (Required) Mandatory, if any User Alert Group is added for escalation Downtime duration for escalation in mins.
    escalation_wait_time = 5 
    // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.    
    notification_medium = [
      1,
		  2,
		  6
    ]
    // (Optional) List of third-party services through which you’d wish to receive the notification.
    third_party_services = [
        "123456000024411001",
        "123456000024899001",
    ]
  }

  // (Optional) Configuration to receive persistent notifications after a specific number of errors.
  escalation_levels { 
    // (Required) User group ID for downtime escalation.
    user_group_id = "123456000000025009"            
    // (Required) Mandatory, if any User Alert Group is added for escalation Downtime duration for escalation in mins.
    escalation_wait_time = 3 
    // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.    
    notification_medium = [
      1,
		  2,
		  6
    ]
    // (Optional) List of third-party services through which you’d wish to receive the notification.
    third_party_services = [
        "123456000024411001",
        "123456000024899001",
    ]
  }

  // (Optional) Execute configured IT automations during an escalation.
  escalation_automations = [
      "123456000000047001",
  ]

}

```

## Attributes Reference

### Required

* `profile_name` (String) Display name for the notification profile.

### Optional

* `rca_needed` (Boolean) Configuration to send root cause analysis when the monitor is down. Default is true.
* `notify_after_executing_actions` (Boolean) Configuration to raise alerts for downtime only after executing the pre-configured monitor actions. Default is false.
* `template_id` (String) Email template ID for notification.
* `suppress_automation` (Boolean) Configuration to stop automation from being executed on the dependent monitors. Default is true.
* `alert_configuration` (Map) Configuration to alert the user. All alerts will be sent through the notification mode of your preference. You can also configure the business hours and the status for which you would like to receive an alert. If you do not set any specific business hours or status preferences, you'll receive alerts for all the status changes throughout the day. (see [below for map schema](#nestedblock--alert_configuration))
* `notification_delay_configuration` (Map) You can choose to delay and receive Down, Trouble, or Critical notifications if the monitor remains in the same state for a specific number of polls. If you haven't configured any Notification Delay for a specific period, you'll receive alerts immediately. (see [below for map schema](#nestedblock--notification_delay_configuration))
* `persistent_alert_configuration` (Map) Persistent alerts provide continuous notifications until you acknowledge the Down/Critical/Trouble alarm. You will be receiving alerts until you acknowledge the alarms, at the frequency you've configured in the Notify Every Field. (see [below for map schema](#nestedblock--persistent_alert_configuration))
* `escalation_levels` (Map) Configuration to receive persistent notifications after a specific number of errors. (see [below for map schema](#nestedblock--escalation_levels))
* `escalation_automations` (List of String) Execute configured IT automations during an escalation.

<a id="nestedblock--alert_configuration"></a>

### Map Schema for `alert_configuration`

### Required

* `notification_medium` (List of Number) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.

### Optional

* `status` (Number) Status for which alerts should be raised. '-1' denotes 'Any', '0' denotes 'Down', '1' denotes 'Up', '2' denotes 'Trouble' and '3' denotes 'Critical'.
* `business_hours_id` (String) Alerting Period - Predefined business hours during which alerts should be sent. Default value is '-1' and it denotes 'All Hours'.
* `outside_business_hours` (String) To specify whether the user would receive alerts within or beyond business hours. Default value is '0' and it denotes 'Time within the business_hours_id configured', '1' denotes 'Time outside the business_hours_id configured'.

<a id="nestedblock--notification_delay_configuration"></a>

### Map Schema for `notification_delay_configuration`

### Optional

* `status` (Number) Status for which alerts should be raised. '0' denotes 'Down', '2' denotes 'Trouble' and '3' denotes 'Critical'.
* `business_hours_id` (String) Alerting Period - Predefined business hours during which alerts should be sent. Default value is '-1' and it denotes 'All Hours'.
* `notification_delay` (Number) Notify based on the downtime delay constants define here - https://www.site24x7.com/help/api/#notification-profile-constants. Default value is '1' and it denotes 'Notify immediately after failure'.
* `outside_business_hours` (String) To specify whether the user would receive alerts within or beyond business hours. Default value is '0' and it denotes 'Time within the business_hours_id configured', '1' denotes 'Time outside the business_hours_id configured'.

<a id="nestedblock--persistent_alert_configuration"></a>

### Map Schema for `persistent_alert_configuration`

### Required

* `notify_every` (Number) Denotes the number of times the error has to be ignored before sending a notification. Value ranges from 0-60.
* `notification_medium` (List of Number) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.

### Optional

* `third_party_services` (List of String) Third-party services through which you’d wish to receive the notification.

<a id="nestedblock--escalation_levels"></a>

### Map Schema for `escalation_levels`

### Required

* `user_group_id` (String) User group ID for downtime escalation.
* `escalation_wait_time` (Number) Mandatory, if any User Alert Group is added for escalation Downtime duration for escalation in mins.
* `notification_medium` (List of Number) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.

### Optional

* `third_party_services` (List of String) Third-party services through which you’d wish to receive the notification.

Refer [API documentation](https://www.site24x7.com/help/api/#enhanced-notification-profiles) for more information about attributes.


