package main

import (
	"fmt"
	"time"
)

/*
介绍了如何使用及存储本地资源，有时需要通过转换函数来实现，有时通过lang来设置，但是最终都是通过key-value的方式来存储Locale对应的数据，在需要时取出相应于Locale的信息后，如果是文本信息就直接输出，如果是时间日期或者货币，则需要先通过fmt.Printf或其他格式化函数来处理，而对于不同Locale的视图和资源则是最简单的，只要在路径里面增加lang就可以实现了。
 */

func init() {

    fmt.Println("本地化资源  前面小节我们介绍了如何设置Locale，设置好Locale之后我们需要解决的问题就是如何存储相应的Locale对应的信息呢？这里面的信息包括：文本信息、时间和日期、货币值、图片、包含文件以及视图等资源。那么接下来我们将对这些信息一一进行介绍，Go语言中我们把这些格式信息存储在JSON中，然后通过合适的方式展现出来。接下来以中文和英文两种语言对比举例,存储格式文件en.json和zh-CN.json")
}

func main() {


	//本地化文本消息
	textDemo()


    //本地化日期和时间
    //因为时区的关系，同一时刻，在不同的地区，表示是不一样的，而且因为Locale的关系，时间格式也不尽相同，例如中文环境下可能显示：2012年10月24日 星期三 23时11分13秒 CST，而在英文环境下可能显示:Wed Oct 24 23:11:13 CST 2012。这里面我们需要解决两点:
	//时区问题
	//格式问题
    timeDemo()


    // 本地化货币值
    moneyDemo()


	 // 本地化视图和资源
	  //我们可能会根据Locale的不同来展示视图，这些视图包含不同的图片、css、js等各种静态资源。那么应如何来处理这些信息呢？首先我们应按locale来组织文件信息
	  localeViewAndResource()




}

func localeViewAndResource() {
	//views
	//|--en  //英文模板
	//  |--images     //存储图片信息
    	//|--js         //存储JS文件
    	//|--css        //存储css文件
	//  index.tpl     //用户首页
	//  login.tpl     //登陆首页
	//|--zh-CN //中文模板
	//  |--images
	//  |--js
	//  |--css
	//  index.tpl
	//  login.tpl
	//这个目录结构后我们就可以在渲染的地方这样来实现代码
	//   todo
	//s1, _ := template.ParseFiles("views/"+lang+"/index.tpl")
	//VV.Lang=lang
	//s1.Execute(os.Stdout, VV)

	  //    todo  而对于里面的index.tpl里面的资源设置如下：
	// js文件
	//<script type="text/javascript" src="views/{{.Lang}}/js/jquery/jquery-1.8.0.min.js"></script>
	//// css文件
	//<link href="views/{{.Lang}}/css/bootstrap-responsive.min.css" rel="stylesheet">
	//// 图片文件
	//<img src="views/{{.Lang}}/images/btn.png">





}
/*
各个地区的货币表示不一样，处理的方式和日期差不多
 */
func moneyDemo() {
	str1:="USD %d"
	str2:="￥%d元"
	// todo  这一句话  感觉打印不出来啊啊啊啊
	fmt.Sprintf(str1,100)
	fmt.Println(fmt.Sprintf(str1,100))
	fmt.Println(fmt.Sprintf(str2,100))

}

/*
$GOROOT/lib/time包中的timeinfo.zip含有locale对应的时区的定义，为了获得对应于当前locale的时间，我们应首先使用time.LoadLocation(name string)获取相应于地区的locale，比如Asia/Shanghai或America/Chicago对应的时区信息，然后再利用此信息与调用time.Now获得的Time对象协作来获得最终的时间
 */
func timeDemo() {
    loc,err:= time.LoadLocation("America/Chicago")
    fmt.Println("loc===",loc)
    fmt.Println("err===",err)
     t:=time.Now()
     t=t.In(loc)
     fmt.Println("现在美国的芝加哥的时间是=",t.Format(time.RFC3339))

	loc1,err1:= time.LoadLocation("Asia/Shanghai")
	fmt.Println("loc===",loc1)
	fmt.Println("err===",err1)
	t1:=time.Now()
	t1=t1.In(loc1)
	fmt.Println("现在亚洲上海的时间是=",t1.Format(time.RFC3339))
     //  todo  了获得对应于当前locale的时间，我们应首先使用time.LoadLocation(name string)获取相应于地区的locale，比如Asia/Shanghai或America/Chicago对应的时区信息，然后再利用此信息与调用time.Now获得的Time对象协作来获得最终的时间
     //  感觉没有这个区啊  我日
	//loc2,err2:= time.LoadLocation("Asia/Shengzhen")
	//fmt.Println("loc===",loc2)
	//fmt.Println("err===",err2)
	//t2:=time.Now()
	//  todo
	//loc=== UTC
	//err=== cannot find Asia/Shengzhen in zip file E:\Go\lib\time\zoneinfo.zip
	//t2=t2.In(loc2)
	//fmt.Println("现在亚洲深圳的时间是=",t2.Format(time.RFC3339))


	//timeFormat1:="%Y-%m-%d %H:%M:%S"
	timeFormatNew:="%d-%d-%d %d:%d:%d"
	timeFormat2:="%d年%d月%d日 %d时%d分%d秒"
    nowTime:=time.Now()
    data(timeFormatNew,nowTime)
    data(timeFormat2,nowTime)


}
func data(format string, nowTime time.Time) {
	year, month, day :=nowTime.Date()
	hour, min, sec := nowTime.Clock()
	//2018 July 24 14 42 15
	fmt.Println(year, month, day ,hour, min, sec)
	str:=fmt.Sprintf(format,year, month, day ,hour, min, sec)
    //时间是多少，注意有个地方打印不出来= 2018年7月24日 14时54分4秒   有点意思   todo
	fmt.Println("时间是多少，注意有个地方打印不出来=",str)
}
/*
文本信息是编写Web应用中最常用到的，也是本地化资源中最多的信息，想要以适合本地语言的方式来显示文本信息，可行的一种方案是:建立需要的语言相应的map来维护一个key-value的关系，在输出之前按需从适合的map中去获取相应的文本
 */
 var locales map[string]map[string]string
func textDemo() {
	locales=make(map[string]map[string]string,2)
	en:=make(map[string]string,10)
	en["name"]="shiming"
	en["age"]="18"

	en["who am i"]="i am shiming,%d years old"
	locales["en"]=en
	//如果呢，输出的map为null  里面什么都没有，如果有数据的话，map的集合
    fmt.Println(locales)

	cn:=make(map[string]string,10)
	cn["姓名仕明"]="仕明"
	cn["年级"]="19岁"
	locales["cn"]=cn
	fmt.Println("填充了两次最新的值=",locales)

	lang:="cn"

	fmt.Println(msg(lang,"shiming"))
	fmt.Println(msg(lang,"姓名仕明"))

	fmt.Println(msg("en","who am i"))
	//结合fmt.Printf函数来实现   来实现特殊的功能
	fmt.Printf(msg("en","who am i"),22)

   fmt.Println()
}
func msg(lang,key string) string{
    if v,ok:= locales[lang];ok {
    	fmt.Println("v=",v)
		if v1, ok := v[key]; ok {
			return v1
		}
	}
	return "没有哦"
}
