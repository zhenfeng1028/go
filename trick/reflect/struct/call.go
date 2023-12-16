package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello：", name)
}

func main() {
	u := User{1, "zs", 20}
	v := reflect.ValueOf(u)
	// 获取方法
	m := v.MethodByName("Hello")
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("lizhenfeng")}
	// 调用方法，需要传入方法的参数
	m.Call(args)
}
