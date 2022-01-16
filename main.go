package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/site24x7/terraform-provider-site24x7/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
