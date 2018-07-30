package main

import (
	"fmt"
	"os"
	"net"
	"time"
)

func init() {
	fmt.Println("Socket 起源于Unix Unix基本哲学就是一切都是文件  ，都可以用“ 打开open-读写 write/read  关闭close ”   Socket就是该模式的一个实现，网络的Scoket 是一个特殊的io，socket也是一种文件的描述符。流式Socket和数据报式Socket ，流式的一种面向连接的Socket，针对于面向连接TCP服务的应用，数据报式是一种无连接的Socket，对应UDP服务")
}
func main() {
   //Socket起源于Unix，而Unix基本哲学之一就是“一切皆文件”，都可以用“打开open –> 读写write/read –> 关闭close”模式来操作。Socket就是该模式的一个实现，网络的Socket数据传输是一种特殊的I/O，Socket也是一种文件描述符。Socket也具有一个类似于打开文件的函数调用：Socket()，该函数返回一个整型的Socket描述符，随后的连接建立、数据传输等操作都是通过该Socket实现的。
	//常用的Socket类型有两种：流式Socket（SOCK_STREAM）和数据报式Socket（SOCK_DGRAM）。流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用。

    // todo  Socket如何通信     注意图片8.1.socket
	//网络中的进程之间如何通过Socket通信呢？首要解决的问题是如何唯一标识一个进程，否则通信无从谈起！在本地可以通过进程PID来唯一标识一个进程，但是在网络中这是行不通的。其实TCP/IP协议族已经帮我们解决了这个问题，网络层的“ip地址”可以唯一标识网络中的主机，而传输层的“协议+端口”可以唯一标识主机中的应用程序（进程）。这样利用三元组（ip地址，协议，端口）就可以标识网络的进程了，网络中需要互相通信的进程，就可以利用这个标志在他们之间进行交互。请看下面这个TCP/IP协议结构图

	//使用TCP/IP协议的应用程序通常采用应用编程接口：UNIX BSD的套接字（socket）和UNIX System V的TLI（已经被淘汰），来实现网络进程之间的通信。就目前而言，几乎所有的应用程序都是采用socket，而现在又是网络时代，网络中进程通信是无处不在，这就是为什么说“一切皆Socket”。

	// Socket 基础
	//通过上面的介绍我们知道Socket有两种：TCP Socket和UDP Socket，TCP和UDP是协议，而要确定一个进程的需要三元组，需要IP地址和端口。

   //IPv4地址
	//目前的全球因特网所采用的协议族是TCP/IP协议。IP是TCP/IP协议中网络层的协议，是TCP/IP协议族的核心协议。目前主要采用的IP协议的版本号是4(简称为IPv4)，发展至今已经使用了30多年。
	//IPv4的地址位数为32位，也就是最多有2的32
	//地址格式类似这样：127.0.0.1 172.122.121.111

	//IPv6地址
	//IPv6是下一版本的互联网协议，也可以说是下一代互联网的协议，它是为了解决IPv4在实施过程中遇到的各种问题而被提出的，IPv6采用128位地址长度，几乎可以不受限制地提供地址。按保守方法估算IPv6实际可分配的地址，整个地球的每平方米面积上仍可分配1000多个地址。在IPv6的设计过程中除了一劳永逸地解决了地址短缺问题以外，还考虑了在IPv4中解决不好的其它问题，主要有端到端IP连接、服务质量（QoS）、安全性、多播、移动性、即插即用等。
	//地址格式类似这样：2002:c0e8:82e7:0:0:0:c0e8:82e7


	//Go支持的IP类型
	goIpDemo()

	//TCP Socket 网络端口访问一个服务时，那么我们能够做什么呢？作为客户端来说，我们可以通过向远端某台机器的的某个网络端口发送一个请求，然后得到在机器的此端口上监听的服务反馈的信息。作为服务端，我们需要把服务绑定到某个指定端口，并且在此端口上监听，当有客户端来访问时能够读取信息并且写入反馈信息。
	TCPDemo()


   //	TCP client Go语言中通过net包中的DialTCP函数来建立一个TCP连接，并返回一个TCPConn类型的对象，当连接建立时服务器端也创建一个同类型的对象，此时客户端和服务器段通过各自拥有的TCPConn对象来进行数据交换。一般而言，客户端通过TCPConn对象将请求信息发送到服务器端，读取服务器端响应的信息。服务器端读取并解析来自客户端的请求，并返回应答信息，这个连接只有当任一端关闭了连接之后才失效，不然这连接可以一直在使用。
    TCPClient()


	//TCP server

	TCPServer()

}
/*
上面的服务跑起来之后，它将会一直在那里等待，直到有新的客户端请求到达。当有新的客户端请求到达并同意接受Accept该请求的时候他会反馈当前的时间信息。值得注意的是，在代码中for循环里，当有错误发生时，直接continue而不是退出，是因为在服务器端跑代码的时候，当有错误发生的情况下最好是由服务端记录错误，然后当前连接的客户端直接报错而退出，从而不会影响到当前服务端运行的整个服务。
 */
