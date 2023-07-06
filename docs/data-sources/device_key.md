---
layout: "site24x7"
page_title: "Site24x7: site24x7_device_key"
sidebar_current: "docs-site24x7-data-source-device-key"
description: |-
  Fetch the device key in Site24x7.
---

# Data Source: site24x7\_device\_key

Use this data source to retrieve device key in Site24x7.

## Example Usage

```hcl

// Data source to fetch the device key
data "site24x7_device_key" "s247devicekey" {}

// Displays the device key
output "s247_device_key" {
  description = "Device Key : "
  value       = data.site24x7_device_key.s247devicekey.id
}

```

## Attributes Reference

### Read-Only

* `id` (String) Denotes the device key of your account.


 