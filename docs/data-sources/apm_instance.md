---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_instance"
sidebar_current: "docs-site24x7-data-source-apm-instance"
description: |-
  Get information about an APM Insight instance in Site24x7.
---

# Data Source: site24x7\_apm\_instance

Use this data source to retrieve information about a single APM Insight instance.

An instance is one reporting agent: in Java each JVM, in .NET each app in an IIS server, in Ruby each Rails server, and in PHP each PHP server.

~> **Instance IDs are not stable across autoscaling.** New instances get new IDs when they come up, so pinning Terraform configuration to a specific instance ID works only for long-lived hosts. For a fleet that scales, filter with [`site24x7_apm_instances`](apm_instances.md) instead.

~> **Performance metrics are not exposed.** See the note on [`site24x7_apm_application`](apm_application.md) for why.

## Example Usage

```hcl
data "site24x7_apm_instance" "app_server" {
  instance_id = "101071000000034035"
}

// Or by name, which is usually "host:port".
data "site24x7_apm_instance" "by_name" {
  name_regex = "^192\\.168\\.1\\.1:8080$"
}

output "agent_version" {
  description = "APM Insight agent version deployed at this instance"
  value       = data.site24x7_apm_instance.app_server.agent_version
}
```

## Attributes Reference

### Optional

Exactly one of `instance_id` or `name_regex` must be set.

* `instance_id` (String) ID of the APM Insight instance to look up.
* `name_regex` (String) Regular expression matched against instance names. It is an error for the expression to match more than one instance.
* `time_window` (String) Time window used when querying the API. Defaults to `H`.

### Read-Only

* `id` (String) The ID of this resource, which is the instance ID.
* `instance_name` (String) Name of the instance, usually `host:port`.
* `application_id` (String) ID of the application this instance reports into.
* `application_name` (String) Name of the application this instance reports into.
* `host` (String) Host address the instance reports from.
* `port` (Number) Port number of the instance.
* `ins_type` (String) Type of APM Insight agent, for example `JAVA`, `PHP`, `RUBY`, `DOTNET` or `NODEJS`.
* `agent_version` (String) Version of the APM Insight agent deployed at this instance.
* `availability` (String) Current availability status, for example `AVAILABLE` or `NOTAVAILABLE`.
* `managed_state` (Boolean) True when the instance is managed rather than suspended.
* `under_maintenance` (Boolean) True when the instance is currently under maintenance.
* `is_cloud_app` (Boolean) True when the instance is hosted in a cloud environment such as AWS or Azure.
* `auto_scale` (Boolean) True when autoscaling is enabled for the instance.
