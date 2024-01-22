package tutorial

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/blueprint"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
	"github.com/blueprint-uservices/blueprint/plugins/golang/gocode"
	"github.com/blueprint-uservices/blueprint/plugins/golang/gogen"
	"golang.org/x/exp/slog"
)

// Blueprint IRNode for representing the wrapper node that adds an additional call parameter and an additional return parameter to every server-side method in the node that gets wrapped
type HelloParamServerWrapper struct {
	golang.Service
	golang.GeneratesFuncs
	golang.Instantiable

	InstanceName string
	Wrapped      golang.Service

	outputPackage string
}

// Implements ir.IRNode
func (node *HelloParamServerWrapper) ImplementsGolangNode() {}

// Implements ir.IRNode
func (node *HelloParamServerWrapper) Name() string {
	return node.InstanceName
}

// Implements ir.IRNode
func (node *HelloParamServerWrapper) String() string {
	return node.Name() + " = HelloParamServerWrapper(" + node.Wrapped.Name() + ")"
}

// Implements golang.ProvidesInterface
func (node *HelloParamServerWrapper) AddInterfaces(builder golang.ModuleBuilder) error {
	iface, err := node.genInterface(builder)
	if err != nil {
		return err
	}
	err = generateClientSideParamInterfaces(builder, iface, node.outputPackage)
	if err != nil {
		return err
	}
	return node.Wrapped.AddInterfaces(builder)
}

func newHelloParamServerWrapper(name string, server ir.IRNode) (*HelloParamServerWrapper, error) {
	serverNode, ok := server.(golang.Service)
	if !ok {
		return nil, blueprint.Errorf("tutorial server wrapper requires %s to be a golang service but got %s", server.Name(), reflect.TypeOf(server).String())
	}

	node := &HelloParamServerWrapper{}
	node.InstanceName = name
	node.Wrapped = serverNode
	node.outputPackage = "tutorial"

	return node, nil
}

func (node *HelloParamServerWrapper) genInterface(ctx ir.BuildContext) (*gocode.ServiceInterface, error) {
	iface, err := golang.GetGoInterface(ctx, node.Wrapped)
	if err != nil {
		return nil, err
	}
	module_ctx, valid := ctx.(golang.ModuleBuilder)
	if !valid {
		return nil, blueprint.Errorf("Tutorial expected build context to be a ModuleBuilder, got %v", ctx)
	}
	i := gocode.CopyServiceInterface(fmt.Sprintf("%v_TutorialParam", iface.BaseName), module_ctx.Info().Name+"/"+node.outputPackage, iface)
	for name, method := range i.Methods {
		method.AddArgument(gocode.Variable{Name: "extraparam", Type: &gocode.BasicType{Name: "string"}})
		method.AddRetVar(gocode.Variable{Name: "", Type: &gocode.BasicType{Name: "string"}})
		i.Methods[name] = method
	}
	return i, nil
}

// Implements service.ServiceNode
func (node *HelloParamServerWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	return node.genInterface(ctx)
}

// Implements golang.GeneratesFuncs
func (node *HelloParamServerWrapper) GenerateFuncs(builder golang.ModuleBuilder) error {
	service, err := golang.GetGoInterface(builder, node.Wrapped)
	if err != nil {
		return err
	}
	iface, err := golang.GetGoInterface(builder, node)
	if err != nil {
		return err
	}
	err = generateServerParamHandler(builder, iface, service, node.outputPackage)
	if err != nil {
		return err
	}
	return nil
}

// Implements golang.Instantiable
func (node *HelloParamServerWrapper) AddInstantiation(builder golang.NamespaceBuilder) error {
	if builder.Visited(node.InstanceName) {
		return nil
	}

	iface, err := golang.GetGoInterface(builder, node.Wrapped)
	if err != nil {
		return err
	}

	constructor := &gocode.Constructor{
		Package: builder.Module().Info().Name + "/" + node.outputPackage,
		Func: gocode.Func{
			Name: fmt.Sprintf("New_%v_TutorialParamServerWrapper", iface.BaseName),
			Arguments: []gocode.Variable{
				{Name: "ctx", Type: &gocode.UserType{Package: "context", Name: "context"}},
				{Name: "service", Type: iface},
			},
		},
	}

	return builder.DeclareConstructor(node.InstanceName, constructor, []ir.IRNode{node.Wrapped})
}

