terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest

    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  # oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
  data_center = "US"

  // The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

// Site24x7 APM Insight API doc - https://www.site24x7.com/help/api/#apm-insight-applications

// ============================================================================
// IMPORTANT: APM Insight applications cannot be created from Terraform.
//
// The APM Insight API has no create endpoint. An application comes into
// existence when an APM Insight agent is deployed and first reports in. The
// site24x7_apm_application resource ADOPTS an application that already exists
// so that its managed/suspended lifecycle can be driven from code.
// ============================================================================

// ========================
// Data Source - Read Only
// ========================

// Look up an application by name rather than hardcoding its ID.
data "site24x7_apm_application" "checkout" {
  // (Optional) Regular expression matched against application names.
  // It is an error for this to match more than one application.
  name_regex = "^checkout-api$"

  // (Optional) Time window used when querying the API. Defaults to "H".
  // This only affects the request path - the identity and configuration
  // attributes read here are the same for every window.
  # time_window = "H"
}

// List every application whose name starts with "prod-".
data "site24x7_apm_applications" "production" {
  name_regex = "^prod-"

  // (Optional) Filter by managed state: "managed", "unmanaged" or "any".
  managed_state = "managed"
}

// Every Java instance reporting into the checkout application.
data "site24x7_apm_instances" "checkout_jvms" {
  application_id = data.site24x7_apm_application.checkout.application_id
  ins_type       = "JAVA"
}

// ====================================
// Adopt an application's lifecycle
// ====================================
resource "site24x7_apm_application" "checkout" {
  // (Required) ID of an APM Insight application that already exists.
  application_id = data.site24x7_apm_application.checkout.application_id

  // (Optional) Whether the application should be managed (true) or
  // suspended (false). Defaults to true.
  managed = true

  // (Optional) Whether `terraform destroy` should also DELETE the application
  // in Site24x7. Defaults to false, which only releases it from Terraform
  // state. Setting this to true permanently deletes the application and its
  // accumulated monitoring history, and cannot be undone.
  delete_on_destroy = false
}

// ====================================
// Suspend a legacy application
// ====================================
resource "site24x7_apm_application" "legacy_batch" {
  application_id = "101071000000034999"

  // Unmanaging stops alerting without deleting any history.
  managed = false
}

// ========================
// Outputs
// ========================

output "checkout_application_name" {
  description = "Name of the checkout APM application"
  value       = data.site24x7_apm_application.checkout.application_name
}

output "checkout_instance_ids" {
  description = "Instances reporting into the checkout application"
  value       = data.site24x7_apm_application.checkout.instance_ids
}

output "production_application_names" {
  description = "Names of the production APM applications"
  value       = [for app in data.site24x7_apm_applications.production.applications : app.application_name]
}

output "checkout_jvm_agent_versions" {
  description = "Agent version deployed at each checkout JVM"
  value = {
    for instance in data.site24x7_apm_instances.checkout_jvms.instances :
    instance.instance_name => instance.agent_version
  }
}
