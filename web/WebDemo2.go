package main

import (
	"fmt"
	"net/http"
	"net"
	"time"
	"context"
	"runtime"
	"crypto/tls"
)

//2、Go如何使Web工作，Go的Web服务工作离不开我们Web工作的方式  todo    https://www.jianshu.com/p/84aa55a8a7eb
func init() {
	//Web 工作方式的几个概念
	//Request : 用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息
	//Response:服务器需要反馈给客户端的信息
	// Conn:用户的每次请求链接
	// Handler :处理请求和生成返回信息的处理逻辑


}
func main() {
      fmt.Println(" 注意查看图3.3http.png 的图")


      /*
      Go实现Web服务的工作模式的流程图
      http包执行的流程
      1、创建Listen Socket ，监听指定的端口，等待客户端请求的到来
      2、Listen Socket接收到客户端的请求，得到Client Socket ，接下来通过Client Socket 与客户端通信
      3、处理客户端的请求，首先从Client Socket读取HTTP 请求的协议头，如果是POST方法，还可能读取客户端提交的数据，然后交给响应的handler处理请求，handler处理完毕准备好客户端需要的数据，通过Client Socket写给客户端

       */


       //1、  如何监听端口
	//设置监听的端口
	err:= http.ListenAndServe(":9090",nil)
	if err!=nil{
		println(err)
	}

}
// a、位于Server.go中的  package http中 ，
func ListenAndServe(addr string, handler Handler) error {
	//返回地址值,初始化一个server  handler其实就是方法
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
// b 位于Server.go中的  package http中 ，
//监听TCP netWork的网络地址，然后调用serve来处理将要进来的连接。
//接受连接的配置能够使 TCP  keep-alives ，如果说地址是个"" 那么会给它默认上一个":http"
//这个方法始终返回的是一个非空的错误----》
func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	//然后调用 net.Listen(),也就是底层用TCP协议搭建的一个服务，然后监控我们设置的端口
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	//最终正常的情况下会走到这里来
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
//  c、 位于Server.go中的  package http中 ，最终调用了 Serve和这个函数
//Server接受侦听器L上的传入连接，创建一个
//每个新的服务GOODUTIN。服务GORATONE读取请求和
//然后调用SRV处理程序来回复它们。
//对于HTTP/2支持，SRV.TLSCONFIG应该被初始化为
//在调用服务之前提供侦听器的TLS配置。如果
//Srv.TLSCONFIG是非零，并且不包含字符串“H2”。
//CONT.NEXPROSTS，HTTP／2支持未启用。
//Service总是返回非零错误。关闭或关闭后，
//RealError错误被关闭。
//   todo  这个函数处理 接受客户端的请求的信息
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()  //Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题
	if fn := testHookServerServe; fn != nil {
		fn(srv, l)
	}
	// 多长的时间接受失败的去sleep
	var tempDelay time.Duration // how long to sleep on accept failure

	if err := srv.setupHTTP2_Serve(); err != nil {
		return err
	}

	srv.trackListener(l, true)
	defer srv.trackListener(l, false)

	baseCtx := context.Background() // base is always background, per Issue 16220
	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	for {// 无限的循环 ---》
		rw, e := l.Accept()  // 首先通过Listener 接受请求
		if e != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		//其次就是创建一个Conn了
		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew) // before Serve can return
		//最后单独的开了一个  goroutine 把这个请求的数据当做参数扔给这个conn服务:
		go c.serve(ctx) //  这就是高并发的体现，用户每次的请求都是一个新的  goroutine  ，并且相互不影响
	}
}


