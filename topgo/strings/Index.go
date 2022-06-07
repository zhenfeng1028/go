package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world hello world"
	str := "wo"

	// 返回子串str在字符串s中第一次出现的位置。
	// 如果找不到则返回-1；如果str为空，则返回0
	index := strings.Index(s, str)
	fmt.Println(index) // 6
}
