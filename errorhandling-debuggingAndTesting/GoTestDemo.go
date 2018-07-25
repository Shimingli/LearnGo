package main

import "fmt"

func main() {
	 fmt.Println("Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架类似，你可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例")


    //建议安装gotests插件自动生成测试代码:
	//go get -u -v github.com/cweill/gotests/...

	//如何编写测试用例
	//由于go test命令只能在一个相应的目录下执行所有文件，所以我们接下来新建一个项目目录gotest,这样我们所有的代码和测试代码都在这个目录下。
	//
	//接下来我们在该目录下面创建两个文件：gotest.go和gotest_test.go
	// todo  注意看  gotest  包下的函数   gotest.go和gotest_test.go


	// todo 如何编写压力测试
    //	go test -test.bench=".*"

    /*

 E:\new_code\GoDemo\gotest2>go test -test.bench=".*"
goos: windows
goarch: amd64
Benchmark_Division-4                    2000000000               1.19 ns/op
Benchmark_TimeConsumingFunction-4       2000000000               0.75 ns/op
PASS
ok      _/E_/new_code/GoDemo/gotest2    4.105s
     */


}
