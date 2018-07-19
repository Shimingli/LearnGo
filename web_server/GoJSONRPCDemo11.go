package main

import (
	"fmt"
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
)

//使用Go提供的json-rpc 标准包
func init() {
	fmt.Println("JSON RPC 采用了JSON，而不是 gob编码，和RPC概念一模一样，")
}
type Work struct {
	Who,DoWhat string
}

type DemoM string

func (m *DemoM) DoWork(w *Work,whoT *string) error  {
    *whoT="是谁："+w.Who+"，在做什么---"+w.DoWhat
	return nil
}

func main() {
    str:=new(DemoM)
    rpc.Register(str)

    tcpAddr,err:=net.ResolveTCPAddr("tcp",":8080")
	if  err!=nil{
		fmt.Println("大哥发生错误了啊，请看错误 ResolveTCPAddr err=",err)
	}

    listener,err:=net.ListenTCP("tcp",tcpAddr)
	if err!=nil {
		fmt.Println("发生错误了--》err=",err)
	}

	for  {
		 conn,err:= listener.Accept()
		if err!=nil {
			continue
		}
		jsonrpc.ServeConn(conn)

	}








}

