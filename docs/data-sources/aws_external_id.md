---
layout: "site24x7"
page_title: "Site24x7: site24x7_aws_external_id"
sidebar_current: "docs-site24x7-data-source-aws-external-id"
description: |-
  Fetch the AWS External ID in Site24x7.
---

# Data Source: site24x7\_aws\_external\_id

Use this data source to retrieve AWS External ID in Site24x7.

## Example Usage

```hcl

// Data source to fetch the AWS External ID
data "site24x7_aws_external_id" "s247aws" {}

// Displays AWS External ID
output "s247_external_id" {
  description = "AWS External ID : "
  value       = data.site24x7_aws_external_id.s247aws.id
}

```

## Attributes Reference

### Read-Only

* `id` (String) Denotes the AWS External ID of your account.


 