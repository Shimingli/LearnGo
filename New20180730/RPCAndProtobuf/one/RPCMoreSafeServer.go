package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
)

func init() {
	fmt.Println("更加安全的 RPCDemo的服务端")
}

func main() {
   fmt.Println("在涉及RPC的应用中，作为开发人员一般至少有三种角色：首选是服务端实现RPC方法的开发人员，其次是客户端调用RPC方法的人员，最后也是最重要的是制定服务端和客户端RPC接口规范的设计人员")

	//注册 RegisterHelloService注册服务时，编译器会要求传入的对象满足HelloServiceInterface接口
	//var _helloServiceInterface = (*HelloServiceInterface)(nil)
	//var serv=HelloServiceInterface(nil)
   RegisterHelloService(new(HelloServiceSafe))
   listen,err:= net.Listen("tcp",":12345")
	if err!=nil {
		log.Fatal("RPC Error:",err)
	}
	// 添加循环 一直接受着我的信息
	for   {
		conn,err:=listen.Accept()
		if err!=nil {
			log.Fatal("Accept Error=",err)
		}
		// 使用协程
		go rpc.ServeConn(conn)

	}
	//conn,err:= listen.Accept()
	//if err!=nil {
	//	log.Fatal("Accept Error=",err)
	//}
	//rpc.ServeConn(conn)

}
//重构HelloService服务，第一步需要明确服务的名字和接口

const ServiceName  = "path/to/pkg.HelloService"//为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀（这个是RPC服务抽象的包路径，并非完全等价Go语言的包路径）
//RPC服务的接口规范分为三个部分：首先是服务的名字，然后是服务要实现的详细方法列表，最后是注册该类型服务的函数
type HelloServiceInterface = interface {
  Hello(request string,reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error  {
	return rpc.RegisterName(ServiceName,svc)
}
// 实现了 HelloServiceInterface 这个对象
type HelloServiceSafe struct {

}

func (h *HelloServiceSafe) Hello(request string,reply *string) error  {
	*reply = "hello: HelloServiceSafe ---" + request
	fmt.Println("safe 执行了啊")
	return nil
}
