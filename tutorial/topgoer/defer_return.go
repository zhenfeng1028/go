package main

import "fmt"

func foo() (i int) {

	i = 0
	defer func() {
		fmt.Println(i)
	}()

	return 2
}

func main() {
	foo()
}

// 执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。所以defer closure 输出结果为 2 而不是 1。
