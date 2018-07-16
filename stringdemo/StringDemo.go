package main

import (
	"fmt"
	"strings"
	"strconv"
)
// 是 main包下面的  才会进行高亮
func main() {
	fmt.Println("Go标准库中的strings 和 strconv两个包中的额函数精心有效的快速的的操作")

	//true
	//true
	//false
	//true
	//true 字符串s中是否包含substr，返回bool值
    fmt.Println(strings.Contains("shiming","shi"))
    fmt.Println(strings.Contains("shiming","s"))
    fmt.Println(strings.Contains("shiming","sm"))
    fmt.Println(strings.Contains("shiming",""))
    fmt.Println(strings.Contains("",""))


    s:=[]string{"s","m","today"}
    //字符串链接，把slice a通过sep链接起来 原来的字符数组 不做改变
	ss:=strings.Join(s,"s")
    fmt.Println(ss)
	fmt.Println(s)

     //在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	//4
	//0  开始的角标为0
	fmt.Println(strings.Index("chicken","ken"))
	fmt.Println(strings.Index("chicken","c"))
	fmt.Println(strings.Index("chicken","s"))

	//重复s字符串count次，最后返回重复的字符串
	//shiming**shiming**
	fmt.Println(strings.Repeat("shiming**",2))


	//在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	//lishiming  lishiming  shiming
	fmt.Println(strings.Replace("shiming  shiming  shiming ","shi","lishi",2))
	//shiming  shiming  shiming
 	fmt.Println(strings.Replace("shiming  shiming  shiming ","shi","lishi",0))

	fmt.Println(strings.Replace("shiming","shi","lishi",0))

    //[sh m ng]  把s字符串按照sep分割，返回slice
	fmt.Println(strings.Split("shiming","i"))
	//[s h i m i n g] 这个的输入有点意思哦
	fmt.Println(strings.Split("shiming",""))
   //[]
	fmt.Println(strings.Split("","ssssss"))

	//在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Trim("s*********s","s"))
	fmt.Println(strings.Trim("s*********","s"))
	fmt.Println(strings.Trim("s*********","*"))
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
    // 输出的结果如下
	//*********
	//*********
	//s
	//["Achtung"]


	fmt.Println()
	fmt.Println(strings.Fields("sss dfdfd fdfasdfadffsd fdfaf fa ddf sa sdf "))
    fmt.Println(" Fields are: %q",strings.Fields("s s s"))
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//[sss dfdfd fdfasdfadffsd fdfaf fa ddf sa sdf]
	//Fields are: %q [s s s]
	//Fields are: ["foo" "bar" "baz"]



	strconvDemo()
}
// 字符串的转换
func strconvDemo() {
	str:=make([]byte,0,100)
	strconv.

}
