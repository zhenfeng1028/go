// defer碰上闭包
package main

import "fmt"

func main() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// 由于闭包引用的局部变量的值一直保存在内存中，
// 第一次延迟调用的 i = 4，所以此后每次的延迟调用中的 i 都是 4

// 注意：Go 1.22 之后修复了该问题
