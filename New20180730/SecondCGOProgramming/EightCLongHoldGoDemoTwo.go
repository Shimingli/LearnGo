package main

/*
#include <stdlib.h>

void printString(const char* s) {
	char* gs = NewGoStringT(s);
	printf("我是 printString 方法中的输出的语句 \n");
	PrintGoStringT(gs);
	FreeGoStringT(gs);
}
*/
//char* NewGoStringT(char* );
//void FreeGoStringT(char* );
//void PrintGoStringT(char* );
import "C"
import "unsafe"

//export NewGoStringT
func NewGoStringT(s *C.char) *C.char {
	gs := C.GoString(s)
	id := NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoStringT
func FreeGoStringT(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoStringT
func PrintGoStringT(s *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(s)))
	gs := id.Get().(string)
	print(gs)
}

func main() {
	C.printString("hello")
}