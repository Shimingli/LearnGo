package main

import (
	"net"
	"time"
	"fmt"
	"os"
	"strings"
	"strconv"
)
/*
通过对TCP和UDP Socket编程的描述和实现，可见Go已经完备地支持了Socket编程，而且使用起来相当的方便，Go提供了很多函数，通过这些函数可以很容易就编写出高性能的Socket应用。
 */
func main() {
     // TCP server
     //，也可以通过net包来创建一个服务器端程序，在服务器端我们需要绑定服务到指定的非激活端口，并监听此端口，当有客户端请求到达的时候可以接收到来自客户端连接的请求
	 // demo1() [Fiddler] ReadResponse() failed: The server did not return a complete response for this request. Server returned 50 bytes.
     //改造以使它支持多并发呢？Go里面有一个goroutine机制
	// demo2()

	 // 以上的两个Demo都没有处理客服单实际的请求的内容，如果我们需要通过从客户端发送不同的请求来获取不同的时间的格式，而且需要以一个长连接
	// demo3()


	 // 控制Tcp的连接
	 tcpDemo111()

    //UDP Socket  go run WebSocketDemo3.go 5545   找到本地文件去运行才可以  而不是Java那样子的情况
    //   todo 一个UDP的客户端
     //UDPSocketDemo()

     //  todo  UDP服务器端
     UDPSocketDemoServer()

}

func UDPSocketDemoServer() {
	service := ":1200"
	udpAddr,err:=net.ResolveUDPAddr("udp4",service)
	checkErrorDe(err)
	conn,err:=net.ListenUDP("udp",udpAddr)
	checkErrorDe(err)
	for  {
		handleUDPClient(conn)
	}
}
func handleUDPClient(conn *net.UDPConn) {
	var  buf [512]byte
	fmt.Println(buf)
	ddd:=&buf
	fmt.Println("ddd=",ddd)
	//	aSlice = array[0:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素
	// 等于说全部的copy下来
	_,addr,err:=conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Println("err====",err)
		return
	}
	dayTime :=time.Now().String()
	conn.WriteToUDP([]byte(dayTime),addr)


}
/*
Go语言包中处理UDP Socket和TCP Socket不同的地方就是在服务器端处理多个客户端请求数据包的方式不同,UDP缺少了对客户端连接请求的Accept函数。其他基本几乎一模一样，只有TCP换成了UDP而已
 */
