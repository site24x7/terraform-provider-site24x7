terraform {
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      source = "site24x7/site24x7"
      # Update to the latest version from the Terraform Registry
    }
  }
}

// Authentication configuration for Site24x7
provider "site24x7" {
  // (Required) The client ID used for OAuth authentication.
  oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Required) The client secret used for OAuth authentication.
  oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"

  // (Required) The refresh token obtained from Site24x7.
  oauth2_refresh_token = "<SITE24X7_OAUTH2_REFRESH_TOKEN>"

  // (Required) Specify the Site24x7 data center. Options: US, EU, IN, AU, CN, JP, CA.
  data_center = "US"

  // (Optional) Maximum retries for Site24x7 API requests.
  max_retries = 4

  // (Optional) Minimum wait time (in seconds) between retries.
  retry_min_wait = 1

  // (Optional) Maximum wait time (in seconds) between retries.
  retry_max_wait = 30
}

// Site24x7 Business Hour API example
resource "site24x7_business_hour" "business_hour_basic" {
  // (Required) Name of the business hour configuration.
  display_name = "General Shift"

  // (Optional) Description of the business hour configuration.
  description = "General shift 5 days a week, 9 AM to 5 PM"

  // (Required) Business hour time configuration.
  time_config = [
    {
      // (Required) Day of the week (1 = Sunday, 7 = Saturday).
      day = 2  // Monday
      
      // (Required) Business hour start time in HH:MM format.
      start_time = "09:00"

      // (Required) Business hour end time in HH:MM format.
      end_time = "17:00"
    },
    {
      day = 3  // Tuesday
      start_time = "09:00"
      end_time = "17:00"
    },
    {
      day = 4  // Wednesday
      start_time = "09:00"
      end_time = "17:00"
    },
    {
      day = 5  // Thursday
      start_time = "09:00"
      end_time = "17:00"
    },
    {
      day = 6  // Friday
      start_time = "09:00"
      end_time = "17:00"
    }
  ]
}
