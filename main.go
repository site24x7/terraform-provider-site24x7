package main

import (
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: site24x7.Provider,
	})
}
