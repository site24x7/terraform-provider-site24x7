---
layout: "site24x7"
page_title: "Site24x7: site24x7_rest_api_monitor"
sidebar_current: "docs-site24x7-resource-rest-api-monitor"
description: |-
  Create and manage a REST API monitor in Site24x7.
---

# Resource: site24x7\_rest\_api\_monitor

Use this resource to create, update and delete a REST API monitor in Site24x7.

## Example Usage

```hcl

// Site24x7 Rest API Monitor API doc - https://www.site24x7.com/help/api/#rest-api
resource "site24x7_rest_api_monitor" "rest_api_monitor_us" {
  // (Required) Display name for the monitor
  display_name = "REST API Monitor - terraform"
  // (Required) Website address to monitor.
  website = "https://dummy.restapiexample.com/"
  // (Optional) Name of the Location Profile that has to be associated with the monitor.
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"
  // (Optional) Name of the notification profile that has to be associated with the monitor.
  // Profile name matching works for both exact and partial match.
  // Either specify notification_profile_id or notification_profile_name.
  // If notification_profile_id and notification_profile_name are omitted,
  // the first profile returned by the /api/notification_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-notification-profiles) will be
  // used.
  notification_profile_name = "Terraform Profile"

  // (Optional) List of monitor group IDs to associate the monitor to.
  monitor_groups = [
    "123",
    "456"
  ]

  // (Optional) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
  dependency_resource_ids = [
    "123",
    "456"
  ]

  // (Optional) List if user group IDs to be notified on down.
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_ids = [
    "123",
  ]

  // (Optional) List if user group names to be notified on down.
  // Either specify user_group_ids or user_group_names. If omitted, the
  // first user group returned by the /api/user_groups endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-user-groups) will be used.
  user_group_names = [
    "Terraform",
    "Network",
    "Admin",
  ]

  // (Optional) List if tag IDs to be associated to the monitor.
  tag_ids = [
    "123",
  ]

  // (Optional) List of tag names to be associated to the monitor. Tag name matching works for both exact and
  //  partial match. Either specify tag_ids or tag_names.
  tag_names = [
    "Terraform",
    "Network",
  ]

  // (Optional) List of Third Party Service IDs to be associated to the monitor.
  third_party_service_ids = [
    "4567"
  ]

  // (Optional) Map of custom HTTP headers to send.
  custom_headers = {
    "Accept" = "application/json"
  }

  // (Optional) Check for the keyword in the website response.
  matching_keyword = {
 	  severity= 2
 	  value= "aaa"
 	}

  // (Optional) Check for non existence of keyword in the website response.
  unmatching_keyword = {
 	  severity= 2
 	  value= "bbb"
 	}

  // (Optional) Match the regular expression in the website response.
  match_regex = {
 	  severity= 2
 	  value= ".*aaa.*"
 	}

  // (Optional) Map of HTTP response headers to check.
  response_headers_severity = 0 // Can take values 0 or 2. '0' denotes Down and '2' denotes Trouble.
  response_headers = {
    "Content-Encoding" = "gzip"
    "Connection" = "Keep-Alive"
  }

  //(Optional) Credential Profile to associate the website with 
  credential_profile_id = "123"

  // HTTP Configuration

  //(Optional) Credential Profile to associate the website with 
  credential_profile_id = "123"

  // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response.
  // You can specify individual status codes, as well as ranges separated with a colon.
  up_status_codes = "400:500"

  // ================ JSON ASSERTION ATTRIBUTES
  // (Optional) Response content type. Default value is 'T'
  // 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML
  // https://www.site24x7.com/help/api/#res_content_type
  response_content_type = "J"
  // (Optional) Provide multiple JSON Path expressions to enable evaluation of JSON Path expression assertions.
  // The assertions must successfully parse the JSON Path in the JSON. JSON expression assertions fails if the expressions does not match.
  match_json_path = [
    "$.store.book[*].author",
    "$..author",
    "$.store.*"
  ]
  // (Optional) Trigger an alert when the JSON path assertion fails during a test.
  // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
  match_json_path_severity = 0
  // (Optional) JSON schema to be validated against the JSON response.
  json_schema = "{\"test\":\"abcd\"}"
  // (Optional) Trigger an alert when the JSON schema assertion fails during a test.
  // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
  json_schema_severity = 2
  // (Optional) JSON Schema check allows you to annotate and validate all JSON endpoints for your web service.
  json_schema_check = true
  // JSON ASSERTION ATTRIBUTES ================

  // ================ HTTP POST with request body
  // (Optional) HTTP Method to be used for accessing the website.  Default value is 'G'. 'G' denotes GET, 'P' denotes POST, 'U' denotes PUT and 'D' denotes DELETE. HEAD is not supported.
  http_method = "P"
  // (Optional) Provide content type for request params when http_method is 'P'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML and 'F' denotes FORM
  request_content_type = "J"
  // (Optional) Provide the content to be passed in the request body while accessing the website.
  request_body = "{\"user_name\":\"joe\"}"
  // (Optional) Map of custom HTTP headers to send.
  request_headers = {
    "Accept" = "application/json"
  }
  // HTTP POST with request body ================

  // ================ GRAPHQL ATTRIBUTES
  // (Optional) Provide content type for request params.
  request_content_type = "G"
  // (Optional) Provide the GraphQL query to get specific response from GraphQL based API service. request_content_type = "G"
  graphql_query = "query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}"
  // (Optional) Provide the GraphQL variables to get specific response from GraphQL based API service. request_content_type = "G"
  graphql_variables = "{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}"
  // GRAPHQL ATTRIBUTES ================
}

```

