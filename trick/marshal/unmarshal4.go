// 如果不想指定Class变量为具体的类型，仍想保留interface{}类型，
// 但又希望该变量可以解析到struct Class对象中，这时候该怎么办呢？
// 办法还是有的，我们可以将该变量定义为json.RawMessage类型
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Stu struct {
	Name  interface{}
	Age   interface{}
	Class json.RawMessage `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	data := "{\"name\":\"张三\",\"Age\":18,\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str := []byte(data)
	stu := Stu{}
	json.Unmarshal(str, &stu)

	// 注意这里，二次解析
	cla := new(Class)
	json.Unmarshal(stu.Class, cla)

	fmt.Println(stu)
	fmt.Println(string(stu.Class))
	fmt.Println(cla)
	printType(&stu)
}

func printType(stu *Stu) {
	nameType := reflect.TypeOf(stu.Name)
	ageType := reflect.TypeOf(stu.Age)
	classType := reflect.TypeOf(stu.Class)

	fmt.Println("nameType:", nameType)
	fmt.Println("ageType:", ageType)
	fmt.Println("classType:", classType)
}
