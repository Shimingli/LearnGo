package main

import (
	"fmt"
	"net/http"
	"html/template"
	"regexp"
)

func init() {
	fmt.Println("确保输入的过滤了  过滤用户数据是Web应用安全的基础。它是验证数据合法性的过程。通过对所有的输入数据进行过滤，可以避免恶意数据在程序中被误信或误用。大多数Web应用的漏洞都是因为没有对用户输入的数据进行恰当过滤所引起的。")

}

func main() {


	//我们介绍的过滤数据分成三个步骤：

	//1、识别数据，搞清楚需要过滤的数据来自于哪里
	//2、过滤数据，弄明白我们需要什么样的数据
	//3、区分已过滤及被污染数据，如果存在攻击数据那么保证过滤之后可以让我们使用更安全的数据


	//   todo   识别数据
	//“识别数据”作为第一步是因为在你不知道“数据是什么，它来自于哪里”的前提下，你也就不能正确地过滤它。这里的数据是指所有源自非代码内部提供的数据。例如:所有来自客户端的数据，但客户端并不是唯一的外部数据源，数据库和第三方提供的接口数据等也可以是外部数据源。
	//
	//由用户输入的数据我们通过Go非常容易识别，Go通过r.ParseForm之后，把用户POST和GET的数据全部放在了r.Form里面。其它的输入要难识别得多，例如，r.Header中的很多元素是由客户端所操纵的。常常很难确认其中的哪些元素组成了输入，所以，最好的方法是把里面所有的数据都看成是用户输入。(例如r.Header.Get("Accept-Charset")这样的也看做是用户输入,虽然这些大多数是浏览器操纵的)


	//   todo   过滤数据
	//在知道数据来源之后，就可以过滤它了。过滤是一个有点正式的术语，它在平时表述中有很多同义词，如验证、清洁及净化。尽管这些术语表面意义不同，但它们都是指的同一个处理：防止非法数据进入你的应用。
	//
	//过滤数据有很多种方法，其中有一些安全性较差。最好的方法是把过滤看成是一个检查的过程，在你使用数据之前都检查一下看它们是否是符合合法数据的要求。而且不要试图好心地去纠正非法数据，而要让用户按你制定的规则去输入数据。历史证明了试图纠正非法数据往往会导致安全漏洞。这里举个例子：“最近建设银行系统升级之后，如果密码后面两位是0，只要输入前面四位就能登录系统”，这是一个非常严重的漏洞。

	//    todo  过滤数据主要采用如下一些库来操作：
	//
	//strconv包下面的字符串转化相关函数，因为从Request中的r.Form返回的是字符串，而有些时候我们需要将之转化成整/浮点数，Atoi、ParseBool、ParseFloat、ParseInt等函数就可以派上用场了。
	//string包下面的一些过滤函数Trim、ToLower、ToTitle等函数，能够帮助我们按照指定的格式获取信息。
	//regexp包用来处理一些复杂的需求，例如判定输入是否是Email、生日之类。
	//过滤数据除了检查验证之外，在特殊时候，还可以采用白名单。即假定你正在检查的数据都是非法的，除非能证明它是合法的。使用这个方法，如果出现错误，只会导致把合法的数据当成是非法的，而不会是相反，尽管我们不想犯任何错误，但这样总比把非法数据当成合法数据要安全得多。


     //  todo  区分过滤数据
	//如果完成了上面的两步，数据过滤的工作就基本完成了，但是在编写Web应用的时候我们还需要区分已过滤和被污染数据，因为这样可以保证过滤数据的完整性，而不影响输入的数据。我们约定把所有经过过滤的数据放入一个叫全局的Map变量中(CleanMap)。这时需要用两个重要的步骤来防止被污染数据的注入：
	//
	//每个请求都要初始化CleanMap为一个空Map。
	//加入检查及阻止来自外部数据源的变量命名为CleanMap。




    http.HandleFunc("/mapDemo",mapDemo)
     http.ListenAndServe(":9090",nil)
}

func mapDemo(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("请求的方式",r.Method)
    t,_:=template.ParseFiles("securityAndEncryption/test2.gtpl")
    t.Execute(w,nil)
    r.ParseForm()
    name:=r.Form.Get("name")
    fmt.Println("name=",name)
     // 创建一个map
     newMap:= make(map[string]string,0)
    if name=="1"{
    	newMap["name"]="shiming"
	}
	if name=="2"{
		newMap["name"]="daming"
	}
	if name=="3"{
		newMap["name"]="xiaoming"
	}
    fmt.Println("newMAp=",newMap)
	var str  string
	for key,V := range newMap{
		fmt.Println("key=",key)
		fmt.Println("v=",V)
		str=V
	}

    template.HTMLEscape(w,[]byte("你选着完了，选择的内容是：name="+str))



   // 数据过滤在Web安全中起到一个基石的作用，大多数的安全问题都是由于没有过滤数据和验证数据引起的

     //字母和数字的组合
     if ok,_:=regexp.MatchString("^[a-zA-Z0-9]+$","shiming");ok{
     	fmt.Println("shiming 是否满足啊 ok=",ok)
	 }

}


