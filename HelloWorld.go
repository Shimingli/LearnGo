//每一个可独立运行的Go程序，必定包含一个package main，在这个main包中必定包含一个入口函数main，而这个函数既没有参数，也没有返回值。
package  main

import ("fmt"
	"errors"
)

//除了main包之外，其它的包最后都会生成*.a文件（也就是包文件）并放置在$GOPATH/pkg/$GOOS_$GOARCH
//为了打印Hello, world...，我们调用了一个函数Printf，这个函数来自于fmt包，所以我们在第三行中导入了系统级别的fmt包：import "fmt"。
/*
main函数是没有任何的参数的，我们接下来就学习如何编写带参数的、返回0个或多个值的函数。
 */

 /*
 Go之所以会那么简洁，是因为它有一些默认的行为：
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
  */
func main() {
   //我们调用了fmt包里面定义的函数Printf。大家可以看到，这个函数是通过<pkgName>.<funcName>的方式调用的，这一点和Python十分相似。
	fmt.Printf("Hello, world   我爱你 小姐姐 ")
	fmt.Println("shiming")
	//看到我们输出的内容里面包含了很多非ASCII码字符。实际上，Go是天生支持UTF-8的，任何字符都可以直接输出，你甚至可以用UTF-8中的任何字符作为标识符。
    //Go使用package（和Python的模块类似）来组织代码。main.main()函数(这个函数位于主包）是每一个独立的可运行程序的入口点。Go使用UTF-8字符串和标识符(因为UTF-8的发明者也就是Go的发明者之一)，所以它天生支持多语言。
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
    name=12456

	//_ , dddd:=  34,35
	do()

    //常量
    demo1()//所谓常量，也就是在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。
    // 在Go程序中，常量可定义为数值、布尔值或字符串等类型。
    test()
    //数值类型
    /*
    整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称。
     */
    numDemo()
  //浮点数的类型有float32和float64两种（没有float类型），默认是float64。
    floatDemo()

    //我们在上一节中讲过，Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string。
    stringDemo()


	//错误类型 Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：

	err := errors.New("emit macho dwarf: elf header corrupted")
	newErr :=errors.New("ddd")
	newErrNull :=errors.New("")
	//判断的语句 可以不要括号啊啊啊啊
	if err != nil {
		fmt.Print(err)
		fmt.Println(nil)
		fmt.Println(newErr)
		fmt.Println(newErr != nil)
		fmt.Println(newErrNull != nil)
		fmt.Println("shiming newErr")

	}

	//下面这张图来源于Russ Cox Blog中一篇介绍Go数据结构的文章，大家可以看到这些基础类型底层都是分配了一块内存，然后存储了相应的值。
	// todo 图片在  image 中的  2.2.basic.png


	//一些技巧

	//在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明。

	//import "fmt"
	//import "os"
	//
	//const i = 100
	//const pi = 3.1415
	//const prefix = "Go_"
	//
	//var i int
	//var pi float32
	//var prefix string

	//可以分组写成如下形式：
	//
	//import(
	//	"fmt"
	//"os"
	//)

	const(
		i = 100
		pi = 3.1415
		prefix = "Go_"
	)

	var(
		i1 int
		pi1 float32
		prefix1 string
	)
	fmt.Println("shiming ee")
	fmt.Println( i1)
	fmt.Println( pi1)
	fmt.Println( prefix1)//其实呢 这个是有去打印的，但是  这个值为  ""
	fmt.Println("shiming ee")

	//iota枚举
    //Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1：
    iotaDemo()
	//注意大写开头相当于，共有的
	//TestMe() todo  还不知道 咋个做  哈哈
	//testMeme()

	arrayDemo()
}
//array就是数组，它的定义方式如下：
func arrayDemo() {
	//var arr [10]int//在[n]type中，n表示数组的长度，type表示存储元素的类型
    //对数组的操作和其它语言类似，都是通过[]来进行读取或赋值：
	var arr [10]int  // 声明了一个int类型的数组
	arr[0] = 42      // 数组下标是从0开始的
	arr[1] = 13      // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The first element is %d\n", arr[1])  // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0
   /*
   由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型了。
    */
	//数组可以使用另一种:=来声明

	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
    //直接就可以把数组全部都打印出来  牛逼
//[1 2 3]
//[1 2 3 0 0 0 0 0 0 0]
//[4 5 6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	//Go支持嵌套数组，即多维数组  第一个的元素为 这个数组有多长，第二个为这个数组每个元素有多长
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][5]int{{1, 2, 3, 4}, {5, 6, 7, 8,1}}

	fmt.Println(doubleArray)
	fmt.Println(easyArray)
	//如果后续使用到了 数组的这一项的功能的话，回来再来看看 有点惊喜哦
    fmt.Println(doubleArray[0][1])
	fmt.Println(len(doubleArray[1]))

}
func iotaDemo() {
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)

	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

	const (
		h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
	)

	const (
		a       = iota //a=0
		b       = "B"
		//ff       =  iota
		c       = iota             //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3
		g       = iota             //g = 4
	)
	//fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v,ff)
	//除非被显式设置为其它值或iota，每个const分组的第一个常量被默认设置为它的0值，第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是iota，则它也被设置为iota。
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}


//Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string。
var shiming string
var emptyString string=""
var anotherString string=`ddd`//这样也可以定义 字符串
func stringDemo() {
	fmt.Println()
	fmt.Println(anotherString)
	no, yes, maybe :="no","yes","maybe"//简短申明，同时申明多个变量
	fmt.Println(no+yes+maybe)
	one :="one"//一个声明
	fmt.Println(one)
	shiming="a"//常规的赋值的操作

	//在Go中字符串是不可变的
    //var shiming1 ="shiming"
    //shiming1[0]='c' //todo cannot assign to shiming1[0]
    //如果真的想改s的值

	s := "hello"
	fmt.Println("改变前的值是"+s)
	c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c)  // 再转换回 string 类型
	//注意这两个的输入是一样的哦  有点意思
	fmt.Printf("%s\n", s2)
	fmt.Printf("%v\n", s2)

	//Go中可以使用+操作符来连接两个字符串：
	s1 := "hello,"
	m := " world"
	a := s1 + m
	fmt.Printf("%s\n", a)

	//修改字符串也可写为
	sss:="shiming"
	fmt.Println("切片等于"+sss[:1])//保留前面
	fmt.Println("切片等于"+sss[1:])//字符串虽不能更改，但可进行切片操作 保留后面
	sss ="c"+sss[1:]
	fmt.Println("%s\n",sss)

	//如果要声明一个多行的字符串怎么办  ` 括起的字符串为Raw字符串，
	mddd := `hello
	       world`
	fmt.Println(mddd)

}

func floatDemo() {
	var  a  float32
	var b  float64
	fmt.Println(a)
	fmt.Println(b)
	//Go还支持复数。它的默认类型是complex128（64位实数+64位虚数）。如果需要小一些的，也有complex64(32位实数+32位虚数)。复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位。下面是一个使用复数的例子：
	var c complex64 = 5+5i
	var d  complex128=50+5454554554
	//output: (5+5i)
	fmt.Printf("Value is: %v", c)//Value is: (5+5i)
	fmt.Println()
	fmt.Printf("Value is: %v", d)//Value is: (5.454554604e+09+0i)



}

func numDemo() {
	var a  int8
	var  b  int32
	var cc  rune
	//尽管int的长度是32 bit, 但int 与 int32并不可以互用。
	//c:=a+b//需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。
	ccc:=b+cc//。其中rune是int32的别称
	fmt.Println(a)
	fmt.Println(ccc)
}


var isActive bool//全局变量的声明  在Go中，布尔值的类型为bool，值是true或false，默认为false。
var isV bool
var enabled,disable=true,false//忽略类型的声明
func test() {
	var availbale  bool
	availbale=true//变量 一定得使用 要不然要报错
	valid := false//简短的申明

    fmt.Println("没有打印" )
    fmt.Println(availbale )
	println(availbale)
	println(valid)
	fmt.Println("真的没有打印")
	fmt.Println(valid)
}
func demo1() {
	//需要的话，指明常量的类型
	//Go 常量和一般程序语言不同的是，可以指定相当多的小数位数(例如200位)， 若指定給float32自动缩短为32bit，指定给float64自动缩短为64bit
	const Pi float32 = 3.1415926
	const i  = 1000
	const MaxThread  =10
	const name  ="shiming"

}
/*

 */
func do()  {
	//不过它有一个限制，那就是它只能用在函数内部；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量。
	vname11111, vname22222, vname33333 := vname11, vname11, vname33
	//.\HelloWorld.go:25:58: vname11111 declared and not used  todo 没有使用的话 就会报错 牛逼的方法,假如没有使用的话
	fmt.Println("shiming 你他妈的要去使用这个变量，不使用的话，就会报错啊"+vname11111)
	fmt.Println(vname22222)
	fmt.Println(vname33333)
    //_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34
	_, b := 34, 35
	fmt.Println("shiming  我得使用了这个变量 才不会报错 ")
	fmt.Println(b)
    //Go对于已声明但未使用的变量会在编译阶段报错，比如下面的代码就会产生一个错误：声明了i但未使用。
	var i int//默认值为0
	fmt.Println(i)
}


//var关键字是Go最基本的定义变量方式，与C语言不同的是Go把变量类型放在变量名后面：
var name  int

//定义一个名称为“variableName”，类型为string的变量
var variableName string
//定义三个类型都是“type”的变量
var vname1, vname2, vname3 string

//初始化“variableName”的变量为“value”值，类型是“type”
var variableName1 string = "shiming"
var variableName11 = "shiming"

/*
	定义三个类型都是"type"的变量,并且分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
*/
var vname11, vname22, vname33 string= "1", "2", "3"

//你是不是觉得上面这样的定义有点繁琐？没关系，因为Go语言的设计者也发现了，有一种写法可以让它变得简单一点。我们可以直接忽略类型声明，那么上面的代码变成这样了：
/*
	定义三个变量，它们分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
	然后Go会根据其相应值的类型来帮你初始化它们
*/
var vname111, vname222, vname333= 1, 2, 3

//你觉得上面的还是有些繁琐？好吧，我也觉得。让我们继续简化：
/*
	定义三个变量，它们分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
	编译器会根据初始化的值自动推导出相应的类型
*/
 //name1dsafdfdsaf , vname22222, vname3333  :=  vname111, vname222, vname333

