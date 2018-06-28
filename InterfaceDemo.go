package main

import (
	"fmt"
	"strconv"
	"sort"
	"reflect"
)

func init() {
	fmt.Println("Go语言中设计最精妙的应该是interface了，它让面向对象，内容组织实现非常的方便")
	fmt.Println("什么事interface  ：简单的说，interface是一组method签名的组合，我们通过interface来定义对象的一组行为。")
}
func main()  {
     s:=StudentF{HumanF{"shiming",22,"15018531365"},"beida",22}
	//dddd:=HumanF{"shiming",22,"15018531365"}

	fmt.Println(s)
	var m MenF
	//      todo    method has pointer receiver  如果使用这个
	m=&s //解决这个问题也很容易，直接使用&user去代替user调用方法即可：
	fmt.Println("我在这里做 ",m)



	/*
	 todo   通过下面的这个例子，interface 就是一组抽象方法的集合，它必须有其他非interface类型实现，额不能自我实现
	 Go 通过 interface 实现了 duck-typing ：当一只鸟走起来像鸭子，游泳起来像鸭子，叫起来也像鸭子，那么这只鸟就是鸭子
	 */
	//几个对象的初始化
	mike := StudentDemo{HumanDemo{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := StudentDemo{HumanDemo{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := EmployeeDemo{HumanDemo{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := EmployeeDemo{HumanDemo{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i MenDemo

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	//    todo  make用于内建类型（map、slice 和channel）的内存分配
	x := make([]MenDemo, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}





    //空 interface  一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值
    emptyInterface()




	//  interface  函数参数
	// interface 的变量可以持有任意实现该interface 类型的对象，这给我们编写函数（包括method）提供了一些额外的思考，我们是不是可以通过定义interface参数，让函数接受各种类型的参数
     interFaceDemo()


	//interface变量存储的类型    value, ok = element.(T)  switch value := element.(type)
    interfaceDemoTwo()


   //嵌入 interface Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，那么相同的逻辑引入到interface里面，那不是更加完美了。如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。
   interfaceInterface()

    //反射 Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。我们一般用到的包是reflect包。如何运用reflect包，官方的这篇文章详细的讲解了reflect包的实现原理  http://golang.org/doc/articles/laws_of_reflection.html   他妈的 要翻墙
	 reflectDemo()
}
func reflectDemo() {
	//要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)
	// 获取值 或者是获取反射的类型  嘿嘿
	t := reflect.TypeOf("dd")    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(10)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	t1:= reflect.TypeOf(10)
	v1 := reflect.ValueOf("shiming")
	fmt.Println(t)
	fmt.Println(t1)
	fmt.Println(v)
	fmt.Println(v1)


	//转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值
	//p:=person{"shiming",25}    todo   目前还不知道 怎么用
	//tag := 	t.Elem().Field(0).Tag  //获取定义在struct里面的标签
	//name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
	//
	//fmt.Println("tag==",tag)
	//fmt.Println("name==",name)
    fmt.Println("到这里来了么 ")
    //获取反射值能返回相应的类型和数值
	var x float64 = 3.4
	v11 := reflect.ValueOf(x)
	fmt.Println("type:", v11.Type())//type: float64
	fmt.Println("shiming   vll.kind()",v11.Kind()) //shiming   vll.kind() float64
	fmt.Println("kind is float64:", v11.Kind() == reflect.Float64)
	fmt.Println("value:", v11.Float())

	//反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，这个里面也是一样的道理。反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误
    //  todo  下面的写法  会抛出异常  panic: reflect: reflect.Value.SetFloat using unaddressable value
	//var xx float64 = 3.4
	//vv := reflect.ValueOf(xx)
	//fmt.Println(vv)
	//vv.SetFloat(7.1)
    //fmt.Println(vv)

	//如果要修改相应的值，必须这样写
    var a =1234
    fmt.Println(a)
    p:= reflect.ValueOf(&a) //传入的是地址值   <=======>
	vd:= p.Elem()
	vd.SetInt(1544)
	fmt.Println(vd)

	//反射的有点多  但是可以呢  更深入的理解还需要自己在编程中不断的实践。


}
//Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，那么相同的逻辑引入到interface里面，那不是更加完美了。如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。
func interfaceInterface() {

	//源码包container/heap里面有这样的一个定义container
    var d Interface
     fmt.Println("shiming===",d)//shiming=== <nil>

	//另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：
	//io.ReadWriter()
}
type Interface interface {
	//sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了  去看 sort.Interface 的基类的
	sort.Interface //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{} //a Pop elements that pops elements from the heap
}




//我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：
func interfaceDemoTwo() {
    //第一种   Comma-ok  断言
	//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
	//如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
	list := make(List, 3)
	fmt.Println("list的长度",len(list))
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}
	//本来的长度是  3  你他妈的还去赋值的话   ，就会出错啊啊啊啊啊
	//list[3] = Person{"Dennis", 70}

	//  同时你是否注意到了多个if里面，还记得我前面介绍流程时讲过，if里面允许初始化变量
	for index, element := range list {
		// value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
	//断言的类型越多，那么if else也就越多，所以才引出了下面要介绍的switch。

	listTwo := make(List, 3)
	listTwo[0] = 1 //an int
	listTwo[1] = "Hello" //a string
	listTwo[2] = Person{"Dennis", 70}
    fmt.Println("这里是 ")
	for index, element := range list{
		//`element.(type)`语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用`comma-ok`。 这个就有点意思了哦
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}


}

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}


func interFaceDemo() {
	//这个函数可以接受任意类型的值，同时打印出来 ，看基类怎么实现的  非常的有意思
	fmt.Println("sss")
	/*
	func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
	 */
    //  todo
    /*
 // Stringer由具有字符串方法的任何值实现，
//为该值定义“原生”格式。
//String方法用于打印作为操作数传递的值。
//接受任何字符串或未格式化打印机的格式
//打印等。
type Stringer interface {
	String() string
}

     */

	//也就是说，任何实现了String方法的类型都能作为参数被fmt.Println调用,让我们来试一试



   man := HumanT{"shiming",22,"15018531365"}
   fmt.Println(man.Strind())
   fmt.Println(man)
   //实现了error接口的对象（即实现了Error() string的对象），使用fmt输出时，会调用Error()方法，因此不必再定义String()方法了。


}

//现在我们再回顾一下前面的Box示例，你会发现Color结构也定义了一个method：String。其实这也是实现了fmt.Stringer这个interface，即如果需要某个类型能被fmt包以特殊的格式输出，你就必须实现Stringer这个接口。如果没有实现这个接口，fmt将以默认的方式输出。
//String()定义在Color上面，返回Color的具体颜色(字符串格式)
//  todo  为啥打印出来是  字符串的格式   原因 就是在这里
//func (c Color) String() string {
//	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
//	return strings[c]
//}



type HumanT struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h HumanT) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func (h HumanT) Strind() string  {
	return "我是世明"
}







/**
空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。
 */
func emptyInterface() {
	// 定义a为空接口
	var a interface{}
	fmt.Println("打印出来的空的类型",a)
	var flag  =a==nil
	fmt.Println("打印出来的空的类型",flag)
	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	//一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值
	a = s
	fmt.Println("打印出来的空的类型",a)
	a=true
	fmt.Println("打印出来的空的类型",a)
	fmt.Println("a  可以接受  任何的值，所以可以不断的变动，非常的有意思啊啊啊啊  ")
}



/*----------------------------------------------*/

type HumanDemo struct {
	name string
	age int
	phone string
}

type StudentDemo struct {
	HumanDemo //匿名字段
	school string
	loan float32
}

type EmployeeDemo struct {
	HumanDemo //匿名字段
	company string
	money float32
}

//Human实现SayHi方法
func (h HumanDemo) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func (h StudentDemo) SayHi()  {
	fmt.Println("我是学生，但是我有重写了这个方法，所以这个方法会走，我的name",h.name)
}
//Human实现Sing方法
func (h HumanDemo) Sing(lyrics string) {
	fmt.Println("都没有重写    La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e EmployeeDemo) SayHi() {
	fmt.Printf(" 我是 雇员，我重写了Sayhi的方法 Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type MenDemo interface {
	SayHi()
	Sing(lyrics string)
}







/*----------------------------------------------*/
type HumanF struct {
	name string
	age int
	phone string
}

type StudentF struct {
	HumanF //匿名字段Human
	school string
	loan float32
}

//type EmployeeF struct {
//	HumanF //匿名字段Human
//	company string
//	money float32
//}

//Human对象实现Sayhi方法
func (h *HumanF) SayHi1() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *HumanF) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *HumanF) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//// Employee重载Human的Sayhi方法
//func (e *EmployeeF) SayHi1() {
//	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
//		e.company, e.phone) //此句可以分成多行
//}

//Student实现BorrowMoney方法
func (s *StudentF) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

////Employee实现SpendSalary方法
//func (e *EmployeeF) SpendSalary(amount float32) {
//	e.money -= amount // More vodka please!!! Get me through the day!
//}

// 定义interface
type MenF interface {
	SayHi1()//上面的Men interface被Human、Student和Employee实现
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChapF interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGentF interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
//通过上面的代码我们可以知道，interface可以被任意的对象实现。我们看到上面的Men interface被Human、Student和Employee实现。同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。
//
//最后，任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。