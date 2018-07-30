package main

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)

func main() {
	fmt.Println("XML的使用的Demo")
	s:=[]server{}
	n:=xml.Name{"dd","ddd"}
    V:=Recurlyservers{n,"dddd",s,"dfd"}
    fmt.Println("v==",V)
    //打开打开命名文件进行读取。如果成功，返回的文件上的方法可以用于读取；关联的文件描述符具有模式O_RDONLY(RDONLY 只读存储器)。如果有错误，它将是路径错误类型的*PathError。
     file,err:=os.Open("handletext/servers.xml")

	if err!=nil {
		fmt.Printf("error : %v",err)
		return
	}
	defer file.Close()
     //Read全部读取R直到一个错误或EOF并返回它读取的数据。
	//一个成功的调用返回Err= nIL，而不是Err= EOF。因为Read被定义为从SRC读取直到EOF，所以它不将EOF从读作为错误来报告。
    data ,err:=  ioutil.ReadAll(file)// 返回一个字节的数组
	if err!=nil {
		fmt.Printf("error : %v",err)
		return
	}

	v:=Recurlyservers{}
	//data接收的是XML数据流，v是需要输出的结构，定义为interface，也就是可以把XML转换为任意的格式
	err=xml.Unmarshal(data,&v)
	if err!=nil {
		fmt.Printf("error: %v",err)
	}
	fmt.Println("shiming  v==",v)
	//{{ servers} 1 [{{ } Shanghai_VPN } {{ } Beijing_VPN }]
	//	<server>
	//<serverName>Shanghai_VPN</serverName>
	//<serverIP>127.0.0.1</serverIP>
	//</server>
	//<server>
	//<serverName>Beijing_VPN</serverName>
	//<serverIP>127.0.0.2</serverIP>
	//</server>
	//}


}
//Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。  用于定义抽象数据类型
type Recurlyservers struct {
    // 为啥后面有一个 `xml:"servers"`这样的内容呢，这是个struct的特性，他们被称为 struct tag ，他们是用来辅助反射的，
	XMLName xml.Name `xml:"servers"`
	Version  string  `xml:"version,attr"`
	Svs     []server  `xml:"server"`
	Description string `xml:",innerxml"`

}
//  struct 的转换，因为struct和XML 都有类似结构的特征
type server struct {
    XMLName xml.Name `xml:"server"`  //XMLName 如果写错了也解析不出来啊
    ServerName string `xml:"serverName"`
    ServerIP string `xml:"serverIP"`  //如果写错了 就解析不了，文本里面的内容
}

func Demo()  {
	//data 为XML数据流，第二个表示储存的对应类型，目前支持struct slice和 string
	//XML包内部采用了反射来进行数据的映射，所以v里面的字段必须是导出的。Unmarshal解析的时候XML元素和字段怎么对应起来的呢？这是有一个优先级读取流程的，首先会读取struct tag，如果没有，那么就会对应字段名。必须注意一点的是解析的时候tag、字段名、XML元素都是大小写敏感的，所以必须一一对应字段。
	//err=xml.Unmarshal(data,&v)


	//Go 语言反射的机制，可以利用这些tag信息将来自XML文件中的数据反射成对应的struct对象，
	// 遵守的规则如下
	// 1 如果struct的一个字段是string或者是[]byte 类型且tag包含有 “，innerxml”，Unmarshal将会将此字段对应的元素内所有内嵌的原始xml累加到此字段上，如上面的Description 定义
	/*<server>
	<serverName>Shanghai_VPN</serverName>
	<serverIP>127.0.0.1</serverIP>
	</server>
	<server>
	<serverName>Beijing_VPN</serverName>
	<serverIP>127.0.0.2</serverIP>
	</server>
   */
   // 2 如果stuct中有一个叫做XMLName ，且类型为 xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段，如上面中的servers

   // 3 如果某个struct 字段的tag定义中包含XML结构中的 element的名称，那么解析的时候就会把相应的element值赋值给该字段，如servername和serverIP

   // 4 如果某个struct字段的tag定义了中含有“，attr”,那么解析的时候就会将该结构所对应的element的于字段同名的属性的值赋值给该字段，如version

	// 5 如果某个struct字段的tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。
	// 6 如果某个struct字段的tag定义了"-",那么不会为该字段解析匹配任何xml数据。
	// 7 如果struct字段后面的tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
	// 8 如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，这个字段的类型可能是[]byte或string,如果没有这样的字段存在，那么注释将会被抛弃。


	// todo  为了正确的解析，go语言中的xml包要求struct定义中所有的字段必须是可导出的 （即首字母大写 ）
}
