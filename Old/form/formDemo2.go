package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"strconv"
	"regexp"
	"time"
)

func init() {
   fmt.Println("开发Web的一个原则就是，不能信任用户输入的任何信息，所以验证和过滤用户的输入信息就变得非常重要，我们经常会在微博、新闻中听到某某网站被入侵了，存在什么漏洞，这些大多是因为网站对于用户输入的信息没有做严格的验证引起的，所以为了编写出安全可靠的Web程序，验证表单输入的意义重大。")
}

func main() {
  fmt.Println("我们平常编写Web应用主要有两方面的数据验证，一个是在页面端的js验证(目前在这方面有很多的插件库，比如ValidationJS插件)，一个是在服务器端的验证，我们这小节讲解的是如何在服务器端验证。")
  http.HandleFunc("/shiming",loginTwo)
  http.ListenAndServe(":9091",nil)
}

func loginTwo(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("打印的方法：method",r.Method)
	fmt.Println("打印的方法：r.Form",r.Form)

	if r.Method=="GET" {
		// 文件放在 根目录
		t, _ := template.ParseFiles("loginTwo.gtpl")
		// t,_:=template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))

		//		下拉菜单
		//		如果我们想要判断表单里面<select>元素生成的下拉菜单中是否有被选中的项目。有些时候黑客可能会伪造这个下拉菜单不存在的值发送给你，那么如何判断这个值是否是我们预设的值呢？
		//我们的select可能是这样的一些元素
		//  TODO  注意 selct的提交和 下面的这一段代码加上的必要性啊  傻逼！！！ 大傻逼
		r.ParseForm()
		slice:=[]string{"apple","pear","banana"}
		v := r.Form.Get("fruit")
		fmt.Println("下拉菜单","v====",v)
		for _, item := range slice {
			if item == v {
				fmt.Println("下拉菜单",item,"v====",v)
			}
		}


		//单选按钮
		//如果我们想要判断radio按钮是否有一个被选中了，我们页面的输出可能就是一个男、女性别的选择，但是也可能一个15岁大的无聊小孩，一手拿着http协议的书，另一只手通过telnet客户端向你的程序在发送请求呢，你设定的性别男值是1，女是2，他给你发送一个3，你的程序会出现异常吗？因此我们也需要像下拉菜单的判断方式类似，判断我们获取的值是我们预设的值，而不是额外的值。

		slice1:=[]string{"1","2"}

		for _, v := range slice1 {
			if v == r.Form.Get("gender") {
				fmt.Println("单选按钮","v====",v)
			}
		}

		//复选框
		//有一项选择兴趣的复选框，你想确定用户选中的和你提供给用户选择的是同一个类型的数据。
         //对于复选框我们的验证和单选有点不一样，因为接收到的数据是一个slice
		slice2:=[]string{"football","basketball","tennis"}
		mmm:=r.Form["interest"]
		//a:= Slice_diff(mmm,slice2)
		//if a == nil{
		//	fmt.Println("复选框"," a == nil")
		//}
		fmt.Println("slice2=",slice2)
		fmt.Println("mmm=",mmm)

		//身份证号码
		//如果我们想验证表单输入的是否是身份证，通过正则也可以方便的验证，但是身份证有15位和18位，我们两个都需要验



	}else{
		//我们输入用户名和密码之后发现在服务器端是不会打印出来任何输出的，为什么呢？默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。我们修改一下代码，在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译，再次测试输入递交，现在是不是在服务器端有输出你的输入的用户名和密码了。
		r.ParseForm() //这一行代码有点意思  加上我这里才会输出  有点意思哦
		//请求的是登录的数据，那么执行登录的逻辑
		fmt.Println("usename",r.Form["username"])
		fmt.Println("password",r.Form["password"])


		//r.Form对不同类型的表单元素的留空有不同的处理， 对于空文本框、空文本区域以及文件上传，元素的值为空值,而如果是未选中的复选框和单选按钮，则根本不会在r.Form中产生相应条目，如果我们用上面例子中的方式去获取数据时程序就会报错。所以我们需要通过r.Form.Get()来获取值，因为如果字段不存在，通过该方式获取的是空值。但是通过r.Form.Get()只能获取单个的值，如果是map的值，必须通过下面的方式来获取。
		//1 必填字段
		//Go有一个内置函数len可以获取字符串的长度，这样我们就可以通过len来获取数据的长度
		if len(r.Form["username"][0]) != 0 {
			fmt.Println("username 长度不为0")
		}
        // 2 数字
        //如果我们是判断正整数，那么我们先转化成int类型，然后进行处理


        getAge,err := strconv.Atoi(r.Form.Get("age"))
		if err!=nil{
			//数字转化出错了，那么可能就不是数字
			fmt.Println("err====",err)
		}
		//接下来就可以判断这个数字的大小范围了
		if getAge >100 {
			//太大了
			fmt.Println("太大了age:=",getAge)
		}else {
			fmt.Println("小了 ：：",getAge)
		}

		//正则匹配的方式验证是否是数字
		m,_:=regexp.MatchString("^[0-9]+$",r.Form.Get("age"))
		fmt.Println("年级是否是数字：",m)
		//对于性能要求很高的用户来说，这是一个老生常谈的问题了，他们认为应该尽量避免使用正则表达式，因为使用正则表达式的速度会比较慢。但是在目前机器性能那么强劲的情况下，对于这种简单的正则表达式效率和类型转换函数是没有什么差别的。如果你对正则表达式很熟悉，而且你在其它语言中也在使用它，那么在Go里面使用正则表达式将是一个便利的方式。

		//Go实现的正则是RE2，所有的字符都是UTF-8编码的。


		//中文
		//有时候我们想通过表单元素获取一个用户的中文名字，但是又为了保证获取的是正确的中文，我们需要进行验证，而不是用户随便的一些输入。对于中文我们目前有两种方式来验证，可以使用 unicode 包提供的 func Is(rangeTab *RangeTable, r rune) bool 来验证，也可以使用正则方式来验证，这里使用最简单的正则方式，如下代码所示

		username,_ := regexp.MatchString("^\\p{Han}+$",r.Form.Get("username"))
        fmt.Println("验证username是否为汉字",username)

		//英文
		//我们期望通过表单元素获取一个英文值，例如我们想知道一个用户的英文名，应该是astaxie，而不是asta谢。
		//   key=eng
		eng,_:= regexp.MatchString("^[a-zA-Z]+$",r.Form.Get("eng"))
		fmt.Println("验证是否为eng",eng)


		//电子邮件地址
		//你想知道用户输入的一个Email地址是否正确，通过如下这个方式可以验证：
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}

		//手机号码
		//你想要判断用户输入的手机号码是否正确，通过正则也可以验证：
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			fmt.Println("不是手机号码")
		}else {
			fmt.Println("是手机号码")
		}


		//日期和时间
		//你想确定用户填写的日期或时间是否有效。例如 ，用户在日程表中安排8月份的第45天开会，或者提供未来的某个时间作为生日。
		//
		//Go里面提供了一个time的处理包，我们可以把用户的输入年月日转化成相应的时间，然后进行逻辑判断 获取time之后我们就可以进行很多时间函数的操作。具体的判断就根据自己的需求调整。
		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())  //Go launched at 2009-11-11 07:00:00 +0800 CST


		//验证15位身份证，15位的是全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); m {
			fmt.Println("验证15位身份证======",r.Form.Get("usercard"))
		}

		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); m {
			fmt.Println("验证18位身份证======",r.Form.Get("usercard"))
		}
	}
}