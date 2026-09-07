---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_agent_config_profile"
sidebar_current: "docs-site24x7-datasource-apm-agent-config-profile"
description: |-
  Get information about an APM Insight agent configuration profile in Site24x7.
---

# Data Source: site24x7\_apm\_agent\_config\_profile

Reads a single APM Insight agent configuration profile — by ID, by name, or the default profile of an agent type.

## Example Usage

```hcl
// The default profile every new Java agent picks up.
data "site24x7_apm_agent_config_profile" "java_default" {
  agent_type      = "JAVA"
  default_profile = true
}

// A specific profile by name, searched within one agent type.
data "site24x7_apm_agent_config_profile" "php_baseline" {
  agent_type = "PHP"
  name_regex = "^PHP baseline$"
}

// A profile whose ID you already know.
data "site24x7_apm_agent_config_profile" "by_id" {
  profile_id = "1000000009001"
}

output "java_default_trace_threshold" {
  value = data.site24x7_apm_agent_config_profile.java_default.transaction_trace_threshold
}
```

## Argument Reference

Exactly one of `profile_id`, `name_regex` or `default_profile` must be set.

* `profile_id` (String) ID of the configuration profile to look up. The API has no get-by-ID endpoint, so this lists the profiles and filters.
* `name_regex` (String) Regular expression matched against profile names. It is an error for the expression to match more than one profile — set `agent_type` as well to narrow the search.
* `default_profile` (Boolean) Set to `true` to look up the default profile of `agent_type`. Requires `agent_type`.
* `agent_type` (String) Type of APM Insight agent, for example `JAVA`, `DOTNET`, `PHP`, `RUBY`, `NODEJS` or `PYTHON`. Required with `default_profile`, optional with `name_regex` to narrow the search, and ignored with `profile_id`.

## Attributes Reference

* `id` (String) The profile ID.
* `profile_id` (String) The profile ID.
* `profile_name` (String) Display name of the profile.
* `agent_type` (String) Type of APM Insight agent the profile configures.
* `is_default` (Boolean) True when this profile is the default configuration for its agent type.
* `transaction_trace_enabled` (Boolean) Whether transaction traces are collected. Maps to `transaction.trace.enabled`.
* `transaction_trace_threshold` (Number) Tracing threshold in seconds. Maps to `transaction.trace.threshold`.
* `transaction_trace_sql_parametrize` (Boolean) Whether SQL queries in traces are obfuscated. Maps to `transaction.trace.sql.parametrize`.
* `transaction_trace_sql_stacktrace_threshold` (Number) Slow SQL query threshold in seconds. Maps to `transaction.trace.sql.stacktrace.threshold`.
* `transaction_tracking_request_interval` (Number) Web transaction sampling factor. Maps to `transaction.tracking.request.interval`.
* `sql_capture_enabled` (Boolean) Whether SQL queries are captured. Maps to `sql.capture.enabled`.
* `autoupgrade_enabled` (Boolean) Whether agents using this profile upgrade themselves automatically. Maps to `autoupgrade.enabled`.
* `show_instance_port_number` (Boolean) Whether instance names include the port number. Maps to `show.instance.port.number`.
* `apdex_threshold` (Number) Apdex threshold in seconds. Maps to `apdex.threshold`.
* `cloud_instance_cleanup_threshold` (Number) Number of inactive days after which auto-suspended cloud instances are deleted. Maps to `cloud.instance.cleanup.threshold`.
* `last_modified_time` (String) When the profile was last updated, in milliseconds since the epoch. Maps to `last.modified.time`.

## Related

* [`site24x7_apm_agent_config_profile` resource](../resources/apm_agent_config_profile.md) — create and manage a profile.
* [`site24x7_apm_agent_config_profiles` data source](apm_agent_config_profiles.md) — read every profile at once.
