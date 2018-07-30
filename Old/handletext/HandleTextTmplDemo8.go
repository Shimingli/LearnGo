package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	fmt.Println("如果字段里面还有对象，如何来循环")
	// 可以使用 {{with ...}} ...{{end}}和 {{range...}}{{end}}来进行数据的输出
	//1 {{range}} 这个和Go语法里面的range类似，循环操作数据
	// 2 {{with}} 操作是指当前对象的值，类型上下文的感念


	f1:= FriendD{Fname:"shiming11111111111111"}
	f2:= FriendD{Fname:"shiming22222222222222"}
	t:= template.New("tempalte  example")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .FriendD}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	p := PersonD{UserName: "shiming",
		Emails:  []string{"shiming1@beego.me", "shiming2@gmail.com"},
		FriendD: []*FriendD{&f1, &f2}}
	t1:= template.New("tempalte  example")
	t1, _ = t1.Parse(`{{.}}`)
	p1 := PersonD{UserName: "shiming",
		Emails:  []string{"shiming1@beego.me", "shiming2@gmail.com"},
		FriendD: []*FriendD{&f1, &f2}}

	t.Execute(os.Stdout, p)
	//{shiming [shiming1@beego.me shiming2@gmail.com] [0xc042042270 0xc042042280]}
	t1.Execute(os.Stdout, p1)


    // 条件处理
    // 在Go模板里面如果需要进行条件的判断，那么我们可以使用 Go语言的 if -esle 语法类似的方法来处理，如果pipeline为空，那么if 就认为是false
    fmt.Println("Demo start ")
    pipelineDemo()



}

func pipelineDemo() {
 	t:= template.New("template test")
 	//Parse 代替 ParseFiles ,因为Parse可以直接测试一个字符串，而不需要额外的文件
 	t=template.Must(t.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
 	t.Execute(os.Stdout,nil)
    //if里面无法使用条件判断，例如.Mail=="astaxie@gmail.com"，这样的判断是不正确的，if里面只能是bool值
	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)


	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	tIfElse1 := template.New("template test")
	tIfElse1 = template.Must(tIfElse1.Parse("if-else demo: {{if ``}} if部分 {{else}} else部分输出  {{end}}\n"))
	tIfElse1.Execute(os.Stdout, nil)

}



type FriendD struct {
	Fname string
}
type PersonD struct {
	UserName  string
    Emails   []string
    FriendD  []*FriendD
}

