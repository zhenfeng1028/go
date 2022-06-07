// Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。
// func Atoi(s string) (i int, err error)
// Itoa()函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下。
// func Itoa(i int) string

// a的典故
// C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi
	s1 := "001"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
	// Itoa
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}
