package main
/*
#include <hello.c>
/////void SayHello(const char* s);
 */
import "C"
import "fmt"

func init() {
   fmt.Println("C代码模块化")
}


func main() {
	//在编程过程中，抽象和模块化是将复杂问题简化的通用手段。当代码语句变多时，我们可以将相似的代码封装到一个个函数中；当程序中的函数变多时，我们将函数拆分到不同的文件或模块中。而模块化编程的核心是面向程序接口编程（这里的接口并不是Go语言的interface，而是API的概念）


	//`C` 代码的模块化，抽象和模块化是将复杂文件简化的通用手段。模块化编程的核心是面向程序接口编程

    C.SayHello(C.CString("我是在Go中输出的语句 ：SayHelloDemo2  \n"))


}
