package main

import (
	"fmt"
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func init() {
	fmt.Println("跨语言的 RPC")
	// 得益于 RPC的框架的设计，Go语言的RPC其实很容易实现跨语言的支持

	//go 语言的RPC框架有两个比较特色的设计 1、RPC数据打包时候可以通过插件实现自定义的编码和解码 2、RPC建立在抽象的 `io.ReadWriteCloser `接口上，这样我们就可以将RPC假设在不同的通讯协议上
}

func main() {
    ReMeSerivce(new(MrServiceEx))
    l,_:= net.Listen("tcp",":1234")
	for   {
		conn,_:=l.Accept()
		//用rpc.ServeCodec函数替代了rpc.ServeConn函数，传入的参数是针对服务端的json编解码器
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

const ServiceNameMy  ="s/s/shiming"

type MeService = interface {
    MeTest(request string,reply *string) error
}

func ReMeSerivce(me MeService)error  {
	return rpc.RegisterName(ServiceNameMy,me)
}

type MrServiceEx struct {

}

func (p *MrServiceEx)MeTest(request string,reply *string) error {
	*reply ="nihao shiming "+request
	fmt.Println("我是服务器的 MeTest 方法，我已经执行了 ")
	return nil
}
