package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
)

func init() {

}

func main() {
	fmt.Println("RPC 的Hello，World")
	rpcDemo()
}
func rpcDemo() {
	//注册
	rpc.RegisterName("Shiming",new(HelloService))
	listen,err:= net.Listen("tcp",":1234")
	if err!=nil {
		log.Fatal("RPC Error:",err)
	}
	conn,err:= listen.Accept()
	if err!=nil {
		log.Fatal("Accept Error=",err)
	}
	rpc.ServeConn(conn)

}

//Go RPC 的函数只有符合四个条件才能够被远程访问，不然会被忽略
//函数必须是首字母大写（可以导出的）
//必须有两个导出类型的参数
//第一个参数是接受的参数，第二个参数是返回给客户端的参数，而且第二个参数是指针的类型
//函数还要有一个返回值error
type HelloService struct {

}

func (h *HelloService)Hello(request string,reply *string) error  {
	*reply="hello :"+request
	fmt.Println("RPC 服务端 Hello 方法执行了")
	return nil
}