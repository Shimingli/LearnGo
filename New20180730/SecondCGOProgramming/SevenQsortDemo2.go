package main
//extern int go_qsort_compare(void* a, void* b);
import "C"
import (
	"fmt"
	"unsafe"
	"GoDemo/New20180730/qsort"
)

func init() {
	fmt.Println("将qsort函数从Go包导出")
}

func main() {
	fmt.Println("为了方便Go语言的非CGO用户使用qsort函数，我们需要将C语言的qsort函数包装为一个外部可以访问的Go函数。")
	// todo  具体的请看  qsort.go


	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.SortShiming(unsafe.Pointer(&values[0]),
		len(values), int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(C.go_qsort_compare),
	)
	fmt.Println(values)
    fmt.Println("目前已经实现了对C语言的qsort初步包装，并且可以通过包的方式被其它用户使用。但是qsort.Sort函数已经有很多不便使用之处：用户要提供C语言的比较函数，这对许多Go语言用户是一个挑战。下一步我们将继续改进qsort函数的包装函数，尝试通过闭包函数代替C语言的比较函数。")
	//目前已经实现了对C语言的qsort初步包装，并且可以通过包的方式被其它用户使用。但是qsort.Sort函数已经有很多不便使用之处：用户要提供C语言的比较函数，这对许多Go语言用户是一个挑战。下一步我们将继续改进qsort函数的包装函数，尝试通过闭包函数代替C语言的比较函数。
}
//go_qsort_compare是用Go语言实现的，并导出到C语言空间的函数，用于qsort排序时的比较函数
//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}