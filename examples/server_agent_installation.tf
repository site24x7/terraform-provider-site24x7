terraform {
  # Require Terraform version 0.15.x (recommended)
  required_version = "~> 0.15.0"

  required_providers {
    aws = {     
      # Declaring the source location/address where Terraform can download plugins
      source  = "hashicorp/aws"
      # Declaring the version of aws provider as greater than 3.0
      version = "~> 3.0"  
    }
    site24x7 = {
      source  = "site24x7/site24x7"
    }
  }
}

#Define keys and region
provider "aws" {
  access_key = ""
  secret_key = ""
  region     = "us-east-2"
}

# Terraform security group - https://sweetcode.io/deploy-infrastructure-seconds-terraform/
resource "aws_security_group" "site24x7-terraform-sg" {
  name = "site24x7-terraform-sg"
  description = "Site24x7 Terraform Security Group"

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }    

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "t2_micro_instance" {
  ami           = "ami-030bd1caa8425dfe8"
  instance_type = "t2.micro"
  associate_public_ip_address = true
  key_name = "site24x7_key_pair"
  security_groups = ["${aws_security_group.site24x7-terraform-sg.name}"]

  tags = {
    Name = "ubuntu-20.04"
  }
  
  connection {
    type        = "ssh"
    user        = "ubuntu"
    host        = "${self.public_ip}"
    private_key = "${file("~/Downloads/aws/site24x7_key_pair.pem")}"
  #The connection will use the local SSH agent for authentication. - https://stackoverflow.com/questions/35381229/why-cant-terraform-ssh-in-to-ec2-instance-using-supplied-example
   #agent = false
  }

  provisioner "remote-exec" {
    inline = [
      "sudo wget https://staticdownloads.site24x7.com/server/Site24x7InstallScript.sh",
      "sudo bash Site24x7InstallScript.sh -i -key=<key>", 
    ]
  }
}

output "aws_instance_ip" {
  description = "The public ip for ssh access"
  value       = aws_instance.t2_micro_instance.public_ip
}
