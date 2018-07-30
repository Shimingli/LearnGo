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

    //Parse 系列函数把字符串转换为其他类型
	ParseDemo()
}
func ParseDemo() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}
func checkError(e error) {
	if e != nil{
		fmt.Println(e)
	}
}
// 字符串的转换
func strconvDemo() {

	// 卧槽  有点意思啊
	str:=make([]byte,0,100)
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	//追加整数的字符串形式，
	str=strconv.AppendInt(str,5,10)
	fmt.Println("*************** shiming ************")
	fmt.Println("111==========",string(str))
    str=strconv.AppendBool(str,false)
    fmt.Println(string(str))
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))


   //Format 系列函数把其他类型的转换为字符串
	a:= strconv.FormatBool(false)
	fmt.Println(a)
	/*
	bitSize 只能是 32  或者是  64   要不然会抛出恐慌   panic
	这个我感觉 很是尴尬啊
	 */
	b:=strconv.FormatFloat(1111111111111123.1211111111111111111111113,'g',12,64)
	fmt.Println(b)

    //4d2  16 进制的值   后面的是多少进制的意思   没有1进制  最少是2进制 ，注意会爆出恐慌
	c:=strconv.FormatInt(1234,16)
    fmt.Println(c)
    //for 2 <= base <= 36. 感觉和上面的是一样的  如果进制正确的话
	d:=strconv.FormatUint(1234,16)
    fmt.Println(d)

	//Parse 系列函数把字符串转化为其他的类型

	a1,_ := strconv.ParseBool("false")
	fmt.Println(a1)


}
