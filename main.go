package main

import (
	"context"
	"flag"
	"log"
	"terraform-provider-xkcdpass/xkcdpwprovider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), xkcdpwprovider.New, providerserver.ServeOpts{
		Address:         "advian.fi/tf-oss/xkcdpass",
		Debug:           debug,
		ProtocolVersion: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
}
