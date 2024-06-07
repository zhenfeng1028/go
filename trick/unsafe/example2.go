package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 读写结构体内部成员
func main() {
	str1 := "hello world"
	hdr1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	fmt.Printf("str:%s, data addr:%d, len:%d\n", str1, hdr1.Data, hdr1.Len)

	str2 := "abc"
	hdr2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	hdr1.Data = hdr2.Data
	hdr1.Len = hdr2.Len
	fmt.Printf("str:%s, data addr:%d, len:%d\n", str1, hdr1.Data, hdr1.Len)
}
