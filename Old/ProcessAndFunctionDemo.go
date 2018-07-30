package Old

import (
	"fmt"
	"os"
)

//流程控制
//流程控制在编程语言中是最伟大的发明了，因为有了它，你可以通过很简单的流程描述来表达很复杂的逻辑。Go中流程控制分三大类：条件判断，循环控制和无条件跳转。
func main()  {

	fmt.Println("流程控制:流程控制在编程语言中是最伟大的发明了，因为有了它，你可以通过很简单的流程描述来表达很复杂的逻辑。Go中流程控制分三大类：条件判断，循环控制和无条件跳转。")


	//if
	//if也许是各种编程语言中最常见的了，它的语法概括起来就是：如果满足条件就做某事，否则做另一件事。
	ifDemo(10)


	//goto
	//Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签。例如假设这样一个循环：

	gotoDemo()



	//Go里面最强大的一个控制逻辑就是for，它既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作。
	forDemo()



	//有些时候你需要写很多的if-else来实现一些逻辑处理，这个时候代码看上去就很丑很冗长，而且也不易于以后的维护，这个时候switch就能很好的解决这个问题。它的语法如下
	switchDemo()


     //函数是Go里面的核心设计，它通过关键字func来声明
    //那么这三个是可以这样写的 哈哈 我是天才
     var _,_,_=funcDemo(10,50,45)
	 var _,_,a  = funcDemo(0,0,0)
	fmt.Println("a====",a)  //a==== 41

    var b=max(3,5)
	fmt.Println("b=====",b)



    //多个返回值的问题
    //go语言比C 更先进的特点，其中一点就是函数能够返回多个值
    aaa,bbb :=SumAndProduct(10,5)

    fmt.Println(aaa)//15
    fmt.Println(bbb)//50



    // 变参
	//Variable parameter
	//Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参：
    variableParameterDemo(10,454)



    //传值与传指针Transmission and pointer
	//当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。
	transmissionAndPointerDemo(1)
	x := 3
	fmt.Println("x = ", x)  // 应该输出 "x = 3"
	//    todo  x在里面加了一个  1    但是外面的值  是不会改变的  这一点和  java的局部的方法 差不多
	x1 := transmissionAndPointerDemo(x)  //调用add1(x)
	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出"x = 3"

	//理由很简单：因为当我们调用transmissionAndPointerDemo的时候，transmissionAndPointerDemo接收的参数其实是x的copy，而不是x本身。

	//如果真的需要传这个x本身,该怎么办呢？
    //    todo
	//这就牵扯到了所谓的指针。我们知道，变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存。只有transmissionAndPointerDemo函数知道x变量所在的地址，才能修改x变量的值。所以我们需要将x所在地址&x传入函数，并将函数的参数的类型由int改为*int，即改为指针类型，才能在函数中修改x变量的值。此时参数仍然是按copy传递的，只是copy的是一个指针。
	x22 := 3
	fmt.Println("x22 = ", x22)  // 应该输出 "x = 3"
	x222:=transmissionAndPointerDemoAnother(&x22) //  todo   调用 transmissionAndPointerDemoAnother(&x) 传x22的地址
    fmt.Println("x222==",x222)
    fmt.Println("x22==",x22)
	//输出的结果如下  这里是改变了 地址
	//x22 =  3
	//x222== 4
	//x22== 4

	//传入指正的好处
	//1 、传入指正使得多个函数能够操作同一个对象
	//2、传指针比较轻量级（8byte），只是传入内存地址，我们可以用指针传递体积大的结构体，如果用参数值传递的话，每次在copy上面就会花费较多的系统的开销，（内存和时间）。所以当你要传递大的结构体的时候，用指正是一个明智的选择
    //3、go语言中，channel slice ，，map 这三种类型机制类似指针，所以不用去地址后传递指针。但是如果函数需要改变slice的长度，则仍然需要去地址传递指正


	//Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题
	deferDemo()//   todo     defer   也得好好看看




	//函数作为值、类型
	//在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
	fmt.Println()
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)


	//Panic和Recover
	//Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制。一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有panic的东西。这是个强大的工具，请明智地使用它。那么，我们应该如何使用它呢？

	//Panic
	//是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。
	//
	//Recover
	//
	//是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
	throwsPanic(initDemo())



	//main函数和init函数
    mainAndInitDemo()


	importDemo()

}
func importDemo() {
	//我们在写Go代码的时候经常用到import这个命令用来导入包文件，而我们经常看到的方式参考如下：
	//import(
	//	"fmt"
	//)
	//上面这个fmt是Go语言的标准库，其实是去GOROOT环境变量指定目录下去加载该模块，当然Go的import还支持如下两种方式来加载自己写的模块：

	//相对路径
	//
	//import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
	//
	//绝对路径
	//
	//import “shorturl/model” //加载gopath/src/shorturl/model模块
	// 1------------->点操作
	//import(
	//. "fmt"
	//)

	//这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")


    //别名操作
	//别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
	//
	//import(
		//f "fmt"
	//)
	//别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")

	//_操作
	//
	//这个操作经常是让很多人费解的一个操作符，请看下面这个import
	//
	//import (
	//	"database/sql"
	//_ "github.com/ziutek/mymysql/godrv"
	//)
	//_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
}

