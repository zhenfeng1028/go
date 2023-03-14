// %b	表示为二进制
// %c	该值对应的unicode码值
// %d	表示为十进制
// %o	表示为八进制
// %x	表示为十六进制，使用a-f
// %X	表示为十六进制，使用A-F
// %U	表示为Unicode格式：U+1234，等价于”U+%04X”
// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

package main

import "fmt"

func main() {
	n := 65
	fmt.Printf("%b\n", n)	// 1000001
	fmt.Printf("%c\n", n)	// A
	fmt.Printf("%d\n", n)	// 65
	fmt.Printf("%o\n", n)	// 101
	fmt.Printf("%x\n", n)	// 41
	fmt.Printf("%X\n", n)	// 41
}