//  d、 Serve a new connection. 提供新的连接   todo  关键的方法
func (c *conn) serve(ctx context.Context) {
	c.remoteAddr = c.rwc.RemoteAddr().String()
	ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
	defer func() {
		if err := recover(); err != nil && err != ErrAbortHandler {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		if !c.hijacked() {
			c.close()
			c.setState(c.rwc, StateClosed)
		}
	}()

	if tlsConn, ok := c.rwc.(*tls.Conn); ok {
		if d := c.server.ReadTimeout; d != 0 {
			c.rwc.SetReadDeadline(time.Now().Add(d))
		}
		if d := c.server.WriteTimeout; d != 0 {
			c.rwc.SetWriteDeadline(time.Now().Add(d))
		}
		if err := tlsConn.Handshake(); err != nil {
			c.server.logf("http: TLS handshake error from %s: %v", c.rwc.RemoteAddr(), err)
			return
		}
		c.tlsState = new(tls.ConnectionState)
		*c.tlsState = tlsConn.ConnectionState()
		if proto := c.tlsState.NegotiatedProtocol; validNPN(proto) {
			if fn := c.server.TLSNextProto[proto]; fn != nil {
				h := initNPNRequest{tlsConn, serverHandler{c.server}}
				fn(c.server, tlsConn, h)
			}
			return
		}
	}

	// HTTP/1.x from here on.
	// http/1.x从这里开始。
	ctx, cancelCtx := context.WithCancel(ctx)
	c.cancelCtx = cancelCtx
	defer cancelCtx()

	c.r = &connReader{conn: c}
	c.bufr = newBufioReader(c.r)
	c.bufw = newBufioWriterSize(checkConnErrorWriter{c}, 4<<10)

	for {
		//todo     首先解析 c.readRequest()
		w, err := c.readRequest(ctx)
		if c.r.remain != c.server.initialReadLimitSize() {
			// If we read any bytes off the wire, we're active.
			//如果我们读取了什么的字节，我们是活跃的
			c.setState(c.rwc, StateActive)
		}
		if err != nil {
			const errorHeaders = "\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n"

			if err == errTooLarge {
				// Their HTTP client may or may not be
				// able to read this if we're
				// responding to them and hanging up
				// while they're still writing their
				// request. Undefined behavior.
				//他们的HTTP客户端可能会或可能无法读取这一点，如果我们响应他们和挂起，而他们仍然在写他们的请求。未定义的行为。
				const publicErr = "431 Request Header Fields Too Large"
				// 写个请求者的数据
				fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
				c.closeWriteAndWait()
				return
			}
			if isCommonNetReadError(err) {
				return // don't reply
			}

			publicErr := "400 Bad Request"
			if v, ok := err.(badRequestError); ok {
				publicErr = publicErr + ": " + string(v)
			}

			fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
			return
		}

		// Expect 100 Continue support
		req := w.req
		if req.expectsContinue() {
			if req.ProtoAtLeast(1, 1) && req.ContentLength != 0 {
				// Wrap the Body reader with one that replies on the connection
				req.Body = &expectContinueReader{readCloser: req.Body, resp: w}
			}
		} else if req.Header.get("Expect") != "" {
			w.sendExpectationFailed()
			return
		}

		c.curReq.Store(w)

		if requestBodyRemains(req.Body) {
			registerOnHitEOF(req.Body, w.conn.r.startBackgroundRead)
		} else {
			if w.conn.bufr.Buffered() > 0 {
				w.conn.r.closeNotifyFromPipelinedRequest()
			}
			w.conn.r.startBackgroundRead()
		}

		// HTTP cannot have multiple simultaneous active requests.[*]
		// Until the server replies to this request, it can't read another,
		// so we might as well run the handler in this goroutine.
		// [*] Not strictly true: HTTP pipelining. We could let them all process
		// in parallel even if their responses need to be serialized.
		// But we're not going to implement HTTP pipelining because it
		// was never deployed in the wild and the answer is HTTP/2.

		//HTTP不能有多个同时激活的请求。[*]直到服务器回复这个请求，它不能读取另一个请求，所以我们也可以在这个GOODUTE中运行处理程序。[*]并不是严格的：HTTP流水线。我们可以让它们都并行处理，即使它们的响应需要被序列化。但是我们不打算实现HTTP流水线，因为它从来没有部署在野外，答案是HTTP／2。
		//  todo  然后调用的是这里
		serverHandler{c.server}.ServeHTTP(w, w.req)
		w.cancelCtx()
		if c.hijacked() {
			return
		}
		w.finishRequest()
		if !w.shouldReuseConnection() {
			if w.requestBodyLimitHit || w.closedRequestBodyEarly() {
				c.closeWriteAndWait()
			}
			return
		}
		c.setState(c.rwc, StateIdle)
		c.curReq.Store((*response)(nil))

		if !w.conn.server.doKeepAlives() {
			// We're in shutdown mode. We might've replied
			// to the user without "Connection: close" and
			// they might think they can send another
			// request, but such is life with HTTP/1.1.
			return
		}

		if d := c.server.idleTimeout(); d != 0 {
			c.rwc.SetReadDeadline(time.Now().Add(d))
			if _, err := c.bufr.Peek(4); err != nil {
				return
			}
		}
		c.rwc.SetReadDeadline(time.Time{})
	}
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	//然后获取相应的handler，也就是我们在  ListenAndServe时候传入的第二个参数，我们前面传递是一个 nil ，为空 ，所以handler =DEfaultServeMux
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
		/*
		这个 handler变量 来做什么？？？ 其实这个handler 变量就是一个路由器，它用来匹配url跳转到相对于的handle函数，在代码的第一行传入	http.HandleFunc("/",sayHelloName)  这个作用就是注册了请求  / 的路由的规则，当请求的 uri 为"/" ,路由就会转到函数 sayhelloName，DefaultServeMux 会调用 ServeHttp方法，这个方法内部其实就是 调用sayhelloName本身，最后通过 写入 response的信息反馈到客户端

		 */
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)


	//   todo   注意看图片  3.3.illustrator.png
}