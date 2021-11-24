#!/bin/bash

make install

OUT=$?
if [ $OUT -eq 0 ];then

   if [ ! -d "$HOME/.terraform.d/plugins/registry.zoho.io/zoho/site24x7/1.0.0/linux_amd64" ]; then
      sudo mkdir -p $HOME/.terraform.d/plugins/registry.zoho.io/zoho/site24x7/1.0.0/linux_amd64
   fi
   
   #sudo cp -vf terraform-provider-site24x7 $HOME/.terraform.d/plugins/registry.zoho.io/zoho/site24x7/1.0.0/linux_amd64/terraform-provider-site24x7_v1.0.0
   sudo cp -vf terraform-provider-site24x7 $HOME/.terraform.d/plugins/registry.terraform.io/site24x7/site24x7/1.0.0/linux_amd64/terraform-provider-site24x7_v1.0.0

   rm -f .terraform.lock.hcl

   terraform init
else
   echo "Compilation Error"
fi

