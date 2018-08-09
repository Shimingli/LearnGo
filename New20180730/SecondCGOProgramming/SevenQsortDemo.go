package main

import "fmt"

func init() {
	 fmt.Println("qsort快速排序函数是C语言的高阶函数，支持用于自定义排序比较函数，可以对任意类型的数组进行排序。本节我们尝试基于C语言的qsort函数封装一个Go语言版本的qsort函数")
}

func main() {
    //1、认识qsort 函数

	//qsort快速排序函数有<stdlib.h>标准库提供

	//void qsort(
	    //其中base参数是要排序数组的首个元素的地址，num是数组中元素的个数，size是数组中每个元素的大小
	//	void* base, size_t num, size_t size,
	   //最关键是cmp比较函数，用于对数组中任意两个元素进行排序。cmp排序函数的两个指针参数分别是要比较的两个元素的地址，如果第一个参数对应元素大于第二个参数对应的元素将返回结果大于0，如果两个元素相等则返回0，如果第一个元素小于第二个元素则返回结果小于0。
	//	int (*cmp)(const void*, const void*)
	//)



}
