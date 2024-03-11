terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 

    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  # oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  # oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // ZAAID of the customer under a MSP or BU
  zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP).
  data_center = "US"

  // The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

// Site24x7 Rest API Monitor API doc - https://www.site24x7.com/help/api/#rest-api-transaction
resource "site24x7_rest_api_transaction_monitor" "rest_api_transaction_monitor_basic" {
  // (Required) Display name for the monitor
  display_name = "REST API Monitor - terraform"

  // (Required) List of Steps details to be associated to the steps.

  steps {

    // (Required) Display name for the step
    display_name = "RestAPI Transaction Monitor"

    // (Required)  API request details related to this step.
    step_details {

      // (Required) Domain address for the step.
      step_url = "https://www.example1.com"
    }
  }
  // (Optional) Name of the Location Profile that has to be associated with the monitor.
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"
}

variable "step_display_name_list" {
  type    = list(string)
  default = ["Rest Api Transaction step 1", "Rest Api Transaction step 2","Rest Api Transaction step 3"]
}
variable "step_url_list" {
  type    = list(string)
  default = ["https://www.example1.com", "https://www.example1.com","https://www.example1.com"]
}
variable "on_error_list" {
  type    = list(number)
  default = [2,3,4]
}
variable "timeout_list" {
  type    = list(number)
  default = [10,15,20]
}
variable "matching_keyword_list" {
  type    = list(object({
    severity=number
    value=string
  }))
  default = [{
    severity=2
    value="aaa"
  },
  {
    severity =2
    value = "aaa"
  },
  {
    severity = 2
    value="aaa"
  }]
}
variable "unmatching_keyword_list" {
  type    = list(object({
    severity=number
    value=string
  }))
  default = [{
    severity=2
    value="bbb"
  },
  {
    severity =2
    value = "bbb"
  },
  {
    severity = 2
    value="bbb"
  }]
}
variable "match_regex_list" {
  type    = list(object({
    severity=number
    value=string
  }))
  default = [{
    severity=2
    value=".*aaa.*"
  },
  {
    severity =2
    value = ".*aaa.*"
  },
  {
    severity = 2
    value=".*aaa.*"
  }]
}
variable "response_headers_severity_list" {
  type    = list(number)
  default = [0,0,0]
}
variable "response_headers_list" {
  type    = list(object({
    Content-Encoding=string
    Connection=string
  }))
  default = [{
    Content-Encoding = "gzip"
        Connection  = "Keep-Alive"
  },
  {
   Content-Encoding = "gzip"
        Connection  = "Keep-Alive"
  },
  {
   Content-Encoding = "gzip"
        Connection  = "Keep-Alive"
  }]
}
variable "up_status_codes_list" {
  type    = list(string)
  default = ["400:500","400:500","400:500"]
}
variable "response_content_type_list" {
  type    = list(string)
  default = ["J","J","J"]
}
variable "match_json_path_list" {
  type    = list(list(string))
  default = [[
        "$.store.book[*].author",
        "$..author",
        "$.store.*"
      ],[
        "$.store.book[*].author",
        "$..author",
        "$.store.*"
      ],[
        "$.store.book[*].author",
        "$..author",
        "$.store.*"
      ]]
}
variable "match_json_path_severity_list" {
  type    = list(number)
  default = [0,0,0]
}
variable "json_schema_list" {
  type    = list(string)
  default =  ["{test:abcd}","{test:abcd}","{test:abcd}"]
}
variable "json_schema_severity_list" {
  type    = list(number)
  default = [2,2,2]
}
variable "json_schema_check_list" {
  type    = list(bool)
  default = [true,true,true]
}
variable "http_method_list" {
  type    = list(string)
  default = ["P","P","P"]
}
variable "request_content_type_list" {
  type    = list(string)
  default = ["J","J","J"]
}
variable "request_body_list" {
  type    = list(string)
  default = ["{test:abcd}","{test:abcd}","{test:abcd}"]
}
variable "request_headers_list" {
  type    = list(object({
    Accept=string
  }))
  default = [{
    Accept = "application/json"
  },
  {
   Accept = "application/json"
  },
  {
    Accept = "application/json"
  }]
}
variable "graphql_query_list" {
  type    = list(string)
  default = ["query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}","query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}",
  "query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}"]
}
variable "graphql_variables_list" {
  type    = list(string)
  default=["{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}","{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}","{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}"]
}
variable "dynamic_param_response_type_list" {
  type    = list(string)
  default=["J","J","J"]
}
variable "response_variables_list" {
  type    = list(object({
    template_id=string
  }))
  default=[{template_id="$.data.template_id"},{template_id="$.data.template_id"},{template_id="$.data.template_id"}]
}
variable "dynamic_header_params_list" {
  type    = list(object({
    Accept=string
  }))
  default = [{
    Accept = "application/json"
  },
  {
   Accept = "application/json"
  },
  {
    Accept = "application/json"
  }]
}

