package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)
// 1、Go搭建的一个Web的服务器 ： Go语言里面提供了一个完善的net/http包，通过http包可以很方便的就搭建起来一个可以运行的Web服务。同时使用这个包可以很简单地对Web的路由，静态文件，模板，cookie等数据进行设置和操作
func init() {


}
func main() {
	fmt.Println("start")
	/*
	注册回调函数从给定的模式在默认的ServeMux文档中，在ServeMux文档解释了模式是如何的匹配的  DefaultServeMux.HandleFunc(pattern, handler)
	 */
	 //具体就是 调用的 是    server.go  中的Handle方法
	http.HandleFunc("/",sayHelloName)
	//设置监听的端口
	 err:= http.ListenAndServe(":7777",nil)
	if err!=nil {
		log.Fatal("ListenAndServe",err)
	}


	/*
	需要编写一个web服务器很简单，只要调用http包的两个函数就可以了
	Go不需要nginx 和apache 服务器，它直接监听了tcp端口，做了nginx做的事情，然后sayhelloname就是我们写的逻辑的函数，跟php里面的控制层函数类似
	Go就是拥有类似python这样动态语言的特性，写Web应用很方便
	Go服务内部支持高并发的特性
	 */

}
/*

每次访问都会，请求两次 感觉是这样子的
这里的关键的地方是 var err=  r.ParseForm()   这个方法不打开的话  ，后面的输出的内容为 nil
 */
func sayHelloName(w http.ResponseWriter,r *http.Request)  {
	   // 访问的地址
	  //http://localhost:9090/?url_long=111&url_long11=222
     var err=  r.ParseForm()//解析参数，默认是不会解析的 返回一个err
     fmt.Println("发生了错误了么",err)//发生了错误了么 <nil>
	//这个信息是输出到服务器端的打印的信息
     fmt.Println("r.Form====",r.Form)//r.Form==== map[url_long:[111] url_long11:[222]]
     fmt.Println("path===",r.URL.Path)//path=== /
     fmt.Println("scheme====",r.URL.Scheme)//scheme====
     fmt.Println(r.Form["url_long"])//[111] 这个的意思，就是把key=url_long的值取出来
    //遍历一下 这个 map里面的 东西,里面没有这个东西的话，就会不会解析
	for k,v:= range  r.Form  {
		fmt.Println("key=====",k)
		fmt.Println("val:==========",v)
		//连接连接A的元素以创建单个字符串。分隔字符串SEP放置在结果字符串中的元素之间。 说白了  就是把 [111]变成 111 具体看方法的实现
		fmt.Println("strings.join========",strings.Join(v,""))
	}
	fmt.Fprintf(w,"ni hao  shiming")//输入到客户端的内容
}


