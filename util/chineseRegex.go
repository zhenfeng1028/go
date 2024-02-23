package main

import (
	"fmt"
	"regexp"
)

// 中文正则匹配
func main() {
	m, _ := regexp.MatchString("^\\p{Han}+$", "你好")
	fmt.Println(m)

	m, _ = regexp.MatchString("[\u4e00-\u9fa5]", "你好")
	fmt.Println(m)
}
