<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# tutorial

```go
import "github.com/blueprint-uservices/tutorial/plugins/tutorial"
```

Package tutorial implements the following wiring functions:

1. AddHelloMethod: Adds a new method called \`HelloNew\` to the service interface.
2. AddHelloParam: Adds an extra call parameter and an extra return parameter to every method in the service.
3. Instrument: Adds logging statements to both the server and client side of every service.

## Index

- [func AddHelloMethod\(spec wiring.WiringSpec, serviceName string\)](<#AddHelloMethod>)
- [func AddHelloParam\(spec wiring.WiringSpec, serviceName string\)](<#AddHelloParam>)
- [func Instrument\(spec wiring.WiringSpec, serviceName string\)](<#Instrument>)
- [type HelloInstrumentClientWrapper](<#HelloInstrumentClientWrapper>)
  - [func \(node \*HelloInstrumentClientWrapper\) AddInstantiation\(builder golang.NamespaceBuilder\) error](<#HelloInstrumentClientWrapper.AddInstantiation>)
  - [func \(node \*HelloInstrumentClientWrapper\) AddInterfaces\(builder golang.ModuleBuilder\) error](<#HelloInstrumentClientWrapper.AddInterfaces>)
  - [func \(node \*HelloInstrumentClientWrapper\) GenerateFuncs\(builder golang.ModuleBuilder\) error](<#HelloInstrumentClientWrapper.GenerateFuncs>)
  - [func \(node \*HelloInstrumentClientWrapper\) GetInterface\(ctx ir.BuildContext\) \(service.ServiceInterface, error\)](<#HelloInstrumentClientWrapper.GetInterface>)
  - [func \(node \*HelloInstrumentClientWrapper\) ImplementsGolangNode\(\)](<#HelloInstrumentClientWrapper.ImplementsGolangNode>)
  - [func \(node \*HelloInstrumentClientWrapper\) Name\(\) string](<#HelloInstrumentClientWrapper.Name>)
  - [func \(node \*HelloInstrumentClientWrapper\) String\(\) string](<#HelloInstrumentClientWrapper.String>)
- [type HelloInstrumentServerWrapper](<#HelloInstrumentServerWrapper>)
  - [func \(node \*HelloInstrumentServerWrapper\) AddInstantiation\(builder golang.NamespaceBuilder\) error](<#HelloInstrumentServerWrapper.AddInstantiation>)
  - [func \(node \*HelloInstrumentServerWrapper\) AddInterfaces\(builder golang.ModuleBuilder\) error](<#HelloInstrumentServerWrapper.AddInterfaces>)
  - [func \(node \*HelloInstrumentServerWrapper\) GenerateFuncs\(builder golang.ModuleBuilder\) error](<#HelloInstrumentServerWrapper.GenerateFuncs>)
  - [func \(node \*HelloInstrumentServerWrapper\) GetInterface\(ctx ir.BuildContext\) \(service.ServiceInterface, error\)](<#HelloInstrumentServerWrapper.GetInterface>)
  - [func \(node \*HelloInstrumentServerWrapper\) ImplementsGolangNode\(\)](<#HelloInstrumentServerWrapper.ImplementsGolangNode>)
  - [func \(node \*HelloInstrumentServerWrapper\) Name\(\) string](<#HelloInstrumentServerWrapper.Name>)
  - [func \(node \*HelloInstrumentServerWrapper\) String\(\) string](<#HelloInstrumentServerWrapper.String>)
- [type HelloMethodWrapper](<#HelloMethodWrapper>)
  - [func \(node \*HelloMethodWrapper\) AddInstantiation\(builder golang.NamespaceBuilder\) error](<#HelloMethodWrapper.AddInstantiation>)
  - [func \(node \*HelloMethodWrapper\) AddInterfaces\(builder golang.ModuleBuilder\) error](<#HelloMethodWrapper.AddInterfaces>)
  - [func \(node \*HelloMethodWrapper\) GenerateFuncs\(builder golang.ModuleBuilder\) error](<#HelloMethodWrapper.GenerateFuncs>)
  - [func \(node \*HelloMethodWrapper\) GetInterface\(ctx ir.BuildContext\) \(service.ServiceInterface, error\)](<#HelloMethodWrapper.GetInterface>)
  - [func \(node \*HelloMethodWrapper\) ImplementsGolangNode\(\)](<#HelloMethodWrapper.ImplementsGolangNode>)
  - [func \(node \*HelloMethodWrapper\) Name\(\) string](<#HelloMethodWrapper.Name>)
  - [func \(node \*HelloMethodWrapper\) String\(\) string](<#HelloMethodWrapper.String>)
- [type HelloParamClientWrapper](<#HelloParamClientWrapper>)
  - [func \(node \*HelloParamClientWrapper\) AddInstantiation\(builder golang.NamespaceBuilder\) error](<#HelloParamClientWrapper.AddInstantiation>)
  - [func \(node \*HelloParamClientWrapper\) AddInterfaces\(builder golang.ModuleBuilder\) error](<#HelloParamClientWrapper.AddInterfaces>)
  - [func \(node \*HelloParamClientWrapper\) GenerateFuncs\(builder golang.ModuleBuilder\) error](<#HelloParamClientWrapper.GenerateFuncs>)
  - [func \(node \*HelloParamClientWrapper\) GetInterface\(ctx ir.BuildContext\) \(service.ServiceInterface, error\)](<#HelloParamClientWrapper.GetInterface>)
  - [func \(node \*HelloParamClientWrapper\) ImplementsGolangNode\(\)](<#HelloParamClientWrapper.ImplementsGolangNode>)
  - [func \(node \*HelloParamClientWrapper\) Name\(\) string](<#HelloParamClientWrapper.Name>)
  - [func \(node \*HelloParamClientWrapper\) String\(\) string](<#HelloParamClientWrapper.String>)
- [type HelloParamServerWrapper](<#HelloParamServerWrapper>)
  - [func \(node \*HelloParamServerWrapper\) AddInstantiation\(builder golang.NamespaceBuilder\) error](<#HelloParamServerWrapper.AddInstantiation>)
  - [func \(node \*HelloParamServerWrapper\) AddInterfaces\(builder golang.ModuleBuilder\) error](<#HelloParamServerWrapper.AddInterfaces>)
  - [func \(node \*HelloParamServerWrapper\) GenerateFuncs\(builder golang.ModuleBuilder\) error](<#HelloParamServerWrapper.GenerateFuncs>)
  - [func \(node \*HelloParamServerWrapper\) GetInterface\(ctx ir.BuildContext\) \(service.ServiceInterface, error\)](<#HelloParamServerWrapper.GetInterface>)
  - [func \(node \*HelloParamServerWrapper\) ImplementsGolangNode\(\)](<#HelloParamServerWrapper.ImplementsGolangNode>)
  - [func \(node \*HelloParamServerWrapper\) Name\(\) string](<#HelloParamServerWrapper.Name>)
  - [func \(node \*HelloParamServerWrapper\) String\(\) string](<#HelloParamServerWrapper.String>)


<a name="AddHelloMethod"></a>
## func [AddHelloMethod](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/wiring.go#L17>)

```go
func AddHelloMethod(spec wiring.WiringSpec, serviceName string)
```

[AddHelloMethod](<#AddHelloMethod>) can be called from the wiring specification to add a \`HelloNew\` method to Service with name \`serviceName\`.

<a name="AddHelloParam"></a>
## func [AddHelloParam](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/wiring.go#L45>)

```go
func AddHelloParam(spec wiring.WiringSpec, serviceName string)
```

[AddHelloParam](<#AddHelloParam>) can be called from the wiring specification to add an extra call parameter and an extra return parameter to every method exposed by the Service with name \`serviceName\`.

<a name="Instrument"></a>
## func [Instrument](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/wiring.go#L88>)

```go
func Instrument(spec wiring.WiringSpec, serviceName string)
```

[Instrument](<#Instrument>) can be called from the wiring specification to add logging statements to every method in both the server and client side code of Service with \`serviceName\`.

<a name="HelloInstrumentClientWrapper"></a>
## type [HelloInstrumentClientWrapper](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L148-L157>)

Blueprint IRNode for representing the wrapper node that instruments every client\-side method in the node that gets wrapped

```go
type HelloInstrumentClientWrapper struct {
    golang.Service
    golang.GeneratesFuncs
    golang.Instantiable

    InstanceName string
    Wrapped      golang.Service
    // contains filtered or unexported fields
}
```

<a name="HelloInstrumentClientWrapper.AddInstantiation"></a>
### func \(\*HelloInstrumentClientWrapper\) [AddInstantiation](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L206>)

```go
func (node *HelloInstrumentClientWrapper) AddInstantiation(builder golang.NamespaceBuilder) error
```

Implements golang.Instantiable

<a name="HelloInstrumentClientWrapper.AddInterfaces"></a>
### func \(\*HelloInstrumentClientWrapper\) [AddInterfaces](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L173>)

```go
func (node *HelloInstrumentClientWrapper) AddInterfaces(builder golang.ModuleBuilder) error
```

Implements golang.ProvidesInterface

<a name="HelloInstrumentClientWrapper.GenerateFuncs"></a>
### func \(\*HelloInstrumentClientWrapper\) [GenerateFuncs](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L197>)

```go
func (node *HelloInstrumentClientWrapper) GenerateFuncs(builder golang.ModuleBuilder) error
```

Implements golang.GeneratesFuncs

<a name="HelloInstrumentClientWrapper.GetInterface"></a>
### func \(\*HelloInstrumentClientWrapper\) [GetInterface](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L192>)

```go
func (node *HelloInstrumentClientWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error)
```

Implements service.ServiceNode

<a name="HelloInstrumentClientWrapper.ImplementsGolangNode"></a>
### func \(\*HelloInstrumentClientWrapper\) [ImplementsGolangNode](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L160>)

```go
func (node *HelloInstrumentClientWrapper) ImplementsGolangNode()
```

Implements ir.IRNode

<a name="HelloInstrumentClientWrapper.Name"></a>
### func \(\*HelloInstrumentClientWrapper\) [Name](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L163>)

```go
func (node *HelloInstrumentClientWrapper) Name() string
```

Implements ir.IRNode

<a name="HelloInstrumentClientWrapper.String"></a>
### func \(\*HelloInstrumentClientWrapper\) [String](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L168>)

```go
func (node *HelloInstrumentClientWrapper) String() string
```

Implements ir.IRNode

<a name="HelloInstrumentServerWrapper"></a>
## type [HelloInstrumentServerWrapper](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L18-L27>)

Blueprint IRNode for representing the wrapper node that instruments every server\-side method in the node that gets wrapped

```go
type HelloInstrumentServerWrapper struct {
    golang.Service
    golang.GeneratesFuncs
    golang.Instantiable

    InstanceName string
    Wrapped      golang.Service
    // contains filtered or unexported fields
}
```

<a name="HelloInstrumentServerWrapper.AddInstantiation"></a>
### func \(\*HelloInstrumentServerWrapper\) [AddInstantiation](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L76>)

```go
func (node *HelloInstrumentServerWrapper) AddInstantiation(builder golang.NamespaceBuilder) error
```

Implements golang.Instantiable

<a name="HelloInstrumentServerWrapper.AddInterfaces"></a>
### func \(\*HelloInstrumentServerWrapper\) [AddInterfaces](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L43>)

```go
func (node *HelloInstrumentServerWrapper) AddInterfaces(builder golang.ModuleBuilder) error
```

Implements golang.ProvidesInterface

<a name="HelloInstrumentServerWrapper.GenerateFuncs"></a>
### func \(\*HelloInstrumentServerWrapper\) [GenerateFuncs](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L67>)

```go
func (node *HelloInstrumentServerWrapper) GenerateFuncs(builder golang.ModuleBuilder) error
```

Implements golang.GeneratesFuncs

<a name="HelloInstrumentServerWrapper.GetInterface"></a>
### func \(\*HelloInstrumentServerWrapper\) [GetInterface](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L62>)

```go
func (node *HelloInstrumentServerWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error)
```

Implements service.ServiceNode

<a name="HelloInstrumentServerWrapper.ImplementsGolangNode"></a>
### func \(\*HelloInstrumentServerWrapper\) [ImplementsGolangNode](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L30>)

```go
func (node *HelloInstrumentServerWrapper) ImplementsGolangNode()
```

Implements ir.IRNode

<a name="HelloInstrumentServerWrapper.Name"></a>
### func \(\*HelloInstrumentServerWrapper\) [Name](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L33>)

```go
func (node *HelloInstrumentServerWrapper) Name() string
```

Implements ir.IRNode

<a name="HelloInstrumentServerWrapper.String"></a>
### func \(\*HelloInstrumentServerWrapper\) [String](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_instrument.go#L38>)

```go
func (node *HelloInstrumentServerWrapper) String() string
```

Implements ir.IRNode

<a name="HelloMethodWrapper"></a>
## type [HelloMethodWrapper](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L18-L27>)

Blueprint IRNode for representing the wrapper node that adds a \`Hello\` method to the wrapped IRNode.

```go
type HelloMethodWrapper struct {
    golang.Service
    golang.GeneratesFuncs
    golang.Instantiable

    InstanceName string
    Wrapped      golang.Service
    // contains filtered or unexported fields
}
```

<a name="HelloMethodWrapper.AddInstantiation"></a>
### func \(\*HelloMethodWrapper\) [AddInstantiation](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L109>)

```go
func (node *HelloMethodWrapper) AddInstantiation(builder golang.NamespaceBuilder) error
```

Implements golang.Instantiable

<a name="HelloMethodWrapper.AddInterfaces"></a>
### func \(\*HelloMethodWrapper\) [AddInterfaces](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L43>)

```go
func (node *HelloMethodWrapper) AddInterfaces(builder golang.ModuleBuilder) error
```

IMplements golang.ProvidesInterface

<a name="HelloMethodWrapper.GenerateFuncs"></a>
### func \(\*HelloMethodWrapper\) [GenerateFuncs](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L92>)

```go
func (node *HelloMethodWrapper) GenerateFuncs(builder golang.ModuleBuilder) error
```

Implements golang.GeneratesFuncs

<a name="HelloMethodWrapper.GetInterface"></a>
### func \(\*HelloMethodWrapper\) [GetInterface](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L87>)

```go
func (node *HelloMethodWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error)
```

Implements service.ServiceNode

<a name="HelloMethodWrapper.ImplementsGolangNode"></a>
### func \(\*HelloMethodWrapper\) [ImplementsGolangNode](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L30>)

```go
func (node *HelloMethodWrapper) ImplementsGolangNode()
```

Implements ir.IRNode

<a name="HelloMethodWrapper.Name"></a>
### func \(\*HelloMethodWrapper\) [Name](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L33>)

```go
func (node *HelloMethodWrapper) Name() string
```

Implements ir.IRNode

<a name="HelloMethodWrapper.String"></a>
### func \(\*HelloMethodWrapper\) [String](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_method.go#L38>)

```go
func (node *HelloMethodWrapper) String() string
```

Implements ir.IRNode

<a name="HelloParamClientWrapper"></a>
## type [HelloParamClientWrapper](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L203-L212>)

Blueprint IRNode for representing the wrapper node that adds an additional call parameter and an additional return parameter to every client\-side method in the node that gets wrapped

```go
type HelloParamClientWrapper struct {
    golang.Service
    golang.GeneratesFuncs
    golang.Instantiable

    InstanceName string
    Wrapped      golang.Service
    // contains filtered or unexported fields
}
```

<a name="HelloParamClientWrapper.AddInstantiation"></a>
### func \(\*HelloParamClientWrapper\) [AddInstantiation](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L287>)

```go
func (node *HelloParamClientWrapper) AddInstantiation(builder golang.NamespaceBuilder) error
```

Implements golang.Instantiable

<a name="HelloParamClientWrapper.AddInterfaces"></a>
### func \(\*HelloParamClientWrapper\) [AddInterfaces](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L228>)

```go
func (node *HelloParamClientWrapper) AddInterfaces(builder golang.ModuleBuilder) error
```

Implements golang.ProvidesInterface

<a name="HelloParamClientWrapper.GenerateFuncs"></a>
### func \(\*HelloParamClientWrapper\) [GenerateFuncs](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L270>)

```go
func (node *HelloParamClientWrapper) GenerateFuncs(builder golang.ModuleBuilder) error
```

Implements golang.GeneratesFuncs

<a name="HelloParamClientWrapper.GetInterface"></a>
### func \(\*HelloParamClientWrapper\) [GetInterface](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L265>)

```go
func (node *HelloParamClientWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error)
```

Implements service.ServiceNode

<a name="HelloParamClientWrapper.ImplementsGolangNode"></a>
### func \(\*HelloParamClientWrapper\) [ImplementsGolangNode](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L215>)

```go
func (node *HelloParamClientWrapper) ImplementsGolangNode()
```

Implements ir.IRNode

<a name="HelloParamClientWrapper.Name"></a>
### func \(\*HelloParamClientWrapper\) [Name](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L218>)

```go
func (node *HelloParamClientWrapper) Name() string
```

Implements ir.IRNode

<a name="HelloParamClientWrapper.String"></a>
### func \(\*HelloParamClientWrapper\) [String](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L223>)

```go
func (node *HelloParamClientWrapper) String() string
```

Implements ir.IRNode

<a name="HelloParamServerWrapper"></a>
## type [HelloParamServerWrapper](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L18-L27>)

Blueprint IRNode for representing the wrapper node that adds an additional call parameter and an additional return parameter to every server\-side method in the node that gets wrapped

```go
type HelloParamServerWrapper struct {
    golang.Service
    golang.GeneratesFuncs
    golang.Instantiable

    InstanceName string
    Wrapped      golang.Service
    // contains filtered or unexported fields
}
```

<a name="HelloParamServerWrapper.AddInstantiation"></a>
### func \(\*HelloParamServerWrapper\) [AddInstantiation](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L110>)

```go
func (node *HelloParamServerWrapper) AddInstantiation(builder golang.NamespaceBuilder) error
```

Implements golang.Instantiable

<a name="HelloParamServerWrapper.AddInterfaces"></a>
### func \(\*HelloParamServerWrapper\) [AddInterfaces](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L43>)

```go
func (node *HelloParamServerWrapper) AddInterfaces(builder golang.ModuleBuilder) error
```

Implements golang.ProvidesInterface

<a name="HelloParamServerWrapper.GenerateFuncs"></a>
### func \(\*HelloParamServerWrapper\) [GenerateFuncs](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L93>)

```go
func (node *HelloParamServerWrapper) GenerateFuncs(builder golang.ModuleBuilder) error
```

Implements golang.GeneratesFuncs

<a name="HelloParamServerWrapper.GetInterface"></a>
### func \(\*HelloParamServerWrapper\) [GetInterface](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L88>)

```go
func (node *HelloParamServerWrapper) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error)
```

Implements service.ServiceNode

<a name="HelloParamServerWrapper.ImplementsGolangNode"></a>
### func \(\*HelloParamServerWrapper\) [ImplementsGolangNode](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L30>)

```go
func (node *HelloParamServerWrapper) ImplementsGolangNode()
```

Implements ir.IRNode

<a name="HelloParamServerWrapper.Name"></a>
### func \(\*HelloParamServerWrapper\) [Name](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L33>)

```go
func (node *HelloParamServerWrapper) Name() string
```

Implements ir.IRNode

<a name="HelloParamServerWrapper.String"></a>
### func \(\*HelloParamServerWrapper\) [String](<https://github.com/blueprint-uservices/tutorial/blob/main/plugins/tutorial/ir_param.go#L38>)

```go
func (node *HelloParamServerWrapper) String() string
```

Implements ir.IRNode

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
