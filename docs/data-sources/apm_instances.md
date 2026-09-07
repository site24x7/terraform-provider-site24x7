---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_instances"
sidebar_current: "docs-site24x7-data-source-apm-instances"
description: |-
  Get information about multiple APM Insight instances in Site24x7.
---

# Data Source: site24x7\_apm\_instances

Use this data source to retrieve a list of APM Insight instances, optionally filtered by name, application, agent type or managed state.

This is the right data source for a fleet that autoscales: filter by application and agent type rather than pinning to instance IDs that change when hosts are replaced.

~> **Performance metrics are not exposed.** See the note on [`site24x7_apm_application`](apm_application.md) for why.

## Example Usage

```hcl
// Every Java instance reporting into one application.
data "site24x7_apm_instances" "checkout_jvms" {
  application_id = data.site24x7_apm_application.checkout.application_id
  ins_type       = "JAVA"
}

// Find agents that are still on an old version, across the whole account.
data "site24x7_apm_instances" "all" {}

output "outdated_agents" {
  description = "Instances still running an agent older than 2.0"
  value = [
    for instance in data.site24x7_apm_instances.all.instances :
    instance.instance_name if tonumber(instance.agent_version) < 2.0
  ]
}

output "checkout_jvm_ids" {
  value = data.site24x7_apm_instances.checkout_jvms.ids
}
```

## Attributes Reference

### Optional

* `name_regex` (String) Regular expression matched against instance names. When omitted, every APM Insight instance is returned.
* `application_id` (String) Only return instances reporting into this application.
* `ins_type` (String) Only return instances running this agent type, for example `JAVA` or `PHP`. Matched case-insensitively.
* `managed_state` (String) Filter by managed state. Accepts `managed`, `unmanaged` or `any`. Defaults to `any`.
* `time_window` (String) Time window used when querying the API. Defaults to `H`.

### Read-Only

* `id` (String) The ID of this resource, derived from the filters.
* `ids` (List of String) IDs of the matching instances, sorted.
* `ids_and_names` (List of String) Matching instances as `<id>__<name>`, sorted by ID.
* `instances` (List of Object) Matching instances, sorted by instance ID. Each entry has `instance_id`, `instance_name`, `application_id`, `application_name`, `host`, `port`, `ins_type`, `agent_version`, `availability`, `managed_state`, `under_maintenance`, `is_cloud_app` and `auto_scale`.
