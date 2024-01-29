// Package tutorial implements the following wiring functions:
//  1. AddHelloMethod: Adds a new method called `HelloNew` to the service interface.
//  2. AddHelloParam: Adds an extra call parameter and an extra return parameter to every method in the service.
//  3. Instrument: Adds logging statements to both the server and client side of every service.
package tutorial

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/blueprint"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/pointer"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
	"golang.org/x/exp/slog"
)

// [AddHelloMethod] can be called from the wiring specification to add a `HelloNew` method to Service with name `serviceName`.
func AddHelloMethod(spec wiring.WiringSpec, serviceName string) {
	// Define the name for the wrapper node we are adding to the Blueprint IR
	wrapper_name := serviceName + ".hello.method"

	// Get the pointer for the serviceName to ensure that the newly defined wrapper IR node will be attached to the node chain of the desired service
	ptr := pointer.GetPointer(spec, serviceName)
	if ptr == nil {
		slog.Error("Unable to add hello method to " + serviceName + " as it is not a pointer. Did you forget to define " + serviceName + "? You can define serviceName using `workflow.Service`")
		return
	}

	// Attach the Hello wrapper node to the server-side node chain of the desired service
	serverNext := ptr.AddDstModifier(spec, wrapper_name)

	// Define the IRNode for the wrapper node and add it to the wiring specification
	spec.Define(wrapper_name, &HelloMethodWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var server golang.Service
		if err := ns.Get(serverNext, &server); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, serverNext, err)
		}

		// Instantiate the IRNode
		return newHelloMethodWrapper(wrapper_name, server)
	})
}

// [AddHelloParam] can be called from the wiring specification to add an extra call parameter and an extra return parameter to every method exposed by the Service with name `serviceName`.
func AddHelloParam(spec wiring.WiringSpec, serviceName string) {
	// Define the names for the wrapper nodes we are adding to the Blueprint IR
	wrapper_name := serviceName + ".hello.param.server"
	client_wrapper_name := serviceName + ".hello.param.client"

	// Get the pointer for the serviceName to ensure that the newly defined wrapper IR node will be attached to the node chain of the desired service
	ptr := pointer.GetPointer(spec, serviceName)
	if ptr == nil {
		slog.Error("Unable to add hello param to " + serviceName + " as it is not a pointer. Did you forget to define " + serviceName + "? You can define serviceName using `workflow.Service`")
		return
	}

	// Attach the Hello wrapper node to the server-side node chain of the desired service
	serverNext := ptr.AddDstModifier(spec, wrapper_name)

	// Define the IRNode for the wrapper node and add it to the wiring specification
	spec.Define(wrapper_name, &HelloParamServerWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var server golang.Service
		if err := ns.Get(serverNext, &server); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, serverNext, err)
		}

		// Instantiate the IRNode
		return newHelloParamServerWrapper(wrapper_name, server)
	})

	// Attach the Hello wrapper node to the client-side node chain of the desired service
	clientNext := ptr.AddSrcModifier(spec, client_wrapper_name)

	// Define the IRNode for the wrapper node and add it to the wiring specification
	spec.Define(client_wrapper_name, &HelloParamClientWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var client golang.Service
		if err := ns.Get(clientNext, &client); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, clientNext, err)
		}

		return newHelloParamClientWrapper(client_wrapper_name, client)
	})
}

// [Instrument] can be called from the wiring specification to add logging statements to every method in both the server and client side code of Service with `serviceName`.
func Instrument(spec wiring.WiringSpec, serviceName string) {
	// Define the names for the wrapper nodes we are adding to the Blueprint IR
	wrapper_name := serviceName + ".hello.instrument.server"
	client_wrapper_name := serviceName + ".hello.instrument.client"

	// Get the pointer for the serviceName to ensure that the newly defined wrapper IR node will be attached to the node chain of the desired service
	ptr := pointer.GetPointer(spec, serviceName)
	if ptr == nil {
		slog.Error("Unable to add instrument " + serviceName + " as it is not a pointer. Did you forget to define " + serviceName + "? You can define serviceName using `workflow.Service`")
		return
	}

	// Attach the Hello wrapper node to the server-side node chain of the desired service
	serverNext := ptr.AddDstModifier(spec, wrapper_name)

	// Define the IRNode for the wrapper node and add it to the wiring specification
	spec.Define(wrapper_name, &HelloInstrumentServerWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var server golang.Service
		if err := ns.Get(serverNext, &server); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, serverNext, err)
		}

		// Instantiate the IRNode
		return newHelloInstrumentServerWrapper(wrapper_name, server)
	})

	// Attach the Hello wrapper node to the client-side node chain of the desired service
	clientNext := ptr.AddSrcModifier(spec, client_wrapper_name)

	// Define the IRNode for the wrapper node and add it to the wiring specification
	spec.Define(client_wrapper_name, &HelloInstrumentClientWrapper{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		// Get the IRNode that will be wrapped by HelloWrapper
		var client golang.Service
		if err := ns.Get(clientNext, &client); err != nil {
			return nil, blueprint.Errorf("Tutorial Plugin %s expected %s to be a golang.Service, but encountered %s", wrapper_name, clientNext, err)
		}

		return newHelloInstrumentClientWrapper(client_wrapper_name, client)
	})
}
