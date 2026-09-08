---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_application"
sidebar_current: "docs-site24x7-data-source-apm-application"
description: |-
  Get information about an APM Insight application in Site24x7.
---

# Data Source: site24x7\_apm\_application

Use this data source to retrieve information about a single APM Insight application.

APM Insight applications are created when an agent first reports in, not through the API, so this data source is the usual way to reference one from Terraform: look the application up by name or ID, then wire its ID into resources you do manage.

~> **Performance metrics are not exposed.** The APM Insight API returns response times, apdex scores, throughput and error counts for the requested time window. Those are recomputed on every request, so storing them in Terraform state would report drift on every refresh that no `apply` could resolve. This data source exposes identity and configuration only.

## Example Usage

```hcl
// Look up an APM Insight application by name.
data "site24x7_apm_application" "checkout" {
  name_regex = "^checkout-api$"
}

// Or by ID, if you already know it.
data "site24x7_apm_application" "billing" {
  application_id = "101071000000034001"
}

// Tag the application alongside the rest of your monitors.
resource "site24x7_tag" "apm_owner" {
  tag_name  = "team-payments"
  tag_value = data.site24x7_apm_application.checkout.application_name
}

output "checkout_instance_ids" {
  description = "Instances reporting into the checkout application"
  value       = data.site24x7_apm_application.checkout.instance_ids
}

output "checkout_rum_app_id" {
  description = "Linked RUM application, empty when none is linked"
  value       = data.site24x7_apm_application.checkout.rum_app_id
}
```

## Attributes Reference

### Optional

Exactly one of `application_id` or `name_regex` must be set.

* `application_id` (String) ID of the APM Insight application to look up.
* `name_regex` (String) Regular expression matched against application names. It is an error for the expression to match more than one application; refine it until it matches exactly one.
* `time_window` (String) Time window used when querying the API, for example `H` for the last hour. Defaults to `H`. This only affects the request path — the attributes read here are the same for every window.

### Read-Only

* `id` (String) The ID of this resource, which is the application ID.
* `application_name` (String) Name of the APM Insight application.
* `instance_ids` (List of String) IDs of all instances reporting into this application, sorted.
* `instance_count` (Number) Number of instances reporting into this application.
* `host_count` (Number) Number of hosts running this application.
* `hosts` (List of String) Host addresses running this application, sorted.
* `instances` (List of Object) Instances reporting into this application, sorted by instance ID. Each entry has `instance_id`, `instance_name`, `host`, `port`, `ins_type` and `agent_version`.
* `rum_app_id` (String) ID of the linked Real User Monitoring application, empty when none is linked.
* `rum_app_key` (String) Key of the linked RUM application. This key is embedded in browser-side JavaScript and is not a secret.
* `rum_app_name` (String) Name of the linked RUM application.
* `availability` (String) Current availability status, for example `AVAILABLE` or `NOTAVAILABLE`.
* `managed_state` (Boolean) True when the application is managed rather than suspended.
* `under_maintenance` (Boolean) True when the application is currently under maintenance.
