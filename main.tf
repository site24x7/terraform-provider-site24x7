terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    site24x7 = {
      # source = "site24x7/site24x7"
      // Uncomment for local build
      source  = "registry.terraform.io/site24x7/site24x7"
      version = "1.0.0"
    }
  }
}

// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // (Required) The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "1000.0D485VEM89FTBJT1P2ORWJZO9ABAFO"

  // (Required) The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "7203f0267db4943c44105b99f2d582ccbf5e58e698"

  // (Required) The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "1000.335dbfb8d379481a6904e6f46af675ee.86a2db587e30443a4177e0ea9e134c68"

  // (Optional) The access token will be looked up in the SITE24X7_OAUTH2_ACCESS_TOKEN
  // environment variable if the attribute is empty or omitted. You need not configure oauth2_access_token
  // when oauth2_refresh_token is set.
  # oauth2_access_token = "<SITE24X7_OAUTH2_ACCESS_TOKEN>"

  // (Optional) oauth2_access_token expiry in seconds. Specify access_token_expiry when oauth2_access_token is configured.
  # access_token_expiry = "0"

  // (Optional) ZAAID of the customer under a MSP or BU
  # zaaid = "1234"

  // (Required) Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
  data_center = "US"

  // (Optional) The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // (Optional) The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // (Optional) Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4
}



# // DomainExpiry Server API doc: https://www.site24x7.com/help/api/#domain-expiry
# resource "site24x7_domain_expiry_monitor" "domain_expiry_monitor_basic3" {
#   // (Required) Display name for the monitor
#   display_name = "ignore registry 1"
#   host_name = "file.com"
#   domain_name="whios.iana.org"
#   port=443
#   timeout = 50
#   expire_days=40
#   use_ipv6=true
#   //matching_keyword={2="aaa"}
#   location_profile_name = "North America"
#   ignore_registry_date = false
#   type = "DOMAINEXPIRY"
#   actions={0=456418000001238121}
#    on_call_schedule_id = "456418000001258016"
  
# }
# // DomainExpiry Server API doc: https://www.site24x7.com/help/api/#domain-expiry
# resource "site24x7_domain_expiry_monitor" "domain_expiry_monitor_basic" {
#   // (Required) Display name for the monitor
#   display_name = "ignore registry 2"
#   host_name = "file.com"
#   domain_name="whios.iana.org"
#   port=443
#   timeout = 50
#   expire_days=40
#   use_ipv6=false
#   match_case=false 
#   matching_keyword = {
#  	  severity= 2
#  	  value= "aaa"
#  	}
#   unmatching_keyword = {
#  	  severity= 0
#  	  value= "bbb"
#  	}
#   match_regex = {
#  	  severity= 0
#  	  value= "test(.*)\\d"
#  	}
#   perform_automation=true
#   location_profile_name = "North America"
#   ignore_registry_date = false
#   type = "DOMAINEXPIRY"
#   actions={1=456418000001238121}
#   notification_profile_name="System Generated"
  
# }


# // DomainExpiry Server API doc: https://www.site24x7.com/help/api/#domain-expiry
# resource "site24x7_domain_expiry_monitor" "domain_expiry_monitor_basic1" {
#   // (Required) Display name for the monitor
#   display_name = "ignore registry 2"
#   host_name = "file.com"
# }

// WebTransactionBrowserMonitor Server API doc: https://www.site24x7.com/help/api/#web-transaction(Browser)
#  resource "site24x7_web_transaction_browser_monitor" "site24x7_web_transaction_browser_monitor_basic77" {
#    // (Required) Display name for the monitor
#      display_name="RBM-www.javatpoint.com Terraform"
#      type= "REALBROWSER"
#      async_dc_enabled = false
#      base_url= "https://www.demoqa.com/"
#      selenium_script="{\"id\":\"4bab335d-6abe-4450-9ca7-ab160de8fcb6\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://demoqa.com\",\"tests\":[{\"id\":\"742a0bbc-e523-461a-afeb-0636271a4361\",\"name\":\"\",\"commands\":[{\"id\":\"0621950a-7705-4aed-a9b7-74fcb968cac3\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://demoqa.com/\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"af22dd1f-886c-4251-8214-98e774f19579\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"b58182de-36a4-4575-9839-a4b0793891b2\",\"comment\":\"\",\"command\":\"click\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:position\"],[\"css=.card:nth-child(2) h5\",\"css:finder\"],[\"xpath=//h5[contains(text(),'Forms')]\",\"xpath:link\"],[\"xpath=//div[@id='app']/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div[2]/div/div[3]/h5\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"3c4eac2b-536a-4de1-9ce1-18fa72eab52c\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"click Practice Form\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/forms\",\"stepCount\":2,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"7673bc49-3288-495b-93bc-39e524edf2e7\",\"comment\":\"\",\"command\":\"clickAndWait\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"xpath:position\"],[\"css=.show .text\",\"css:finder\"],[\"xpath=//span[contains(text(),'Practice Form')]\",\"xpath:link\"],[\"xpath=(//li[@id='item-0']/span)[2]\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div/div/div[2]/div/ul/li/span\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"c026cc61-1301-428c-b7ca-3e4a76ee1b1d\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"742a0bbc-e523-461a-afeb-0636271a4361\"]}],\"urls\":[\"https://demoqa.com/\"],\"plugins\":[]}"
#      location_profile_name="Asia Pacific"
#      check_frequency="20"
#      browser_version=10101
#      think_time=3
#      page_load_time=60
#      script_type="txt"
#      resolution="1600,900"
#      threshold_profile_id="456418000000210001"
#      on_call_schedule_id = "456418000001258016"
 
