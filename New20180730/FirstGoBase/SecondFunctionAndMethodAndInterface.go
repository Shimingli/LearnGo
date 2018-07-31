package main

import (
	"fmt"
	"GoDemo/New20180730/FirstGoBase/firstpkg"
	"GoDemo/New20180730/FirstGoBase/secondpkg"
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