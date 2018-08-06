package main



/*

#include <errno.h>//增加一个错误的结果 因为C语言不支持返回多个结果，因此<errno.h>标准库提供了一个errno宏用于返回错误状态。我们可以近似地将errno看成一个线程安全的全局变量，可以用于记录最近一次错误的状态码
static int divError(int a,int b){
     if(b==0){
     	errno = 10;
       return 0;
     }
     return a/b;
}

static int add(int a,int b){
   return a+b+10;
 }
static int div(int a,int b){
   return  a/b;
}
 */
//static void noreturn() {}
import "C"
import "fmt"

func init() {
	fmt.Println("函数是c语言编程的核心，通过CGO技术我们不仅仅可以在Go语言中调用C语言函数，也可以将Go语言函数导出为C语言函数")
}

func main() {
	fmt.Println("函数的调用")
	// Go调用C函数
	demoAdd()

	//C 函数有返回值，正常取返回值
	demoDiv()

	//因为C语言不支持返回多个结果，因此<errno.h>标准库提供了一个errno宏用于返回错误状态。我们可以近似地将errno看成一个线程安全的全局变量，可以用于记录最近一次错误的状态码
	demoDivError()


   //void函数的返回值
   voidDemo()


	//C调用Go导出函数
	CUserGoCode()


}

//CGO还有一个强大的特性：将Go函数导出为C语言函数。这样的话我们可以定义好C语言接口，然后通过Go语言实现
func CUserGoCode() {

}
// add函数名以小写字母开头，对于Go语言来说是包内的私有函数 但是从C语言角度来看，导出的add函数是一个可全局访问的C语言函数。如果在两个不同的Go语言包内，都存在一个同名的要导出为C语言函数的add函数，那么在最终的链接阶段将会出现符号重名的问题。
//export addTwo
//func addTwo(a, b C.int) C.int {
//	return a+b
//}




/*
C语言函数还有一种没有返回值类型的函数，用void表示返回值类型。一般情况下，我们无法获取void类型函数的返回值，因为没有返回值可以获取。前面的例子中提到，cgo对errno做了特殊处理，可以通过第二个返回值来获取C语言的错误状态。对于void类型函数，这个特性依然有效。
 */
func voidDemo() {
	v, _ := C.noreturn()
	fmt.Printf("%#v\n", v)
	//可以看出C语言的void类型对应的是当前的main包中的_Ctype_void类型。其实也将C语言的noreturn函数看作是返回_Ctype_void类型的函数，这样就可以直接获取void类型函数的返回值
	fmt.Println("可以看出C语言的void类型对应的是当前的main包中的_Ctype_void类型。其实也将C语言的noreturn函数看作是返回_Ctype_void类型的函数，这样就可以直接获取void类型函数的返回值")

	//todo 其实在CGO生成的代码中，_Ctype_void类型对应一个0长的数组类型[0]byte，因此fmt.Println输出的是一个表示空数值的方括弧。
	fmt.Println("C.noreturn()=",C.noreturn())


}
func demoDivError() {
	//CGO也针对<errno.h>标准库的errno宏做的特殊支持：在CGO调用C函数时如果有两个返回值，那么第二个返回值将对应errno错误状态。
	v,err:=C.divError(10,0)
	//  todo  #include <errno.h>  注意不要导错包了啊
    fmt.Println(v,"err=",err)
}
func demoDiv() {
	c:=C.div(10,2)
	fmt.Println("我觉得应该是5 c=",c)

}
func demoAdd() {
	c:=C.add(10,20)
	fmt.Println("我知道值肯定等于40，  c=",c)
}
