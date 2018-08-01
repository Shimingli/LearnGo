package main

import (
	"fmt"
	"GoDemo/New20180730/FirstGoBase/firstpkg"
	"GoDemo/New20180730/FirstGoBase/secondpkg"
	"image/color"
	"sync"
	"io"
	"bytes"
	"os"
)

func init() {
	fmt.Println("函数、方法和接口")

}
func main() {
	// 注意包一定在 go的src目录下 才可以 导入进来
	firstpkg.FirstPgk()
	secondpkg.FirstPgk()
	//FisrtPkg init
	//secondpkg init
	//函数、方法和接口
	//FisrtPkg doing
	//secondpkg doing

	// 通过日记输出，可以看出，先执行所有的pkg的init的方法，然后执行mian.init ，然后执行main.main 最后执行方法导入到方法
	// todo  imag
	//  要注意的是，在main.main函数执行之前所有代码都运行在同一个goroutine，也就是程序的主系统线程中。因此，如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到


    funcDemo()

    defer fmt.Println("我会最后执行-----------")



	method()

	interfaceDemo()

}
/*
Go语言之父Rob Pike曾说过一句名言：那些试图避免白痴行为的语言最终自己变成了白痴语言（Languages that try to disallow idiocy become themselves idiotic）。一般静态编程语言都有着严格的类型系统，这使得编译器可以深入检查程序员没有作出什么出格的举动。但是，过于严格的类型系统却会使得编程太过繁琐，让程序员把大好的青春都浪费在了和编译器的斗争中。Go语言试图让程序员能在安全和灵活的编程之间取得一个平衡。它在提供严格的类型检查的同时，通过接口类型实现了对鸭子类型的支持，使得安全动态的编程变得相对容易。
 */
func interfaceDemo() {
    // 般静态编程语言都有着严格的类型系统。过于严格的编译系统，会导致编程的效率过低 ---  go在其中取得平衡
    // 鸭子类型：当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子



	//接口在Go语言中无处不在，在“Hello world”的例子中，fmt.Printf函数的设计就是完全基于接口的，它的真正功能由fmt.Fprintf函数完成。用于表示错误的error类型更是内置的接口类型。在C语言中，printf只能将几种有限的基础数据类型打印到文件对象中。但是Go语言灵活接口特性，fmt.Fprintf却可以向任何自定义的输出流对象打印，可以打印到文件或标准输出、也可以打印到网络、甚至可以打印到一个压缩文件；同时，打印的数据也不仅仅局限于语言内置的基础类型，任意隐式满足fmt.Stringer接口的对象都可以打印，不满足fmt.Stringer接口的依然可以通过反射的技术打印。fmt.Fprintf函数的签名如下

	 fmt.Printf("go go go")

	//func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
	//其中io.Writer用于输出的接口，error是内置的错误接口，它们的定义如下：
	//
	//type io.Writer interface {
	//Write(p []byte) (n int, err error)
	//}
	//
	//type error interface {
	//Error() string
	//}

	//   todo  定制自己的输出对象，将每个字符转化为大写的字符
    fmt.Println()
    fmt.Fprintln(&UpWriter{os.Stdout},"shi , ming")
    fmt.Println("[]byte{'s','m'}====",string(bytes.ToUpper([]byte{'s','m'})))
	//SHI , MING
	//[]byte{'s','m'}==== SM


	 //  我们也可以定义自己的打印格式来实现将每个字符转为大写字符后输出的效果。对于每个要打印的对象，如果满足了fmt.Stringer接口，则默认使用对象的String方法返回的结果打印




}

//type UpperString string
//
//type fmt.Stringer interface{
//	String() string
//}
//
//func (s UpperString) string() string {
//    return strings.ToUpper(string(s))
//}



type UpWriter struct {
	io.Writer
}
//  todo 如果把大小写写错了  比如说把 Write 写成 write的话！  是不会输入正确的-----！！！
func (p *UpWriter) Write(data []byte) (n int,err error){
	return p.Writer.Write(bytes.ToUpper(data))
}









/*
OOP: Object Oriented Programming,面向对象的程序设计。所谓“对象”在显式支持面向对象的语言中，一般是指类在内存中装载的实例，具有相关的成员变量和成员函数（也称为：方法）。面向对象的程序设计完全不同于传统的面向过程程序设计，它大大地降低了软件开发的难度，使编程就像搭积木一样简单，是当今电脑编程的一股势不可挡的潮流。 方法一般是面向对象编程(OOP)的一个特性
 */
