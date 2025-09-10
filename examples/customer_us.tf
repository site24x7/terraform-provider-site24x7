terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Security recommendation - Store your credentials in a Vault)
  oauth2_client_id     = "<SITE24X7_OAUTH2_CLIENT_ID>"
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // Data center (US/EU/IN/AU/CN/JP/CA)
  data_center = "US"

  // (Optional) ZAAID of the customer under MSP or BU
  zaaid = "1234"

  retry_min_wait = 1
  retry_max_wait = 30
  max_retries    = 4
}

// Customer API doc: https://www.site24x7.com/help/api/#msp_customers
resource "site24x7_customer" "customer_creation" {
  country_code = "US"
  timezone = "Asia/Kolkata"
  language_code = "en"
  industry = "15"
  roletitle = "4"
  invite = false
  customer_groups = [
    "37152000000043029"
  ]
  digest = "1_C_797c6d2644b53cb62763de6ba0980fb01a9a10188ae30dbd313ecfbe1c2417f28f469b5447f67a731fb1359731fcd21259416e5e4025f887b3b9656800f22130"
  zuids = [
    "75086549"
  ]
  customer_company  = "w3schools"
  display_name      = "phillips"
  customer_website  = "https://www.w3schools.com"
  email_address     = "selvalakshmi.m+aug18@zohotest.com"
  portal_name       = "w3schools"
  captcha           = "D6EF1P"
}

resource "site24x7_customer" "customer_creation" {
  country_code     = "US"
  timezone         = "Asia/Kolkata"
  language_code    = "en"
  industry         = "15"
  roletitle        = "4"
  invite           = falseAC

  customer_groups = [
    "37152000000043029"
  ]

  digest = "1_C_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

  zuids = [
    "12345678"
  ]

  customer_company  = "Example Inc"
  display_name      = "example"
  customer_website  = "https://www.example.com"
  email_address     = "test@example.com"
  portal_name       = "example"
  captcha           = "XYZ123"
}
