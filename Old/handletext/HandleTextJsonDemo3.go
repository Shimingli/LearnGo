package main

import (
	"fmt"
	"encoding/json"
	_ "github.com/bitly/go-simplejson"
	"github.com/bitly/go-simplejson"
)

func main() {

	fmt.Println("JSON（Javascript Object Notation）是一种轻量级的数据交换语言，以文字为基础，具有自我描述性且易于让人阅读。尽管JSON是Javascript的一个子集，但JSON是独立于语言的文本格式，并且采用了类似于C语言家族的一些习惯。JSON与XML最大的不同在于XML是一个完整的标记语言，而JSON不是。JSON由于比XML更小、更快，更易解析,以及浏览器的内建快速解析支持,使得其更适用于网络数据传输领域。目前我们看到很多的开放平台，基本上都是采用了JSON作为他们的数据交互的接口。")

	//定义了与json数据对应的结构
	var s Serverslice
	//数组对应slice，字段名对应JSON里面的KEY
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	fmt.Println("start s===", s)
	json.Unmarshal([]byte(str), &s)
	fmt.Println("end  s===", s)

	//解析到interface,如果我们不知道被解析的数据的格式，第一点我们知道 interface{} 可以用来储存任意数据类型的对象，这种数据结构正好用于储存解析的未知结构的json的数据结构， JSON 包中采用map[string]interface{}结构来储存任意的JSON对象和数组。
	//bool 代表 JSON booleans,
	//	float64 代表 JSON numbers,
	//	string 代表 JSON strings,
	//	nil 代表 JSON null.

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{}

	json.Unmarshal(b, &f)
	fmt.Println("new b==", f) //map[Name:Wednesday Age:6 Parents:[Gomez Morticia]]

	// 如何来访问这里面的数据呢，通过断言来
	//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
	m, _ := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string ", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64 ", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println("i===", i, "   u==", u)
			}
		default:
			fmt.Println("老子不知道啊")
		}
	}

	//官方提供的解决方案，其实很多时候我们通过类型断言，操作起来不是很方便，目前bitly公司开源了一个叫做simplejson的包,在处理未知结构体的JSON时相当方便
	js, _ := simplejson.NewJson([]byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	}
}`))
	arr, _ := js.Get("test").Get("array").Array()
	fmt.Println("arr   ==== ",arr)
	i, _ := js.Get("test").Get("int").Int()
	fmt.Println("i   ==== ",i)
	ms := js.Get("test").Get("string").MustString()
	fmt.Println("ms   ==== ",ms)
}

//如何将json数据与struct字段相匹配呢？例如JSON的key是Foo，那么怎么找对应的字段呢？
//
//首先查找tag含有Foo的可导出的struct字段(首字母大写)
//其次查找字段名是Foo的导出字段
//最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段

//{[{Shanghai_VPN} {Beijing_VPN}]}  把下面的注释掉一行  当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题
type ServerJson struct {
	serverName string // 如果你不想要这个名称的话，就把它全部小写 就行了
	//ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []ServerJson
}
