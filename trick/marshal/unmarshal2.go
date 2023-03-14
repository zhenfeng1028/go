package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Stu struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	Local interface{}
	sex   interface{}
	Class Class `json:"class"`
	// Class *Class `json:"class"`
	Test interface{}
}

type Class struct {
	Name  string
	Grade int
}

func main() {

	data := "{\"name\":\"张三\",\"Age\":18,\"local\":true,\"sex\":\"男\",\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str := []byte(data)

	stu := Stu{}
	printType(&stu) // 打印json解析前变量类型
	err := json.Unmarshal(str, &stu)
	fmt.Println("--------------json解析后----------------")

	if err != nil {
		fmt.Println(err)
	}

	printType(&stu) // 打印json解析后变量类型
	fmt.Println(stu)
}

// 利用反射，打印变量类型
func printType(stu *Stu) {
	nameType := reflect.TypeOf(stu.Name)
	ageType := reflect.TypeOf(stu.Age)
	localType := reflect.TypeOf(stu.Local)
	sexType := reflect.TypeOf(stu.sex)
	classType := reflect.TypeOf(stu.Class)
	testType := reflect.TypeOf(stu.Test)

	fmt.Println("nameType:", nameType)
	fmt.Println("ageType:", ageType)
	fmt.Println("localType:", localType)
	fmt.Println("sexType:", sexType)
	fmt.Println("classType:", classType)
	fmt.Println("testType:", testType)
}

// 从结果中可以看出
// 一、interface{}类型变量在json解析前，打印出的类型都为nil，就是没有具体类型，这是空接口（interface{}类型）的特点。
// 二、json解析后，json串中value，只要是”简单数据”，都会按照默认的类型赋值，如”张三”被赋值成string类型到Name变量中，数字18对应float64，true对应bool类型。
// 	“简单数据”：是指不能再进行二次json解析的数据，如”name”:”张三”只能进行一次json解析。
// 	“复合数据”：类似”CLASS\”:{\”naME\”:\”1班\”,\”GradE\”:3}这样的数据，是可进行二次甚至多次json解析的，因为它的value也是个可被解析的独立json。即第一次解析key为CLASS的value，第二次解析value中的key为naME和GradE的value
// 三、对于”复合数据”，如果接收体中配的项被声明为interface{}类型，go都会默认解析成map[string]interface{}类型。如果我们想直接解析到struct Class对象中，可以将接受体对应的项定义为该struct类型。如下所示：
// 	//普通struct类型
// 	Class Class `json:"class"`
// 	//指针类型
// 	Class *Class `json:"class"`
// stu打印结果
// Class类型：{张三 18 true <nil> {1班 3} <nil>}
// *Class类型：{张三 18 true <nil> 0xc42008a0c0 <nil>}