// Site24x7 Rest API Transaction Monitor API doc - https://www.site24x7.com/help/api/#rest-api-transaction
resource "site24x7_rest_api_transaction_monitor" "rest_api_transaction_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "RestAPI Transaction Monitor"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = "5"

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

  // (Required) List of Steps details to be associated to the steps.

  count = length(var.step_display_name_list)
  steps {

    // (Required) Display name for the step
    display_name = var.step_display_name_list[count.index]
    // (Required)  API request details related to this step.
    step_details {

      // (Required) Domain address for the step.
      step_url = var.step_url_list[count.index]

      // (Optional) Stop on Error severity for the step. '0' means Stop and Notify, '1' means Proceed , '2' means Notify and Proceed.
      on_error = var.on_error_list[count.index]

      // (Optional)  Timeout for connecting to REST API Default value is 10. Range 1 - 45.
      timeout = var.timeout_list[count.index]

      // (Optional) Check for the keyword in the website response.
      matching_keyword = {
        severity = var.matching_keyword_list[count.index].severity
        value    =  var.matching_keyword_list[count.index].value
      }

      // (Optional) Check for non existence of keyword in the website response.
      unmatching_keyword = {
        severity =  var.unmatching_keyword_list[count.index].severity
        value    =  var.unmatching_keyword_list[count.index].value
      }

      // (Optional) Match the regular expression in the website response.
      match_regex = {
        severity =  var.match_regex_list[count.index].severity
        value    =  var.match_regex_list[count.index].value
      }

      // (Optional) Map of HTTP response headers to check.
      response_headers_severity =  var.response_headers_severity_list[count.index] // Can take values 0 or 2. '0' denotes Down and '2' denotes Trouble.
      response_headers = {
        Content-Encoding =  var.response_headers_severity_list[count.index].Content-Encoding
        Connection       = var.response_headers_severity_list[count.index].Connection
      }

      // HTTP Configuration
      // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response.
      // You can specify individual status codes, as well as ranges separated with a colon.
      up_status_codes = var.up_status_codes_list[count.index]

      // ================ JSON ASSERTION ATTRIBUTES
      // (Optional) Response content type. Default value is 'T'
      // 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML
      // https://www.site24x7.com/help/api/#res_content_type
      response_content_type = var.response_content_type_list[count.index]
      // (Optional) Provide multiple JSON Path expressions to enable evaluation of JSON Path expression assertions.
      // The assertions must successfully parse the JSON Path in the JSON. JSON expression assertions fails if the expressions does not match.
      match_json_path = var.match_json_path_list[count.index]
      // (Optional) Trigger an alert when the JSON path assertion fails during a test.
      // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
      match_json_path_severity = var.match_json_path_severity_list[count.index]
      // (Optional) JSON schema to be validated against the JSON response.
      json_schema = var.json_schema_list[count.index]
      // (Optional) Trigger an alert when the JSON schema assertion fails during a test.
      // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
      json_schema_severity = var.json_schema_severity_list[count.index]
      // (Optional) JSON Schema check allows you to annotate and validate all JSON endpoints for your web service.
      json_schema_check = var.json_schema_check_list[count.index]
      // JSON ASSERTION ATTRIBUTES ================

      // ================ HTTP POST with request body
      // (Optional) HTTP Method to be used for accessing the website. Default value is 'G'. 'G' denotes GET, 'P' denotes POST, 'U' denotes PUT and 'D' denotes DELETE. HEAD is not supported.
      http_method = var.http_method_list[count.index]
      // (Optional) Provide content type for request params when http_method is 'P'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML and 'F' denotes FORM
      request_content_type = var.request_content_type_list[count.index]
      // (Optional) Provide the content to be passed in the request body while accessing the website.
      request_body = var.request_body_list[count.index]
      // (Optional) Map of custom HTTP headers to send.
      request_headers = {
        Accept = var.request_headers_list[count.index].Accept
      }
      // HTTP POST with request body ================

      // ================ GRAPHQL ATTRIBUTES
      // (Optional) Provide content type for request params.
      // request_content_type = "G"
      // (Optional) Provide the GraphQL query to get specific response from GraphQL based API service. request_content_type = "G"
      graphql_query = var.graphql_query_list[count.index]
      // (Optional) Provide the GraphQL variables to get specific response from GraphQL based API service. request_content_type = "G"
      graphql_variables = var.graphql_variables_list[count.index]
      // GRAPHQL ATTRIBUTES ================

      // ================ PARAMETER FORWARDING ATTRIBUTES

      // (Optional) Provide the Response Format for Parameter Forwarding.

      dynamic_param_response_type = var.dynamic_param_response_type_list[count.index]

      // (Optional) Provide the Response Variable for parameter forwarding in Map format. 

      response_variables = {
        template_id = var.response_variables_list[count.index].template_id
      }

      // (Optional) Provide the Response Header/Cookies for parameter forwarding in Map format.

      dynamic_header_params = {
        Accept =var.dynamic_header_params_list[count.index].Accept
      }

      //  PARAMETER FORWARDING ATTRIBUTES ================
    }
  }

  
}


