package main

import (
	"github.com/dtrouillet/terraform-provider-bitbucketserver/bitbucket"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bitbucket.Provider})
}
