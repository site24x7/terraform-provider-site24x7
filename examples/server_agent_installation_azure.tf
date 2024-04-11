terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
	site24x7 = {
      source  = "site24x7/site24x7"
      # Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      
    }
  }
}
// Az login before running this template to get the accesskey from azure for this template deployment
provider "azurerm" {
  features {}
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
	// OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN/JP/CA).
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

data "site24x7_device_key" "s247devicekey" {}

// Displays the device key
output "s247_device_key" {
  description = "Device Key : "
  value       = data.site24x7_device_key.s247devicekey.id
}
// resource Resource group
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}
// resource virtual network 
resource "azurerm_virtual_network" "example" {
  name                = "site24x7vn"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_subnet" "example" {
  name                 = "site24x7sub"
  resource_group_name  = azurerm_resource_group.example.name
  virtual_network_name = azurerm_virtual_network.example.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_network_interface" "example" {
  name                = "site24x7ni"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  ip_configuration {
    name                          = "testconfiguration1"
    subnet_id                     = azurerm_subnet.example.id
    private_ip_address_allocation = "Dynamic"
  }
}

// resource virtual machine
resource "azurerm_virtual_machine" "example" {
  name                  = "site24x7vm"
  location              = azurerm_resource_group.example.location
  resource_group_name   = azurerm_resource_group.example.name
  network_interface_ids = [azurerm_network_interface.example.id]
  vm_size               = "standard_b2s"

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name          = "myosdisk1"
    caching       = "ReadWrite"
    create_option = "FromImage"
	  managed_disk_type = "Standard_LRS"
  }

  os_profile {
    computer_name  = "site24x7vm"
    admin_username = "<Username>"
    admin_password = "<Strong_Password>"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags = {
    environment = "staging"
  }
}
//specify site24x7 server agent installation through azure extension block
resource "azurerm_virtual_machine_extension" "example" {
  name                 = "Site24x7.site24x7-linuxserver-agent-linux"
  virtual_machine_id   = azurerm_virtual_machine.example.id
  publisher            = "Site24x7"
  type                 = "Site24x7LinuxServerExtn"
  type_handler_version = "1.8"

  protected_settings = <<SETTINGS
 {
  "site24x7LicenseKey": "${data.site24x7_device_key.s247devicekey.id}"
 }
SETTINGS


  tags = {
    environment = "Staging"
 }
}