func mainAndInitDemo() {
	/*
Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。
	 */


	//Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

	//程序的初始化和执行都起始于main包。如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。下图详细地解释了整个执行过程：
	//   todo   注意这张图片   2.3.init.png

}
func initDemo() func() {
	fmt.Println("这个方法执行了啊 initDemo")
	panic("no value for $USER")

}

var user = os.Getenv("USER")
//初始化方法  比  main方法  还早的运行
func init() {
	fmt.Println("这个方法会执行",user)
	//if user == "" {     todo  直接程序会停止掉
	//	panic("no value for $USER")
	//}
}
//下面这个函数检查作为其参数的函数在执行时是否会产生panic：   todo  不明白啊啊啊
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

//函数当做值和类型在我们写一些通用接口的时候非常有用，通过上面例子我们看到testInt这个类型是一个函数类型，然后两个filter函数的参数和返回值与testInt类型是一样的，但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活。

//   todo   把函数传递进去了    有点意思  嘿嘿
type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	//range  这里是用来读取数据的  来用的
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}


func deferDemo() {
	//原来是这样子去做的
	//file.Open("file")
	//// 做一些工作
	//if failureX {
	//	file.Close()
	//	return false
	//}
	//if failureY {
	//	file.Close()
	//	return false
	//}
	//file.Close()
	//return true

	//我们看到上面有很多重复的代码，Go的defer有效解决了这个问题。使用它后，不但代码量减少了很多，而且程序变得更优雅。在defer后指定的函数会在函数退出前调用。
	//file.Open("file")
	//defer file.Close()
	//if failureX {
	//	return false
	//}
	//if failureY {
	//	return false
	//}
	//return true

	//如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("--------->>>>>>%d ", i)  //执行的时候，有点晚啊啊啊啊啊
		if i==4 {
			fmt.Println()
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("------------>>>%d ", i)
	}
	fmt.Println("fdfd")
	fmt.Println("我会去先自行，然后后面的才会慢慢的去自行")
}

func transmissionAndPointerDemoAnother(i *int) int {
 	*i=*i+1// 修改了a的值  注意   这里的使用
 	return *i
}
//简单的一个函数，实现了参数+1的操作
func transmissionAndPointerDemo(i int) int {
   i=i+1// 我们改变了a的值
   return i//返回一个新值
}
//arg ...int告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice
func variableParameterDemo(arg...int) {
     for _,n:=range  arg{
     	fmt.Println("这个n是多少啊",n)
	 }
}

//看到直接返回了两个参数，当然我们也可以命名返回参数的变量，这个例子里面只是用了两个类型，我们也可以改成如下这样的定义，然后返回的时候不用带上变量名，因为直接在函数里面初始化了。但如果你的函数是导出的(首字母大写)，官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差

func SumAndProductAnother(i int, i2 int)(A int ,B int) {
	A=i+i2
	B=i*i2
	return  //这样官方的建议是  使用这种的模式去做 ，才会读起来 更加的牛逼
}
func SumAndProduct(i int, i2 int)(int ,int) {
	return i+i2,i*i2
}
func max(i int, i2 int) int {
	if i>i2 {
		return i
	}else {
		return i2
	}
}
//可以看到max函数有两个参数，它们的类型都是int，那么第一个变量的类型可以省略（即 a,b int,而非 a int, b int)，默认为离它最近的类型，同理多于2个同类型的变量或者返回值。同时我们注意到它的返回值就是一个类型，这个就是省略写法。
func maxAnother(i ,i2 int) int {
	if i>i2 {
		return i
	}else {
		return i2
	}
}

