package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var deviceKeyDataSourceSchema = map[string]*schema.Schema{}

func DataSourceSite24x7DeviceKey() *schema.Resource {
	return &schema.Resource{
		Read:   deviceKeyDataSourceRead,
		Schema: deviceKeyDataSourceSchema,
	}
}

// deviceKeyDataSourceRead fetches the device key from Site24x7
func deviceKeyDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	deviceKey, err := client.DeviceKey().Get()
	if err != nil {
		return err
	}

	updateDeviceKeyDataSourceResourceData(d, deviceKey)

	return nil
}

func updateDeviceKeyDataSourceResourceData(d *schema.ResourceData, deviceKey *api.DeviceKey) {
	d.SetId(deviceKey.ID)
}
