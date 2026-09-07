---
layout: "site24x7"
page_title: "Site24x7: site24x7_apm_application"
sidebar_current: "docs-site24x7-resource-apm-application"
description: |-
  Manage the lifecycle of an existing APM Insight application in Site24x7.
---

# Resource: site24x7\_apm\_application

Manages the lifecycle of an APM Insight application that already exists.

## This resource adopts, it does not create

The APM Insight API has no create or update endpoint. An application comes into existence when an APM Insight agent is deployed and first reports in; there is no way to conjure one from Terraform. What the API does offer is `manage`, `unmanage` and `delete`.

So this resource takes an `application_id` that already exists and brings its lifecycle under Terraform:

* **Create** verifies the application exists and reconciles its managed state. It does not create anything in Site24x7.
* **Update** flips between managed and suspended via the `manage` and `unmanage` endpoints. `managed` is the only attribute with a remote effect.
* **Delete** removes the resource from Terraform state and, only if you opt in with `delete_on_destroy`, deletes the application in Site24x7.

If you only need to read an application's attributes, use the [`site24x7_apm_application` data source](../data-sources/apm_application.md) instead.

!> **`delete_on_destroy = true` is irreversible.** `DELETE /apminsight/app/{id}` permanently removes the application and its accumulated monitoring history. It defaults to `false`, so an ordinary `terraform destroy` releases the application from state and leaves it running.

## Example Usage

```hcl
// Adopt an application an agent already registered, and keep it managed.
resource "site24x7_apm_application" "checkout" {
  application_id = "101071000000034001"
  managed        = true
}

// Suspend a legacy application without deleting its history.
resource "site24x7_apm_application" "legacy_batch" {
  application_id = "101071000000034999"
  managed        = false
}

// Resolve the ID by name rather than hardcoding it.
data "site24x7_apm_application" "billing" {
  name_regex = "^billing-api$"
}

resource "site24x7_apm_application" "billing" {
  application_id = data.site24x7_apm_application.billing.application_id
  managed        = true
}

output "checkout_instance_count" {
  value = site24x7_apm_application.checkout.instance_count
}
```

## Import

Existing applications can be imported by application ID:

```shell
terraform import site24x7_apm_application.checkout 101071000000034001
```

`delete_on_destroy` is not stored in Site24x7, so an imported resource starts with the safe default of `false`.

## Argument Reference

### Required

* `application_id` (String) ID of the APM Insight application to adopt. Changing this forces a new resource, which adopts a different application — it does not move or rename anything in Site24x7.

### Optional

* `managed` (Boolean) Whether the application should be managed (`true`) or suspended (`false`). Defaults to `true`.
* `delete_on_destroy` (Boolean) Whether `terraform destroy` should also delete the application in Site24x7. Defaults to `false`, which removes it from state only. Setting this to `true` makes destroy permanently delete the application and its monitoring history.
* `time_window` (String) Time window used when querying the API. Defaults to `H`. Only affects the request path.

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
* `under_maintenance` (Boolean) True when the application is currently under maintenance.

## What is deliberately absent

**Performance metrics.** The APM Insight API returns response times, apdex scores, throughput, error counts and CPU time for the requested time window. They are recomputed on every request, so keeping them in Terraform state would report drift on every refresh that no `apply` could ever resolve. This mirrors how other observability providers handle it — Datadog's provider, for instance, manages metric *metadata* and *tag configuration* but never metric values.

**An instance resource.** The instance endpoints support `manage`, `unmanage` and `delete` too, but instance IDs are not stable: autoscaling replaces instances and each new one gets a new ID, so a resource pinned to an instance ID would churn. Instances are exposed through the [`site24x7_apm_instance`](../data-sources/apm_instance.md) and [`site24x7_apm_instances`](../data-sources/apm_instances.md) data sources instead.
