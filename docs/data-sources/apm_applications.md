---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_applications"
sidebar_current: "docs-site24x7-data-source-apm-applications"
description: |-
  Get information about multiple APM Insight applications in Site24x7.
---

# Data Source: site24x7\_apm\_applications

Use this data source to retrieve a list of APM Insight applications, optionally filtered by name or managed state.

Unlike [`site24x7_apm_application`](apm_application.md), an empty result is not an error here — you asked for whatever matches, and nothing matching is a valid answer.

~> **Performance metrics are not exposed.** See the note on [`site24x7_apm_application`](apm_application.md) for why.

## Example Usage

```hcl
// Every APM Insight application whose name starts with "prod-".
data "site24x7_apm_applications" "production" {
  name_regex = "^prod-"
}

// Only applications that are currently suspended.
data "site24x7_apm_applications" "suspended" {
  managed_state = "unmanaged"
}

output "production_application_ids" {
  description = "IDs of the production APM applications"
  value       = data.site24x7_apm_applications.production.ids
}

// The `applications` attribute carries more than the IDs, so you can branch
// on it without a second lookup per application.
output "production_application_names" {
  description = "Names of the production APM applications"
  value       = [for app in data.site24x7_apm_applications.production.applications : app.application_name]
}
```

## Attributes Reference

### Optional

* `name_regex` (String) Regular expression matched against application names. When omitted, every APM Insight application is returned.
* `managed_state` (String) Filter by managed state. Accepts `managed`, `unmanaged` or `any`. Defaults to `any`.
* `time_window` (String) Time window used when querying the API, for example `H` for the last hour. Defaults to `H`.

### Read-Only

* `id` (String) The ID of this resource, derived from the filters.
* `ids` (List of String) IDs of the matching applications, sorted.
* `ids_and_names` (List of String) Matching applications as `<id>__<name>`, sorted by ID.
* `applications` (List of Object) Matching applications, sorted by application ID. Each entry has `application_id`, `application_name`, `instance_count`, `host_count`, `availability`, `managed_state` and `under_maintenance`.
