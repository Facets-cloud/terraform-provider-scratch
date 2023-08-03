package main

import (
	"context"
	"log"

	"terraform-provider-scratch/scratch"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	// Example version string that can be overwritten by a release process
	version string = "0.4.0"
	commit  string = ""
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/hashicorp/scratch",
	}

	err := providerserver.Serve(context.Background(), scratch.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