#      ip_type = 2
#     // ignore_cert_err = true
#      user_agent="Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) Gecko/20100101 Firefox/101.0 Site24x7"
#      browser_type = 1
#     // parallel_polling = false
#       perform_automation=true
#      # actions={1=456418000001238121}
#      dependency_resource_ids = ["456418000000081003"]
#      proxy_details={webProxyUrl: "www.javatpoint", webProxyUname: "sdcsdcsdc", webProxyPass: "sadasds"}
#      auth_details={userName: "12345",  password: "12345"}
#      cookies={"234234"="3423432"}
#    }
resource "site24x7_web_transaction_browser_monitor" "site24x7_web_transaction_browser_monitor_basic89" {
  display_name="RBM-www.javatpoint.com Terraform"
    base_url= "https://www.demoqa.com/"
    //selenium_script="{\"id\":\"4bab335d-6abe-4450-9ca7-ab160de8fcb6\",\"version\":\"1.1\",\"name\":\"\",\"url\":\"https://demoqa.com\",\"tests\":[{\"id\":\"742a0bbc-e523-461a-afeb-0636271a4361\",\"name\":\"\",\"commands\":[{\"id\":\"0621950a-7705-4aed-a9b7-74fcb968cac3\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"Loading - https://demoqa.com/\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/\",\"stepCount\":1,\"stepTime\":\"0\",\"actionName\":\"\"},{\"id\":\"af22dd1f-886c-4251-8214-98e774f19579\",\"comment\":\"\",\"command\":\"open\",\"target\":\"/\",\"targets\":[],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"b58182de-36a4-4575-9839-a4b0793891b2\",\"comment\":\"\",\"command\":\"click\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:position\"],[\"css=.card:nth-child(2) h5\",\"css:finder\"],[\"xpath=//h5[contains(text(),'Forms')]\",\"xpath:link\"],[\"xpath=//div[@id='app']/div/div/div[2]/div/div[2]/div/div[3]/h5\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div[2]/div/div[3]/h5\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"3c4eac2b-536a-4de1-9ce1-18fa72eab52c\",\"comment\":\"\",\"command\":\"newStep\",\"target\":\"click Practice Form\",\"targets\":[],\"value\":\"\",\"URL\":\"https://demoqa.com/forms\",\"stepCount\":2,\"stepTime\":0,\"actionName\":\"\"},{\"id\":\"7673bc49-3288-495b-93bc-39e524edf2e7\",\"comment\":\"\",\"command\":\"clickAndWait\",\"target\":\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"targets\":[[\"xpath=/html/body/div[2]/div/div/div[2]/div/div/div/div[2]/div/ul/li/span\",\"xpath:position\"],[\"css=.show .text\",\"css:finder\"],[\"xpath=//span[contains(text(),'Practice Form')]\",\"xpath:link\"],[\"xpath=(//li[@id='item-0']/span)[2]\",\"xpath:idRelative\"],[\"xpath=/html/body/div/div/div/div/div/div/div/div[2]/div/ul/li/span\",\"xpath1:position\"]],\"value\":\"\",\"URL\":\"\",\"stepCount\":0,\"stepTime\":0,\"actionName\":\"\"}]}],\"suites\":[{\"id\":\"c026cc61-1301-428c-b7ca-3e4a76ee1b1d\",\"name\":\"\",\"persistSession\":false,\"parallel\":false,\"timeout\":300,\"tests\":[\"742a0bbc-e523-461a-afeb-0636271a4361\"]}],\"urls\":[\"https://demoqa.com/\"],\"plugins\":[]}"
    //script_type="txt"
    think_time = 1
    check_frequency="20"
    threshold_profile_id="456418000000210001"
    page_load_time=60
    resolution = "1024,768"
    proxy_details={webProxyUrl: "www.javatpoint", webProxyUname: "uname", webProxyPass: "sadasds"}
    cookies={"234234"="3423432"}
    
  }