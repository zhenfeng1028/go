package main

import (
	"fmt"
	"reflect"
)

// 就像函数名所述，reflect.Indirect()函数是用来获取reflect.Value所指向的基础值的。当使用reflect.ValueOf()函数获取一个指针类型变量的reflect.Value时，
// 默认情况下返回的reflect.Value的Kind()是reflect.Ptr而不是原始类型。这时，我们可以使用reflect.Indirect()在反射中获取原始类型。

type MyStruct struct {
	ID int
}

func main() {
	s := &MyStruct{ID: 123}

	v := reflect.ValueOf(s) // v.Kind() is reflect.Ptr
	fmt.Println(v.Kind())   // ptr

	v = reflect.ValueOf(s).Elem() // get the actual value
	fmt.Println(v.Kind())         // struct

	v1 := reflect.ValueOf(s)
	fmt.Println("v1.Kind():", v1.Kind())
	fmt.Println("v1.Type():", v1.Type())
	fmt.Println("v1.CanSet():", v1.CanSet())
	fmt.Println("v1.Elem().CanSet():", v1.Elem().CanSet())

	v2 := reflect.Indirect(v1)
	fmt.Println("v2.Kind():", v2.Kind())
	fmt.Println("v2.Type():", v2.Type())
	fmt.Println("v2.CanSet():", v2.CanSet())
}