//关键字func用来声明一个函数funcName
//函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
//函数可以返回多个值
//上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
//如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
//如果没有返回值，那么就直接省略最后的返回信息
//如果有返回值， 那么必须在函数的外层添加return语句
func funcDemo(i int, i2 int, i3 int) (int, int, int) {
	//这里是处理逻辑代码
	//返回多个值
	return 10,50,41
}

func switchDemo() {
	//switch sExpr {
	//case expr1:
	//	some instructions
	//case expr2:
	//	some other instructions
	//case expr3:
	//	some other instructions
	//default:
	//	other code
	//}

	//sExpr和expr1、expr2、expr3的类型必须一致。Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true。

	i := 3
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

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

	//可以使用fallthrough强制执行后面的case代码。
	//The integer was <= 6
	//The integer was <= 7
	//The integer was <= 8
	//default case


}
func forDemo() {
	//  它的语法如下：
	//for expression1; expression2; expression3 {
	//	//...
	//}

	//expression1、expression2和expression3都是表达式，其中expression1和expression3是变量声明或者函数调用返回值之类的，expression2是用来条件判断，expression1在循环开始之前调用，expression3在每轮循环结束之时调用。
	sum := 0
	for index:=0; index < 10 ; index++ {
		sum += index
	}
	fmt.Println("最后的结果是 ", sum)


	//有些时候需要进行多个赋值操作，由于Go里面没有,操作符，那么可以使用平行赋值i, j = i+1, j-1

	//有些时候如果我们忽略expression1和expression3：
	sum1 := 1
	for ; sum1 < 10;  {  //另外一种实现的方式 45   和上面的结果是一样的啦  感觉是
		fmt.Println("另外一种实现的方式  每次改变的值是",sum1)
	    sum1 += sum1
	}
	fmt.Println("另外一种实现的方式",sum1)

	//其中 ； 也是可以省略的 ，变成了下面的代码  ，也就是  while的功能
	sum2 := 1
	for sum2 < 10 {
		fmt.Println("每次改变的值是",sum2)
		sum2 += sum2
	}
	fmt.Println("最后得到的值是===",sum2)


	//在循环里面有两个关键操作break和continue	,break操作是跳出当前循环，continue是跳过本次循环。当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：

	for index := 10; index>0; index-- {
		if index == 5{
			break // 或者continue
		}
		fmt.Println("break输出的index====》" ,index)
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1
	for index := 10; index>0; index-- {
		if index == 5{
			continue // 或者continue 跳过本次的循环
		}
		fmt.Println("continue输出的index====》" ,index)
	}

	//reak和continue还可以跟着标号，用来跳到多重循环中的外层循环

	//for配合range可以用于读取slice和map的数据：
     slice :=[]int{1,54,47,71,5,4,857,1,5,415,45,45,54}
	for k,v:=range slice {
		fmt.Println("slice's key:",k)
		fmt.Println("slice's val:",v)
	}
	//
	fmt.Println("开始打印了   打印 map的分隔线了")
	mapArr :=map[string]int{"s":1,"h":5}
	for k1,v1 :=range mapArr { //再次 验证了 map是无序的  牛逼啊
		fmt.Println("maps key",k1)
		fmt.Println("maps  val:",v1)
	}
	//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值 例如
	for _, v2 := range mapArr{//_ 的意思就是这个 返回回来 我不想用 ，所以说  这里不想报错，可以用下划线代替
		fmt.Println("map's val:", v2)
	}
	//_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34
	_, b := 34, 35
	fmt.Println("shiming  我得使用了这个变量 才不会报错 ",b)

}
//标签名是大小写敏感的。这个  就比较牛逼了啊  啊哈哈哈哈哈哈哈哈 哈哈哈啊哈哈哈哈哈 哈哈哈哈哈
func gotoDemo() {
	i:=0
	Here://标记的意思 牛逼 标签名是大小写敏感的。
		fmt.Println("gotoDemo===",i)
	i++
	if i>100 { //这里为了方便一下  我就改下一点
          return
	}
	goto Here
}
func ifDemo(i int) {
	if i>11{
		fmt.Println("传入的值是大于 11")
	}else {
		fmt.Println("传入的值是小于 11")
	}

	//Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，如下所示
	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := computedValue(); x > 10 {//相当于局部的变量
		fmt.Println("x is greater than 10   x======",x)
	} else {
		fmt.Println("x is less than 10    x======",x)
	}
	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	//fmt.Println(x)

	//多个条件的时候

	if i==3 {
	}else if i<10 {
	}else {
	}

}
//定义了一个函数 ，返回了一个int
func computedValue() int {
	return 22
}


