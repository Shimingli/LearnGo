package generator

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
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
	p.P("// TODO: import code")
}
//自定义的genImportCode和genServiceCode方法只是输出一行简单的注释
func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}