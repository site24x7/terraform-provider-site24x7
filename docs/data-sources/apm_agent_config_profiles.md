---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_agent_config_profiles"
sidebar_current: "docs-site24x7-datasource-apm-agent-config-profiles"
description: |-
  Get information about all APM Insight agent configuration profiles in Site24x7.
---

# Data Source: site24x7\_apm\_agent\_config\_profiles

Reads every APM Insight agent configuration profile, optionally narrowed to one agent type or to names matching an expression. An empty result is not an error.

## Example Usage

```hcl
// Every profile of every agent type.
data "site24x7_apm_agent_config_profiles" "all" {}

// Only the Java profiles.
data "site24x7_apm_agent_config_profiles" "java" {
  agent_type = "JAVA"
}

// Every profile whose name starts with "Default profile".
data "site24x7_apm_agent_config_profiles" "defaults" {
  name_regex = "^Default profile"
}

output "java_profile_names" {
  value = [for profile in data.site24x7_apm_agent_config_profiles.java.profiles : profile.profile_name]
}

// Which agent types trace every request.
output "fully_sampled_agent_types" {
  value = [
    for profile in data.site24x7_apm_agent_config_profiles.all.profiles :
    profile.agent_type if profile.transaction_tracking_request_interval == 1
  ]
}
```

## Argument Reference

* `agent_type` (String) Return only the profiles of this agent type, for example `JAVA`. When omitted, the profiles of every agent type are returned. Setting it uses the API's per-agent-type endpoint rather than filtering the full list.
* `name_regex` (String) Regular expression matched against profile names. When omitted, every profile is returned.

## Attributes Reference

* `id` (String) A hash of the arguments used for the lookup.
* `ids` (List of String) IDs of the matching profiles, sorted.
* `ids_and_names` (List of String) Matching profiles as `<id>__<name>`, sorted by ID.
* `profiles` (List of Object) Matching profiles with their full configuration, sorted by profile ID. Each entry carries the same attributes as the [`site24x7_apm_agent_config_profile` data source](apm_agent_config_profile.md): `profile_id`, `profile_name`, `agent_type`, `is_default`, `transaction_trace_enabled`, `transaction_trace_threshold`, `transaction_trace_sql_parametrize`, `transaction_trace_sql_stacktrace_threshold`, `transaction_tracking_request_interval`, `sql_capture_enabled`, `autoupgrade_enabled`, `show_instance_port_number`, `apdex_threshold`, `cloud_instance_cleanup_threshold` and `last_modified_time`.

Results are sorted by profile ID because the API does not promise an ordering, so an unsorted list would show a diff on every refresh.

## Related

* [`site24x7_apm_agent_config_profile` resource](../resources/apm_agent_config_profile.md) — create and manage a profile.
* [`site24x7_apm_agent_config_profile` data source](apm_agent_config_profile.md) — read one profile, including the default for an agent type.
