---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_agent_config_profile"
sidebar_current: "docs-site24x7-resource-apm-agent-config-profile"
description: |-
  Create and manage an APM Insight agent configuration profile in Site24x7.
---

# Resource: site24x7\_apm\_agent\_config\_profile

Creates and manages an APM Insight agent configuration profile: a named set of agent settings — tracing thresholds, SQL capture, sampling, auto-upgrade — applied to the instances of one agent type.

Unlike [`site24x7_apm_application`](apm_application.md), which can only adopt an application an agent brought into existence, profiles have full create, update and delete endpoints, so this is an ordinary Terraform resource.

## Attribute names map to the API's dotted keys

The API nests the settings in an `agent_config` object whose keys are dotted. Each one is exposed here as a top-level attribute named after the wire key with the dots replaced by underscores, so `transaction.trace.enabled` is `transaction_trace_enabled`. Nothing else is renamed.

Every setting except `cloud_instance_cleanup_threshold` is mandatory in the API, so each attribute defaults to the value Site24x7 uses in its own default profiles and is always sent — you only declare what you want to differ.

## Example Usage

```hcl
// A profile for Java agents with tighter tracing than the default.
resource "site24x7_apm_agent_config_profile" "java_verbose" {
  profile_name = "Configuration Profile - Java"
  agent_type   = "JAVA"

  // Trace anything slower than one second, and sample every request.
  transaction_trace_enabled             = true
  transaction_trace_threshold           = 1
  transaction_tracking_request_interval = 1

  // Capture SQL, obfuscated, with stack traces for slow queries.
  sql_capture_enabled                        = true
  transaction_trace_sql_parametrize          = true
  transaction_trace_sql_stacktrace_threshold = 2

  apdex_threshold = 0.5
}

// Make a profile the default that every new PHP agent picks up.
resource "site24x7_apm_agent_config_profile" "php_default" {
  profile_name = "PHP baseline"
  agent_type   = "PHP"
  is_default   = true

  // Sample one request in five on a busy fleet.
  transaction_tracking_request_interval = 5

  autoupgrade_enabled = true

  // Delete auto-suspended cloud instances after three inactive days.
  cloud_instance_cleanup_threshold = 3
}

// Start from the settings Site24x7 ships for Java.
data "site24x7_apm_agent_config_profile" "java_default" {
  agent_type      = "JAVA"
  default_profile = true
}

resource "site24x7_apm_agent_config_profile" "java_staging" {
  profile_name = "Java - staging"
  agent_type   = "JAVA"

  transaction_trace_threshold = data.site24x7_apm_agent_config_profile.java_default.transaction_trace_threshold
  apdex_threshold             = data.site24x7_apm_agent_config_profile.java_default.apdex_threshold
}
```

## Import

Existing profiles can be imported by profile ID:

```shell
terraform import site24x7_apm_agent_config_profile.java_verbose 1000000009001
```

## Argument Reference

### Required

* `profile_name` (String) Display name of the configuration profile.
* `agent_type` (String) Type of APM Insight agent the profile configures, for example `JAVA`, `DOTNET`, `PHP`, `RUBY`, `NODEJS` or `PYTHON`. Not validated locally, so agent types Site24x7 adds after this provider release can still be used.

### Optional

* `is_default` (Boolean) Whether this profile is the default configuration for its agent type. Defaults to `false`. See the caveat below.
* `transaction_trace_enabled` (Boolean) Whether transaction traces are collected. Defaults to `true`. Maps to `transaction.trace.enabled`.
* `transaction_trace_threshold` (Number) Tracing threshold in seconds: transactions slower than this are traced. Defaults to `2`. Maps to `transaction.trace.threshold`.
* `transaction_trace_sql_parametrize` (Boolean) Whether SQL queries in traces are obfuscated by replacing literal values with parameters. Defaults to `true`. Maps to `transaction.trace.sql.parametrize`.
* `transaction_trace_sql_stacktrace_threshold` (Number) Slow SQL query threshold in seconds: queries slower than this get a stack trace. Defaults to `3`. Maps to `transaction.trace.sql.stacktrace.threshold`.
* `transaction_tracking_request_interval` (Number) Web transaction sampling factor: `1` tracks every request, `5` tracks one in five. Defaults to `1`. Maps to `transaction.tracking.request.interval`.
* `sql_capture_enabled` (Boolean) Whether SQL queries are captured. Defaults to `true`. Maps to `sql.capture.enabled`.
* `autoupgrade_enabled` (Boolean) Whether agents using this profile upgrade themselves automatically. Defaults to `false`. Maps to `autoupgrade.enabled`.
* `show_instance_port_number` (Boolean) Whether instance names include the port number. Defaults to `true`. Maps to `show.instance.port.number`.
* `apdex_threshold` (Number) Apdex threshold in seconds, the response time below which a request counts as satisfying. Defaults to `0.5`. Maps to `apdex.threshold`.
* `cloud_instance_cleanup_threshold` (Number) Number of inactive days after which auto-suspended cloud instances are deleted. Must be between `1` and `15`. Defaults to `7`. Maps to `cloud.instance.cleanup.threshold`.

### Read-Only

* `id` (String) The ID of this resource, which is the profile ID.
* `last_modified_time` (String) When the profile was last updated, in milliseconds since the epoch. Maps to `last.modified.time`, which the API maintains and this provider never sends.

## Only one profile per agent type is the default

Setting `is_default = true` demotes whichever profile was previously the default for that agent type. Terraform cannot see that side effect: the other profile's state still says `is_default = true` until it is refreshed, and the next plan will then propose promoting it back. If you manage several profiles for one agent type, keep `is_default = true` on exactly one of them.

## Related

* [`site24x7_apm_agent_config_profile` data source](../data-sources/apm_agent_config_profile.md) — read one profile, including the default for an agent type.
* [`site24x7_apm_agent_config_profiles` data source](../data-sources/apm_agent_config_profiles.md) — read every profile, or every profile of one agent type.
