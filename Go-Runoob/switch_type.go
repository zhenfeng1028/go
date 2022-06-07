package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case func(int):
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
