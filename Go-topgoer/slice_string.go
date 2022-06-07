// 要修改字符串，需要先将其转换成[]rune或者byte[]，完成后再转换成string。
package main

import (
	"fmt"
)

// 英文字符串
func ConvertAlpha() {
	str := "Hello world"
	s := []byte(str) // 中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}

// 中文字符串
func ConvertChinese() {
	str := "你好，世界！hello world！"
	s := []rune(str)
	s[3] = '够'
	s[4] = '浪'
	s[12] = 'g'
	s = s[:14]
	str = string(s)
	fmt.Println(str)
}

func main() {
	ConvertAlpha()
	ConvertChinese()
}
