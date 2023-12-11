package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

LOOP: // 循环
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
