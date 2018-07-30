package main

import (
	"github.com/yulebron/terraform-provider/alicloud"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: alicloud.Provider,
	})
}
