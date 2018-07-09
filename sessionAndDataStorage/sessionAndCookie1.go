package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"crypto/md5"
	"io"
	"text/template"
)

func main() {
    fmt.Println("当用户来到微博登陆页面，输入用户名和密码之后点击“登录”后浏览器将认证信息POST给远端的服务器，服务器执行验证逻辑，如果验证通过，则浏览器会跳转到登录用户的微博首页，在登录成功后，服务器如何验证我们对其他受限制页面的访问呢？因为HTTP协议是无状态的，所以很显然服务器不可能知道我们已经在上一次的HTTP请求中通过了验证。当然，最简单的解决方案就是所有的请求里面都带上用户名和密码，这样虽然可行，但大大加重了服务器的负担（对于每个request都需要到数据库验证），也大大降低了用户体验(每个页面都需要重新输入用户名密码，每个页面都带有登录表单)。既然直接在请求中带上用户名与密码不可行，那么就只有在服务器或客户端保存一些类似的可以代表身份的信息了，所以就有了cookie与session")


   //  cookie ： 简单的说就是本地的计算机保存一些用户操作的历史信息（也包括登录的信息），并通过用户再次的访问的时候，通过HTTP协议将本地的cookie内容发送给服务器，从而完成验证，或者是继续上一步操作

   //  客户端发起请求，服务端赋值cookie，客户端携带着cookie一起发起请求
   // 注意图片6.1.cookie2.png

   fmt.Println("cookie 是有时间的限制的，根据生命周期的不同分为两种： 会话的cookie和持久的cookie")

    // 会话的cookie是有时间的限制的，则表示这个cookie的生命周期为从创建到浏览器关闭位置，只要关闭了浏览器，这些的cookiue就消失了。这种生命周期为浏览会话期的cookie，被称为：会话的cookie，会话的cookie一般不保存在硬盘上而是保存到内存里

    // 如果设置了过期的时间（setMaxAge（606024）），浏览器就会把cookie保存到硬盘上，关闭后再次打开浏览器，这些cookie依然有效直到超过设定的过期时间，存储在硬盘上的cookie可以在不同的浏览器进程间共享，比如两个IE窗口，而对于这个保存在内存中的cookie，不同的浏览器有不同的处理的方式







   // session  简单的说就是，服务器上保存用户操作的历史的信息，服务器使用session id来标识session，session id由服务器负责产生，保证随机性与唯一性，相当于一个秘钥，避免在握手或者是传输过程中暴露用户真实密码，但该方式下，仍然需要将发送请求的客户端与session进行对应，所以可以借助cookie机制来获取客户端的标识（即session id），也可以通过GET方式将id提交给服务器。
   //session机制是一种服务器端的机制，服务器使用一种类似于散列表的结构来保存信息，每一个网站访客都会被分配给一个唯一的标志符,即sessionID,它的存放形式无非两种:要么经过url传递,要么保存在客户端的cookies里.当然,你也可以将Session保存到数据库里,这样会更安全,但效率方面会有所下降。

   // 1、访问并建立session对应的关系  2、访问网站通过session ID获取对应内容， 3、返回内容
   // 注意看图片 6.1.session.png


	http.HandleFunc("/cookieDemo",cookieDemo)
	//设置监听的端口
     http.ListenAndServe(":9090",nil)

}
/*
session
//session，中文经常翻译为会话，其本来的含义是指有始有终的一系列动作/消息，比如打电话是从拿起电话拨号到挂断电话这中间的一系列过程可以称之为一个session。然而当session一词与网络协议相关联时，它又往往隐含了“面向连接”和/或“保持状态”这样两个含义。
//
//session在Web开发环境下的语义又有了新的扩展，它的含义是指一类用来在客户端与服务器端之间保持状态的解决方案。有时候Session也用来指这种解决方案的存储结构。
//
//session机制是一种服务器端的机制，服务器使用一种类似于散列表的结构(也可能就是使用散列表)来保存信息。
//
//但程序需要为某个客户端的请求创建一个session的时候，服务器首先检查这个客户端的请求里是否包含了一个session标识－称为session id，如果已经包含一个session id则说明以前已经为此客户创建过session，服务器就按照session id把这个session检索出来使用(如果检索不到，可能会新建一个，这种情况可能出现在服务端已经删除了该用户对应的session对象，但用户人为地在请求的URL后面附加上一个JSESSION的参数)。如果客户请求不包含session id，则为此客户创建一个session并且同时生成一个与此session相关联的session id，这个session id将在本次响应中返回给客户端保存。
//
//session机制本身并不复杂，然而其实现和配置上的灵活性却使得具体情况复杂多变。这也要求我们不能把仅仅某一次的经验或者某一个浏览器，服务器的经验当作普遍适用的。
 */
func cookieDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("打印的方法：method",r.Method)
	fmt.Println("打印的方法：r.Form",r.Form)
	//session和cookie的目的相同，都是为了克服http协议无状态的缺陷，但完成的方法不同。session通过cookie，在客户端保存session id，而将用户的其他会话消息保存在服务端的session对象中，与此相对的，cookie需要将所有信息都保存在客户端。因此cookie存在着一定的安全隐患，例如本地cookie中保存的用户名密码被破译，或cookie被其他网站收集（例如：1. appA主动设置域B cookie，让域B cookie获取；2. XSS，在appA上通过javascript获取document.cookie，并传递给自己的appB）。
	if r.Method=="GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("设置的token是==",token)
		t, _ := template.ParseFiles("sessionAndDataStorage/upload.gtpl")
		t.Execute(w, token)

		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)
		cookie := http.Cookie{Name: "cookie", Value: "astaxie", Expires: expiration}
		fmt.Println("设置的cooklie是==",cookie)
		http.SetCookie(w, &cookie)
	}else {
		r.ParseForm() //这一行代码有点意思  加上我这里才会输出  有点意思哦
		//请求的是登录的数据，那么执行登录的逻辑

		cookie, _ := r.Cookie("cookie")
		fmt.Fprint(w, cookie)
		fmt.Println("得到的cooklie是==",cookie)
		template.HTMLEscape(w, []byte("欢迎仕明大哥登录")) //输出到客户端  反馈给浏览器
	}
}


//s