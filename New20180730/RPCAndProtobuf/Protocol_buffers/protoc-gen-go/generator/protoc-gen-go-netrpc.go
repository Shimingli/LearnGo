package generator

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"bytes"
	"log"
	"html/template"
)

func init() {
	fmt.Println("设计 Net RPC 的插件")
	//该插件需要先通过generator.RegisterPlugin函数注册插件，可以在init函数中完成
	//generator.RegisterPlugin(new(netrpcPlugin))
}

type netrpcPlugin struct{ *generator.Generator }

//首先Name方法返回插件的名字
func (p *netrpcPlugin) Name() string   {
	return "netrpc"
}

//netrpcPlugin插件内置了一个匿名的*generator.Generator成员，然后在Init初始化的时候用参数g进行初始化，
// 因此插件是从g参数对象继承了全部的公有方法
func (p *netrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
	}
//其中GenerateImports方法调用自定义的genImportCode函数生成导入代码
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}
//Generate方法调用自定义的genServiceCode方法生成每个服务的代码。
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}
//自定义的genImportCode和genServiceCode方法只是输出一行简单的注释
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	// todo 现在开始继续完善netrpcPlugin插件，最终目标是生成RPC安全接口。
	//p.P("// TODO: import code")

	p.P(`import "net/rpc"`)
}
//自定义的genImportCode和genServiceCode方法只是输出一行简单的注释
//然后要在自定义的genServiceCode方法中为每个服务生成相关的代码。分析可以发现每个服务最重要的是服务的名字，然后每个服务有一组方法。而对于服务定义的方法，最重要的是方法的名字，还有输入参数和输出参数类型的名字。
func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	//p.P("// TODO: service code, Name = " + svc.GetName())
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())

}
//定义了一个ServiceSpec类型，用于描述服务的元信息：
type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}
// 新建一个buildServiceSpec方法用来解析每个服务的ServiceSpec元信息
/*
输入参数是*descriptor.ServiceDescriptorProto类型，完整描述了一个服务的所有信息。然后通过svc.GetName()就可以获取Protobuf文件中定义的服务的名字。Protobuf文件中的名字转为Go语言的名字后，需要通过generator.CamelCase函数进行一次转换。类似的，在for循环中我们通过m.GetName()获取方法的名字，然后再转为Go语言中对应的名字。比较复杂的是对输入和输出参数名字的解析：首先需要通过m.GetInputType()获取输入参数的类型，然后通过p.ObjectNamed获取类型对应的类对象信息，最后获取类对象的名字
 */
func (p *netrpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}

	return spec
}

const tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
	{{- end}}
}

func Register{{.ServiceName}}(
	srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
	if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
		return err
	}
	return nil
}

type {{.ServiceName}}Client struct {
	*rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}}(network, address string) (
	*{{.ServiceName}}Client, error,
) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(
	in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}},
) error {
	return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`