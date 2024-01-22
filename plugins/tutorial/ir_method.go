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

// Blueprint IRNode for representing the wrapper node that adds a `Hello` method to the wrapped IRNode.
type HelloMethodWrapper struct {
	golang.Service
	golang.GeneratesFuncs
	golang.Instantiable

	InstanceName string
	Wrapped      golang.Service

	outputPackage string
}

// Implements ir.IRNode
func (node *HelloMethodWrapper) ImplementsGolangNode() {}

// Implements ir.IRNode
func (node *HelloMethodWrapper) Name() string {
	return node.InstanceName
}

// Implements ir.IRNode
func (node *HelloMethodWrapper) String() string {
	return node.Name() + " = HelloMethodWrapper(" + node.Wrapped.Name() + ")"
}

// IMplements golang.ProvidesInterface
func (node *HelloMethodWrapper) AddInterfaces(builder golang.ModuleBuilder) error {
	iface, err := node.genInterface(builder)
	if err != nil {
		return err
	}
	err = generateClientSideInterfaces(builder, iface, node.outputPackage)
	if err != nil {
		return err
	}
	return node.Wrapped.AddInterfaces(builder)
}

func newHelloMethodWrapper(name string, server ir.IRNode) (*HelloMethodWrapper, error) {
	serverNode, is_callable := server.(golang.Service)
	if !is_callable {
		return nil, blueprint.Errorf("tutorial server wrapper requires %s to be a golang service but got %s", server.Name(), reflect.TypeOf(server).String())
	}

	node := &HelloMethodWrapper{}
	node.InstanceName = name
	node.Wrapped = serverNode
	node.outputPackage = "tutorial"

	return node, nil
}

func (node *HelloMethodWrapper) genInterface(ctx ir.BuildContext) (*gocode.ServiceInterface, error) {
	iface, err := golang.GetGoInterface(ctx, node.Wrapped)
	if err != nil {
		return nil, err
	}
	module_ctx, valid := ctx.(golang.ModuleBuilder)
	if !valid {
		return nil, blueprint.Errorf("Tutorial expected build context to be a ModuleBuilder, got %v", ctx)
	}
	i := gocode.CopyServiceInterface(fmt.Sprintf("%v_TutorialMethod", iface.BaseName), module_ctx.Info().Name+"/"+node.outputPackage, iface)
	health_check_method := &gocode.Func{}
	health_check_method.Name = "Hello"
	health_check_method.Returns = append(health_check_method.Returns, gocode.Variable{Type: &gocode.BasicType{Name: "string"}})
	i.AddMethod(*health_check_method)
	return i, nil
}

// Implements service.ServiceNode
func (node *HelloMethodWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	return node.genInterface(ctx)
}

// Implements golang.GeneratesFuncs
func (node *HelloMethodWrapper) GenerateFuncs(builder golang.ModuleBuilder) error {
	service, err := golang.GetGoInterface(builder, node.Wrapped)
	if err != nil {
		return err
	}
	iface, err := golang.GetGoInterface(builder, node)
	if err != nil {
		return err
	}
	err = generateServerHandler(builder, iface, service, node.outputPackage)
	if err != nil {
		return err
	}
	return nil
}

// Implements golang.Instantiable
func (node *HelloMethodWrapper) AddInstantiation(builder golang.NamespaceBuilder) error {
	// Only generate instantiation code for this instance once
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
			Name: fmt.Sprintf("New_%v_TutorialMethodImpl", iface.BaseName),
			Arguments: []gocode.Variable{
				{Name: "ctx", Type: &gocode.UserType{Package: "context", Name: "Context"}},
				{Name: "service", Type: iface},
			},
		},
	}

	return builder.DeclareConstructor(node.InstanceName, constructor, []ir.IRNode{node.Wrapped})
}

func generateServerHandler(builder golang.ModuleBuilder, iface *gocode.ServiceInterface, wrapped_service *gocode.ServiceInterface, outputPackage string) error {
	pkg, err := builder.CreatePackage(outputPackage)
	if err != nil {
		return err
	}

	server := &serverArgs{
		Package:   pkg,
		Service:   wrapped_service,
		Iface:     iface,
		Name:      wrapped_service.BaseName + "_TutorialMethodImpl",
		IfaceName: iface.Name,
		Imports:   gogen.NewImports(pkg.Name),
	}

	server.Imports.AddPackages("context")

	slog.Info(fmt.Sprintf("Generating %v/%v", server.Package.PackageName, iface.Name))
	outputFile := filepath.Join(server.Package.Path, iface.Name+".go")
	return gogen.ExecuteTemplateToFile("Tutorial", serverTemplate, server, outputFile)
}

type serverArgs struct {
	Package   golang.PackageInfo
	Service   *gocode.ServiceInterface
	Iface     *gocode.ServiceInterface
	Name      string
	IfaceName string
	Imports   *gogen.Imports
}

func generateClientSideInterfaces(builder golang.ModuleBuilder, iface *gocode.ServiceInterface, outputPackage string) error {
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
	return gogen.ExecuteTemplateToFile("Tutorial", clientTemplate, server, outputFile)
}

var serverTemplate = `// Blueprint: Auto-generated by Tutorial Plugin
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
{{ range $_, $f := .Service.Methods }}
func (handler *{{$receiver}}) {{$f.Name -}} ({{ArgVarsAndTypes $f "ctx context.Context"}}) ({{RetTypes $f "error"}}) {
	return handler.Service.{{$f.Name}}({{ArgVars $f "ctx"}})
}
{{end}}
func (handler *{{$receiver}}) Hello(ctx context.Context) (string, error) {
	return "Hello!", nil
}
`

var clientTemplate = `// Blueprint: Auto-generated by HealthChecker plugin
package {{.Package.ShortName}}

{{.Imports}}

type {{.IfaceName}} interface {
	{{range $_, $f := .Iface.Methods -}}
	{{Signature $f}}
	{{end}}
}
`
