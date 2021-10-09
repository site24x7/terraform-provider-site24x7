---
layout: "site24x7"
page_title: "Provider: Site24x7"
description: |-
  Site24x7 offers a performance monitoring solution for DevOps and IT Operations enabling developers and network engineers to diagnose and fix application/network problems in real time.
---

# Site24x7 Provider

The [Site24x7](https://www.site24x7.com/) provider helps you to manage Site24x7 resources via terraform, thereby giving deep visibility into critical performance parameters of your resources and proactive insight into areas that could become an issue.

You must configure the provider with the proper credentials before you can use it. 

Use the navigation to the left to read about the available resources.

Refer [examples](https://github.com/site24x7/terraform-provider-site24x7/tree/main/examples) for adding Site24x7 resources via terraform.

## Example Usage

```terraform
# Terraform 0.13+ uses the Terraform Registry:

terraform {
  required_version = "~> 0.13.0"
  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      version = "~> 1.0.0"
    }
  }
}

# Configure the Site24x7 provider
// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_CLIENT_ID>"

  // The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_CLIENT_SECRET>"

  // The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_REFRESH_TOKEN>"

  // Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
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

// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_example" {
  // (Required) Display name for the monitor
  display_name = "Example Monitor"

  // (Required) Website address to monitor.
  website = "https://www.example.com"

  // (Optional) Interval at which your website has to be monitored.
  // See https://www.site24x7.com/help/api/#check-interval for all supported values.
  check_frequency = 1

  // (Optional) Name of the Location Profile that has to be associated with the monitor. 
  // Either specify location_profile_id or location_profile_name.
  // If location_profile_id and location_profile_name are omitted,
  // the first profile returned by the /api/location_profiles endpoint
  // (https://www.site24x7.com/help/api/#list-of-all-location-profiles) will be
  // used.
  location_profile_name = "North America"

}

```


## Provider Attributes

| Attribute              | Type    | Required? | Description                                                                                                                                                                 |
| ---------------------- | ------- | --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `oauth2_client_id`     | String  | Required  | Client ID obtained during Client Registration. The `SITE24X7_OAUTH2_CLIENT_ID` environment variable can also be used.                                                       |
| `oauth2_client_secret` | String  | Required  | Client Secret obtained during Client Registration. The `SITE24X7_OAUTH2_CLIENT_SECRET` environment variable can also be used.                                               |
| `oauth2_refresh_token` | String  | Required  | Refresh Token using which a new access token has to be generated. The `SITE24X7_OAUTH2_REFRESH_TOKEN` environment variable can also be used.                                |
| `data_center`          | String  | Optional  | The region for the data center from which OAuth 2.0 client credentials and refresh token were generated. Valid values are `US` or `EU` or `AU` or `IN` or `CN`.             |
| `max_retries`          | Number  | Optional  | Maximum number of Site24x7 API request retries to perform until giving up.                                                                                                  |
| `retry_max_wait`       | Number  | Optional  | The maximum time to wait in seconds before retrying failed Site24x7 API requests. This is the upper limit for the wait duration with exponential backoff.                   |
| `retry_min_wait`       | Number  | Optional  | The minimum time to wait in seconds before retrying failed Site24x7 API requests.                                                                                           |


## Debugging

Additional debugging information can be generated by exporting the `TF_LOG` environment variable when running Terraform commands. See [Debugging Terraform](https://www.terraform.io/docs/internals/debugging.html) for more information. 

```shell
export TF_LOG=TRACE
```