#!/bin/bash

make install

OUT=$?
if [ $OUT -eq 0 ];then
   sudo cp terraform-provider-site24x7 /usr/local/lib/terraform/registry.zoho.io/zoho/site24x7/1.0.0/linux_amd64/terraform-provider-site24x7_v1.0.0

   rm -f .terraform.lock.hcl

   terraform init
else
   echo "Compilation Error"
fi

