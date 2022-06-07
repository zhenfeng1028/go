// struct{} 是一个无元素的结构体类型，通常在没有信息存储时使用。优点是大小为0，不需要内存来存储 struct{} 类型的值。
// struct{}{} 是一个复合字面量，它构造了一个struct{}类型的值，该值也是空。

// 比如我们可以用map[string]struct{}来当作成一个set来用

package main

import "fmt"

func main() {
	var set map[string]struct{}
	set = make(map[string]struct{})

	set["red"] = struct{}{} // struct{}{} 构造了一个struct{}类型的值
	set["blue"] = struct{}{}

	_, ok := set["red"]
	fmt.Println("Is red in the map?", ok)
	_, ok = set["green"]
	fmt.Println("Is green in the map?", ok)
}

// map可以通过“comma ok”机制来获取该key是否存在, _, ok := map["key"],
// 如果没有对应的值,ok为false, 这样可以通过定义成map[string]struct{}的形式,值不再占用内存。
