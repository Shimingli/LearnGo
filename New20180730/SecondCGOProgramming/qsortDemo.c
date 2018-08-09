#include <stdio.h>
#include <stdlib.h>
#include<windows.h>  // 记得导入包

#define DIM(x) (sizeof(x)/sizeof((x)[0]))
//qsort快速排序函数有<stdlib.h>标准库提供

	//void qsort(
	    //其中base参数是要排序数组的首个元素的地址，num是数组中元素的个数，size是数组中每个元素的大小
	//	void* base, size_t num, size_t size,
	   //最关键是cmp比较函数，用于对数组中任意两个元素进行排序。cmp排序函数的两个指针参数分别是要比较的两个元素的地址，如果第一个参数对应元素大于第二个参数对应的元素将返回结果大于0，如果两个元素相等则返回0，如果第一个元素小于第二个元素则返回结果小于0。
	//	int (*cmp)(const void*, const void*)
	//)
static int cmp(const void* a, const void* b) {
	const int* pa = (int*)a;
	const int* pb = (int*)b;
	return *pa - *pb;
}

int main() {
	int values[] = { 42, 8, 109, 97, 23, 25 };
	int i;
//    int j;
	qsort(values, DIM(values), sizeof(values[0]), cmp);
//    for( j = 0;j<10000;j++){
     	for(i = 0; i < DIM(values); i++) {
	    printf("-----shiming------- \n"+i);
		printf ("%d ",values[i]);
     	}
//	1
//   输入的结果为 8 23 25 42 97 109
    //线程睡了 100s
    Sleep(100000);
	return 0;
}