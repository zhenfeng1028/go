package main

import (
	"fmt"
	"unsafe"
)

// 指针类型转换
func main() {
	i := int64(1)
	var iPtr *int
	iPtr = (*int)(unsafe.Pointer(&i))
	fmt.Printf("%d\n", *iPtr)
}
