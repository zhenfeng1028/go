package main

import "unsafe"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)
}

// unsafe.Sizeof(a) 返回数据类型大小
// string 在 go 中不是直存类型，它是一个结构体类型
// type StringHeader struct {
//     Data uintptr // 8字节
//     Len  int     // 8字节
// }
