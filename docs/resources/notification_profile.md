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
// Site24x7 notification profile API doc - https://www.site24x7.com/help/api/#notification-profiles
resource "site24x7_notification_profile" "notification_profile_all_attributes_us" {
  // (Required) Display name for the notification profile.
  profile_name = "Notification Profile All Attributes - Terraform"

  // (Optional) Settings to send root cause analysis when monitor goes down. Default is true.
  rca_needed= true

  // (Optional) Settings to downtime only after executing configured monitor actions.
  notify_after_executing_actions = true

  // (Optional) Configuration for delayed notification. Default value is 1. Can take values 1, 2, 3, 4, 5.
  downtime_notification_delay = 2

  // (Optional) Settings to receive persistent notification after number of errors. Can take values 1, 2, 3, 4, 5.
  persistent_notification = 1

  // (Optional) User group ID for downtime escalation.
  escalation_user_group_id = "123456000000025005"

  // (Optional) Duration of Downtime before Escalation. Mandatory if any user group is added for escalation.
  escalation_wait_time = 30

  // (Optional) Email template ID for notification
  template_id = 123456000024578001

  // (Optional) Settings to stop an automation being executed on the dependent monitors.
  suppress_automation = true

  // (Optional) Execute configured IT automations during an escalation.
  escalation_automations = [
    "123456000000047001"
  ]

  // (Optional) Invoke and manage escalations in your preferred third party services.
  escalation_services = [
    "123456000008777001"
  ]
}
```

## Attributes Reference

### Required

* `profile_name` (String) Display name for the notification profile.

### Optional

* `rca_needed` (Boolean) Settings to send root cause analysis when monitor goes down. Default value is true.
* `notify_after_executing_actions` (Boolean) Settings to downtime only after executing configured monitor actions. Default value is true.
* `escalation_wait_time` (Number) Duration of Downtime before Escalation. Mandatory if any user group is added for escalation.
* `downtime_notification_delay` (Number) Configuration for delayed notification. Default value is 1.
* `persistent_notification` (Number) Settings to receive persistent notification after number of errors.
* `escalation_user_group_id` (String) User group ID for downtime escalation.
* `template_id` (String) Email template ID for notification.
* `suppress_automation` (Boolean) Settings to stop an automation being executed on the dependent monitors.
* `escalation_automations` (List of String) Execute configured IT automations during an escalation.
* `escalation_services` (List of String) Invoke and manage escalations in your preferred third party services.

Refer [API documentation](https://www.site24x7.com/help/api/#notification-profiles) for more information about attributes.


