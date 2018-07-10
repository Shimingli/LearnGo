package main

import (
	"fmt"
	"net/http"
	"net/url"
	"html/template"
	//导入session包
	"github.com/astaxie/beego/session"
)
var provides =make(map[string]Provider)

func init() {
    // fmt.Println("provides=",provides) //provides= map[]
    // globalSessions,error:=NewManager("memory","goSessionID",3600)

     //fmt.Println("globalSessions===",globalSessions)
     //fmt.Println("error===",error)
     //fmt.Println("globalSessions.sessionId()==",globalSessions.sessionId())
	//
	//globalSessions, _ = session.NewManager("memory", {cookieName:"gosessionid",gclifetime:3600})
	//go globalSessions.GC()


}
func main() {
	http.HandleFunc("/shiming",login)
	http.ListenAndServe(":9092",nil)

	//session是在服务器实现的一种用户和服务器之间的认证的解决方案，但是目前Go标准包没有为session提供任何的支持。需要自己的去实现
	fmt.Println("Go如何使用 session")
	// session的创建过程
	//   session的基本原理由服务器为每一个会话维护一份信息数据，客服端和服务端靠一个全局唯一的标识来访问这份数据，以达到交互的目的
	// 服务端程序创建需要的session的过程
	//  1 生成唯一的标识符（session id）
	//   2 开辟数据存储的空间。一般会在内存中创建相应的数据结构，但是这种情况下，系统一旦断电的话，所有的会话的数据就会丢失，对电子商务类的网站，这将会赵成严重的后果。通常的做法是将会话的数据写入到文件里或者是数据库中，这样会增加I/O开销，但是它可以实现某种程度的session的持久化，也利于session的共享、//
	//  3  将session的全局唯一标识发送给客户端
	fmt.Println("最关键的session步骤就是如何发送这个session的标识符传送到客户端，考虑到HTTP协议的定义，数据无非可以放到请求行、头域或者是body里，所以一般来说会有两种的常用的方式：cookie和URL重写")
	//1 Cookie服务端通过设置Set-Cookie头就可以将session的标识符传送到客户端，而客户端此后的每一次请求都会带上这个标识符，另外一般包含session信息的cookie会将失效的时间设置为0（会话的cookie），就是浏览器进程的有效的时间。每个浏览器器都有自己的处理这个0的方式
	//2 URL重写 所谓URL重写，就是在返回给用户的页面里的所有的URL后面追加session标识符，这样用户在收到响应之后，无论点击响应页面里的那个链接或者是提交表单，都会自动带上session标识符，从而就实现了会话的保持，虽然这种做法比较的麻烦，但是客户端禁用cookie的话，这种方法就是首选

	fmt.Println("Go 实现session的管理器")
	// 结合到session的生命周期 （lifecycle）
	// 1 全局的session管理器， 保证sessionid的全局的唯一性 为每个用户管理一个session session的存储（内存，文件，数据库）session的过期的处理
}

var globalSessions *session.Manager

func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("sessionAndDataStorage/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
func (m *Manager) SessionStart(w http.ResponseWriter,r *http.Request) (session Session){
	m.lock.Lock()
	defer  m.lock.Unlock()
	cookie,err:= r.Cookie(m.cookieName)
	if err!=nil||cookie.Value=="" {
		sid:=m.sessionId()
		session,_:=m.provider.SessionInit(sid)
		cookie:=http.Cookie{Name: m.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(m.maxLifeTime)}
       http.SetCookie(w,&cookie)
		return session
	}else {
		sid,_:=url.QueryUnescape(cookie.Value)
		session,_:=m.provider.SessionRead(sid)
		return session

	}
}



