---
layout: "site24x7"
page_title: "Site24x7: site24x7_user"
sidebar_current: "docs-site24x7-resource-user"
description: |-
  Create and manage a user in Site24x7.
---

# Resource: site24x7\_user

Use this resource to create, update and delete a user in Site24x7.

## Example Usage

```hcl

// User API doc: https://www.site24x7.com/help/api/#users
resource "site24x7_user" "user_basic" {

  // (Required) Name of the User.
  display_name = "User - Terraform"

  // (Required) Email address of the user. Email verification has to be done manually.
  email_address = "jim@example.com"

  // (Required) Phone number configurations to receive alerts.
  mobile_settings = {
    "country_code" = "93"
    "mobile_number"= "434388234"
  }

  // (Required) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  notification_medium = [
    1,
  ]

  // (Required) Role assigned to the user for accessing Site24x7. Role will be updated only after the user accepts the invitation. Refer https://www.site24x7.com/help/api/#site24x7_user_constants
  user_role = 10
  
  // (Required) Medium through which you’d wish to receive the down alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  down_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the critical alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  critical_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the trouble alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  trouble_notification_medium = [
    1,
  ]

  // (Required) Medium through which you’d wish to receive the up alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
  up_notification_medium = [
    1,
  ]
}

```

## Attributes Reference


### Required

* `display_name` (String) Name of the user.
* `email_address` (String) Email address of the user. Email verification has to be done manually.
* `user_role` (Number) Role assigned to the user for accessing Site24x7. Role will be updated only after the user accepts the invitation. Refer https://www.site24x7.com/help/api/#site24x7_user_constants.
* `mobile_settings` (Map) Phone number configurations to receive alerts. {country_code: $country_code, mobile_number: $mobile_number, sms_provider_id: $sms_providers, call_provider_id: $voice_call_provider}. Refer https://www.site24x7.com/help/api/#alerting_constants
* `notification_medium` (List of Number) Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'. 
* `down_notification_medium` (List of Number) Medium through which you’d wish to receive the Down alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
* `critical_notification_medium` (List of Number) Medium through which you’d wish to receive the Critical alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
* `trouble_notification_medium` (List of Number) Medium through which you’d wish to receive the Trouble alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.
* `up_notification_medium` (List of Number) Medium through which you’d wish to receive the Up alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.


### Optional

* `id` (String) The ID of this resource.
* `job_title` (Number) Provide your job title to be added in Site24x7. Refer https://www.site24x7.com/help/api/#job_title
* `selection_type` (Number) Resource type associated to this user. Default value is '0'. Can take values 0|1. '0' denotes 'All Monitors', '1' denotes 'Monitor Group'. 'monitor_groups' attribute is mandatory when the 'selection_type' is '1'.
* `monitor_groups` (List of String) List of monitor groups to which the user has access to. 'monitor_groups' attribute is mandatory when the 'selection_type' is '1'.
* `user_group_ids` (List of String) List of groups to be associated for the user for receiving alerts.
* `alerting_period_start_time` (String) Define a time window so you can receive Voice/SMS status alerts during this period alone. You can't define this window for email or IM based notifications.
* `alerting_period_end_time` (String) Define a time window so you can receive Voice/SMS status alerts during this period alone. You can't define this window for email or IM based notifications.
* `email_format` (Number) Denotes the email format. '0' - Text, '1' - HTML.


Refer [API documentation](https://www.site24x7.com/help/api/#users) for more information about attributes.
 