# // Site24x7 Rest API Transaction Monitor API doc - https://www.site24x7.com/help/api/#rest-api-transaction
variable "steps_config" {
  type = list(object({
    display_name = list(string)
    step_details = object({
      step_url                 = list(string)
      on_error                 = list(number)
      timeout                  = list(number)
      matching_keyword         = list(map(any))
      unmatching_keyword       = list(map(any))
      match_regex              = list(map(any))
      response_headers_severity =list( number)
      response_headers         = list(map(string))
      up_status_codes          = list(string)
      response_content_type    = list(string)
      match_json_path          = list(list(string))
      match_json_path_severity = list(number)
      json_schema              = list(string)
      json_schema_severity     = list(number)
      json_schema_check        = list(bool)
      http_method              = list(string)
      request_content_type     = list(string)
      request_body             = list(string)
      request_headers          = list(map(string))
      graphql_query            = list(string)
      graphql_variables        = list(string)
      dynamic_param_response_type = list(string)
      response_variables       = list(map(string))
      dynamic_header_params     =list(map(string))
    })
  }))
  default =  [
    {
      display_name =  ["Rest Api Transaction step 1", "Rest Api Transaction step 2","Rest Api Transaction step 3"]
      step_details = {
        step_url                 =["https://www.example1.com", "https://www.example2.com","https://www.example3.com"]
        on_error                 =  [2,3,4]
        timeout                  = [10,15,20]
        matching_keyword         = [{severity=2  
                                     value="aaa1"
                                    },
                                    {
                                      severity =2
                                      value = "aaa2"
                                    },
                                    {
                                      severity = 2
                                      value="aaa3"
                                    }]
        unmatching_keyword       = [{
                                      severity=2
                                      value="bbb"
                                    },
                                    {
                                      severity =2
                                      value = "bbb"
                                    },
                                    {
                                      severity = 2
                                      value="bbb"
                                    }]  # Add your default values
        match_regex              = [{
                                      severity=2
                                      value=".*aaa.*"
                                    },
                                    {
                                      severity =2
                                      value = ".*aaa.*"
                                    },
                                    {
                                      severity = 2
                                      value=".*aaa.*"
                                    }]  # Add your default values
        response_headers_severity = [0,0,0]
        response_headers         = [{
                                      Content-Encoding = "gzip"
                                          Connection  = "Keep-Alive"
                                    },
                                    {
                                    Content-Encoding = "gzip"
                                          Connection  = "Keep-Alive"
                                    },
                                    {
                                    Content-Encoding = "gzip"
                                          Connection  = "Keep-Alive"
                                    }]
        up_status_codes          = ["400:500","400:500","400:500"]
        response_content_type    = ["J","J","J"]
        match_json_path          = [[
                                    "$.store.book[*].author",
                                    "$..author",
                                    "$.store.*"
                                  ],[
                                    "$.store.book[*].author",
                                    "$..author",
                                    "$.store.*"
                                  ],[
                                    "$.store.book[*].author",
                                    "$..author",
                                    "$.store.*"
                                  ]]
        match_json_path_severity = [0,0,0]
        json_schema              =["{test:abcd}","{test:abcd}","{test:abcd}"]
        json_schema_severity     = [2,2,2]
        json_schema_check        = [true,true,true]
        http_method              = ["P","P","P"]
        request_content_type     = ["J","J","J"]
        request_body             = ["{test:abcd}","{test:abcd}","{test:abcd}"]
        request_headers          = [{
                                      Accept = "application/json"
                                    },
                                    {
                                    Accept = "application/json"
                                    },
                                    {
                                      Accept = "application/json"
                                    }]
        graphql_query            = ["query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}","query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}",
                                      "query GetFlimForId($FilmId:ID!){\n        film(id:$FilmId){\n            id\n            title\n            director\n            producers\n        }\n}"]
        graphql_variables        = ["{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}","{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}","{\n    \"FilmId\":\"ZmlsbXM6NQ==\"\n}"]
        dynamic_param_response_type = ["J","J","J"]
        response_variables       = [{template_id="$.data.template_id"},{template_id="$.data.template_id"},{template_id="$.data.template_id"}]
        dynamic_header_params     =[{
                                        Accept = "application/json"
                                      },
                                      {
                                      Accept = "application/json"
                                      },
                                      {
                                        Accept = "application/json"
                                      }]
                                          }
                                      } ]
}

