package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

//防止多次递交表单
func main() {
	http.HandleFunc("/shiming",loginForth)
	http.ListenAndServe(":9092",nil)
}
func loginForth(w http.ResponseWriter,r *http.Request) {
	fmt.Println("打印的方法：method", r.Method)
	fmt.Println("打印的方法：r.Form", r.Form)


	//第一次，会走 get的方法 ，把 get的方法 请求进来  就会走这个方法
	if r.Method == "GET" {
        //获取token   把 token 写入到这里来
		crutime := time.Now().Unix()//一个时间值
		h := md5.New()//新的返回一个新的散列。哈希计算MD5校验和。
		// FormatInt returns the string representation of i in the given base,
		// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
		// for digit values >= 10.
		//formatint返回的字符串表示形式，给出了第一的基础上，
		//2＜＜＝36＝基础。的结果的情况下使用字母“Z”
		//数字值≥10。
		io.WriteString(h, strconv.FormatInt(crutime, 10))

		//和将当前哈希追加到B并返回结果切片。
		//它不会改变潜在的哈希状态。
		token := fmt.Sprintf("%x", h.Sum(nil))

		// 文件放在 根目录
		t, _ := template.ParseFiles("loginForth.gtpl")
		fmt.Println("这个token 是多少====",token)
		log.Println(t.Execute(w, token))
	} else {
		//我们输入用户名和密码之后发现在服务器端是不会打印出来任何输出的，为什么呢？默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。我们修改一下代码，在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译，再次测试输入递交，现在是不是在服务器端有输出你的输入的用户名和密码了。
		r.ParseForm() //这一行代码有点意思  加上我这里才会输出  有点意思哦
		//请求的是登录的数据，那么执行登录的逻辑
		fmt.Println("usename", r.Form["username"])
		fmt.Println("password", r.Form["password"])



		//复选框
		//有一项选择兴趣的复选框，你想确定用户选中的和你提供给用户选择的是同一个类型的数据。
		//对于复选框我们的验证和单选有点不一样，因为接收到的数据是一个slice
		slice2:=[]string{"football","basketball","tennis"}
		mmm:=r.Form["interest"]
		fmt.Println("slice2=",slice2)
		fmt.Println("mmm=",mmm)
		template.HTMLEscape(w, []byte("提交完成，请手动返回，重新的提交一下")) //输出到客户端  反馈给浏览器


		//我们在模版里面增加了一个隐藏字段token，这个值我们通过MD5(时间戳)来获取唯一值，然后我们把这个值存储到服务器端(session来控制，我们将在第六章讲解如何保存)，以方便表单提交时比对判定。
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性

			// 可以 通过 view-source : 加上网页的网址 ，这里就可以 看到网页的源码

		} else {
			//不存在token报错
		}
		fmt.Println("点击提交了以后的token-=====》",token)
		//token  这个token，在同一个浏览器，是一样的
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

     //  token已经有输出值，你可以不断的刷新，可以看到这个值在不断的变化。这样就保证了每次显示form表单的时候都是唯一的，用户递交的表单保持了唯一性。
		//我们的解决方案可以防止非恶意的攻击，并能使恶意用户暂时不知所措，然后，它却不能排除所有的欺骗性的动机，对此类情况还需要更复杂的工作。

	}
}
