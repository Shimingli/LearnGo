package qsort
////typedef int (*qsort_cmp_func_t)(const void* a,const void* b);
//import "C"
//import "unsafe"
////把这个放到
//func SortShiming(base unsafe.Pointer,num,size C.size_t,cmp C.qsort_func_t)  {
//	C.qsort(base, num, size, cmp)
//}
//
//虽然Sort函数已经导出了，但是对于qsort包之外的用户依然不能直接使用该函数——Sort函数的参数还包含了虚拟的C包提供的类型。 在CGO的内部机制一节中我们已经提过，虚拟的C包下的任何名称其实都会被映射为包内的私有名字。比如C.size_t会被展开为_Ctype_size_t，C.qsort_cmp_func_t类型会被展开为_Ctype_qsort_cmp_func_t。
//
//被CGO处理后的Sort函数的类型如下：
//
//func Sort(
//	base unsafe.Pointer, num, size _Ctype_size_t,
//	cmp _Ctype_qsort_cmp_func_t,
//)
//这样将会导致包外部用于无法构造_Ctype_size_t和_Ctype_qsort_cmp_func_t类型的参数而无法使用Sort函数。因此，导出的Sort函数的参数和返回值要避免对虚拟C包的依赖。
//
//重新调整Sort函数的参数类型和实现如下

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"
import "unsafe"

type CompareFunc C.qsort_cmp_func_t

func SortShiming(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}