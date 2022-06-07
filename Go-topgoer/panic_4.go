package main

import "fmt"

func test() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()
	defer recover()              // 无效！
	defer fmt.Println(recover()) // 无效！
	defer func() {
		func() {
			println("defer inner")
			recover() // 无效！
		}()
	}()

	panic("test panic")
}

func main() {
	test()
}

// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。
// 任何未捕获的错误都会沿调用堆栈向外传递。
