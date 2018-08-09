package main

import (
	"fmt"
	"sort"
	"unsafe"
	"GoDemo/New20180730/qsort"
)

func init() {
	fmt.Println("改进：闭包函数作为比较函数")
}

func main() {
	fmt.Println("标准库的sort.Slice因为支持通过闭包函数指定比较函数")


	values := []int32{10,8,15,14,15,356,15,10}
	fmt.Println(values)
	sort.Slice(values,  func(i, j int) bool{
	    return values[i]<values[j]
	})
	fmt.Println("排序了以后 values=",values)

	//我们也尝试将C语言的qsort函数包装为以下格式的Go语言函数：
	newvalues := []int32{42, 9, 101, 95, 27, 25}
    /*
    调用的是 qsort3.go中的qsort.SortNewShiming()的方法
     */
	qsort.SortNewShiming(unsafe.Pointer(&newvalues[0]), len(newvalues), int(unsafe.Sizeof(newvalues[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)

	fmt.Println("排序了以后 newvalues=",newvalues)



}