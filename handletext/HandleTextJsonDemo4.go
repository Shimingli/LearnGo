package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {

	fmt.Println("JSON包里面通过Marshal函数来处理")
	var s ServersliceJson
	s.Servers = append(s.Servers, ServerJsonD{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, ServerJsonD{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//{"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}  输出的结果
	//我们看到上面的输出字段名的首字母都是大写的，如果你想用小写的首字母怎么办呢？把结构体的字段名改成首字母小写的？JSON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么就会发现什么都不会输出，所以必须通过struct tag定义来实现：


    //如果 ServerIP 为空，则不输出到JSON串中
	sss := ServerDemoDemo{ID:3, ServerName:`Go "1.0" `, ServerName2:`Go "1.0" `, ServerIP:``}
	db, _ := json.Marshal(sss)
	os.Stdout.Write(db)

	//{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}

	sss1 := ServerDemoDemo{ID:3, ServerName:`Go "1.0" `, ServerName2:`Go "1.0" `, ServerIP:`158`}
	db1, _ := json.Marshal(sss1)
	os.Stdout.Write(db1)
    // {"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \"","serverIP":"158"}

	//Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
	//
	//JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
	//Channel, complex和function是不能被编码成JSON的
	//嵌套的数据是不能编码的，不然会让JSON编码进入死循环
	//指针在编码的时候会输出指针指向的内容，而空指针会输出null



}
type ServerJsonD struct {
	//ServerName string
	//ServerIP   string
	//为了把首尾的字母变成小写
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type ServersliceJson struct {

	//Servers []ServerJsonD
	//为了把首尾的字母变成小写
	Servers []ServerJsonD `json:"servers"`
}

//
//针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:
//
//字段的tag是"-"，那么这个字段不会输出到JSON
//tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
//tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
//如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串

//{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}
type ServerDemoDemo struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP   string `json:"serverIP,omitempty"`
}