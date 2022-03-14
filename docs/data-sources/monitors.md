---
layout: "site24x7"
page_title: "Site24x7: site24x7_monitors"
sidebar_current: "docs-site24x7-data-source-monitors"
description: |-
  Get information about monitors in Site24x7.
---

# Data Source: site24x7\_monitors

Use this data source to retrieve IDs and names of monitors in Site24x7.

## Example Usage

```hcl
// Data source to fetch all URL monitors
data "site24x7_monitors" "s247monitors" {
  // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
  monitor_type = "URL"
}
// Displays the monitor IDs
output "s247monitors_ids" {
  description = "Monitor IDs ============================ "
  value       = data.site24x7_monitors.s247monitors.ids
}
// Displays the monitor IDs and names of the monitors
output "s247monitors_ids_and_names" {
  description = "Monitor IDs and Names ============================ "
  value       = data.site24x7_monitors.s247monitors.ids_and_names
}

// Data source to fetch URL monitors starting with the name "zylker"
data "site24x7_monitors" "zylkerMonitorIDs" {
  // (Optional) Regular expression denoting the name of the monitor.
  name_regex = "^zylker"
  // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
  monitor_type = "URL"
}

// Displays the monitor IDs
output "zylkerMonitorIDs_monitor_id" {
  description = "Zylker Monitor IDs ============================ "
  value       = data.site24x7_monitors.zylkerMonitorIDs.ids
}

// Displays the monitor IDs and names of the monitors
output "zylkerMonitorIDs_monitor_id_and_names" {
  description = "Zylker Monitor IDs and Names ============================ "
  value       = data.site24x7_monitors.zylkerMonitorIDs.ids_and_names
}
```

## Attributes Reference


### Optional

* `name_regex` (String) Regular expression denoting the name of the monitor.
* `monitor_type` (String) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.

### Read-Only

* `ids` (List of String) List of monitor IDs.
* `ids_and_names` (List of String) List of monitor IDs and names separated by "__".