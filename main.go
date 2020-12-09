package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/appgate/terraform-provider-appgate/appgate"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: appgate.Provider})
}