func method() {
    // todo oop(面向对象的程序设计)  1、封装 2、继承 3、多态 4、抽象

	//Go语言中的做法是，将CloseFile和ReadFile函数的第一个参数移动到函数名的开头：
	//
	//// 关闭文件
	//func (f *File) CloseFile() error {
	//// ...
	//}
	//
	//// 读文件数据
	//func (f *File) ReadFile(int64 offset, data []byte) int {
	//// ...
	//}

	//CloseFile和ReadFile函数就成了File类型独有的方法了（而不是File对象方法）。它们也不再占用包级空间中的名字资源，同时File类型已经明确了它们操作对象，因此方法名字一般简化为Close和Read

	//// 关闭文件
	//func (f *File) Close() error {
	//// ...
	//}
	//
	//// 读文件数据
	//func (f *File) Read(int64 offset, data []byte) int {
	//// ...
	//}


	//Go语言不仅支持传统面向对象中的继承特性，而是以自己特有的组合方式支持了方法的继承。Go语言中，通过在结构体内置匿名的成员来实现继承


	type Point struct{ X, Y float64 }

	type ColoredPoint struct {
		Point
		Color color.RGBA
	}
   //虽然我们可以将ColoredPoint定义为一个有三个字段的扁平结构的结构体，但是我们这里将Point嵌入到ColoredPoint来提供X和Y这两个字段
	var cp ColoredPoint
	cp.X = 1
	fmt.Println("cp.Point.X=",cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y)       // "2"

	//通过嵌入匿名的成员，我们不仅可以继承匿名成员的内部成员，而且可以继承匿名成员类型所对应的方法。我们一般会将Point看作基类，把ColoredPoint看作是它的继承类或子类。不过这种方式继承的方法并不能实现C++中虚函数的多态特性。所有继承来的方法的接收者参数依然是那个匿名成员本身，而不是当前的变量


	//在传统的面向对象语言(eg.C++或Java)的继承中，子类的方法是在运行时动态绑定到对象的，因此基类实现的某些方法看到的this可能不是基类类型对应的对象，这个特性会导致基类方法运行的不确定性。而在Go语言通过嵌入匿名的成员来“继承”的基类方法，this就是实现该方法的类型的对象，Go语言中方法是编译时静态绑定的。如果需要虚函数的多态特性，我们需要借助Go语言接口来实现

}
//Cache结构体类型通过嵌入一个匿名的sync.Mutex来继承它的Lock和Unlock方法
type Cache struct {
	m map[string]string
	sync.Mutex
}
func (p *Cache) Lookup(key string) string {

	//但是在调用p.Lock()和p.Unlock()时, p并不是Lock和Unlock方法的真正接收者, 而是会将它们展开为p.Mutex.Lock()和p.Mutex.Unlock()调用. 这种展开是编译期完成的, 并没有运行时代价
	p.Lock()
	defer p.Unlock()
    //在传统的面向对象语言(eg.C++或Java)的继承中，子类的方法是在运行时动态绑定到对象的，因此基类实现的某些方法看到的this可能不是基类类型对应的对象，这个特性会导致基类方法运行的不确定性。而在Go语言通过嵌入匿名的成员来“继承”的基类方法，this就是实现该方法的类型的对象，Go语言中方法是编译时静态绑定的。如果需要虚函数的多态特性，我们需要借助Go语言接口来实现

	return p.m[key]
}
/*

 */
