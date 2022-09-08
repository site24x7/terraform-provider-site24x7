terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
	// (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
	// environment variable if the attribute is empty or omitted.
	oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"
  
	// (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
	// environment variable if the attribute is empty or omitted.
	oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
  
	// (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
	// environment variable if the attribute is empty or omitted.
	oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"
  
	// (Required) Specify the data center from which you have obtained your
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
	data_center = "US"
	
	// (Optional) ZAAID of the customer under a MSP or BU
	zaaid = "1234"
  
	// (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
	retry_min_wait = 1
  
	// (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
	// requests. This is the upper limit for the wait duration with exponential
	// backoff.
	retry_max_wait = 30
  
	// (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
	max_retries = 4
  
  }

// Webhook API doc: https://www.site24x7.com/help/api/#create-webhook
resource "site24x7_webhook_integration" "webhook_integration_basic" {
  // (Required) Display name for the Webhook integration
  name                            = "Webhook - Terraform"
  // (Required) Hook URL to which the message will be posted.
  url                             = "http://example.com" 
}

// Webhook API doc: https://www.site24x7.com/help/api/#create-webhook
resource "site24x7_webhook_integration" "webhook_integration" {
  // (Required) Display name for the Webhook integration
  name                            = "Test Webhook"
  // (Required) Hook URL to which the message will be posted.
  url                             = "http://example.com"
  // (Required) The amount of time a connection waits to time out. Default value is 30. Range 1 - 45.
  timeout                         = 30
  // (Optional) HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported. Default value is 'G'.
  // https://www.site24x7.com/help/api/#http_methods
  method                          = "P"
  // (Optional) Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'.
  selection_type = 0
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'.
  trouble_alert = true
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  critical_alert = false
  // (Optional) Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.
  down_alert = false
  // (Optional) Boolean indicating whether it is an On-Premise Poller based Webhook.
  is_poller_webhook               = false
  // (Optional) Denotes On-Premise Poller ID
  poller                          = "113770000023231022"
  // (Optional) Configuration to send incident parameters while executing the action
  send_incident_parameters        = false
  // (Optional) Configuration to send custom parameters while executing the action
  send_custom_parameters          = true
  // (Optional) Mandatory when send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL.
  custom_parameters               = "{\"test\":\"abcd\"}"
  // (Optional) Configuration to enable json format for post parameters
  send_in_json_format             = true
  // (Optional) Authentication method to access the action url.
  // https://www.site24x7.com/help/api/#auth_method
  auth_method                     = "B"
  // (Optional) User name for authentication
  user_name                        = "username"
  // (Optional) Password for authentication
  password                        = "password"
  // (Optional) Provider ID of the OAuth Provider to be associated with the action
  // https://www.site24x7.com/help/api/#list-oauth-providers
  oauth2_provider                 = "113770000023231001"
  // (Optional) User Agent to be used while monitoring the website
  user_agent                      = "Mozilla"
  // (Optional) Map of custom HTTP headers to send.
  custom_headers = {
    "Accept" = "application/json"
  }
  // (Optional) Monitors to be associated with the integration when the selection_type = 2.
  monitors                        = ["756"]
  // (Optional) Tags to be associated with the integration when the selection_type = 3.
  tags                        = ["345"]
  // (Optional) List of tag IDs to be associated with the integration.
  alert_tags_id                   = ["123"]
  // (Optional) Configuration to handle ticketing based integration
  manage_tickets                  = false
  // (Optional) URL to be invoked to update the request
  update_url                      = "http://test.tld"
  // (Optional) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  update_method                   = "P"
  // (Optional) Configuration to send incident parameters while updating the ticket.
  update_send_incident_parameters = false
  // (Optional) Configuration to send custom parameters while updating the ticket.
  update_send_custom_parameters   = false
  // (Optional) Mandatory when update_send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL
  update_custom_parameters        = "param=value"
  // (Optional) Configuration to post in JSON format while updating the ticket.
  update_send_in_json_format = true
  // (Optional) URL to be invoked to close the request
  close_url                       = "http://test.tld"
  // (Optional) HTTP Method to access the URL
  // https://www.site24x7.com/help/api/#http_methods
  close_method                    = "P"
  // (Optional) Configuration to send incident parameters while closing the ticket.
  close_send_incident_parameters  = false
  // (Optional) Configuration to send custom parameters while closing the ticket.
  close_send_custom_parameters    = false
  // (Optional) Mandatory when close_send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL
  close_custom_parameters         = "param=value"
  // (Optional) Configuration to post in JSON format while closing the ticket.
  close_send_in_json_format = true
}
