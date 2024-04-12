package main

import (
	"fmt"
)

// 如果申请的内存大小为0，则直接返回一个固定地址zerobase

func main() {
	var (
		// 0内存对象
		a struct{}
		b [0]int

		// 100个0内存struct{}
		c [100]struct{}

		// 100个0内存struct{},make申请形式
		d = make([]struct{}, 100)
	)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c[50])   // 取任意元素
	fmt.Printf("%p\n", &(d[50])) // 取任意元素
}
