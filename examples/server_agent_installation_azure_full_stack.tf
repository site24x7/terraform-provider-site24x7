terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.0"
    }
  }
}

provider "azurerm" {
  features {}
}

variable "site24x7_device_key" {
  description = "Site24x7 device key"
  type        = string
  sensitive   = true
}

variable "admin_ssh_public_key" {
  description = "SSH public key for the VM"
  type        = string
}

resource "azurerm_resource_group" "demo" {
  name     = "rg-site24x7-demo"
  location = "Central India"
}

resource "azurerm_virtual_network" "demo" {
  name                = "vnet-site24x7-demo"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.demo.location
  resource_group_name = azurerm_resource_group.demo.name
}

resource "azurerm_subnet" "demo" {
  name                 = "default"
  resource_group_name  = azurerm_resource_group.demo.name
  virtual_network_name = azurerm_virtual_network.demo.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_public_ip" "demo" {
  name                = "pip-site24x7-demo"
  location            = azurerm_resource_group.demo.location
  resource_group_name = azurerm_resource_group.demo.name
  allocation_method   = "Static"
  sku                 = "Standard"
}

resource "azurerm_network_interface" "demo" {
  name                = "nic-site24x7-demo"
  location            = azurerm_resource_group.demo.location
  resource_group_name = azurerm_resource_group.demo.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.demo.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.demo.id
  }
}

resource "azurerm_linux_virtual_machine" "demo" {
  name                = "vm-site24x7-demo"
  resource_group_name = azurerm_resource_group.demo.name
  location            = azurerm_resource_group.demo.location
  size                = "Standard_B2s"

  admin_username = "azureuser"

  network_interface_ids = [
    azurerm_network_interface.demo.id
  ]

  admin_ssh_key {
    username   = "azureuser"
    public_key = var.admin_ssh_public_key
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "ubuntu-24_04-lts"
    sku       = "server"
    version   = "latest"
  }

  # Runs during VM provisioning
  custom_data = base64encode(<<-EOF
    #!/bin/bash
    set -e

    echo "Installing Site24x7 Full Stack Agent..."

    cd /tmp

    wget -q \
      https://staticdownloads.site24x7.in/server/Site24x7FullStackAgent_LinuxIns.sh

    chmod +x Site24x7FullStackAgent_LinuxIns.sh

    bash Site24x7FullStackAgent_LinuxIns.sh \
      -i \
      -key='${var.site24x7_device_key}' \
      -automation=true \
      -apm_insight=false \
      -ebpf=false

    echo "Site24x7 Full Stack Agent installation completed."
  EOF
  )
}

output "vm_public_ip" {
  value = azurerm_public_ip.demo.ip_address
}
