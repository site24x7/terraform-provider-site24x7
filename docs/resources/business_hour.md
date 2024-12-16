layout: "site24x7"
page_title: "Site24x7: site24x7_business_hour"
sidebar_current: "docs-site24x7-business-hour"
description: |-
Create and manage Business Hour configurations in Site24x7.

Resource: site24x7_business_hour

Use this resource to create, update, and delete Business Hour configurations in Site24x7.

Example Usage

// Site24x7 Business Hour API doc - https://www.site24x7.com/help/api/#business-hours
resource "site24x7_business_hour" "business_hour_basic" {
  // (Required) Display name for the business hour configuration.
  display_name = "General Shift"

  // (Optional) Description for the business hour configuration.
  description  = "General shift 5 days a week 9 AM to 5 PM"

  // (Required) Business hour time configuration for each day.
  time_config {
    day         = 2
    start_time  = "09:00"
    end_time    = "17:00"
  }
  time_config {
    day         = 3
    start_time  = "09:00"
    end_time    = "17:00"
  }
  time_config {
    day         = 4
    start_time  = "09:00"
    end_time    = "17:00"
  }
  time_config {
    day         = 5
    start_time  = "09:00"
    end_time    = "17:00"
  }
  time_config {
    day         = 6
    start_time  = "09:00"
    end_time    = "17:00"
  }
}

Attributes Reference

Required

display_name (String) - Display name for the business hour configuration.

time_config (Block List) - Business hour configuration for each day.

day (Number) - Day of the week (1 for Sunday, 7 for Saturday).

start_time (String) - Start time in HH:mm format.

end_time (String) - End time in HH:mm format.

Optional

description (String) - Description for the business hour configuration.

Refer to the API documentation(https://www.site24x7.com/help/api/#business-hours) for more details about attributes and supported values.