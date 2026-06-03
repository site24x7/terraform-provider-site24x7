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

// Site24x7 Server Threshold Profile API doc - https://www.site24x7.com/help/api/#threshold-server
resource "site24x7_threshold_profile" "server_threshold_profile_us" {
  // (Required) Name of the profile
  profile_name = "Server Threshold Profile - Terraform"
  // (Required) Type of the profile - Must be "SERVER" for server monitor threshold
  type = "SERVER"
  // (Optional) Threshold profile types
  // 1 - Static Threshold, 2 - AI-based Threshold
  profile_type = 1

  // (Optional) CPU Usage - Trouble threshold
  cpu_trouble_threshold = {
    // https://www.site24x7.com/help/api/#threshold_severity
    // 2 - Trouble
    severity = 2
    // https://www.site24x7.com/help/api/#constants
    // 1 - Greater than (>), 2 - Less than (<), 3 - Greater than or equal to (>=), 4 - Less than or equal to (<=), 5 - Equal to (=)
    comparison_operator = 1
    // CPU usage percentage threshold value
    value = 80
    // https://www.site24x7.com/help/api/#threshold_strategy
    // 1 - Poll Count, 2 - Poll Average, 3 - Time Range, 4 - Average Time
    strategy = 2
    // Poll Check Value
    polls_check = 5
  }

  // (Optional) CPU Usage - Critical threshold
  cpu_critical_threshold = {
    severity            = 3
    comparison_operator = 1
    value               = 95
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Memory Usage - Trouble threshold
  memory_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 80
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Memory Usage - Critical threshold
  memory_critical_threshold = {
    severity            = 3
    comparison_operator = 1
    value               = 95
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Disk Usage - Trouble threshold
  disk_usage_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 80
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Disk Usage - Critical threshold
  disk_usage_critical_threshold = {
    severity            = 3
    comparison_operator = 1
    value               = 95
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Disk Partition - Trouble threshold
  disk_partition_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 80
    strategy            = 1
    polls_check         = 1
  }

  // (Optional) Process CPU Usage - Trouble threshold
  process_cpu_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 70
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Process Memory Usage - Trouble threshold
  process_memory_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 70
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Network Error Packet - Trouble threshold
  network_error_packet_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 10
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Process Down Alert - Triggers alert if process is down
  process_down_alert = {
    value       = true
    severity    = 2
    strategy    = 1
    polls_check = 5
  }

  // (Optional) Notify if a resource check fails
  server_resource_down_alert = {
    value    = true
    severity = 2
  }

  // (Optional) DC alert
  dc_alert = {
    value    = true
    severity = 2
  }

  // (Optional) Disk status threshold
  disk_status_threshold = {
    value    = true
    severity = 2
  }

  // (Optional) Service status threshold
  service_status_threshold = {
    value       = true
    severity    = 2
    strategy    = 1
    polls_check = 5
  }

  // (Optional) Network status threshold
  nw_status_threshold = {
    value    = true
    severity = 2
  }

  // (Optional) Disk Used Size threshold
  disk_used_size = {
    trouble             = 50
    comparison_operator = 1
    polls_check         = 5
    strategy            = 2
    unit_id             = 1
  }

  // (Optional) Disk Free Size threshold
  disk_free_size = {
    trouble             = 10
    comparison_operator = 2
    polls_check         = 5
    strategy            = 2
    unit_id             = 1
  }

  // (Optional) Server Uptime threshold
  server_uptime = {
    trouble             = 60
    comparison_operator = 2
    polls_check         = 1
    strategy            = 1
    unit_id             = 2
  }

  // ========== Linux Specific Attributes ==========

  // (Optional) System Load 1 Min - Trouble threshold (Linux only)
  system_load_1min_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 5
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) System Load 5 Min - Trouble threshold (Linux only)
  system_load_5min_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 3
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) System Load 15 Min - Trouble threshold (Linux only)
  system_load_15min_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 2
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Running Process Count - Trouble threshold (Linux only)
  process_running_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 200
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Total Process Count - Trouble threshold (Linux only)
  total_process_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 500
    strategy            = 2
    polls_check         = 5
  }

  // (Optional) Blocked Process Count - Trouble threshold (Linux only)
  blocked_process_trouble_threshold = {
    severity            = 2
    comparison_operator = 1
    value               = 10
    strategy            = 2
    polls_check         = 5
  }

  // ========== Windows Specific Attributes ==========
  // Uncomment below for Windows servers

  // (Optional) Running Process Count - Trouble threshold (Windows only)
  # running_process_trouble_threshold = {
  #   severity            = 2
  #   comparison_operator = 1
  #   value               = 200
  #   strategy            = 2
  #   polls_check         = 5
  # }

  // (Optional) Total Service Count - Trouble threshold (Windows only)
  # total_service_trouble_threshold = {
  #   severity            = 2
  #   comparison_operator = 1
  #   value               = 100
  #   strategy            = 2
  #   polls_check         = 5
  # }

  // (Optional) Processor Queue Length - Trouble threshold (Windows only)
  # process_queue_length_trouble_threshold = {
  #   severity            = 2
  #   comparison_operator = 1
  #   value               = 5
  #   strategy            = 2
  #   polls_check         = 5
  # }
}