## Attributes Reference


### Required

* `display_name` (String) Display Name for the monitor.
* `website` (String) Website address to monitor.

### Optional

* `id` (String) The ID of this resource.
* `check_frequency` (String) Interval at which your website has to be monitored. Default value is 1 minute.
* `timeout` (Number) Timeout for connecting to website. Default value is 10. Range 1 - 45.
* `use_ipv6` (Boolean) Select IPv6 for monitoring the websites hosted with IPv6 address. If you choose non IPv6 supported locations, monitoring will happen through IPv4.
* `http_method` (String) HTTP Method to be used for accessing the website. Default value is 'G'. 'G' denotes GET, 'P' denotes POST, 'U' denotes PUT and 'D' denotes DELETE. HEAD is not supported.
* `request_content_type` (String) Provide content type for request params when http_method is 'P'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML, 'F' denotes FORM and 'G' denotes GRAPHQL.
* `request_body` (String) Provide the content to be passed in the request body while accessing the website.
* `request_headers` (Map of String) A Map of request header name and value.
* `graphql_query` (String) Provide the GraphQL query to get specific response from GraphQL based API service. request_content_type should be "G"
* `graphql_variables` (String) Provide the GraphQL variables to get specific response from GraphQL based API service. request_content_type should be "G"
* ~~`request_param` (String) Provide parameters to be passed while accessing the website.~~ (Deprecated: https://github.com/site24x7/terraform-provider-site24x7/pull/94/files#diff-48dba37a89bbad21af6c4d8b66fd20583aadfca584594b57793cdd14f4d6330fL262)
* `ssl_protocol` (String) Specify the version of the SSL protocol. If you are not sure about the version, use Auto.
* `use_alpn` (Boolean) Enable ALPN to send supported protocols as part of the TLS handshake.
* `http_method` (String) HTTP Method to be used for accessing the website.  Default value is 'G'. 'G' denotes GET, 'P' denotes POST, 'U' denotes PUT and 'D' denotes DELETE. HEAD is not supported.
* `http_protocol` (String) Specify the version of the HTTP protocol. Default value is H1.1.
* `client_certificate_password` (String) Password of the uploaded client certificate.
* `auth_pass` (String) Authentication user name to access the website.
* `auth_user` (String) Authentication password to access the website.
* `credential_profile_id` (String)Credential Profile to associate the website with. Notes: If you're using Auth user and Auth password, you can't configure Credential Profile
* `oauth2_provider` (String) Provider ID of the OAuth Provider to be associated with the monitor.
* `jwt_id` (String) Token ID of the Web Token to be associated with the monitor.
* `up_status_codes` (String) Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.
* `use_name_server` (Boolean) Resolve the IP address using Domain Name Server.
* `user_agent` (String) User Agent to be used while monitoring the website.
* `response_content_type` (String) Response content type. Default value is 'T'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML. Refer [API documentation](https://www.site24x7.com/help/api/#res_content_type) for more information.
* `match_json_path` (List of String) Provide multiple JSON Path expressions to enable evaluation of JSON Path expression assertions. The assertions must successfully parse the JSON Path in the JSON. JSON expression assertions fails if the expressions does not match.
* `match_json_path_severity` (Number) Trigger an alert when the JSON path assertion fails during a test. Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
* `json_schema` (String) JSON schema to be validated against the JSON response.
* `json_schema_severity` (Number) Trigger an alert when the JSON schema assertion fails during a test. Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
* `json_schema_check` (Boolean) Enable this option to perform the JSON schema check. JSON Schema check allows you to annotate and validate all JSON endpoints for your web service.
* `matching_keyword` (Map of String) Check for the keyword in the website response.
* `unmatching_keyword` (Map of String) Check for non existence of keyword in the website response.
* `match_case` (Boolean) Perform case sensitive keyword search or not.
* `match_regex` (Map of String) Match the regular expression in the website response.
* `response_headers` (Map of String) A Map of Header name and value.
* `response_headers_severity` (Number) Severity with which alert has to raised when the header is found in the website response. Default value is 2. '0' denotes Down and '2' denotes Trouble.
* `location_profile_id` (String) Location profile to be associated with the monitor.
* `location_profile_name` (String) Name of the location profile to be associated with the monitor.
* `notification_profile_id` (String) Notification profile to be associated with the monitor. Either specify notification_profile_id or notification_profile_name. If notification_profile_id and notification_profile_name are omitted, the first profile returned by the /api/notification_profiles endpoint will be used.
* `notification_profile_name` (String) Name of the notification profile to be associated with the monitor. Profile name matching works for both exact and partial match.
* `threshold_profile_id` (String) Threshold profile to be associated with the monitor.
* `monitor_groups` (List of String) List of monitor groups to which the monitor has to be associated.
* `dependency_resource_ids` (List of String) List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.
* `user_group_ids` (List of String) List of user groups to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `user_group_names` (List of String) List of user group names to be notified when the monitor is down. Either specify user_group_ids or user_group_names. If omitted, the first user group returned by the /api/user_groups endpoint will be used.
* `tag_ids` (List of String) List of tags IDs to be associated to the monitor. Either specify tag_ids or tag_names.
* `tag_names` (List of String) List of tag names to be associated to the monitor. Tag name matching works for both exact and partial match. Either specify tag_ids or tag_names.
* `third_party_service_ids` (List of String) List of Third Party Service IDs to be associated to the monitor.
* `actions` (Map of String) Action to be performed on monitor IT Automation templates. 

Refer [API documentation](https://www.site24x7.com/help/api/#rest-api) for more information about attributes.
