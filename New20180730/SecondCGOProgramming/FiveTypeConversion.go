package main
/*
#include <stdint.h>

union B1 {
	int i;
	float f;
};

union B2 {
	int8_t i8;
	int64_t i64;
};

union B {
	int i;
	float f;
};
struct A {
	int i;
	float f;
	int   type;  // type 是 Go 语言的关键字
	float _type; // 将屏蔽CGO对 type 成员的访问

	int   size: 10; // 位字段无法访问
	float arr[];    // 零长的数组也无法访问
};

enum C {
	ONE,
	TWO,
	THIRD,
	FOURTH,

};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func init() {
   fmt.Println("最初CGO是为了达到方便从Go语言函数调用C语言函数以复用C语言资源这一目的而出现的(因为C语言还会涉及回调函数，自然也会涉及到从C语言函数调用Go语言函数)。现在，它已经演变为C语言和Go语言双向通讯的桥梁。要想利用好CGO特性，自然需要了解此二语言类型之间的转换规则")
}

func main() {
	var  uni  uint8
	fmt.Println(uni)
    // 数值类型
	//   C语言类型               | CGO类型      | Go语言类型
	//   ---------------------- | ----------- | ---------
	//   char                   | C.char      | byte
	//   singed char            | C.schar     | int8
	//   unsigned char          | C.uchar     | uint8
	//   short                  | C.short     | int16
	//   unsigned short         | C.ushort     | uint16
	//   int                    | C.int       | int32
	//   unsigned int           | C.uint      | uint32
	//   long                   | C.long      | int32
	//   unsigned long          | C.ulong     | uint32
	//   long long int          | C.longlong  | int64
	//   unsigned long long int | C.ulonglong | uint64
	//   float                  | C.float     | float32
	//   double                 | C.double    | float64
	//   size_t                 | C.size_t    | uint

	//虽然在C语言中int、short等类型没有明确定义内存大小，但是在CGO中它们的内存大小是确定的。在CGO中，C语言的int和long类型都是对应4个字节的内存大小，size_t类型可以当作Go语言uint无符号整数类型对待

	//CGO中，虽然C语言的int固定为4字节的大小，但是Go语言自己的int和uint却在32位和64位系统下分别对应4个字节和8个字节大小。如果需要在C语言中访问Go语言的int类型，可以通过GoInt类型访问，GoInt类型在CGO工具生成的_cgo_export.h头文件中定义。其实在_cgo_export.h头文件中，每个基本的Go数值类型都定义了对应的C语言类型，它们一般都是以单词Go为前缀。下面是64位环境下，_cgo_export.h头文件生成的Go数值类型的定义，其中GoInt和GoUint类型分别对应GoInt64和GoUint64


	//除了GoInt和GoUint之外，我们并不推荐直接访问GoInt32、GoInt64等类型。更好的做法是通过C语言的C99标准引入的<stdint.h>头文件。为了提高C语言的可移植性，在<stdint.h>文件中，不但每个数值类型都提供了明确内存大小，而且和Go语言的类型命名更加一致。


   //Go 字符串和切片
	//在CGO生成的_cgo_export.h头文件中还会为Go语言的字符串、切片、字典、接口和管道等特有的数据类型生成对应的C语言类型

	//typedef struct { const char *p; GoInt n; } GoString;
	//typedef void *GoMap;
	//typedef void *GoChan;
	//typedef struct { void *t; void *v; } GoInterface;
	//typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;


	//不过需要注意的是，其中只有字符串和切片在CGO中有一定的使用价值，因为此二者可以在Go调用C语言函数时马上使用;而CGO并未针对其他的类型提供相关的辅助函数，且Go语言特有的内存模型导致我们无法保持这些由Go语言管理的内存指针，所以它们C语言环境并无使用的价值。
	//
	//在导出的C语言函数中我们可以直接使用Go字符串和切片。假设有以下两个导出函数：
	//
	////export helloString
	//func helloString(s string) {}
	//
	////export helloSlice
	//func helloSlice(s []byte) {}
	//CGO生成的_cgo_export.h头文件会包含以下的函数声明：
	//
	//extern void helloString(GoString p0);
	//extern void helloSlice(GoSlice p0);
	//不过需要注意的是，如果使用了GoString类型则会对_cgo_export.h头文件产生依赖，而这个头文件是动态输出的。
	//
	//Go1.10针对Go字符串增加了一个_GoString_预定义类型，可以降低在cgo代码中可能对_cgo_export.h头文件产生的循环依赖的风险。我们可以调整helloString函数的C语言声明为：
	//
	//extern void helloString(_GoString_ p0);
	//因为_GoString_是预定义类型，我们无法通过此类型直接访问字符串的长度和指针等信息。Go1.10同时也增加了以下两个函数用于获取字符串结构中的长度和指针信息：
	//
	//size_t _GoStringLen(_GoString_ s);
	//const char *_GoStringPtr(_GoString_ s);
	//更严谨的做法是为C语言函数接口定义严格的头文件，然后基于稳定的头文件实现代码。



	// 结构体、联合、枚举类型
	sructDemo()




}
func sructDemo() {

	/*
struct A {
	int i;
	float f;
};
*/
	//import "C"
	var a C.struct_A
	// 在这里我给他赋值上数字
	a.i=10
	fmt.Println("结构体、联合、枚举类型 :a.i",a.i)
	fmt.Println(a.f)
	a._type=10.2
	fmt.Println(a._type) // _type 对应 _type
	//在go中找不到这个type的类型--》
	//fmt.Println(a.type) // _type 对应 _type


	//fmt.Println(a.size) // 错误: 位字段无法访问 a.size undefined (type C.struct_A has no field or method size)
	//fmt.Println(a.arr)  // 错误: 零长的数组也无法访问 a.arr undefined (type C.struct_A has no field or method arr)


	//对于联合类型，我们可以通过C.union_xxx来访问C语言中定义的union xxx类型。但是Go语言中并不支持C语言联合类型，它们会被转为对应大小的字节数组
	var b1 C.union_B1
	fmt.Printf("%T\n", b1) // [4]uint8

	var b2 C.union_B2
	fmt.Printf("%T\n", b2) // [8]uint8


	//需要操作C语言的联合类型变量，一般有三种方法：第一种是在C语言中定义辅助函数；第二种是通过Go语言的"encoding/binary"手工解码成员(需要注意大端小端问题)；第三种是使用unsafe包强制转型为对应类型(这是性能最好的方式)。下面展示通过unsafe包访问联合类型成员的方式
	var b C.union_B;
	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))
	//虽然unsafe包访问最简单、性能也最好，但是对于有嵌套联合类型的情况处理会导致问题复杂化。对于复杂的联合类型，推荐通过在C语言中定义辅助函数的方式处理   在C语言中，枚举类型底层对应int类型，支持负数类型的值。我们可以通过C.ONE、C.TWO等直接访问定义的枚举值
	fmt.Println("枚举的开始----》》》 在C语言中，枚举类型底层对应int类型，支持负数类型的值。我们可以通过C.ONE、C.TWO等直接访问定义的枚举值")
	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
	fmt.Println(C.THIRD)
	fmt.Println(C.FOURTH)
	fmt.Println(C.FOURTH)
}