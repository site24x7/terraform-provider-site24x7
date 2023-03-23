terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}
// Az login before running this template to get the accesskey from azure for this template deployment
provider "azurerm" {
  features {}
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
// resource storage account
resource "azurerm_storage_account" "example" {
  name                     = "site24x7pm001"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "staging"
  }
}

resource "azurerm_storage_container" "example" {
  name                  = "site24x7pm001"
  storage_account_name  = azurerm_storage_account.example.name
  container_access_type = "private"
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
    vhd_uri       = "${azurerm_storage_account.example.primary_blob_endpoint}${azurerm_storage_container.example.name}/myosdisk1.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "site24x7vm"
    admin_username = "testadmin"
    admin_password = "<strong password>"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags = {
    environment = "staging"
  }
}
// Extension block to define the extention installation
resource "azurerm_virtual_machine_extension" "example" {
  name                 = "Site24x7.site24x7-linuxserver-agent-linux"
  virtual_machine_id   = azurerm_virtual_machine.example.id
  publisher            = "Site24x7"
  type                 = "Site24x7LinuxServerExtn"
  type_handler_version = "1.8"

  protected_settings = <<SETTINGS
 {
  "site24x7LicenseKey":"<Key_from_Site_24x7>"
 }
SETTINGS


  tags = {
    environment = "Production"
 }
}