func funcDemo() {
	//Go语言中，函数是第一类对象，我们可以将函数保持到变量中。函数主要有具名和匿名之分，包级函数一般都是具名函数，具名函数是匿名函数的一种特例。当然，Go语言中每个类型还可以有自己的方法，方法其实也是函数的一种
	fmt.Println(Add(10,10))
	fmt.Println(AddTwo(100,10))

	//Go语言中的函数可以有多个输入参数和多个返回值，输入参数和返回值都是以传值的方式和被调用者交换数据。在语法上，函数还支持可变数量的参数，可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数

	fmt.Println(Swap(10,20))
	fmt.Println(Sum(10,10))
	fmt.Println(Sum(10,10,20))
	fmt.Println(Sum(10,-10,50))

	//可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果

	var a = []interface{}{123, "abc"}

	//rint调用时传入的参数是a...，等价于直接调用Print(123, "abc")
	Print(a...) // 123 abc
	//第二个Print调用传入的是为解包的a，等价于直接调用Print([]interface{}{123, "abc"})
	Print(a)    // [123 abc]

	//var m map[int]int
	var m =make(map[int]int)
	m[1]=1
	m[2]=2
	m[3]=3

	fmt.Println(Find(m,2))

	fmt.Println(Inc())
    //等循环执行完了，才会开始执行 defer里面的东西
	for i := 0; i < 3; i++ {
		fmt.Println("每次执行循环i=",i)
		defer func(){ fmt.Println(i) } ()
	}

	//1 修复的思路是在每轮迭代中为每个defer函数生成独有的变量
    //在循环体内部再定义一个局部变量，这样每次迭代defer语句的闭包函数捕获的都是不同的变量，这些变量的值对应迭代时的值
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func(){ fmt.Println("修复defer的独有的变量i=",i) } ()
	}
   // 是将迭代变量通过闭包函数的参数传入，defer语句会马上对调用参数求值
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		fmt.Println("i====",i)
		defer func(i int){ fmt.Println("语句会马上对调用参数求值 i=",i) } (i)
	}
	// todo  for循环内部执行defer语句并不是一个好的习惯


    //Go语言函数的递归调用深度逻辑上没有限制，函数调用的栈是不会出现溢出错误的，因为Go语言运行时会根据需要动态地调整函数栈的大小。每个goroutine刚启动时只会分配很小的栈（4或8KB，具体依赖实现），根据需要动态调整栈的大小，栈最大可以达到GB级（依赖具体实现）。在Go1.4以前，Go的动态栈采用的是分段式的动态栈，通俗地说就是采用一个链表来实现动态栈，每个链表的节点内存位置不会发生变化。但是链表实现的动态栈对某些导致跨越链表不同节点的热点调用的性能影响较大，因为相邻的链表节点它们在内存位置一般不是相邻的，这会增加CPU高速缓存命中失败的几率。为了解决热点调用的CPU缓存命中率问题，Go1.4之后改用连续的动态栈实现，也就是采用一个类似动态数组的结构来表示栈。不过连续动态栈也带来了新的问题：当连续栈动态增长时，需要将之前的数据移动到新的内存空间，这会导致之前栈中全部变量的地址发生变化。虽然Go语言运行时会自动更新引用了地址变化的栈变量的指针，但最重要的一点是要明白Go语言中指针不再是固定不变的了（因此不能随意将指针保持到数值变量中，Go语言的地址也不能随意保存到不在GC控制的环境中，因此使用CGO时不能在C语言中长期持有Go语言对象的地址）

    fmt.Println("Go语言函数的栈不会溢出，所以普通Go程序员已经很少需要关心栈的运行机制的。在Go语言规范中甚至故意没有讲到栈和堆的概念。我们无法知道函数参数或局部变量到底是保存在栈中还是堆中")

    fmt.Println( f(10))
    fmt.Println( g())




}
//函数直接返回了函数参数变量的地址——这似乎是不可以的，因为如果参数变量在栈上的话，函数返回之后栈变量就失效了，返回的地址自然也应该失效了。但是Go语言的编译器和运行时比我们聪明的多，它会保证指针指向的变量在合适的地方
func f(x int) *int {
	return &x
}
//内部虽然调用new函数创建了*int类型的指针对象，但是依然不知道它具体保存在哪里。对于有C/C++编程经验的程序员需要强调的是：不用关心Go语言中函数栈和堆的问题，编译器和运行时会帮我们搞定；同样不要假设变量在内存中的位置是固定不变的，指针随时可能会变化，特别是在你不期望它变化的时候
func g() int {
	 x := new(int)
	return *x
}
//通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值
func Inc() (v int) {
	//通过defer语句在return语句之后修改返回值
	//其中defer语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量v，这种函数我们一般叫闭包。闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问
	defer func(){
		v++
        v=v*10
    }()
	return 100
}

//不仅函数的输入参数可以有名字，也可以给函数的返回值命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}


func Print(a ...interface{}) {
	fmt.Println(a...)
}

// 多个输入参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}


// 具名函数
func Add(a, b int) int {
	return a+b
}

// 匿名函数
var AddTwo = func(a, b int) int {
	return a+b
}