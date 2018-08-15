package main

import (
	"fmt"
	"net/rpc"
	"log"
)

func init() {
	fmt.Println("用Go来模拟 客户端的请求")
}

func main() {
	client,err:=rpc.Dial("tcp",":1234")
	if err!=nil {
		log.Fatal("Dial Error=",err)
	}
	var  reply string
	//调用client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，第二和第三个参数分别我们定义RPC方法的两个参数。
	errr:=client.Call("Shiming.Hello","shiming",&reply)
	if errr!=nil {
		log.Fatal(errr)
	}

	fmt.Println("客户端接收到的信息：",reply)




}