func UDPSocketDemo() {
     fmt.Println("输入的是什么  用户的输入的是什么 ",os.Args)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErrorDe(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErrorDe(err)
	_, err = conn.Write([]byte("anything"))
	checkErrorDe(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkErrorDe(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)

}
func tcpDemo111() {
	fmt.Println("多看文字")
	//TCP有很多连接控制函数，我们平常用到比较多的有如下几个函数：
	//
	//func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)
	//设置建立连接的超时时间，客户端和服务器端都适用，当超过设置时间时，连接自动关闭。
	//
	//func (c *TCPConn) SetReadDeadline(t time.Time) error
	//func (c *TCPConn) SetWriteDeadline(t time.Time) error
	//用来设置写入/读取一个连接的超时时间。当超过设置时间时，连接自动关闭。
	//
	//func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
	//设置keepAlive属性，是操作系统层在tcp上没有数据和ACK的时候，会间隔性的发送keepalive包，操作系统可以通过该包来判断一个tcp连接是否已经断开，在windows上默认2个小时没有收到数据和keepalive包的时候人为tcp连接已经断开，这个功能和我们通常在应用层加的心跳包的功能类似。
	//更多的内容请查看net包的文档。
}

func demo3() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErrorDe(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErrorDe(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClientChangLianjie(conn)
	}
}
// todo request 里面的东西
//request== GET / HTTP/1.1
//Host: localhost:1200
//Connection: keep-alive
//Cache-Control: max-age=0
//Upgrade-Insecure-Requests: 1
//User-Agent
//read_le=== 128
//request== : Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 SE 2.X MetaSr 1.
//read_le=== 128
//request== 0
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
//Accept-Encoding: gzip, deflate, sdch, br
//read_le=== 128
//request==
//Accept-Language: zh-CN,zh;q=0.8
//Cookie: _ga=GA1.1.27955907.1529919744
//
//,*/*;q=0.8
//Accept-Encoding: gzip, deflate, sdch, br
//read_le=== 75
//request==
//Accept-Language: zh-CN,zh;q=0.8
//Cookie: _ga=GA1.1.27955907.1529919744
//
//,*/*;q=0.8
//Accept-Encoding: gzip, deflate, sdch, br


//在这个例子中，我们使用conn.Read()不断读取客户端发来的请求。由于我们需要保持与客户端的长连接，所以不能在读取完一次请求后就关闭连接。由于conn.SetReadDeadline()设置了超时，当一定时间内客户端无请求发送，conn便会自动关闭，下面的for循环即会因为连接已关闭而跳出。需要注意的是，request在创建时需要指定一个最大长度以防止flood attack；每次读取到请求处理完毕后，需要清理request，因为conn.Read()会将新读取到的内容append到原内容之后。
func handleClientChangLianjie(conn net.Conn) {
	//设置2分钟时间超时
	conn.SetReadDeadline(time.Now().Add(2*time.Minute))
	//设置最大 设置最大请求长度到128B以防止洪水攻击
	request:=make([]byte,128)
	//退出之前关闭连接
	defer conn.Close()
	for  {
		//EOF是当没有更多的输入可用时由Read返回的错误
		read_len,err:=conn.Read(request)
		fmt.Println("request==",string(request))
		if err != nil {
			fmt.Println("err！！！===",err)
			break
		}
		fmt.Println("read_le===",read_len )
		if read_len==0 {
			//说明已经连接了，关闭client
			break
		}else if strings.TrimSpace(string(request[:read_len]))=="timestamp" {
            dattime:=strconv.FormatInt(time.Now().Unix(),10)
			fmt.Println("daytime==",dattime)
			conn.Write([]byte(dattime))
		}else {
			daytime:=time.Now().String()
			conn.Write([]byte(daytime))
		}
	}

}
//就是一个监听的作用
func demo2() {
   server:= ":1200"
   tcpAddr,err:=net.ResolveTCPAddr("tcp4",server)
   checkErrorDe(err)
   listener,err:= net.ListenTCP("tcp",tcpAddr)
   checkErrorDe(err)

	for  {
		fmt.Println("新的for 循环的开始   ")
		 conn,err:=listener.Accept()
		if err != nil {
			continue
		}
		/*
		过把业务处理分离到函数handleClient，我们就可以进一步地实现多并发执行了。看上去是不是很帅，增加go关键词就实现了服务端的多并发，从这个小例子也可以看出goroutine的强大之处
		 */
		go handlerClient(conn)
	}
}
func handlerClient(conn net.Conn) {
	defer  conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client

}
/*
上面的服务跑起来之后，它将会一直在那里等待，直到有新的客户端请求到达。当有新的客户端请求到达并同意接受Accept该请求的时候他会反馈当前的时间信息。值得注意的是，在代码中for循环里，当有错误发生时，直接continue而不是退出，是因为在服务器端跑代码的时候，当有错误发生的情况下最好是由服务端记录错误，然后当前连接的客户端直接报错而退出，从而不会影响到当前服务端运行的整个服务。
 */
func demo1() {
	service := ":7777"
	tcpAddr,err :=net.ResolveTCPAddr("tcp4",service)
	fmt.Println("开始接收到消息了")
	checkErrorDe(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	checkErrorDe(err)
	for   {
		fmt.Println("for  循环开始了")
		conn,err:=listener.Accept()
		if err!=nil {
			continue
		}
		//2018-07-16 17:01:28.5196 +0800 CST m=+101.170000001
		dayTime:=time.Now().String()
		fmt.Println("day Time",dayTime)
		conn.Write([]byte(dayTime))
		conn.Close()
	}
}

func checkErrorDe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
