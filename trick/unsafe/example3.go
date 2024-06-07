package main

import (
	"fmt"
	"unsafe"
)

// 指针运算
func main() {
	data := []byte("abcd")
	for i := 0; i < len(data); i++ {
		ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + uintptr(i)*unsafe.Sizeof(data[0]))
		fmt.Printf("%c,", *(*byte)(unsafe.Pointer(ptr)))
	}
	fmt.Printf("\n")
}
