package main

import (
	"errors"
	"fmt"
)

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) {
		fmt.Printf("second defer err %v\n", err)
	}(err)
	defer func() {
		fmt.Printf("third defer err %v\n", err)
	}()

	if b == 0 {
		err = errors.New("divided by zero!")
		return
	} else {
		i = a / b
		return
	}
}

func main() {
	foo(2, 0)
}

// 参考 defer_5.go 第一个和第二个 defer 中的 err 已经用 nil 赋值
// 第三个 defer 中的 err 为闭包引用
