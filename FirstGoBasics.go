package main

import "fmt"

func init() {

}
func main() {
	/*
	一共有 25个关键字
	 */
	 //break
	 //default
	 //func
//interface   用于定义接口，参考2.6小节



//	select {   用于选择不同类型的通讯
	//多个channel 的情况，Go提供了一个关键字 select ，通过select可以监听 channel 上的数据流动
		//go  routine  运行在相同的地址，因此访问共享内存必须做好同步，那么go routine 是如何进行数据的通讯呢，Go提供了一个很好的通信机制channel
		//Channel是Go中的一个核心类型,你可以把它看成一个管道, 可以通过 channel 发送和接受值，这些值只能是 channle的类型，定义一个channel时，也需要定义
		//发送到channel的值的类型，注意必须使用make创建channel
       //}
       //case
       // defer  Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题
       //如果有很多调用defer，那么defer是采用后进先出模式   用于类似析构函数




       //go   goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便   用于并发

       //map  map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型
	//使用map过程中需要注意的几点：
	//map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	//map的长度是不固定的，也就是和slice一样，也是一种引用类型
	//内置的len函数同样适用于map，返回map拥有的key的数量
	//map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	//map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	//   todo   这里不太明白   ---  》
	//map的初始化可以通过key:val的方式初始化值，同时map内置有判断是否存在key的方式

	//struct Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。  用于定义抽象数据类型

   //chan  用于channel通讯

   //else
   //goto  Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签

	i:=0
    Here://标记的意思 牛逼 标签名是大小写敏感的。
	fmt.Println("gotoDemo===",i)
	i++
	if i>100 { //这里为了方便一下  我就改下一点

	}else{
		goto Here
	}


   //package

	switch "ca" {
	case "s":
		fmt.Println("ss")
		break
	default:
		fmt.Println("default")
		break

	}
	list := make(ListDemo, 3)
	fmt.Println("list的长度",len(list))
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	//list[2] = Person{"Dennis", 70}
	for index, element := range list{
		//`element.(type)`语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用`comma-ok`。 这个就有点意思了哦
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		//case Person:
		//	fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}

	var aSlice, bSlice []int
    var a=append(aSlice, 12)
    var b=append(a, 150)
	//append(aSlice, 45)
	//append(aSlice, 25)
	//append(aSlice, 25,45,45)
	fmt.Println(b)
	fmt.Println(bSlice)
	fmt.Println(aSlice)


	//把很多值聚合在了一个case里面，同时，Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough  //是可以使用fallthrough强制执行后面的case代码。
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
	//sExpr和expr1、expr2、expr3的类型必须一致。Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true。

	iSwitch := 3
	switch iSwitch {
	case 1:  //下面三个的类型必须和iSwitch 是一样的
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}


	//fallthrough  //是可以使用fallthrough强制执行后面的case代码。

	const name =33
	//const定义常量的
	//if 判断

	//range	   用于读取slice、map、channel数据               for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了   用于读取slice、map、channel数据

	//type  定义数据类型 来用的  声明自定义类型

	//continue在循环里面有两个关键操作break和continue	,break操作是跳出当前循环，continue是跳过本次循环。当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：

	//for
	// import  package和import已经有过短暂的接触
	// return  用于从函数返回
	//  var   var和const参考2.2Go语言基础里面的变量和常量申明



}

type ListDemo []ElementDemo//申明一个元素ElementDemo元素的集合 由于这个歌属性是ElementDemo的组合，所以说，可以这样说，这个集合什么类型的都可以接受
type ElementDemo interface{}