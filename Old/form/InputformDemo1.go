package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"net/url"
)

func init() {

}

func main() {
    http.HandleFunc("/",sayMe)//设置访问的路由
    http.HandleFunc("/login",login)
    http.ListenAndServe(":9090",nil)

}

func sayMe(w  http.ResponseWriter, r  *http.Request)  {
	r.ParseForm()
	fmt.Println("nihao shiming")
	fmt.Fprintf(w,"nihaogege")
}


func login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("打印的方法：method",r.Method)
	fmt.Println("打印的方法：r.Form",r.Form)


	//request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息，下面展示了可以对form数据进行的一些操作:
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
	//request本身也提供了FormValue()函数来获取用户提交的参数。如r.Form["username"]也可写成r.FormValue("username")。调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。




	if r.Method=="GET" {
		// 文件放在 根目录
		t, _ := template.ParseFiles("login.gtpl")
		// t,_:=template.ParseFiles("login.gtpl")
		 log.Println(t.Execute(w,nil))
	 }else{
		//我们输入用户名和密码之后发现在服务器端是不会打印出来任何输出的，为什么呢？默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。我们修改一下代码，在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译，再次测试输入递交，现在是不是在服务器端有输出你的输入的用户名和密码了。
		r.ParseForm() //这一行代码有点意思  加上我这里才会输出  有点意思哦
	 	//请求的是登录的数据，那么执行登录的逻辑
	 	fmt.Println("usename",r.Form["username"])
	 	fmt.Println("password",r.Form["password"])
	}

    //<form action="/login?username=astaxie" method="post">  当login.gtpl 这里面这样的修改的话，输入下面的结果
	//usename [李仕明 astaxie]
	//password [lishiming]

	//<form action="/login" method="post">  当login.gtpl 这里面这样的修改的话，输入下面的结果
	//usename [李仕明]
	//password [lishiming]
}