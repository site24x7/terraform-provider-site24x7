
resource "site24x7_server_monitor" "SERVER_123456000025786003" { 
 perform_automation = true 
 log_needed = true 
 notification_profile_id = "123456000000029001" 
 tag_ids = ["123456000024829001", "123456000024829005"] 
 poll_interval = 1 
 monitor_groups = ["123456000000120011"] 
 threshold_profile_id = "123456000000029003" 
 user_group_ids = ["123456000000025005", "123456000000025009"] 
 display_name = "ubuntu-server"
}


resource "site24x7_server_monitor" "SERVER_123456000027570003" { 
 log_needed = true 
 notification_profile_id = "123456000000029001" 
 poll_interval = 1 
 threshold_profile_id = "123456000000029003" 
 user_group_ids = ["123456000000025005"] 
 display_name = "sdp-w10-2305"
}

