package main

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func init() {
	fmt.Println("SQL 注入实例    ")
}

func main() {
	fmt.Println("很多Web开发者没有意识到SQL查询是可以被篡改的，从而把SQL查询当作可信任的命令。殊不知，SQL查询是可以绕开访问控制，从而绕过身份验证和权限检查的。更有甚者，有可能通过SQL查询去运行主机系统级的命令。")

	demoSQL()
}


func demoSQL() {


	http.HandleFunc("/sqldemo",SQLDemo)
	http.ListenAndServe(":9090",nil)

}
func checkErr(e error) {
	if e != nil {
		fmt.Println("发生错误了哦 e===",e)
	}
}

func SQLDemo(w http.ResponseWriter,r *http.Request)  {

    t,_:=	template.ParseFiles("securityAndEncryption/sql.stpl")

    t.Execute(w,nil)


    r.ParseForm()
	username:=r.Form.Get("username")
	password:=r.Form.Get("password")

	fmt.Println("password",password)
	fmt.Println("username",username)


	template.HTMLEscape(w,[]byte("你输入的是：password="+password+"\r\n"+ "username="+username))
	db, err := sql.Open("mysql", "root:App123@tcp(localhost:3306)/godbdemo?charset=utf8")
	checkErr(err)
	//查询数据  db.Query()函数用来直接执行Sql返回Rows结果。
	rows, err := db.Query("SELECT * FROM userinfo  WHERE username='shimingupdate' AND uid='5'")
	checkErr(err)
	fmt.Println("rows=",rows)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid=",uid)
		fmt.Println("username=",username)
		fmt.Println("department=",department)
		fmt.Println("created=",created)
	}



	//如何预防SQL注入
	//也许你会说攻击者要知道数据库结构的信息才能实施SQL注入攻击。确实如此，但没人能保证攻击者一定拿不到这些信息，一旦他们拿到了，数据库就存在泄露的危险。如果你在用开放源代码的软件包来访问数据库，比如论坛程序，攻击者就很容易得到相关的代码。如果这些代码设计不良的话，风险就更大了。目前Discuz、phpwind、phpcms等这些流行的开源程序都有被SQL注入攻击的先例。
	//
	//这些攻击总是发生在安全性不高的代码上。所以，永远不要信任外界输入的数据，特别是来自于用户的数据，包括选择框、表单隐藏域和 cookie。就如上面的第一个例子那样，就算是正常的查询也有可能造成灾难。
	//
	//SQL注入攻击的危害这么大，那么该如何来防治呢?下面这些建议或许对防治SQL注入有一定的帮助。
	//
	//严格限制Web应用的数据库的操作权限，给此用户提供仅仅能够满足其工作的最低权限，从而最大限度的减少注入攻击对数据库的危害。
	//检查输入的数据是否具有所期望的数据格式，严格限制变量的类型，例如使用regexp包进行一些匹配处理，或者使用strconv包对字符串转化成其他基本类型的数据进行判断。
	//对进入数据库的特殊字符（'"\尖括号&*;等）进行转义处理，或编码转换。Go 的text/template包里面的HTMLEscapeString函数可以对字符串进行转义处理。
	//所有的查询语句建议使用数据库提供的参数化查询接口，参数化的语句使用参数而不是将用户输入变量嵌入到SQL语句中，即不要直接拼接SQL语句。例如使用database/sql里面的查询函数Prepare和Query，或者Exec(query string, args ...interface{})。
	//在应用发布之前建议使用专业的SQL注入检测工具进行检测，以及时修补被发现的SQL注入漏洞。网上有很多这方面的开源工具，例如sqlmap、SQLninja等。
	//避免网站打印出SQL错误信息，比如类型错误、字段不匹配等，把代码里的SQL语句暴露出来，以防止攻击者利用这些错误信息进行SQL注入。



}