func generateServerParamHandler(builder golang.ModuleBuilder, iface *gocode.ServiceInterface, wrapped_service *gocode.ServiceInterface, outputPackage string) error {
	pkg, err := builder.CreatePackage(outputPackage)
	if err != nil {
		return err
	}

	server := &serverArgs{
		Package:   pkg,
		Service:   wrapped_service,
		Iface:     iface,
		Name:      wrapped_service.BaseName + "_TutorialParamServerWrapper",
		IfaceName: iface.Name,
		Imports:   gogen.NewImports(pkg.Name),
	}

	server.Imports.AddPackages("context")

	slog.Info(fmt.Sprintf("Generating %v/%v", server.Package.PackageName, wrapped_service.BaseName+"_TutorialParamServerWrapper"))
	outputFile := filepath.Join(server.Package.Path, wrapped_service.BaseName+"_TutorialParamServerWrapper.go")
	return gogen.ExecuteTemplateToFile("Tutorial", serverParamTemplate, server, outputFile)
}

func generateClientSideParamInterfaces(builder golang.ModuleBuilder, iface *gocode.ServiceInterface, outputPackage string) error {
	pkg, err := builder.CreatePackage(outputPackage)
	if err != nil {
		return err
	}

	server := &serverArgs{
		Package:   pkg,
		Iface:     iface,
		IfaceName: iface.Name,
		Imports:   gogen.NewImports(pkg.Name),
	}

	server.Imports.AddPackages("context")
	slog.Info(fmt.Sprintf("Generating %v/%v", server.Package.PackageName, iface.Name))
	outputFile := filepath.Join(server.Package.Path, iface.Name+".go")
	// Re-use the template from ir_method
	return gogen.ExecuteTemplateToFile("Tutorial", clientTemplate, server, outputFile)
}

var serverParamTemplate = `// Blueprint: Auto-generated by Tutorial Plugin
package {{.Package.ShortName}}

{{.Imports}}

type {{.IfaceName}} interface {
	{{range $_, $f := .Iface.Methods -}}
	{{Signature $f}}
	{{end}}
}

type {{.Name}} struct {
	Service {{.Imports.NameOf .Service.UserType}}
}

func New_{{.Name}}(ctx context.Context, service {{.Imports.NameOf .Service.UserType}}) (*{{.Name}}, error) {
	handler := &{{.Name}}{}
	handler.Service = service
	return handler, nil
}

{{$service := .Service.Name -}}
{{$receiver := .Name -}}
{{ range $_, $f := .Service.Methods}}
func (handler *{{$receiver}}) {{$f.Name -}} ({{ArgVarsAndTypes $f "ctx context.Context"}}, extraparam string) ({{RetVarsAndTypes $f "retparam string" "err error"}}) {
	{{RetVars $f "err"}} = handler.Service.{{$f.Name}}(ArgVars $f "ctx")
	retparam = extraparam
	return
}
{{end}}
`

// Blueprint IRNode for representing the wrapper node that adds an additional call parameter and an additional return parameter to every client-side method in the node that gets wrapped
type HelloParamClientWrapper struct {
	golang.Service
	golang.GeneratesFuncs
	golang.Instantiable

	InstanceName string
	Wrapped      golang.Service

	outputPackage string
}

// Implements ir.IRNode
func (node *HelloParamClientWrapper) ImplementsGolangNode() {}

// Implements ir.IRNode
func (node *HelloParamClientWrapper) Name() string {
	return node.InstanceName
}

// Implements ir.IRNode
func (node *HelloParamClientWrapper) String() string {
	return node.Name() + " = HelloParamServerWrapper(" + node.Wrapped.Name() + ")"
}

// Implements golang.ProvidesInterface
func (node *HelloParamClientWrapper) AddInterfaces(builder golang.ModuleBuilder) error {
	iface, err := node.genInterface(builder)
	if err != nil {
		return err
	}
	err = generateClientSideParamInterfaces(builder, iface, node.outputPackage)
	if err != nil {
		return err
	}
	return node.Wrapped.AddInterfaces(builder)
}

func (node *HelloParamClientWrapper) genInterface(ctx ir.BuildContext) (*gocode.ServiceInterface, error) {
	iface, err := golang.GetGoInterface(ctx, node.Wrapped)
	if err != nil {
		return nil, err
	}
	module_ctx, valid := ctx.(golang.ModuleBuilder)
	if !valid {
		return nil, blueprint.Errorf("TutorialParamClientWrapper expected build context to be a ModuleBuilder, got %v", ctx)
	}
	i := gocode.CopyServiceInterface(fmt.Sprintf("%v_TutorialParamClientWrapperInterface", iface.BaseName), module_ctx.Info().Name+"/"+node.outputPackage, iface)
	for name, method := range i.Methods {
		method.Arguments = method.Arguments[:len(method.Arguments)-1]
		method.Returns = method.Returns[:len(method.Returns)-1]
		i.Methods[name] = method
	}
	return i, nil
}

