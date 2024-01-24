package specs

import (
	"strings"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/simple"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/tutorial/examples/helloworld/workflow/servicea"
	"github.com/blueprint-uservices/tutorial/examples/helloworld/workflow/serviceb"
)

var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with gRPC, uses an in-memorycache as the cache backend.",
	Build:       makeDockerSpec,
}

func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	cache := simple.Cache(spec, "cache")
	applyLoggerDefaults := func(service_name string) string {

		procName := strings.ReplaceAll(service_name, "service", "process")
		cntrName := strings.ReplaceAll(service_name, "service", "container")
		http.Deploy(spec, service_name)
		goproc.CreateProcess(spec, procName, service_name)
		return linuxcontainer.CreateContainer(spec, cntrName, procName)
	}
	serviceb := workflow.Service[*serviceb.ServiceBImpl](spec, "b_service", cache)
	servicea := workflow.Service[*servicea.ServiceAImpl](spec, "a_service", serviceb)
	cntrb := applyLoggerDefaults(serviceb)
	cntra := applyLoggerDefaults(servicea)
	return []string{cntra, cntrb}, nil
}
