package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

func main() {

      fmt.Println("模板处理")
      // 注意看图   7.4 template.png
      // Web应用反馈给客户端的信息中的大部分内容都是静态的，不变的，而另外少部分的内容更具的是用户的请求动态生成的，例如显示用户的访问记录列表。用户之间的只有记录数据是不同的，而列表的样式则是固定的，此时采用模板可以复用很多的代码

      //字段的操作
      templateDemo1()




}
func templateDemo1() {
	//新分配一个新的HTML模板与给定的名称。
	t:= template.New("shiming demo 11ddddddddddd")
	//Parse 代替 ParseFiles ,因为Parse可以直接测试一个字符串，而不需要额外的文件
	/*
	Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类似，如果要访问当前对象的字段通过{{.FieldName}}，但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的)，否则在渲染的时候就会报错
	 */
	//t,_=t.Parse(" hell0{{.UserName}}")
//	t,_=t.Parse(" hell0{{.UserName}}{{.}}")// hell0世明{世明 737141437@qq.com}
	t,_=t.Parse(" hell0{{.}}")//  hell0{世明 737141437@qq.com}
	p:=Person{"世明","737141437@qq.com"}

	// 使用 os.Stdout 是因为实现了 io.Writer接口
	t.Execute(os.Stdout,p)

}

type Person struct {
	UserName string
   // too few values in struct initializer     报错了  如下的信息
	//email	string  //未导出的字段，首字母是小写的
	Email	string  //未导出的字段，首字母是小写的
}


//为了演示和测试代码的方便，我们在接下来的例子中采用如下格式的代码
//使用Parse代替ParseFiles，因为Parse可以直接测试一个字符串，而不需要额外的文件
//不使用handler来写演示代码，而是每个测试一个main，方便测试
//使用os.Stdout代替http.ResponseWriter，因为os.Stdout实现了io.Writer接口
// Go 语言的模板操作非常的简单方便，其他语言非常的类似，都是先获取数据，然后渲染数据
func handler(w http.ResponseWriter, r *http.Request) {
	//t := template.New("some template") //创建一个模板
	//t, _ = t.ParseFiles("tmpl/welcome.html")  //解析模板文件
	//user := GetUser() //获取当前用户信息
	//t.Execute(w, user)  //执行模板的merger操作
}