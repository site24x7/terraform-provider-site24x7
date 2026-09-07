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

// Site24x7 APM Insight API doc - https://www.site24x7.com/help/api/#apm-insight-agent-configurations

// ============================================================================
// Every setting below is named after the API's dotted agent_config key with the
// dots replaced by underscores, so "transaction.trace.enabled" becomes
// "transaction_trace_enabled".
//
// All of them are optional: each defaults to the value Site24x7 uses in its own
// default profiles, so you only declare what you want to differ.
// ============================================================================

// ========================
// Data Sources - Read Only
// ========================

// The default profile every new Java agent picks up.
data "site24x7_apm_agent_config_profile" "java_default" {
  // (Required with default_profile) Type of APM Insight agent.
  agent_type = "JAVA"

  // (Optional) Look up the default profile for agent_type.
  default_profile = true
}

// A specific profile by name, searched within one agent type.
data "site24x7_apm_agent_config_profile" "php_baseline" {
  agent_type = "PHP"

  // (Optional) It is an error for this to match more than one profile.
  name_regex = "^PHP baseline$"
}

// Every profile of every agent type.
data "site24x7_apm_agent_config_profiles" "all" {
  // (Optional) Narrow to one agent type.
  # agent_type = "JAVA"

  // (Optional) Narrow by profile name.
  # name_regex = "^Default profile"
}

// ====================================
// A Java profile with tighter tracing
// ====================================
resource "site24x7_apm_agent_config_profile" "java_verbose" {
  // (Required) Display name of the profile.
  profile_name = "Configuration Profile - Java"

  // (Required) Type of APM Insight agent, for example JAVA, DOTNET, PHP,
  // RUBY, NODEJS or PYTHON.
  agent_type = "JAVA"

  // (Optional) Whether this profile is the default for its agent type.
  // Only one profile per agent type can be the default. Defaults to false.
  is_default = false

  // (Optional) Whether transaction traces are collected. Defaults to true.
  transaction_trace_enabled = true

  // (Optional) Tracing threshold in seconds. Defaults to 2.
  transaction_trace_threshold = 1

  // (Optional) Sampling factor: 1 tracks every request, 5 one in five.
  // Defaults to 1.
  transaction_tracking_request_interval = 1

  // (Optional) Whether SQL queries are captured. Defaults to true.
  sql_capture_enabled = true

  // (Optional) Whether captured SQL is obfuscated by replacing literal
  // values with parameters. Defaults to true.
  transaction_trace_sql_parametrize = true

  // (Optional) Slow SQL query threshold in seconds, above which a query gets
  // a stack trace. Defaults to 3.
  transaction_trace_sql_stacktrace_threshold = 2

  // (Optional) Whether agents upgrade themselves automatically.
  // Defaults to false.
  autoupgrade_enabled = false

  // (Optional) Whether instance names include the port number.
  // Defaults to true.
  show_instance_port_number = true

  // (Optional) Apdex threshold in seconds. Defaults to 0.5.
  apdex_threshold = 0.5

  // (Optional) Inactive days after which auto-suspended cloud instances are
  // deleted. Between 1 and 15. Defaults to 7.
  cloud_instance_cleanup_threshold = 7
}

// =========================================
// A default profile for a busy PHP fleet
// =========================================
resource "site24x7_apm_agent_config_profile" "php_default" {
  profile_name = "PHP baseline"
  agent_type   = "PHP"

  // Every new PHP agent picks this up.
  is_default = true

  // Sample one request in five.
  transaction_tracking_request_interval = 5

  autoupgrade_enabled = true

  // Reclaim auto-suspended cloud instances after three inactive days.
  cloud_instance_cleanup_threshold = 3
}

// =================================================
// Start from the settings Site24x7 ships for Java
// =================================================
resource "site24x7_apm_agent_config_profile" "java_staging" {
  profile_name = "Java - staging"
  agent_type   = "JAVA"

  transaction_trace_threshold = data.site24x7_apm_agent_config_profile.java_default.transaction_trace_threshold
  apdex_threshold             = data.site24x7_apm_agent_config_profile.java_default.apdex_threshold
}

// ========================
// Outputs
// ========================

output "java_default_profile_name" {
  description = "Name of the default Java configuration profile"
  value       = data.site24x7_apm_agent_config_profile.java_default.profile_name
}

output "php_baseline_sampling_factor" {
  description = "Sampling factor of the PHP baseline profile"
  value       = data.site24x7_apm_agent_config_profile.php_baseline.transaction_tracking_request_interval
}

output "all_profile_names_by_agent_type" {
  description = "Configuration profile names grouped by agent type"
  value = {
    for profile in data.site24x7_apm_agent_config_profiles.all.profiles :
    profile.profile_name => profile.agent_type
  }
}

output "fully_sampled_agent_types" {
  description = "Agent types whose profiles trace every request"
  value = [
    for profile in data.site24x7_apm_agent_config_profiles.all.profiles :
    profile.agent_type if profile.transaction_tracking_request_interval == 1
  ]
}
