layout: "site24x7"
page_title: "Site24x7: site24x7_customer"
sidebar_current: "docs-site24x7-customer"
description: |-
Create and manage Customer configurations in Site24x7 MSP.

Resource: site24x7_customer

Use this resource to create, update, and delete Customers in Site24x7 MSP.

Example Usage

// Site24x7 Customer API doc - https://www.site24x7.com/help/api/#msp-create-customer
resource "site24x7_customer" "customer_creation {
  // (Required) Company name of the customer.
  customer_company = "Customer Name1"

  // (Required) Display name for the customer.
  display_name     = "Main 007"

  // (Required) Unique portal name.
  portal_name      = "testacc007"

  // (Optional) Captcha string, if applicable.
  captcha          = "abcd1234"

  // (Optional) Customer groups to which this customer belongs.
  customer_groups  = ["47627000000087001"]

  // (Optional) Zuids associated with the customer.
  zuids            = ["83625125"]
}

Attributes Reference

Required

customer_company (String) - Company name of the customer.

display_name (String) - Display name for the customer.

portal_name (String) - Unique portal name for the customer.

Optional

captcha (String) - Captcha string if required by the API.

customer_groups (List of String) - IDs of customer groups to which this customer belongs.

zuids (List of String) - ZUIDs associated with this customer.

Read-Only

customer_id (String) - Unique identifier for the customer, returned by the API after creation.

Refer to the API documentation (https://www.site24x7.com/help/api/#msp-create-customer) for more details about attributes and supported values.
