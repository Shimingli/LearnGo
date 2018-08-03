package main


/*
#include <stdio.h>
#include <shiming.c>
//方法的名称可以自己用  非常的nice
static void SayHelloShimingDemo(const char* s) {
	puts(s);
}
void SayHelloDemo2(const char* s);
 */
import "C"
import "fmt"

func init() {
	fmt.Println("你好 CGO  开始了新的Demo")
	//fmt.Println("GO编程：C语言作为一个通用语言，很多库会选择提供一个C兼容的API，然后用其他不同的编程语言实现。Go语言通过自带的一个叫CGO的工具来支持C语言函数调用，同时我们可以用Go语言导出C动态库接口给其它语言使用。开发区块链，就是通过Go生成so库给安卓使用，有点意思")
}
func main() {

    //第一个 CGO Demo
	firstDemo()

	//使用自己的C的函数
	secondDemo()


	thirdDemo()


}
//我们也可以将SayHelloDemo2函数放到当前目录下的一个C语言源文件中（后缀名必须是shiming.c）。因为是编写在独立的C文件中，为了允许外部引用，所以需要去掉函数的static修饰符
func thirdDemo() {
	C.SayHelloDemo2(C.CString("我是在Go中输出的语句 ：SayHelloDemo2  \n"))
}


func secondDemo() {
	C.SayHelloShimingDemo(C.CString("Hello, World----secondDemo  -\n"))
}
func firstDemo() {
	//println("hello cgo")
	//我们不仅仅通过import "C"语句启用CGO特性，同时包含C语言的<stdio.h>头文件。然后通过CGO包的C.CString函数将Go语言字符串转为C语言字符串，最后调用C语言的C.puts函数向标准输出窗口打印转换后的C字符串
	C.puts(C.CString("我真正的是一个CGO的程序哦！！！\n"))
	//没有释放使用C.CString创建的C语言字符串会导致内存泄露。但是对于这个小程序来说，这样是没有问题的，因为程序退出后操作系统会自动回收程序的所有资源

}
