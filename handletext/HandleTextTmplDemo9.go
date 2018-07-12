package main

import (
	"fmt"
	"os"
	"html/template"
	"strings"
)

func main() {
	fmt.Println("模板函数")

	//模板在输出对象的字段值时，采用了fmt包把对象转化成了字符串。但是有时候我们的需求可能不是这样的，例如有时候我们为了防止垃圾邮件发送者通过采集网页的方式来发送给我们的邮箱信息，我们希望把@替换成at例如：astaxie at beego.me，如果要实现这样的功能，我们就需要自定义函数来做这个功能

	f1 := FriendDD{Fname: "minux.ma"}
	f2 := FriendDD{Fname: "xushiwei"}
	t := template.New("fieldname example")
	//每一个模板函数都有一个唯一值的名字，然后与一个Go函数关联，通过如下的方式来关联
	//type FuncMap map[string]interface{}
	//我们想要的email函数的模板函数名是emailDeal，它关联的Go函数名称是EmailDealWith,那么我们可以通过下面的方式来注册这个函数
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
				{{end}}
				`)
	p := PersonDD{UserName: "Astaxie",
		Emails:  []string{"astaxie  @    beego.me", "astaxie  @    gmail.com"},
		Friends: []*FriendDD{&f1, &f2}}
	t.Execute(os.Stdout, p)


   //Must操作
	mustDemo()

   //嵌套模板
   // 开发Web应用的时候，会遇到一些模板有些部分是固定不变的，然后可以抽取出来作为一个独立的部分，例如一个博客的头部和尾部是不变的，而唯一改变的就是中间的内容的部分，可以定义三个 header / content /footer 三个部分
     //Go语言中通过如下的语法来申明
	//{{define "子模板名称"}}内容{{end}}
   //通过如下方式来调用
	//{{template "子模板名称"}}
	demo()
}

/*
可以看到通过template.ParseFiles把所有的嵌套模板全部解析到模板里面，其实每一个定义的{{define}}都是一个独立的模板，他们相互独立，是并行存在的关系，内部其实存储的是类似map的一种关系(key是模板的名称，value是模板的内容)，然后我们通过ExecuteTemplate来执行相应的子模板内容，我们可以看到header、footer都是相对独立的，都能输出内容，content 中因为嵌套了header和footer的内容，就会同时输出三个的内容。但是当我们执行s1.Execute，没有任何的输出，因为在默认的情况下没有默认的子模板，所以不会输出任何的东西
 */
func demo() {
	fmt.Println("****************************************************")
	s1, _ := template.ParseFiles("handletext/header.tmpl", "handletext/content.tmpl", "handletext/footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

//Must操作
//模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
func mustDemo() {
	fmt.Println("------------------ ---------------------- ------------------- ---------------")
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	//panic: template: check parse error with Must:1: unexpected "}" in command
	//template.Must(tErr.Parse(" some static text {{ .Name }"))
	template.Must(tErr.Parse(" some static text {{ .Name }}"))


}
func EmailDealWith(args ... interface{}) string {
	ok := false
	var s string
	fmt.Println("args shiming  ===",args)
	if len(args)==1 {
		//s:=args[0]
		s,ok=args[0].(string)
		fmt.Println("s===",s)
		fmt.Println("ok===",ok)
	}
	if !ok {
		s=fmt.Sprint(args...)
	}
	substrs:=strings.Split(s,"@")
	if len(substrs)!=2 {
		return s
	}
	return substrs[0] +"at"+substrs[1]



}
type FriendDD struct {
	Fname string

}

type PersonDD struct {
	UserName string
    Emails  []string
    Friends []*FriendDD
}
