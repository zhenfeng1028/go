// %v		值的默认格式表示
// %+v		类似%v，但输出结构体时会添加字段名
// %#v		值的Go语法表示
// %T		打印值的类型
// %%		百分号

package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}