resource "site24x7_rest_api_transaction_monitor" "rest_api_transaction_monitor_multiple_steps_example" {
  count = length(var.steps_config)
   // (Required) Display name for the monitor
   display_name            = "RestAPI Transaction Monitor Multiple steps - terraform"

   // (Optional) Interval at which your website has to be monitored.
   // See https://www.site24x7.com/help/api/#check-interval for all supported values.
   check_frequency         = "5"

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
    // (Required) List of Steps details to be associated to the steps.
  dynamic "steps" {
        for_each = var.steps_config[count.index].step_details.step_url
        content {
          // (Required) Display name for the step
          display_name = var.steps_config[count.index].display_name[steps.key]
          // (Required)  API request details related to this step.
          step_details {
            // (Required) Domain address for the step.
            step_url                 = var.steps_config[count.index].step_details.step_url[steps.key]
            
            // (Optional) Stop on Error severity for the step. '0' means Stop and Notify, '1' means Proceed , '2' means Notify and Proceed.
            on_error                 = var.steps_config[count.index].step_details.on_error[steps.key]
            
            // (Optional)  Timeout for connecting to REST API Default value is 10. Range 1 - 45.
            timeout                  = var.steps_config[count.index].step_details.timeout[steps.key]
            
            // (Optional) Check for the keyword in the website response.
            matching_keyword         = var.steps_config[count.index].step_details.matching_keyword[steps.key]
            
            // (Optional) Check for non existence of keyword in the website response.
            unmatching_keyword       = var.steps_config[count.index].step_details.unmatching_keyword[steps.key]
            
            // (Optional) Match the regular expression in the website response.
            match_regex              = var.steps_config[count.index].step_details.match_regex[steps.key]
            
            // (Optional) Map of HTTP response headers to check.
            response_headers_severity = var.steps_config[count.index].step_details.response_headers_severity[steps.key]
            response_headers         = var.steps_config[count.index].step_details.response_headers[steps.key]
            
            // HTTP Configuration
            // (Optional) Provide a comma-separated list of HTTP status codes that indicate a successful response.
            // You can specify individual status codes, as well as ranges separated with a colon.
            up_status_codes          = var.steps_config[count.index].step_details.up_status_codes[steps.key]
            
            // ================ JSON ASSERTION ATTRIBUTES
            // (Optional) Response content type. Default value is 'T'
            // 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML
            // https://www.site24x7.com/help/api/#res_content_type
            response_content_type    = var.steps_config[count.index].step_details.response_content_type[steps.key]

            // (Optional) Provide multiple JSON Path expressions to enable evaluation of JSON Path expression assertions.
            // The assertions must successfully parse the JSON Path in the JSON. JSON expression assertions fails if the expressions does not match.
            match_json_path          = var.steps_config[count.index].step_details.match_json_path[steps.key]
            
            // (Optional) Trigger an alert when the JSON path assertion fails during a test.
            // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
            match_json_path_severity = var.steps_config[count.index].step_details.match_json_path_severity[steps.key]
            
            // (Optional) JSON schema to be validated against the JSON response.
            json_schema              = var.steps_config[count.index].step_details.json_schema[steps.key]
            
            // (Optional) Trigger an alert when the JSON schema assertion fails during a test.
            // Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.
            json_schema_severity     = var.steps_config[count.index].step_details.json_schema_severity[steps.key]
            
            // (Optional) JSON Schema check allows you to annotate and validate all JSON endpoints for your web service.
            json_schema_check        = var.steps_config[count.index].step_details.json_schema_check[steps.key]
            
            // (Optional) HTTP Method to be used for accessing the website. Default value is 'G'. 'G' denotes GET, 'P' denotes POST, 'U' denotes PUT and 'D' denotes DELETE. HEAD is not supported.
            http_method              = var.steps_config[count.index].step_details.http_method[steps.key]
            
            // (Optional) Provide content type for request params when http_method is 'P'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML and 'F' denotes FORM
            request_content_type     = var.steps_config[count.index].step_details.request_content_type[steps.key]
            
            // (Optional) Provide the content to be passed in the request body while accessing the website.
            request_body             = var.steps_config[count.index].step_details.request_body[steps.key]
            
            // (Optional) Map of custom HTTP headers to send.
            request_headers          = var.steps_config[count.index].step_details.request_headers[steps.key]
            
            // (Optional) Provide content type for request params.
            // request_content_type = "G"
            // (Optional) Provide the GraphQL query to get specific response from GraphQL based API service. request_content_type = "G"
            graphql_query            = var.steps_config[count.index].step_details.graphql_query[steps.key]
            
            // (Optional) Provide the GraphQL variables to get specific response from GraphQL based API service. request_content_type = "G"
            graphql_variables        = var.steps_config[count.index].step_details.graphql_variables[steps.key]
            
            // (Optional) Provide the Response Format for Parameter Forwarding.
            dynamic_param_response_type = var.steps_config[count.index].step_details.dynamic_param_response_type[steps.key]
            
            // (Optional) Provide the Response Variable for parameter forwarding in Map format. 
            response_variables       = var.steps_config[count.index].step_details.response_variables[steps.key]

            // (Optional) Provide the Response Header/Cookies for parameter forwarding in Map format.
            dynamic_header_params     = var.steps_config[count.index].step_details.dynamic_header_params[steps.key]
          }
      }
    }
}