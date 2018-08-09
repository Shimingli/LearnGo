package main

import (
	"fmt"
	"GoDemo/New20180730/qsort"
)

func init() {
   // 改进：消除用户对unsafe包的依赖
   fmt.Println(" 改进：消除用户对unsafe包的依赖")

}

func main() {
  fmt.Println("qsort.Sort包装函数已经比最初的C语言版本的qsort易用很多，但是依然保留了很多C语言底层数据结构的细节。 现在我们将继续改进包装函数，尝试消除对unsafe包的依赖，并实现一个类似标准库中sort.Slice的排序函数")
  // 我在想go语言的底层
	//Go语言是谷歌推出的一种全新的编程语言，可以在不损失应用程序性能的情况下降低代码的复杂性
	//Go语言的特色
	//简洁、快速、安全
	//并行、有趣、开源，
	//内存管理、数组安全、编译迅速

	values := []int64{42, 9, 101, 95, 27, 25,100,1100,11,0,1,0,10,410,4,0,410,4,0,14,0,1,1,0,1,0,1,0,1,0,1,0,0,1,1,1}

	//自己通过调用 C 语言实现的一种排序的方式
	qsort.SliceFourth(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
