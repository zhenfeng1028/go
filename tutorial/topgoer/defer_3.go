package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}
}

// Close方法定义成结构体指针接收者，因此全部输出 c closed
// 如果想输出 c b a 可以将指针接收者改成值接收者

// 注意：Go 1.22 之后修复了该问题