func TCPServer() {
	//  todo  我们实现一个简单的时间同步服务，监听7777端口   手动的刷新一个地址   http://localhost:7777/
	service := ":7777"
	/*
	net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only), TCP(IPv6-only)或者TCP(IPv4, IPv6的任意一个)。
addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22"。
	 */
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println("tcpAddr",tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		//进入for循环了
		//conn== &{{0xc04207d8c0}}

		/*
		之后，它将会一直在那里等待，直到有新的客户端请求到达。当有新的客户端请求到达并同意接受Accept该请求的时候他会反馈当前的时间信息。值得注意的是，在代码中for循环里，当有错误发生时，直接continue而不是退出，是因为在服务器端跑代码的时候，当有错误发生的情况下最好是由服务端记录错误，然后当前连接的客户端直接报错而退出，从而不会影响到当前服务端运行的整个服务。
上面的代码有个缺点，执行的时候是单任务的，不能同时接收多个请求，那么该如何改造以使它支持多并发呢？Go里面有一个goroutine机制

		 */
		fmt.Println("进入for循环了 ")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=====",err)
			continue
		}
		fmt.Println("conn==",conn)
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func TCPClient() {
	//net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
	//laddr表示本机地址，一般设置为nil
	//raddr表示远程的服务地址
	//net.DialTCP(network string, laddr, raddr *TCPAddr)
//	接下来我们写一个简单的例子，模拟一个基于HTTP协议的客户端请求去连接一个Web服务端。我们要写一个简单的http请求头，格式类似如下：
//
//	"HEAD / HTTP/1.0\r\n\r\n"
//	从服务端接收到的响应信息格式可能如下：
//
//	HTTP/1.0 200 OK
//ETag: "-9985996"
//	Last-Modified: Thu, 25 Mar 2010 17:51:10 GMT
//	Content-Length: 18074
//Connection: close
//Date: Sat, 28 Aug 2010 00:43:48 GMT
//Server: lighttpd/1.4.23
//	我们的客户端代码如下所示：
//
//	package main
//
//	import (
//		"fmt"
//	"io/ioutil"
//	"net"
//	"os"
//	)
//
//	func main() {
//		if len(os.Args) != 2 {
//			fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
//			os.Exit(1)
//		}
//		service := os.Args[1]
//		tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
//		checkError(err)
//		conn, err := net.DialTCP("tcp", nil, tcpAddr)
//		checkError(err)
//		_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
//		checkError(err)
//		result, err := ioutil.ReadAll(conn)
//		checkError(err)
//		fmt.Println(string(result))
//		os.Exit(0)
//	}
//	func checkError(err error) {
//	if err != nil {
//	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
//	os.Exit(1)
//	}
//	}
//	通过上面的代码我们可以看出：首先程序将用户的输入作为参数service传入net.ResolveTCPAddr获取一个tcpAddr,然后把tcpAddr传入DialTCP后创建了一个TCP连接conn，通过conn来发送请求信息，最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息。


}

func TCPDemo() {
	//在Go语言的net包中有一个类型TCPConn，这个类型可以用来作为客户端和服务器端交互的通道，他有两个主要的函数：
	//func (c *TCPConn) Write(b []byte) (int, error)
	//func (c *TCPConn) Read(b []byte) (int, error)
	//TCPConn可以用在客户端和服务器端来读写数据。
	//
	//还有我们需要知道一个TCPAddr类型，他表示一个TCP的地址信息，他的定义如下：
	//
	//type TCPAddr struct {
	//	IP IP
	//	Port int
	//	Zone string // IPv6 scoped addressing zone
	//}
	//在Go语言中通过ResolveTCPAddr获取一个TCPAddr
	//
	//func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
	//net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only), TCP(IPv6-only)或者TCP(IPv4, IPv6的任意一个)。
	//addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22"。




}
func goIpDemo() {

    //在Go的net包中定义了很多类型、函数和方法用来网络编程，其中IP的定义
	type IP []byte

	str:="ddd"
	if len(str) != 2 {
		fmt.Println("shi ming  go start  ")
	}

	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	//	os.Exit(1)
	//}
	//name := os.Args[1]
	//addr := net.ParseIP(name)
	//if addr == nil {
	//	fmt.Println("Invalid address")
	//} else {
	//	fmt.Println("The address is ", addr.String())
	//}
	//os.Exit(0)


}
