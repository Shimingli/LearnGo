// +build   go  1.10
package main

//void SayHelloFourthDemoBetter( _GoString_ s);
import "C"
import "fmt"

func init() {
	fmt.Println("面向C接口的Go编程")
}

func main() {
	/*
	现在版本的CGO代码中C语言代码的比例已经很少了，但是我们依然可以进一步以Go语言的思维来提炼我们的CGO代码。通过分析可以发现SayHello函数的参数如果可以直接使用Go字符串是最直接的。在Go1.10中CGO新增加了一个_GoString_预定义的C语言类型，用来表示Go语言字符串
	 */
	C.SayHelloFourthDemoBetter("  在Go1.10中CGO新增加了一个_GoString_预定义的C语言类型，用来表示Go语言字符串  \n")
}

func SayHelloFourthDemoBetter(s string)  {
     fmt.Println("我是更好的实现的方式  我是去实现的C中的代码，同时我在Go中")
     fmt.Println(s)
}