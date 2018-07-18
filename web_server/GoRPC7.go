package main

import (
	"fmt"
	"net/rpc"
	"net/http"
	"errors"
)

func main() {
	fmt.Println("过学习我们了解了Socket和HTTP采用的是类似\"信息交换\"模式，即客户端发送一条信息到服务端，然后(一般来说)服务器端都会返回一定的信息以表示响应。客户端和服务端之间约定了交互信息的格式，以便双方都能够解析交互所产生的信息。但是很多独立的应用并没有采用这种模式，而是采用类似常规的函数调用的方式来完成想要的功能。")

  //RPC 就想实现函数调用的模式的网络化。客户端就像调用本地函数一样，然后在客户端吧这些参数打包之后通过网络传递到服务端，服务端解包到处理过程中执行，然后执行的结果反馈给客户端

  fmt.Println("RPC 远程过程调用协议，是一种通过网络从计算机程序上请求服务，而不需要了解底层网络技术的协议。它假定某些传输协议的存在，如TCP，以便为通信程序之间携带信息数据，通过它可以使函数调用模式网络化。在OSI网络通信模型中欧冠，RPC跨越了传输层和应用层。RPC使得开发包括网络分布式多程序在内的应用程序更加的容易")

  //  todo  RPC工作原理     注意图片    rpc.png RPC工作流程图
	//1.调用客户端句柄；执行传送参数
	//2.调用本地系统内核发送网络消息
	//3.消息传送到远程主机
	//4.服务器句柄得到消息并取得参数
	//5.执行远程过程
	//6.执行的过程将结果返回服务器句柄
	//7.服务器句柄返回结果，调用远程系统内核
	//8.消息传回本地主机
	//9.客户句柄由内核接收消息
	//10.客户接收句柄返回的数据



	fmt.Println("Go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC。但Go的RPC包是独一无二的RPC，它和传统的RPC系统不同，它只支持Go开发的服务器与客户端之间的交互，因为在内部，它们采用了Gob来编码。")


	//Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
	//
	//函数必须是导出的(首字母大写)
	//必须有两个导出类型的参数，
	//第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
	//函数还要有一个返回值error

	//func (t *T) MethodName(argType T1, replyType *T2) error
	//T、T1和T2类型必须能被encoding/gob包编解码。
	//任何的RPC都需要通过网络来传递数据，Go RPC可以利用HTTP和TCP来传递数据，利用HTTP的好处是可以直接复用net/http里面的一些函数。


     rpcDemo()
}

type Arith int
func rpcDemo() {
	arith:=new(Arith)
	//arith=== 0xc04204e090
	fmt.Println("arith===",arith)
	rpc.Register(arith)
    rpc.HandleHTTP()
	err:=http.ListenAndServe(":1234",nil)
	if err != nil {
		fmt.Println("err=====",err.Error())
	}
}
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}