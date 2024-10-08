---
layout: "site24x7"
page_title: "Site24x7: site24x7_amazon_monitor"
sidebar_current: "docs-site24x7-resource-amazon-monitor"
description: |-
  Create and manage amazon monitors in Site24x7.
---

# Resource: site24x7\_amazon\_monitor

Use this resource to create, update and delete amazon monitors in Site24x7.

## Example Usage

```hcl

# Require aws provider

provider "aws" {
  version = "~> 2.0"
  region  = "us-east-1"
}

# resource and data block to define AWS IAM Role with the name Site24x7Infrastructure-Integrations

resource "aws_iam_role" "site24x7" {
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
  name               = "Site24x7Infrastructure-Integrations"
}

# IAM role policy attachment

resource "aws_iam_role_policy_attachment" "read_only_access" {
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
  role       = aws_iam_role.site24x7.name
}

# IAM role policy definition

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = [
      "sts:AssumeRole"
    ]

    condition {
      test = "StringEquals"

      values = [
        data.site24x7_aws_external_id.s247aws.id
      ]

      variable = "sts:ExternalId"
    }

    effect = "Allow"

# Site24x7 AWS account details 

    principals {
      identifiers = [
        "949777495771"
      ]

      type = "AWS"
    }
  }
}

resource "site24x7_amazon_monitor" "aws_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "aws_added_via_terraform"
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) External ID for the AWS account
  external_id = data.site24x7_aws_external_id.s247aws.id
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) AWS Role ARN
  role_arn = "arn:aws:iam::23243:role/TerraformAdminRole"
  // (Optional) AWS discover frequency
  aws_discovery_frequency = 5
  // (Optional) AWS services to discover. See https://www.site24x7.com/help/api/#aws_discover_services 
  // for knowing service ID.
  aws_discover_services = [1,2,3,4,5,6,8,11,14,15,16,17,18,19,20,21,22,23,25,27,29,30,31,32,33,34,35,38,39,40,41,42,43,45,46,49,48,47,53,59,56,57,58,60,61,62,63,65,66,69,70,68,75,76,79,82,83,85,87,92,95,88]
}

// Site24x7 Amazon Monitor API doc - https://www.site24x7.com/help/api/#amazon-webservice-monitor
resource "site24x7_amazon_monitor" "aws_monitor_site24x7" {
  // (Required) Display name for the monitor
  display_name = "aws_added_via_terraform"
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) External ID for the AWS account
  external_id = data.site24x7_aws_external_id.s247aws.id
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
  // (Required) AWS Role ARN
  role_arn = ""
  // (Required) AWS External ID
  aws_external_id = ""
  // (Optional) AWS discover frequency
  aws_discovery_frequency = 5
  // (Optional) AWS services to discover. See https://www.site24x7.com/help/api/#aws_discover_services 
  // for knowing service ID.
  aws_discover_services = [1]
  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

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
}

# Data block to get the site24x7 external ID and Role ARN details 

data "site24x7_aws_external_id" "s247aws" {}

// Displays AWS External ID
output "s247_external_id" {
  description = "AWS External ID : "
  value       = data.site24x7_aws_external_id.s247aws.id
}

data "aws_iam_role" "role_arn" {
	name = aws_iam_role.site24x7.name
}

// Displays AWS Role ARN
output "rolearn" {
  description = "AWS rolearn : "
  value       = data.aws_iam_role.role_arn.arn
}
```

## Attributes Reference

### Required

* `display_name` (String) Display name for the AWS monitor.
* `role_arn` (String) Role ARN for the AWS account.
* `external_id` (String) External ID for the AWS account.
* `aws_discover_services` (List of String) List of AWS services that needs to be discovered. Please refer [API documentation](https://www.site24x7.com/help/api/#aws_discover_services) for knowing AWS service ID's.

### Optional

* `aws_discovery_frequency` (Number) Rediscovery polling interval for the AWS account. Please refer [API documentation](https://www.site24x7.com/help/api/#aws_discover_frequency) for knowing values that can be configured.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.

### Output

* `id` (String) The ID of this resource.

Refer [API documentation](https://www.site24x7.com/help/api/#amazon-webservice-monitor) for more information about attributes.
