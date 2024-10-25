---
layout: "site24x7"
page_title: "Site24x7: site24x7_soap_monitor"
sidebar_current: "docs-site24x7-resource-port-monitor"
description: |-
  Create and manage a SOAP monitor in Site24x7.
---

# Resource: site24x7\_soap\_monitor

Use this resource to create, update and delete a SOAP monitor in Site24x7.

## Example Usage

```hcl

// Site24x7 SOAP Monitor API doc - https://www.site24x7.com/help/api/#SOAP
resource "site24x7_soap_monitor" "soap_monitor_example" {

  // (Required) (String)Display name for the monitor
  display_name = "SOAP Monitor"

  //(Required)(String)Host name of the monitor
  website="www.example.com"

  //(Optional)(Number)
  //0	Down
  //2	Trouble
  soap_attributes_severity=0

  //(Optional)(Map of String)
  //SOAP attribute name and value in a string array
  soap_attributes={
    "name" = "soap attribute name"
    "value" = "soap attribute value"
  }

  //(Optional)(Boolean)
  //Select IPv6 for monitoring the websites hosted with IPv6 //address. If 
  //you choose non IPv6 supported locations, monitoring will happen through
  //IPv4.
  use_ipv6=true

  //(Optional)(Boolean)
  //Enable ALPN to send supported protocols as part of the TLS handshake.
  use_alpn=true

  //(Optional)(String)
  //Specify the version of the SSL protocol. If you are not sure about the version, use Auto.
  //Default value is Auto
  ssl_protocol="TLSv1.2"

  //(Optional)
  //Check whether the HTTP response headers are present or verify //header and corresponding values against predefined header and //values. 
  //Trigger down or trouble alerts during failure.
  //JSON Format: {value: [{name: “$Header Name”, value: “$Header Value”}], severity: “$alert_type_constant”}
  response_headers_check="{\"value\":[{\"name\":\"HeaderName\",\"value\":\"HeaderValue\"}],\"severity\":2}"

  //(Optional)(String)
  //Provide content type for request params.
  request_content_type=""

  //(Optional)(String)
  //HTTP Method to be used for accessing the website.
  HEAD, PUT, PATCH and DELETE are not supported
  http_method=""

  //(Optional)(Boolean)
  //Resolve the IP address using Domain Name Server.
  use_name_server=false

  //(Optional)(String)
  //Specify the version of the HTTP protocol.
  http_protocol=""

  //(Optional)(String)
  //Response content type. Response Content Types
  response_type=""

  //(Optional)(String)
  //The frequency or interval for monitoring.
  check_frequency="10"

  //(Optional)(String)
  //Web Credential Profile to be associated.
  //Add a new profile or find the ID of your preferred Credential Profile.
  credential_profile_id=""

  //(Optional)(String)
  //Password of the uploaded client certificate.
  client_certificate_password="abc"

  //(Optional)(Number)
  //(Optional)Timeout for connecting to the host.Range 1 - 45.
  timeout=10 

  //(Optional)(String)
  //Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.
  up_status_codes="

  //(Optional)(Boolean)Toggle button to perform automation or not
  perform_automation=true

  //(Optional)(String)if user_group_ids is not choosen
  //On-Call Schedule of your choice.
  //Create new On-Call Schedule or find your preferred On-Call Schedule ID.
  on_call_schedule_id="456418000001258016"

  // (Optional)(Map of String) Map of status to actions that should be performed on monitor
  // status changes. See
  // https://www.site24x7.com/help/api/#action-rule-constants for all available
  // status values.
  actions = {1=465545643755}

  //(Optional)(String)
  //Threshold profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/threshold_profiles
  // endpoint for the HEARTBEAT monitor type (https://www.site24x7.com/help/api/#list-threshold-profiles) will
  // be used.
  threshold_profile_id = "456418000342341"

  // (Optional)(List of String) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]

  //(Optional)(String)
  //Location Profile to be associated with the monitor. If 
  // location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_id = "123"

  //(Optional)(String)
  //Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"
  
  //(Optional)(String)
  //Notification profile to be associated with the monitor. If
  // omitted, the first profile returned by the /api/notification_profiles
  // endpoint (https://www.site24x7.com/help/api/#list-notification-profiles)
  // will be used.
  notification_profile_id = "123"

  //(Optional)(String)
  //Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  //(Optional)(List of String)
  //List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    "123",
    "456"
  ]

  //(Optional)(List of String)
  //List of dependent resource IDs. Suppress alert  when dependent monitor(s) is down.
  dependency_resource_ids = [
      "123",
      "456"
    ]

  //(Optional)(List of String)
  // List if user group IDs to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be 
  //used.
  user_group_ids = [
    "123",
  ]

  //(Optional)(List of String)
  //  List if user group names to be notified on down. 
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]

  //(Optional)(List of String)
  // List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]

  //(Optional)(List of String)
  // List of tag names to be associated to the monitor. Tag name matching works for both exact and 
  //  partial match. Either specify tag_ids or tag_names.
  tag_names = [
    "Terraform",
    "Server",
  ]
}
```
## Attributes Reference

### Required
* `display_name` (String) Display Name for the monitor.
* `website`(String)Registered domain name.
### Optional
* `id` (String) The ID of this resource.
* `type` DOMAINEXPIRY
* `domain_name`(String)Who is server.
* `timeout`(int) Timeout for connecting to the host.
* `use_ipv6`(bool) Prefer IPV6
* `check_frequency` Check interval for monitoring.
* `ssl_protocol` (String) Specify the version of the SSL protocol. If you are not sure about the version, use Auto.
* `perform_automation` (bool) Automating the scheduled maintenance
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `on_call_schedule_id` (String) Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.

Refer [API documentation](https://www.site24x7.com/help/api/#SOAP) for more information about attributes.
