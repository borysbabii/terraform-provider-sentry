package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/borysbabii/terraform-provider-sentry/sentry"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sentry.Provider,
	})
}