func newHelloParamClientWrapper(name string, server ir.IRNode) (*HelloParamServerWrapper, error) {
	serverNode, ok := server.(golang.Service)
	if !ok {
		return nil, blueprint.Errorf("tutorial server wrapper requires %s to be a golang service but got %s", server.Name(), reflect.TypeOf(server).String())
	}

	node := &HelloParamServerWrapper{}
	node.InstanceName = name
	node.Wrapped = serverNode
	node.outputPackage = "tutorial"

	return node, nil
}

// Implements service.ServiceNode
func (node *HelloParamClientWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	return node.genInterface(ctx)
}

// Implements golang.GeneratesFuncs
func (node *HelloParamClientWrapper) GenerateFuncs(builder golang.ModuleBuilder) error {
	service, err := golang.GetGoInterface(builder, node.Wrapped)
	if err != nil {
		return err
	}
	iface, err := golang.GetGoInterface(builder, node)
	if err != nil {
		return err
	}
	err = generateServerParamHandler(builder, iface, service, node.outputPackage)
	if err != nil {
		return err
	}
	return nil
}

// Implements golang.Instantiable
func (node *HelloParamClientWrapper) AddInstantiation(builder golang.NamespaceBuilder) error {
	if builder.Visited(node.InstanceName) {
		return nil
	}

	iface, err := golang.GetGoInterface(builder, node.Wrapped)
	if err != nil {
		return err
	}

	constructor := &gocode.Constructor{
		Package: builder.Module().Info().Name + "/" + node.outputPackage,
		Func: gocode.Func{
			Name: fmt.Sprintf("New_%v_TutorialParamClientWrapper", iface.BaseName),
			Arguments: []gocode.Variable{
				{Name: "ctx", Type: &gocode.UserType{Package: "context", Name: "context"}},
				{Name: "service", Type: iface},
			},
		},
	}

	return builder.DeclareConstructor(node.InstanceName, constructor, []ir.IRNode{node.Wrapped})
}

func generateClientParamHandler(builder golang.ModuleBuilder, iface *gocode.ServiceInterface, wrapped_service *gocode.ServiceInterface, outputPackage string) error {
	pkg, err := builder.CreatePackage(outputPackage)
	if err != nil {
		return err
	}

	server := &serverArgs{
		Package:   pkg,
		Service:   wrapped_service,
		Iface:     iface,
		Name:      wrapped_service.BaseName + "_TutorialParamClientWrapper",
		IfaceName: iface.Name,
		Imports:   gogen.NewImports(pkg.Name),
	}

	server.Imports.AddPackages("context", "log")

	slog.Info(fmt.Sprintf("Generating %v/%v", server.Package.PackageName, wrapped_service.BaseName+"_TutorialParamClientWrapper"))
	outputFile := filepath.Join(server.Package.Path, wrapped_service.BaseName+"_TutorialParamClientWrapper.go")
	return gogen.ExecuteTemplateToFile("Tutorial", clientParamTemplate, server, outputFile)
}

var clientParamTemplate = `// Blueprint: Auto-generated by Tutorial plugin
package {{.Package.ShortName}}

{{.Imports}}

type {{.IfaceName}} interface {
	{{range $_, $f := .Iface.Methods -}}
	{{Signature $f}}
	{{end}}
}

type {{.Name}} struct {
	Client {{.Imports.NameOf .Service.UserType}}
}

func New_{{.Name}}(ctx context.Context, client {{.Imports.NameOf .Service.UserType}}) (*{{.Name}}, error) {
	handler := &{{.Name}}{}
	handler.Client = client
	return handler, nil
}

{{$service := .Service.Name -}}
{{$receiver := .Name -}}
{{ range $_, $f := .Service.Methods }}
func (handler *{{$receiver}}) {{$f.Name -}} ({{ArgVarsAndTypes $f "ctx context.Context"}}) ({{RetVarsAndTypes $f "err error"}}) {
	var retparam string
	{{RetVars $f "retparam" "err"}} handler.Service.{{$f.Name}}({{ArgVars $f "ctx"}}, "Hello!")
	log.Println("Ret param was ", retparam)
	return
}
{{end}}
`
