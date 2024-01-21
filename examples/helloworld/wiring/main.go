// Package main provides the HelloWorld application, a simple application designed to be a tutorial for demonstrating Blueprint
// usage and not as a realistic executable application.
//
// The wiring specs in the [specs] directory illustrate the usage of various Blueprint plugins.
//
// # Usage
//
// To display usage, run
//
//	go run . -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/tutorial/examples/helloworld/wiring/specs"
)

func main() {
	// Configure the location of our workflow spec
	workflow.Init("../workflow")

	// Build a supported wiring spec
	name := "HelloWorld"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
