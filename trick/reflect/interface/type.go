package main

import (
	"fmt"
	"reflect"
)

// 反射获取interface类型信息
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t)
	// Kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("float64")
	case reflect.String:
		fmt.Println("string")
	}
}

func main() {
	var x float64 = 3.4
	reflect_type(x)
}
