package main

import (
	"fmt"
	"encoding/xml"
	"os"
)

func main() {
	fmt.Println("输出 XML")

	v:=&Servers{Version:"2.0"}
	fmt.Println("v==",v)
	v.Svs=append(v.Svs,Server{"shimiing","168.0.0.1"})
	v.Svs=append(v.Svs,Server{"baobei","168.0.0.2"})
	fmt.Println("v==",v)
	//xml包中提供了Marshal和MarshalIndent两个函数，来满足我们的需求。这两个函数主要的区别是第二个函数会增加前缀和缩进
	output,err:=xml.MarshalIndent(v," ","  ")
	if err != nil {
		fmt.Println(err)
	}
	//os.Stdout.Write([]byte(xml.Header)) 这句代码的出现，是因为xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头的，为了生成正确的xml文件，我们使用了xml包预定义的Header变量。
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
	fmt.Println()
	fmt.Println("end     v=== ",v)

   //xml.Marshal(v)
	//如果v是 array或者slice，那么输出每一个元素，类似value
	//如果v是指针，那么会Marshal指针指向的内容，如果指针为空，什么都不输出
	//如果v是interface，那么就处理interface所包含的数据
	//如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息

	//生成的XML文件中的element的名字又是根据什么决定的呢？元素名按照如下优先级从struct中获取：

	//如果v是struct，XMLName的tag中定义的名称
	//类型为xml.Name的名叫XMLName的字段的值
	//通过struct中字段的tag来获取
	//通过struct的字段名用来获取
	//marshall的类型名称


	//我们应如何设置struct 中字段的tag信息以控制最终xml文件的生成呢？
	//
	//XMLName不会被输出
	//tag中含有"-"的字段不会输出
	//tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
	//tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。
	//tag中含有",chardata"，输出为xml的 character data而非element。
	//tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程
	//tag中含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
	//tag中含有"omitempty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string
	//tag中含有"a>b>c"，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出
}
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}
