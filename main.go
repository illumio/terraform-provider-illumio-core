// Copyright 2021 Illumio, Inc. All Rights Reserved.

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/illumio/terraform-provider-illumio-core/illumio-core"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return illumiocore.Provider()
		},
	})
}
