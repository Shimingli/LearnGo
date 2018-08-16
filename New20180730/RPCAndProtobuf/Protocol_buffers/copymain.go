package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/proto"  // https://github.com/golang/protobuf  自己重新下载了一个protobuf
)

func init() {
	//Go语言的包只能静态导入，我们无法向已经安装的protoc-gen-go添加我们新编写的插件。我们将重新克隆protoc-gen-go对应的main函数
	fmt.Println("Go语言的包只能静态导入，我们无法向已经安装的protoc-gen-go添加我们新编写的插件。我们将重新克隆protoc-gen-go对应的main函数")
}
func main() {
	g := generator.New()
    fmt.Println("我是copy的procotc_gen_go的 main.go 的函数")
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}