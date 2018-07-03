package main

import (
	"fmt"
	"net/http"
	"log"
	"net"
	"sync"
)

func init() {
   fmt.Println("Go的http包的详解")
}

func main() {
	http.HandleFunc("/",sayHelloName)
	//设置监听的端口
	err:= http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("ListenAndServe",err)
	}

	//  与一般服务器 编写的 http 不同，Go为了实现高并发和高性能，使用了 goroutines 来处理 conn 的读写事件，保证每个请求都能够保持独立，相互不会阻塞，可以高效的响应网络事件，这就是Go高效的保证

	// 代码如下
	c := srv.newConn(rw)
	c.setState(c.rwc, StateNew) // before Serve can return
	go c.serve(ctx)


}


// Create new connection from rwc. 从RWC创建新连接。
// 每次请求都会创建一个 Conn ，每个Conn里面保存了，然后在传递到相应的handler，改handler 中便可以读取到相应的header信息，这样就保证了请求的独立性
func (srv *Server) newConn(rwc net.Conn) *conn {
	c := &conn{
		server: srv,
		rwc:    rwc,
	}
	if debugServerConnections {
		c.rwc = newLoggingConn("server", c.rwc)
	}
	return c
}

//	http.HandleFunc("/",sayHelloName)  第一次调用
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

type ServeMux struct {
	//锁，由于请求涉及到并发处理，因此这里需要一个锁机制
	mu    sync.RWMutex
	//路由的规则，一个string 对应一个 mux实体，这里的string  就是注册路由表达式
	m     map[string]muxEntry
	//是否在任意的规则中带有host的信息
	hosts bool // whether any patterns contain hostnames
}
type muxEntry struct {
	explicit bool// 是否精确的配置
	h        Handler// 这个路由表达式对应的是那个handler
	pattern  string// 匹配字符串
}
type Handler interface {
	// 路由实现器
	ServeHTTP(ResponseWriter, *Request)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
//Handler是一个接口，但是前一小节中的sayhelloName函数并没有实现ServeHTTP这个接口，为什么能添加呢？原来在http包里面还定义了一个类型HandlerFunc,我们定义的函数sayhelloName就是这个HandlerFunc调用之后的结果，这个类型默认就实现了ServeHTTP这个接口，即我们调用了HandlerFunc(f),强制类型转换f成为HandlerFunc类型，这样f就拥有了ServeHTTP方法
//   todo   这里麻痹就是一个简单的 方法调度啊 麻痹这就看不懂了  ？？
// HandlerFunc类型是允许使用普通函数作为HTTP处理程序的适配器。如果f是具有适当签名的函数，则HandlerFunc（f）是调用f的处理程序。
type HandlerFunc func(ResponseWriter, *Request)
// todo   我终于搞明白了  ，如果不明白的话，注意请看  GoDemoInrerface 这个类

/*
在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
函数作为类型到底有什么好处呢？那就是可以把这个类型的函数当做值来传递
 */

func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}

/**
只有这个类 HandlerFunc 才能调用 这个方法
 */
// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

// 路由器里面存储了相应的路由规则后，那么具体的请求是怎么样的分发的，默认的路由器实现了 ServeHTTP
//记住一点就是  HandlerFunc  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) { 所以这个方法需要看这个
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	//所有的路由器接收到请求之后，如果是 * 那么关闭连接，
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	// 不然的话，调用下面的方法，返回对应设置路由的处理的 Handler，然后执行 h.ServeHttp
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}

// 	h, _ := mux.Handler(r)  调用的就是这个方法
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {

	// CONNECT requests are not canonicalized.
	if r.Method == "CONNECT" {
		return mux.handler(r.Host, r.URL.Path)
	}
	// All other requests have any port stripped and path cleaned
	// before passing to mux.handler.
	host := stripHostPort(r.Host)
	path := cleanPath(r.URL.Path)
	if path != r.URL.Path {
		_, pattern = mux.handler(host, path)
		url := *r.URL
		url.Path = path
		return RedirectHandler(url.String(), StatusMovedPermanently), pattern
	}
    //也有可能调用下面的方法
	return mux.handler(host, r.URL.Path)
}
//原来他是根据用户请求的URL和路由器里面存储的map去匹配的，当匹配到之后返回存储的handler，调用这个handler的ServeHTTP接口就可以执行到相应的函数了。
func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	// Host-specific pattern takes precedence over generic ones
	if mux.hosts { // 里面存储的map去匹配的
		h, pattern = mux.match(host + path)
	}
	if h == nil {
		h, pattern = mux.match(path)
	}
	if h == nil {
		h, pattern = NotFoundHandler(), ""
	}
	return
}
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		return v.h, v.pattern
	}

	// Check for longest valid match.
	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
		}
	}
	return
}