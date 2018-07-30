package main

import (
	"fmt"
	"image"
	"io"
	"image/png"
	"image/jpeg"
)

func main() {

	//Go语言中数组、字符串和切片三者是密切相关的数据结构。 todo 这三种数据类型，在底层原始数据有着相同的内存结构，            在上层，因为语法的限制而有着不同的行为表现。首先，Go语言的数组是一种值类型，虽然数组的元素可以被修改，但是数组本身的赋值和函数传参都是以整体复制的方式处理的。Go语言字符串底层数据也是对应的字节数组，但是字符串的只读属性禁止了在程序中对底层字节数组的元素的修改。字符串赋值只是复制了数据地址和对应的长度，而不会导致底层数据的复制。切片的行为更为灵活，切片的结构和字符串结构类似，但是解除了只读限制。切片的底层数据虽然也是对应数据类型的数组，但是每个切片还有独立的长度和容量信息，切片赋值和函数传参数时也是将切片头信息部分按传值方式处理。因为切片头含有底层数据的指针，所以它的赋值也不会导致底层数据的复制。其实Go语言的赋值和函数传参规则很简单，除了闭包函数以引用的方式对外部变量访问之外，其它赋值和函数传参数都是以传值的方式处理。要理解数组、字符串和切片三种不同的处理方式的原因需要详细了解它们的底层数据结构
	fmt.Println("数组、字符串、切片")

	array()




}
func array() {
	var a [3]int                    // 定义一个长度为3的int类型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义一个长度为3的int类型数组, 元素为 1, 2, 3
	// todo  第三种方式是以索引的方式来初始化数组的元素，因此元素的初始化值出现顺序比较随意。这种初始化方式和map[int]Type类型的初始化语法类似。数组的长度以出现的最大的索引为准，没有明确初始化的元素依然用0值初始化
	var c = [...]int{2: 3, 1: 2,}    // 定义一个长度为3的int类型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义一个长度为6的int类型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Println(a,b,c,d)



	//Go语言中数组是值语义。一个数组变量即表示整个数组，它并不是隐式的指向第一个元素的指针（比如C语言的数组），而是一个完整的值。当一个数组变量被赋值或者被传递的时候，实际上会复制整个数组。如果数组较大的话，数组的赋值也会有较大的开销。为了避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组。

	var aa = [...]int{1, 2, 3} // a 是一个数组
	// 避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组
	var bb = &a                // b 是指向数组的指针
	fmt.Println(aa,"*****",bb)
	fmt.Println("新的打印的日志",aa[0], bb[1])   // 打印数组的前2个元素
	fmt.Println("bb===",bb[0], bb[1])   // 通过数组指针访问数组元素的方式和数组类似
	fmt.Println("开始了range的输入-------------》")
	for i, v := range bb {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	//内置函数len可以用于计算数组的长度，cap函数可以用于计算数组的容量。不过对于数组类型来说，len和cap函数返回的结果始终是一样的，都是对应数组类型的长度。
	//a= [0 0 0]
	//b= [1 2 3]
	//c= [0 2 3]
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	fmt.Println("c=",c)
	for i := range b {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}
	//用for range方式迭代的性能可能会更好一些，因为这种迭代可以保证不会出现数组越界的情形，每轮迭代对数组元素的访问时可以省去对下标越界的判断
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}



	var times [5][0]int
	//用for range方式迭代，还可以忽略迭代时的下标
	for range times {
		fmt.Println("hello")
	}

	for i,v := range times {
		fmt.Println("i=",i,"v=",v)
	}
	// times对应一个[5][0]int类型的数组，虽然第一维数组有长度，但是数组的元素[0]int大小是0，因此整个数组占用的内存大小依然是0。没有付出额外的内存代价，通过for range方式实现了times次快速迭代


	 //字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好", }
    fmt.Println(s1,s2,s3,"*************aaaaaa")


	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line1,line2,line3,"*************bbbbbb")
	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	fmt.Println(decoder1,decoder2,"decoder---->")


	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}
    fmt.Println(unknown1,unknown2,"unknown")
	// 管道数组
	var chanList = [2]chan int{}
	fmt.Println(chanList,"chanList")

	//我们还可以定义一个空的数组：
	var df [0]int       // 定义一个长度为0的数组
	var e = [0]int{}   // 定义一个长度为0的数组
	var f = [...]int{} // 定义一个长度为0的数组
	fmt.Println(df,e,f)

    //  todo 长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间，比如用于管道的同步操作

	c1 := make(chan [0]int)
	//v1, ok1 := <-c1  从c中接收数据，并赋值给v1,ok1为true 表示才有值  ，如果没有的话  ，就表示没有值，这个通道已经关闭了
	go func() {
		fmt.Println("chan 执行了哈---")
		c1 <- [0]int{}
	}()
	<-c1//channel通过操作符<-来接收和发送数据


	//  todo  在这里，我们并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。对于这种场景，我们用空数组来作为管道类型可以减少管道元素赋值时的开销。当然一般更倾向于用无类型的匿名结构体代替
    //  go  routine  运行在相同的地址，因此访问共享内存必须做好同步，那么go routine 是如何进行数据的通讯呢，Go提供了一个很好的通信机制channel
	//     channel是Go中的一个核心类型,你可以把它看成一个管道, 可以通过 channel 发送和接受值，这些值只能是 channle的类型，定义一个channel时，也需要定义
	//    发送到channel的值的类型，注意必须使用make创建channel
	c2 := make(chan struct{

	})
	go func() {
		fmt.Println("chan 执行了啊  ")
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2

     //可以用fmt.Printf函数提供的%T或%#v谓词语法来打印数组的类型和详细信息：
	fmt.Printf("b: %T\n", b)  // b: [3]int
	fmt.Printf("b: %#v\n", b) // b: [3]int{1, 2, 3}


}

