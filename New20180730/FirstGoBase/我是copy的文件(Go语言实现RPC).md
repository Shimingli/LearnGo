####RPC定义，来源于百度百科
* RPC（Remote Procedure Call）—[远程过程调用](https://baike.baidu.com/item/%E8%BF%9C%E7%A8%8B%E8%BF%87%E7%A8%8B%E8%B0%83%E7%94%A8)，它是一种通过[网络](https://baike.baidu.com/item/%E7%BD%91%E7%BB%9C)从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。[RPC协议](https://baike.baidu.com/item/RPC%E5%8D%8F%E8%AE%AE)假定某些[传输协议](https://baike.baidu.com/item/%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)的存在，如TCP或UDP，为通信程序之间携带信息数据。在OSI[网络通信](https://baike.baidu.com/item/%E7%BD%91%E7%BB%9C%E9%80%9A%E4%BF%A1)模型中，RPC跨越了[传输层](https://baike.baidu.com/item/%E4%BC%A0%E8%BE%93%E5%B1%82)和[应用层](https://baike.baidu.com/item/%E5%BA%94%E7%94%A8%E5%B1%82)。RPC使得开发包括网络[分布式](https://baike.baidu.com/item/%E5%88%86%E5%B8%83%E5%BC%8F)多程序在内的应用程序更加容易。

* RPC采用客户机/服务器模式。请求程序就是一个客户机，而服务提供程序就是一个服务器。首先，客户机调用进程发送一个有进程参数的调用信息到服务进程，然后等待应答信息。在服务器端，进程保持睡眠状态直到调用信息到达为止。当一个调用信息到达，服务器获得进程参数，计算结果，发送答复[信息](https://baike.baidu.com/item/%E4%BF%A1%E6%81%AF)，然后等待下一个调用信息，最后，客户端调用进程接收答复信息，获得进程结果，然后调用执行继续进行。

* 有多种 RPC模式和执行。最初由 Sun 公司提出。IETF ONC 宪章重新修订了 Sun 版本，使得 ONC RPC 协议成为 IETF 标准协议。现在使用最普遍的模式和执行是开放式软件基础的分布式计算[环境](https://baike.baidu.com/item/%E7%8E%AF%E5%A2%83)（DCE）。
* 个人的理解：不用管什么底层网络技术的协议，是一种通过网络从计算机程序上请求服务，通俗一点，我们写代码，要在一个地方，比如安卓，就需要在一个工程里面，才可以调用到其他的程序代码执行的过程。Go语言提供RPC支持使得开发网络分布式多程序在内的应用程序更加的`easy`

#### RPC工作流程图


![图片来源于gitHub](https://upload-images.jianshu.io/upload_images/5363507-147298372fc05727.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


* 1.调用客户端句柄；执行传送参数
* 2.调用本地系统内核发送网络消息
* 3.消息传送到远程主机
* 4.服务器句柄得到消息并取得参数
* 5.执行远程过程
* 6.执行的过程将结果返回服务器句柄
* 7.服务器句柄返回结果，调用远程系统内核
* 8.消息传回本地主机
* 9.客户句柄由内核接收消息
* 10.客户接收句柄返回的数据

#### Go语言提供对RPC的支持：`HTTP、TCP、JSPNRPC`,但是在`Go`中`RPC`是独一无二的，它采用了`GoLang Gob`编码,只能支持Go语言！
* GoLang Gob:是Golang包自带的一个数据结构序列化的编码/解码工具。编码使用Encoder，解码使用Decoder。一种典型的应用场景就是RPC(remote procedure calls)。

#### HTTP RPC Demo


* 服务端的代码
```
package main

import (
	"fmt"
	"net/rpc"
	"net/http"
	"errors"
)
func main() {
     rpcDemo()
}
type Arith int
func rpcDemo() {
	arith:=new(Arith)
	//arith=== 0xc04204e090
	fmt.Println("arith===",arith)

	rpc.Register(arith)
	//HandleHTTP将RPC消息的HTTP处理程序注册到Debug服务器
	//DEFAUTUPCPATH和Debug调试路径上的调试处理程序。
	//仍然需要调用http.Services（），通常是在GO语句中。
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

//函数必须是导出的(首字母大写)
//必须有两个导出类型的参数，
//第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
//函数还要有一个返回值error
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("这个方法执行了啊---嘿嘿--- Multiply ",reply)
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	fmt.Println("这个方法执行了啊---嘿嘿--- Divide quo==",quo)
	return nil
}

```
* Go RPC 的函数只有符合四个条件才能够被远程访问，不然会被忽略
  *  函数必须是首字母大写（可以导出的）
  * 必须有两个导出类型的参数 
  * 第一个参数是接受的参数，第二个参数是返回给客户端的参数，而且第二个参数是指针的类型 
  * 函数还要有一个返回值`error` 

```
func (t *T) MethodName(argType T1, replyType *T2) error
```
* T、T1和T2类型必须能被`encoding/gob`包编解码。

* 客户端的代码
```

package main

import (
	"log"
	"fmt"
	"os"
	"net/rpc"
	"strconv"
)

type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {
	// 如果什么都不输入的话 ，就是以下的这个值
	//os***************** [C:\Users\win7\AppData\Local\Temp\go-build669605574\command-
	//line-arguments\_obj\exe\GoRPCWeb.exe 127.0.0.1] **********************

	fmt.Println("os*****************",os.Args,"**********************")
	if len(os.Args) != 4 { //   todo  第二个地址是  我们本地的地址
		fmt.Println("老子要退出了哦 傻逼 一号start--------》》》", os.Args[0], "《《《---------------server  end")
		os.Exit(1)
	}else{
		fmt.Println("长度是多少 "+strconv.Itoa( len(os.Args))+"才是准确的长度 哦---》")
	}
    //获取输入的地址是获取输入得 os 数据的 第一个位置的值
	serverAddress := os.Args[1]
    fmt.Println("severAddress==",serverAddress)
	// //DelayHTTP在指定的网络地址连接到HTTP RPC服务器
	///在默认HTTP RPC路径上监听。
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("发生错误了 在这里地方  DialHTTP", err)
	}
	i1,_:=strconv.Atoi( os.Args[2])
	i2,_:=strconv.Atoi( os.Args[3])
	args := ArgsTwo{i1, i2}
	var reply int
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Call Multiply  发生错误了哦   arith error:", err)
	}
	fmt.Printf("Arith 乘法: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientTwo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith 除法取整数: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
```

![运行的结果图](https://upload-images.jianshu.io/upload_images/5363507-2ebbd9d70053e8e3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
```
E:\new_code\GoDemo\web_server>go run GoRPCWeb8.go 127.0.0.1:1234  20 3
os***************** [C:\Users\win7\AppData\Local\Temp\go-build011170718\command-
line-arguments\_obj\exe\GoRPCWeb8.exe 127.0.0.1:1234 20 3] *********************
*
长度是多少 4才是准确的长度 哦---》
severAddress== 127.0.0.1:1234
Arith 乘法: 20*3=60
Arith 除法取整数: 20/3=6 余数 2
```
* go run GoRPCWeb8.go 127.0.0.1:1234  20 3 
   * go run 运行的命令 
   * GoRPCWeb8.go ：这是文件的名称，需要到指定的目录下启动`cmd` 
    * 127.0.0.1:1234  ： ip地址和端口号
   *  20 3 这是客服端传入的值：一个除数，一个被除数，传入到服务器做乘法运算 乘法: `20*3=60`和除法取整数: `20/3=6` 余数 `2`,怎么做的，客户端不关心，给到服务端去完成
* `os.Args[0]`=`[C:\Users\win7\AppData\Local\Temp\go-build011170718\command-
line-arguments\_obj\exe\GoRPCWeb8.exe 127.0.0.1:1234 20 3]`
```
	if len(os.Args) != 4 { //   todo  第二个地址是  我们本地的地址
		fmt.Println("老子要退出了哦 傻逼 一号start--------》》》", os.Args[0], "《《《---------------server  end")
		os.Exit(1)
	}else{
		fmt.Println("长度是多少 "+strconv.Itoa( len(os.Args))+"才是准确的长度 哦---》")
	}
```


#### TCP RPC Demo

* 基于TCP协议实现的RPC，服务端的代码
```
package main

import (
	"fmt"
	"net/rpc"
	"net"
	"os"
	"errors"
)

func init() {
	fmt.Println("基于TCP协议实现的RPC，服务端的代码如下")
}
type Me struct {
	A,B int
}
type You struct {
	CC,D int
}
type Num int

/*
Go RPC的函数只有符合下面的条件才能够被远程访问，不然会被忽略
1 函数必须是导出的（首字母大写）
2 必须有两个导出类型的参数
3 第一个参数是接受的参数，第二个参数是返回给客户端的参数，第二个参数必须是指正类型的
4 函数还必须有一个返回值error
 */
func (n *Num) M(args *Me,reply *int) error  {
	*reply=args.A * args.B
	return nil
}


func (n *Num) F(args * Me,u *You ) error  {
	if  args.B==0{
		return errors.New("输入不能够为0 被除数")
	}
	u.D=args.A/args.B
	u.CC=args.A % args.B
	return nil
}



func main() {
	//内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
	//new返回指针。
    num:=new(Num)
    rpc.Register(num)
    //ResolveTCPAddr返回TCP端点的地址。
	//网络必须是TCP网络名称。
    tcpAddr,err:=net.ResolveTCPAddr("tcp",":1234")

	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
    listener,err:=net.ListenTCP("tcp",tcpAddr)
	for  {
		// todo   需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc 来处理
		conn,err:=listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

```

* 基于TCP客户端代码
```
package main

import (
	"fmt"
	"os"
	"net/rpc"
	"log"
	"strconv"
)

func main() {
	fmt.Println("客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go")

	if len(os.Args)==4{
		fmt.Println("长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip=",os.Args[1],"嘿嘿,加上后面的被除数os.Args[2]=",os.Args[2],"和除数os.Args[3]=",os.Args[3])
		//os.Exit(1)
	}
    // 获取 ip 地址
    service:= os.Args[1]
    //连接 拨号连接到指定的网络地址的RPC服务器。
    client,err:=rpc.Dial("tcp",service)
	if err!=nil {
		log.Fatal("老子在连接Dial的发生了错误，我要退出了",err)
	}
	num1:=os.Args[2]
	i1,error1:=strconv.Atoi(num1)
	if error1!=nil {
		fmt.Println("自己不知道 自己输入错误了啊 请看error ：",error1)
		os.Exit(1)
	}
	num2:=os.Args[3]
	i2,error2:=strconv.Atoi(num2)
	if error2!=nil {
		fmt.Println("自己不知道 自己输入错误了啊 请看error ：",error2)
		os.Exit(1)
	}
	aa:=AAA{i1,i2}
	var reply  int
	err1:=client.Call("Num.M",aa,&reply)

	if err1 != nil{
		log.Fatal("我要退出了，因为我在Call的时候发生了 错误",err1)
	}
	fmt.Println("我进行正常结果如下")
	fmt.Printf("Num : %d*%d=%d\n",aa.A,aa.B,reply)

	var bb BDemo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err= client.Call("Num.F",aa,&bb)
	if err!=nil {
		log.Fatal("我对这个方法发生了过敏的反应 哈哈哈哈  err=====",err)
	}
	fmt.Printf("Num: %d/%d=%d 余数 %d\n",aa.A,aa.B,bb.DD,bb.CC)
	
}


// 定义两个类，那边需要操作的类
type AAA struct {
	A,B int
}
//记住这里不能够大写 两个连着一起大写 有点意思
//reading body gob: type mismatch: no fields matched compiling decoder for  DDDD
//  todo 为啥 第二个参数  只要是两个连在一起的DDDD   就会报错   reading body gob: type mismatch: no fields matched compiling decoder for
type BDemo struct {
	DD, CC int
}

```

* 运行结果图
![结果图](https://upload-images.jianshu.io/upload_images/5363507-a4168e88f7eaf474.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```
E:\new_code\GoDemo\web_server>go run GoTCPRPCWeb10.go 127.0.0.1:1234  20 1
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:1234 嘿嘿,加上后面
的被除数os.Args[2]= 20 和除数os.Args[3]= 1
我进行正常结果如下
Num : 20*1=20
Num: 20/1=0 余数 0

E:\new_code\GoDemo\web_server>go run GoTCPRPCWeb10.go 127.0.0.1:1234  20 2
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:1234 嘿嘿,加上后面
的被除数os.Args[2]= 20 和除数os.Args[3]= 2
我进行正常结果如下
Num : 20*2=40
Num: 20/2=0 余数 0

E:\new_code\GoDemo\web_server>go run GoTCPRPCWeb10.go 127.0.0.1:1234  20 3
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:1234 嘿嘿,加上后面
的被除数os.Args[2]= 20 和除数os.Args[3]= 3
我进行正常结果如下
Num : 20*3=60
Num: 20/3=0 余数 2
```

* 在定义` BDemo` 的时候，  如果定义的` DD, CC int `和服务端不一样，就会报错  ` reading body gob: type mismatch: no fields matched compiling decoder for` ,其实发现好多种情况，也会出现这种错误，但是目前不知道为啥会这样，后续，等源码深入一点，回来看这个问题    todo2018/07/19 
* 这种`TCP`方式和上面的`HTTP`不同的是
  * HTTP:指定的网络地址连接到HTTP RPC服务器
```
         //DelayHTTP在指定的网络地址连接到HTTP RPC服务器
	///在默认HTTP RPC路径上监听。
	client, err := rpc.DialHTTP("tcp", serverAddress)
```
  * TCP:指定的网络地址连接到HTTP RPC服务器
```
    client,err:=rpc.Dial("tcp",service)
```

#### JSON RPC
* `JSON RPC`是数据编码采用了`JSON`，而不是`gob`编码，其他和上面介绍的`RPC`概念一模一样的。

* 服务端的代码如下

```
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

```
* 客户端的代码 
```
package main

import (
	"fmt"
	"os"
	"net/rpc/jsonrpc"
	"log"
)

func main() {
	fmt.Println("这是客户端，用来启动，通过命令行来启动")

	fmt.Println("客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go")

	if len(os.Args)==4{
		fmt.Println("长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip=",os.Args[1],"嘿嘿,加上后面的被除数os.Args[2]=",os.Args[2],"和除数os.Args[3]=",os.Args[3])
		//os.Exit(1)
	}

	 service:=os.Args[1]
	 client,err:=jsonrpc.Dial("tcp",service)
	if err != nil {
		log.Fatal("Dial 发生了错误了哦 错误的信息为   err=",err)
	}
	send:=Send{os.Args[2],os.Args[3]}
	var  resive  string
	err1:=client.Call("DemoM.DoWork",send,&resive)
	if err1!=nil {
		fmt.Println("shiming call error    ")
		fmt.Println("Call 的时候发生了错误了哦  err=",err1)
	}
	fmt.Println("收到信息了",resive)




}
// 类可以不一样 但是 Who 和DoWhat 要必须一样  要不然接收到不到值，等我在详细的了解了 才去分析下原因  感觉有点蒙蔽啊
type Send struct {
	Who, DoWhat string
}





```
* 运行的结果如下 
![运行结果](https://upload-images.jianshu.io/upload_images/5363507-48090786abf0e69c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```
E:\new_code\GoDemo\web_server>go run GoJSONRPCWeb11.go 127.0.0.1:8080  shiming g
ongzuo
这是客户端，用来启动，通过命令行来启动
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:8080 嘿嘿,加上后面
的被除数os.Args[2]= shiming 和除数os.Args[3]= gongzuo
收到信息了 是谁：shiming，在做什么---gongzuo

E:\new_code\GoDemo\web_server>go run GoJSONRPCWeb11.go 127.0.0.1:8080  shiming q
iaodaima
这是客户端，用来启动，通过命令行来启动
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:8080 嘿嘿,加上后面
的被除数os.Args[2]= shiming 和除数os.Args[3]= qiaodaima
收到信息了 是谁：shiming，在做什么---qiaodaima
```
* `os.Args`是一个数组  `var Args []string`,通过输入获取到，然后把这个客户端输入的内容传送到服务端，服务端做些操作，然后在返回给客户端 
*  `Go`已经提供了`RPC`良好的支持，通过`HTTP` `TCP` `JSONRPC`的实现，可以很方便开发分布式的`Web`应用，但是我目前还不会，在学习中。遗憾的是`Go`没有提供`SOAP RPC`的支持~~~


