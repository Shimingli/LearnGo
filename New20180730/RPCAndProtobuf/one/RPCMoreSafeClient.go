package main

import (
	"fmt"
	"net/rpc"
	"log"
)

func init() {
	fmt.Println("更加安全的RPC的client端")
}

func main() {

	// 原来的client端  我们也要做好更加的 升级的代码
	//client, err := rpc.Dial("tcp", "localhost:12345")
	//if err != nil {
	//	log.Fatal("dialing:", err)
	//}
	//
	//var reply string
	//err = client.Call("path/to/pkg.HelloService"+".Hello", "hello", &reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("客户端接收到的信息：",reply)

	//接口规范部分增加对客户端的简单包装   todo 新的Demo
 	client,err:=DialHelloService("tcp",":12345")
	if err!=nil {
		log.Fatal("Dial Error ",err)
	}
	var reply  string
	client.Hello("我是请求的方 ",&reply)
 	fmt.Println("我是安全的 RPC的客户端 ：： reply=",reply)


}


//为了简化 客户端调用RPC函数  ，在接口规范部分增加对客户端的简单包装
type HelloServiceClient struct {
	*rpc.Client
}

//运算符	描述	实例
//&	返回变量存储地址	&a; 将给出变量的实际地址。
//*	指针变量。	*a; 是一个指针变量           *a 是一个指针的变量
func DialHelloService(network,address string)(*HelloServiceClient ,error)  {
	client, err := rpc.Dial(network,address)
	if err!=nil {
		return nil,err
	}
     //返回变量存储地址
	return &HelloServiceClient{client},nil
}
// "path/to/pkg.HelloService" 是可以抽出来
func (p* HelloServiceClient)Hello(request string,reply *string)error{
	fmt.Println("我是客户端的调用的RPC的方法")
	return p.Client.Call("path/to/pkg.HelloService"+".Hello",request,reply)
}


