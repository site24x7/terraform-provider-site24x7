---
layout: "site24x7"
page_title: "Site24x7: site24x7_msp"
sidebar_current: "docs-site24x7-data-source-msp"
description: |-
  Get information about a MSP in Site24x7.
---

# Data Source: site24x7\_msp

Use this data source to retrieve information about MSP in Site24x7.

## Example Usage

```hcl

// Data source to fetch a MSP customer
data "site24x7_msp" "s247mspcustomer" {
  // (Required) Regular expression denoting the name of the MSP customer.
  customer_name_regex = "a"
  
}

// Displays MSP customer ID
output "s247_msp_id" {
  description = "MSP customer ID : "
  value       = data.site24x7_msp.s247mspcustomer.id
}

// Displays MSP customer name
output "s247_msp_name" {
  description = "MSP customer name : "
  value       = data.site24x7_msp.s247mspcustomer.customer_name
}

// Displays the matching ZAAIDs
output "s247_matching_ids" {
  description = "Matching MSP customer IDs : "
  value       = data.site24x7_msp.s247mspcustomer.matching_zaaids
}

// Displays the matching ZAAIDs and names
output "s247_matching_ids_and_names" {
  description = "Matching MSP customer IDs and names : "
  value       = data.site24x7_msp.s247mspcustomer.matching_zaaids_and_names
}

// Displays ZAAID of the customer
output "s247_msp_customer_zaaid" {
  description = "ZAAID : "
  value       = data.site24x7_msp.s247mspcustomer.zaaid
}

// Displays user ID of the customer
output "s247_msp_customer_userid" {
  description = "User ID : "
  value       = data.site24x7_msp.s247mspcustomer.user_id
}

```

## Attributes Reference

### Required

* `customer_name_regex` (String) Regular expression denoting the name of the MSP customer.

### Read-Only

* `id` (String) ZAAID of the MSP customer.
* `matching_zaaids` (List) List of ZAAIDs matching the attribute `customer_name_regex`.
* `matching_zaaids_and_names` (List) List of ZAAIDs and names matching the attribute `customer_name_regex`.
* `customer_name` (String) Display name for the MSP customer.
* `zaaid` (String) ZAAID of the MSP customer.
* `user_id` (String) User ID of the MSP customer.


 