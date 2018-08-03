package main
//#include <HelloThird.h>
import "C"
import "fmt"

func init() {
	//通过CGO的//export SayHello指令将Go语言实现的函数SayHello导出为C语言函数。为了适配CGO导出的C语言函数，我们禁止了在函数的声明语句中的const修饰符。需要注意的是，这里其实有两个版本的SayHello函数：一个Go语言环境的；另一个是C语言环境的。cgo生成的C语言版本SayHello函数最终会通过桥接代码调用Go语言版本的SayHello函数
	fmt.Println("用 Go 实现C函数")
}

func main() {
    //其实CGO不仅仅用于Go语言中调用C语言函数，还可以用于导出Go语言函数给C语言函数调用
	C.SayThirdHello(C.CString("Hello, World\n"))
}
//export SayThirdHello
func SayThirdHello(s *C.char)  {
	fmt.Println("不要管我  我肯定会执行 ")
   fmt.Println(C.GoString(s))
}