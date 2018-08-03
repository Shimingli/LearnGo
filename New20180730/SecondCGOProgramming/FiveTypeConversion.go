package main

import "fmt"

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

}