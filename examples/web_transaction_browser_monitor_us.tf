terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.13.0"

  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
	// environment variable if the attribute is empty or omitted.
	oauth2_client_id = "<SITE24X7_OAUTH2_CLIENT_ID>"

  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
	// (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
	// environment variable if the attribute is empty or omitted.
	oauth2_client_secret = "<SITE24X7_OAUTH2_CLIENT_SECRET>"
    
  // (Security recommendation - It is always best practice to store your credentials in a Vault of your choice.)
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

// Site24x7 Domain Expiry monitor API doc - https://www.site24x7.com/help/api/#web-transaction(Browser)
resource "site24x7_web_transaction_browser_monitor" "web_transaction_browser_monitor_basic" {
 // (Required) Display name for the monitor
     display_name="RBM-Terraform"
     type= "REALBROWSER"
     async_dc_enabled = false
     base_url= "https://www.demoqa.com/"
     selenium_script="{\"id\":\"4bab335d-6abe-4450-9ca7-ab160de8fcb6\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://demoqa.com\",\"tests\":[{\"id\":\"742a0bbc-e523-461a-afeb-0636271a4361\",\"name\":\"\",\"commands\":[{\"id\":\"0621950a-7705-4aed-a9b7-74fcb968cac3\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://demoqa.com/\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"af22dd1f-886c-4251-8214-98e774f19579\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"b58182de-36a4-4575-9839-a4b0793891b2\",\"comment\":\"\",\"command\":\"click\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:position\"],[\"css=.card:nth-child(2) h5\",\"css:finder\"],[\"xpath=//h5[contains(text(),'Forms')]\",\"xpath:link\"],[\"xpath=//div[@id='app']/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div[2]/div/div[3]/h5\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"3c4eac2b-536a-4de1-9ce1-18fa72eab52c\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"click Practice Form\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/forms\",\"stepCount\":2,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"7673bc49-3288-495b-93bc-39e524edf2e7\",\"comment\":\"\",\"command\":\"clickAndWait\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"xpath:position\"],[\"css=.show .text\",\"css:finder\"],[\"xpath=//span[contains(text(),'Practice Form')]\",\"xpath:link\"],[\"xpath=(//li[@id='item-0']/span)[2]\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div/div/div[2]/div/ul/li/span\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"c026cc61-1301-428c-b7ca-3e4a76ee1b1d\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"742a0bbc-e523-461a-afeb-0636271a4361\"]}],\"urls\":[\"https://demoqa.com/\"],\"plugins\":[]}"
     location_profile_name="Asia Pacific"
     check_frequency="20"
     browser_version=10101
     think_time=3
     page_load_time=60
     script_type="txt"
     resolution="1600,900"
     threshold_profile_id="456418000000210001"
     on_call_schedule_id = "456418000001258016"
 
     ip_type = 2
    // ignore_cert_err = true
     user_agent="Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) Gecko/20100101 Firefox/101.0 Site24x7"
     browser_type = 1
    // parallel_polling = false
      perform_automation=true
     # actions={1=456418000001238121}
     dependency_resource_ids = ["456418000000081003"]
     proxy_details={webProxyUrl: "www.javatpoint", webProxyUname: "sdcsdcsdc", webProxyPass: "sadasds"}
     auth_details={userName: "12345",  password: "12345"}
     cookies={"234234"="3423432"}

}

// Site24x7 Domain Expiry monitor API doc - https://www.site24x7.com/help/api/#web-transaction(Browser)
resource "site24x7_web_transaction_browser_monitor" "web_transaction_browser_monitor_basic2" {
    display_name="RBM-Terraform"
    base_url= "https://www.demoqa.com/"
    selenium_script="{\"id\":\"4bab335d-6abe-4450-9ca7-ab160de8fcb6\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://demoqa.com\",\"tests\":[{\"id\":\"742a0bbc-e523-461a-afeb-0636271a4361\",\"name\":\"\",\"commands\":[{\"id\":\"0621950a-7705-4aed-a9b7-74fcb968cac3\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://demoqa.com/\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"af22dd1f-886c-4251-8214-98e774f19579\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"b58182de-36a4-4575-9839-a4b0793891b2\",\"comment\":\"\",\"command\":\"click\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:position\"],[\"css=.card:nth-child(2) h5\",\"css:finder\"],[\"xpath=//h5[contains(text(),'Forms')]\",\"xpath:link\"],[\"xpath=//div[@id='app']/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div[2]/div/div[3]/h5\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"3c4eac2b-536a-4de1-9ce1-18fa72eab52c\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"click Practice Form\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/forms\",\"stepCount\":2,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"7673bc49-3288-495b-93bc-39e524edf2e7\",\"comment\":\"\",\"command\":\"clickAndWait\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"xpath:position\"],[\"css=.show .text\",\"css:finder\"],[\"xpath=//span[contains(text(),'Practice Form')]\",\"xpath:link\"],[\"xpath=(//li[@id='item-0']/span)[2]\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div/div/div[2]/div/ul/li/span\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"c026cc61-1301-428c-b7ca-3e4a76ee1b1d\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"742a0bbc-e523-461a-afeb-0636271a4361\"]}],\"urls\":[\"https://demoqa.com/\"],\"plugins\":[]}"
    script_type="txt"
    check_frequency="20"
    threshold_profile_id="456418000000210001"
    page_load_time=60
    ignore_cert_err = true
    proxy_details={webProxyUrl: "www.javatpoint", webProxyUname: "uname", webProxyPass: "sadasds"}
    //auth_details={userName: "12345",  password: "12345"}
    cookies={"234234"="3423432"}
}