package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func init() {
	fmt.Println("预防跨站的脚本")
	fmt.Println("现在的网站包含大量的动态的内容已提高用户的体验，比过去复杂的多，动态的内容：就是根据用户环境和需要，Web应用程序能够输出相对应的内容。动态的站点会受到一种名为“跨站脚本的攻击（Cross Site Scripting）Cross Site Script，缩写CSS又叫XSS，中文意思跨站脚本攻击，指的是恶意攻击者往Web页面里插入恶意html代码。”但是静态网站则不受其影响 ")
	fmt.Println("要认真检查您的应用程序是否存在XSS漏洞。必须明确：一切输入都是有害的，不要信任一切输入的数据。 缓和XSS问题的首要法则是确定哪个输入是有效的，并且拒绝所有别的无效输入。替换危险字符，如：&, <, >, ，', /, ?，;, :, %,<SPACE>, =, +。各种语言替换的程度不尽相同，但是基本上能抵御住一般的XSS攻击。")

   fmt.Println("攻击者通常会在有漏洞的程序中插入JavaScript、VBScript、 ActiveX或Flash以欺骗用户。一旦得手，他们可以盗取用户帐户信息，修改用户设置，盗取/污染cookie和植入恶意广告等。对XSS最佳的防护应该结合以下两种方法：一是验证所有输入数据，有效检测攻击(这个我们前面小节已经有过介绍);另一个是对所有输出数据进行适当的处理，以防止任何已成功注入的脚本在浏览器端运行。")



}
func main() {
	http.HandleFunc("/login",loginThree)
	http.ListenAndServe(":9092",nil)
}
func loginThree(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("打印的方法：method",r.Method)
	fmt.Println("打印的方法：r.Form",r.Form)

	if r.Method=="GET" {
		// 文件放在 根目录
		t, _ := template.ParseFiles("loginThree.gtpl")
		// t,_:=template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	}else{
		//我们输入用户名和密码之后发现在服务器端是不会打印出来任何输出的，为什么呢？默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。我们修改一下代码，在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译，再次测试输入递交，现在是不是在服务器端有输出你的输入的用户名和密码了。
		r.ParseForm() //这一行代码有点意思  加上我这里才会输出  有点意思哦
		//请求的是登录的数据，那么执行登录的逻辑
		fmt.Println("usename",r.Form["username"])
		fmt.Println("password",r.Form["password"])

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		//template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		//template.HTMLEscape(w, []byte("欢迎仕明大哥登录")) //输出到客户端  反馈给浏览器

		/**
		如果我们输入的username是<script>alert()</script>,那么我们可以在浏览器上面看到输出如
		 */
		 //&lt;script&gt;alert()&lt;/script&gt;

		//Go的html/template包默认帮你过滤了html标签，但是有时候你只想要输出这个<script>alert()</script>看起来正常的信息，该怎么处理？请使用text/template


		//import "text/template"    todo   注意导的包是 这个里面的啊
		//t, err1 := 	template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)//`` 由于有两个的 “”  所有需要转义下
		//err1 = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
		//fmt.Println(err1)
		//template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		//  Hello, <script>alert('you have been pwned')</script>!   输出的结果

        //使用template.HTML类型
		//t, err1 := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		//err1 = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		//fmt.Println(err1)
		// 输出的结果是  Hello, <script>alert('you have been pwned')</script>!



		//转换成template.HTML后，变量的内容也不会被转义
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
		fmt.Println(err)
		// 输出的结果   Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!

	}

}