package main

//void SayHelloFourthDemo(char* s);
import "C"
import "fmt"

func init() {
	fmt.Println("面向C接口的Go编程")
}

func main() {
	/*
	在开始的例子中，我们的全部CGO代码都在一个Go文件中。然后，通过面向C接口编程的技术将SayHello分别拆分到不同的C文件，而main依然是Go文件。再然后，是用Go函数重新实现了C语言接口的SayHello函数。但是对于目前的例子来说只有一个函数，要拆分到三个不同的文件确实有些繁琐了。
正所谓合久必分、分久必合，我们现在尝试将例子中的几个文件重新合并到一个Go文件
	 */
	C.SayHelloFourthDemo(C.CString("我是在Go中的代码  \n"))
}
//export SayHelloFourthDemo
func SayHelloFourthDemo(s *C.char) {
	fmt.Println("我是在实现C中的代码 ---> 把所有的代码 都在一个文件中去实现     ")
	fmt.Print(C.GoString(s))
}