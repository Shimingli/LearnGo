package main

import (
	"fmt"
	"net"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func init() {

	//无论采用何种的语言，只要遵循同样的json结构，以同样的流程就可以和Go语言编程的RPC服务进行通信
	fmt.Println("Cross Language RPC 的 客户端")
}

func main() {
	//client,_:=Dial("tcp",":1234")
	//
	//var reply string
	//client.MeTest("nihao ",&reply)
	//fmt.Println(reply)

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("s/s/shiming"+".MeTest", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

//type MeClient struct {
//	*net.Conn
//}
//
//func Dial(netWork,address string)(*MeClient,error)  {
//	conn,err:= net.Dial(netWork,address)
//	return &MeClient{conn},err
//}
//
//func (p* MeClient)MeTest(request string,reply *string)error{
//	fmt.Println("我是客户端的调用的RPC的方法")
//	rpc.NewClientWithCodec(jsonrpc.NewClientCodec(p.Client))
//	return p.Client.Call("path/to/pkg.HelloService"+".MeTest",request,reply)
//}