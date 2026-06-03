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
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
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

// Site24x7 OAuth2 Provider API doc - https://www.site24x7.com/help/api/#oauth2-providers

// ========================
// Client Credentials Grant
// ========================
resource "site24x7_oauth2_provider" "oauth_client_credentials" {
  // (Required) A name to identify this OAuth provider in the dashboard.
  provider_name = "OAuth Provider - Client Credentials"

  // (Required) The grant type. Use 2 for Client Credentials.
  oauth2_flow = 2

  // (Required) The client ID obtained from your OAuth provider.
  client_id = "100005498423413561-nlv4ticqf9nom3sh0"

  // (Required) The client secret obtained from your OAuth provider.
  client_secret = "EtU_Te3lvQ5MVJ_pC0kP_2MR"

  // (Required) The API Token Endpoint of the OAuth 2 Provider.
  access_token_uri = "https://accounts.zoho.com/oauth/v2/token"

  // (Optional) Authentication method for sending Client ID and Client Secret.
  // "B" for sending credentials in the request body.
  auth_method = "B"

  // (Required) Specifies how the access token will be sent - via headers or query parameters.
  send_token_as {
    // Method: "H" for Header, "Q" for Query parameter
    method = "H"
    // Header name or query parameter name
    name = "Authorization"
    // Header value or query parameter value
    value = "Bearer $${access.token}"
  }

  // (Optional) Parameters to be sent in the request body.
  request_body {
    name  = "param1"
    value = "value1"
  }

  // (Optional) User groups notified when access token refresh fails after 3 retries.
  user_group_ids = ["8000000000017"]
}

// ==========================================
// Resource Owner Password Credentials Grant
// ==========================================
resource "site24x7_oauth2_provider" "oauth_resource_owner" {
  // (Required) A name to identify this OAuth provider in the dashboard.
  provider_name = "OAuth Provider - Resource Owner"

  // (Required) The grant type. Use 3 for Resource Owner Password Credentials.
  oauth2_flow = 3

  // (Required) The client ID obtained from your OAuth provider.
  client_id = "100005498423413561-nlv4ticqf9nom3sh0"

  // (Required) The client secret obtained from your OAuth provider.
  client_secret = "EtU_Te3lvQ5MVJ_pC0kP_2MR"

  // (Required) The API Token Endpoint of the OAuth 2 Provider.
  access_token_uri = "https://accounts.zoho.com/oauth/v2/token"

  // (Optional) Authentication method for sending Client ID and Client Secret.
  auth_method = "B"

  // (Optional) Resource owner's username. Required for oauth2_flow = 3.
  auth_user = "username"

  // (Optional) Resource owner's password. Required for oauth2_flow = 3.
  auth_pass = "password"

  // (Required) Specifies how the access token will be sent.
  send_token_as {
    method = "H"
    name   = "Authorization"
    value  = "Bearer $${access.token}"
  }

  // (Optional) Parameters to be sent in the request body.
  request_body {
    name  = "scope"
    value = "read write"
  }

  // (Optional) User groups notified when access token refresh fails after 3 retries.
  user_group_ids = ["8000000000017"]
}

// ========================
// Data Source - Read Only
// ========================
data "site24x7_oauth2_provider" "existing" {
  provider_id = site24x7_oauth2_provider.oauth_client_credentials.id
}

// Output the provider details
output "oauth2_provider_name" {
  value = data.site24x7_oauth2_provider.existing.provider_name
}

output "oauth2_provider_expiry" {
  value = data.site24x7_oauth2_provider.existing.expiry_time
}
