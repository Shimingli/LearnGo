package qsort
/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare_fourth(void* a, void* b);
*/
import "C"
import (
	"unsafe"
	"sync"
	"reflect"
	"fmt"
)
//避免在排序过程中，排序数组的上下文信息go_qsort_compare_info_fourth被修改，我们进行了全局加锁。 因此目前版本的qsort.Slice函数是无法并发执行的
var go_qsort_compare_info_fourth struct {
	base     unsafe.Pointer
	elemnum  int
	elemsize int
	less     func(a, b int) bool
	sync.Mutex
}

//export _cgo_qsort_compare_fourth
func _cgo_qsort_compare_fourth(a, b unsafe.Pointer) C.int {
	var (
		// array memory is locked
		base     = uintptr(go_qsort_compare_info_fourth.base)
		elemsize = uintptr(go_qsort_compare_info_fourth.elemsize)
	)

	i := int((uintptr(a) - base) / elemsize)
	j := int((uintptr(b) - base) / elemsize)

	switch {
	case go_qsort_compare_info_fourth.less(i, j): // v[i] < v[j]
		return -1
	case go_qsort_compare_info_fourth.less(j, i): // v[i] > v[j]
		return +1
	default:
		return 0
	}
}

func SliceFourth(slice interface{}, less func(a, b int) bool) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if sv.Len() == 0 {
		return
	}

	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	defer func() {
		go_qsort_compare_info_fourth.base = nil
		go_qsort_compare_info_fourth.elemnum = 0
		go_qsort_compare_info_fourth.elemsize = 0
		go_qsort_compare_info_fourth.less = nil
	}()

	// baseMem = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	// baseMem maybe moved, so must saved after call C.fn
	go_qsort_compare_info_fourth.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	go_qsort_compare_info_fourth.elemnum = sv.Len()
	go_qsort_compare_info_fourth.elemsize = int(sv.Type().Elem().Size())
	go_qsort_compare_info_fourth.less = less

	C.qsort(
		go_qsort_compare_info_fourth.base,
		C.size_t(go_qsort_compare_info_fourth.elemnum),
		C.size_t(go_qsort_compare_info_fourth.elemsize),
		C.qsort_cmp_func_t(C._cgo_qsort_compare_fourth),
	)
}