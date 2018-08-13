package main

/*
#include <stdio.h>
#include <stdlib.h>
void* makeslice(size_t memsize){
   printf("我是 makeslice 方法中的输出的语句 \n");
   return malloc(memsize);
}
void printShiming(const char* s){
  printf("%s",s);
}
void printShimingTwo(const char* s){
  printf("%s",s);
}
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func init() {
	fmt.Println("Start   CGO是架接Go语言和C语言的桥梁，它使二者在二进制接口层面实现了互通")
}
func main() {
	fmt.Println("CGO 内存模型  C语言的内存在分配之后就是稳定的，但是Go语言因为函数栈的动态伸缩可能导致栈中内存地址的移动(这是Go和C内存模型的最大差异)")

	//C语言内存分配后是稳定的，但是Go语言因为函数栈的动态伸缩可能导致栈中内存地址的移动（这是Go和C内存模型的最大差异）

	//如果C语言持有的是移动之前的Go指针，那么以旧指针访问Go对象时会导致程序崩溃。

	//Go访问C内存
	goVisitC_demo()


	//C临时访问传入的Go内存
	CVisitGo_demo()

	// C长期持有Go指针对象
	CLongHoldGo()


   // 导出的C的函数不能返回Go内存
	//在Go语言中，Go是从一个固定的虚拟地址空间分配内存。而C语言分配的内存则不能使用Go语言保留的虚拟内存空间。在CGO环境，Go语言运行时默认会检查导出返回的内存是否是由Go语言分配的，如果是则会抛出运行时异常。
	demo()
}
func demo() {
    fmt.Println("导出的C的函数不能返回Go内存")

	//C.Main()


	//extern int* getGoPtr();
	//static void Main(){
	//	printf("我是 Main 方法中的输出的语句 \n");
	//	int* p = getGoPtr();
	//	*p = 42;
	//}

}
////export getGoPtr
//func getGoPtr() *C.int  {
//	fmt.Println("我开始执行了啊  getGoPtr () ")
//	return  new(C.int)
//}



/*
使用Go调用C函数，其实CGo中，C函数也可以回调Go函数。但是C语言函数调用Go语言函数的时候，C语言函数就成了程序的调用方，Go语言函数返回的Go对象内存的生命周期也就自然超出了Go语言运行时的管理。简言之，我们不能在C语言函数中直接使用Go语言对象的内存。
虽然Go语言禁止在C语言函数中长期持有Go指针对象，但是这种需求是切实存在的。如果需要在C语言中访问Go语言内存对象，我们可以将Go语言内存对象在Go语言空间映射为一个int类型的id，然后通过此id来间接访问和控制Go语言对象。

以下代码用于将Go对象映射为整数类型的ObjectId，用完之后需要手工调用free方法释放该对象ID：
 */
func CLongHoldGo() {
   //  todo   EightCLongHoldGoDemo.go   And  EightCLongHoldGoDemoTwo.go

}
/*
cgo之所以存在的一大因素是为了方便在Go语言中接纳吸收过去几十年来使用C/C++语言软件构建的大量的软件资源。C/C++很多库都是需要通过指针直接处理传入的内存数据的，因此cgo中也有很多需要将Go内存传入C语言函数的应用场景
 */
func CVisitGo_demo() {
	//假设一个极端场景：我们将一块位于某goroutinue的栈上的Go语言内存传入了C语言函数后，在此C语言函数执行期间，此goroutinue的栈因为空间不足的原因发生了扩展，也就是导致了原来的Go语言内存被移动到了新的位置。但是此时此刻C语言函数并不知道该Go语言内存已经移动了位置，仍然用之前的地址来操作该内存——这将将导致内存越界。以上是一个推论（真实情况有些差异），也就是说C访问传入的Go内存可能是不安全的！

  //可以通过 ：RPC远程过程调用的经验的用户可能会考虑通过完全传值的方式处理：借助C语言内存稳定的特性，在C语言空间先开辟同样大小的内存，然后将Go的内存填充到C的内存空间；返回的内存也是如此处理
  s:="shiming c visit go  \n"
  //将Go的字符串传入C语言时，先通过C.CString将Go语言字符串对应的内存数据复制到新创建的C语言内存空间上
  cs:= C.CString(s)
  //  todo 但是这个处理的思路是安全的，但是效率极其低下（因为要多次分配内存并逐个赋值元素），同时也是极其麻烦的
  fmt.Println("cs=",cs)
  defer C.free(unsafe.Pointer(cs))
  //shiming c visit go demo

  C.printShiming(cs)

 //为了简化并高效处理此种向C语言传入Go语言内存的问题，cgo针对该场景定义了专门的规则：在CGO调用的C语言函数返回前，cgo保证传入的Go语言内存在此期间不会发生移动，C语言函数可以大胆地使用Go语言的内存！
   fmt.Println(" new  Demo-----------------------------------------")
    s1:= 'L'
    C.printShimingTwo((*C.char)(unsafe.Pointer(&s1)))
	fmt.Println()


	// todo  任何完美的技术都有被滥用的时候，CGO的这种看似完美的规则也是存在隐患的。我们假设调用的C语言函数需要长时间运行，那么将会导致被他引用的Go语言内存在C语言返回前不能被移动，从而可能间接地导致这个Go内存栈对应的goroutine不能动态伸缩栈内存，也就是可能导致这个goroutine被阻塞。因此，在需要长时间运行的C语言函数（特别是在纯CPU运算之外，还可能因为需要等待其它的资源而需要不确定时间才能完成的函数），需要谨慎处理传入的Go语言内存

}


//Go语言实现的限制，我们无法在Go语言中创建大于2GB内存的切片。不过借助cgo技术，我们可以在C语言环境创建大于2GB的内存，然后转为Go语言的切片使用
func goVisitC_demo() {
	s := makeByteSlize(1<<30+1)
	//fmt.Println("创建s===",s)
	var l= len(s)
	//通过C创建的切片是多长 len= 1073741825
	fmt.Println("通过C创建的切片是多长 len=",l)
	s[l-1] = 1
	s[0]=10
	fmt.Println("最后一个数字是多少 s=",s[l-1])
	freeByteSlice(s)

	// 通过Go去创建一个len= 1073741825 ，就会报错
	slice := make([]string,1073741825)
	fmt.Println("slice 的长度=",len(slice))

	s1 := makeByteSlize(100000000)
	//fmt.Println("创建s===",s)
	var l1= len(s1)
	//通过C创建的切片是多长 len= 1073741825
	fmt.Println("通过C创建的切片是多长 len=",l1)




}
//通过makeByteSlize来创建大于4G内存大小的切片，从而绕过了Go语言实现的限制（需要代码验证
func makeByteSlize(n int)[]byte  {
	p:= C.makeslice(C.size_t(n))
	return ((*[1 << 31]byte)(p))[0:n:n]
}
//而freeByteSlice辅助函数则用于释放从C语言函数创建的切片
func freeByteSlice(p []byte) {
	//Pointer类型用于表示任意类型的指针
	//1) 任意类型的指针可以转换为一个Pointer类型值
	//2) 一个Pointer类型值可以转换为任意类型的指针
	//3) 一个uintptr类型值可以转换为一个Pointer类型值
	//4) 一个Pointer类型值可以转换为一个uintptr类型值
	fmt.Println("p[0]=",p[0])
	fmt.Println("&p[0]=",&p[0])
	fmt.Println("unsafe.Pointer(&p[0])=",unsafe.Pointer(&p[0]))
	//释放 unsafe.Pointer(&p[0]) 所指的内存区
	C.free(unsafe.Pointer(&p[0]))
}
