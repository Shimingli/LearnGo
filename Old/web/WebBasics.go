package main

import (
	"net/http"
	"log"
)

func init() {
	// 如何使用 go来编写 web应用，go目前已经有了成熟的http处理包，这使得编写能够做任何事情的动态的Web程序易于反掌
	// 1、web的工作的方式
	//2、Go搭建一个简单的web服务
	//3、go如何使得Web工作
	//4、go的http包详解
	//5、小结

//     todo    https://www.jianshu.com/p/84aa55a8a7eb

   //Go代码执行的流程
	http.HandleFunc("/",sayHelloName)
	//设置监听的端口
	err:= http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("ListenAndServe",err)
	}
	//1、首先调用的是Http.HandleFunc
	//调用了DefaultServeMux.HandleFunc(pattern, handler)
	//滴啊用了mux.Handle(pattern, HandlerFunc(handler)) 也就是调用了DefaultServeMux的Handle
	//往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则
	//2、其次调用http.ListenAndServe(":9090", nil)
	//1 实例化Server
	//2 调用Server的ListenAndServe()
	//3 调用net.Listen("tcp", addr)监听端口
	//4 启动一个for循环，在循环体中Accept请求
	//5 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
	//6 读取每个请求的内容w, err := c.readRequest()
	//7 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux
	//8 调用handler的ServeHttp
	//9 在这个例子中，下面就进入到DefaultServeMux.ServeHttp
	//10 根据request选择handler，并且进入到这个handler的ServeHTTP   mux.handler(r).ServeHTTP(w, r)
	//11 选择handler：
	//A 判断是否有路由能满足这个request（循环遍历ServeMux的muxEntry）
	//B 如果有路由满足，调用这个路由handler的ServeHTTP
	//C 如果没有路由满足，调用NotFoundHandler的ServeHTTP


}

