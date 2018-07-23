package main

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"io"
	"strconv"
	"time"
	"html/template"
)

func init() {
	fmt.Println("CSRF Demo  CSRF（Cross-site request forgery），中文名称：跨站请求伪造，也被称为：one click attack/session riding，缩写为：CSRF/XSRF")
}
func main() {
	fmt.Println("start Demo  那么CSRF到底能够干嘛呢？你可以这样简单的理解：攻击者可以盗用你的登陆信息，以你的身份模拟发送各种请求。攻击者只要借助少许的社会工程学的诡计，例如通过QQ等聊天软件发送的链接(有些还伪装成短域名，用户无法分辨)，攻击者就能迫使Web应用的用户去执行攻击者预设的操作。例如，当用户登录网络银行去查看其存款余额，在他没有退出时，就点击了一个QQ好友发来的链接，那么该用户银行帐户中的资金就有可能被转移到攻击者指定的帐户中。")

	str:="所以遇到CSRF攻击时，将对终端用户的数据和操作指令构成严重的威胁；当受攻击的终端用户具有管理员帐户的时候，CSRF攻击将危及整个Web应用程序。"
	fmt.Println(str)

	//完成一次 CSRF （跨站请求伪造 ） 需要满足两个条件
    // 1 登录受到信任的网址，并在本地生成cookie
    // 2  在不退出网址的情况下，访问危险的网址
     //  todo 两个条件 不满足一个，就不会受到CSRF的攻击
     //  但是你不能保证以下情况不会发生：
	// 1、 你不能保证你登录了一个网站后，不再打开一个tab页面并访问另外的网站，特别现在浏览器都是支持多tab的。
	// 2、 你不能保证你关闭浏览器了后，你本地的Cookie立刻过期，你上次的会话已经结束。
	// 3、 所谓的攻击网站，可能是一个存在其他漏洞的可信任的经常被人访问的网站。

	//  todo CSRF攻击主要是因为Web的隐式身份验证机制，Web的身份验证机制虽然可以保证一个请求是来自于某个用户的浏览器，但却无法保证该请求是用户批准发送的。


	//如何预防CSRF?
	// CSRF的防御可以从服务端和客户端两方面着手，防御效果是从服务端着手效果比较好，现在一般的CSRF防御也都在服务端进行。

	// todo  1、正确使用GET,POST和Cookie；
	// todo  2、在非GET请求中增加伪随机数；


	//我们上一章介绍过REST方式的Web应用，一般而言，普通的Web应用都是以GET、POST为主，还有一种请求是Cookie方式。我们一般都是按照如下方式设计应用：
	//
	//1、GET常用在查看，列举，展示等不需要改变资源属性的时候；
	//
	//2、POST常用在下达订单，改变一个资源的属性或者做其他一些事情；
	//
	//接下来我就以Go语言来举例说明，如何限制对资源的访问方法：


	//在非GET方式的请求中增加随机数，这个大概有三种方式来进行：
	//
	//为每个用户生成一个唯一的cookie token，所有表单都包含同一个伪随机值，这种方案最简单，因为攻击者不能获得第三方的Cookie(理论上)，所以表单中的数据也就构造失败，但是由于用户的Cookie很容易由于网站的XSS漏洞而被盗取，所以这个方案必须要在没有XSS的情况下才安全。
	//每个请求使用验证码，这个方案是完美的，因为要多次输入验证码，所以用户友好性很差，所以不适合实际运用。
	//不同的表单包含一个不同的伪随机值，我们在4.4小节介绍“如何防止表单多次递交”时介绍过此方案，复用相关代码，实现如下：



	//生成随机数的token
	tokenDemo()


}

func tokenDemo() {



    http.HandleFunc("/token",tokenDD)
    http.ListenAndServe(":9090",nil)
}
func tokenDD( w  http.ResponseWriter,h *http.Request)  {
	fmt.Println("是什么的请求--》",h.Method)
	md5:=md5.New()
	time:=time.Now()
    io.WriteString(md5,strconv.FormatInt(	time.Unix(),10))
	io.WriteString(md5,"shiming")

	token:=fmt.Sprintf("%x",md5.Sum(nil))
    fmt.Println("服务器返回的token==",token)
    t,_:= 	template.ParseFiles("securityAndEncryption/test.gtpl")

    t.Execute(w,token)



    h.ParseForm()
    token1:= h.Form.Get("token")
	if token1!= "" {
		fmt.Println("shiming token==",token1)
		template.HTMLEscape(w, []byte("欢迎仕明大哥登录\n你的token===="+token1)) //输出到客户端  反馈给浏览器
	}

    // 破解token理论上是可以破解，但是实际实际上破解是基本不可能的额 ，暴力破解token需要的时间是 2的11次的时间

    //  todo 跨站请求伪造，即CSRF，是一种非常危险的Web安全威胁，它被Web安全界称为“沉睡的巨人”，其威胁程度由此“美誉”便可见一斑

}