package main

import (
	"encoding/json"
	"fmt"
)

// 当存在匹配的json标签时，其对应的项被赋值。
// {<nil> <nil> 张三}
type Stu1 struct {
	NAme interface{}
	Name interface{}
	NAMe interface{} `json:"name"`
}

// 当匹配的json标签有多个时，标签对应的项都不会被赋值。
// 忽略标签项，从上往下寻找第一个没有标签且匹配的项赋值
// {张三 <nil> <nil> <nil>}
type Stu2 struct {
	NAme interface{}
	Name interface{}
	NAMe interface{} `json:"name"`
	NamE interface{} `json:"name"`
}

// 没有json标签时，从上往下，第一个匹配的项会被赋值
// {张三 <nil>}
type Stu3 struct {
	NAme interface{}
	Name interface{}
}

// 当相同的json标签有多个，且没有不带标签的匹配项时，都不赋值
// {<nil> <nil>}
type Stu4 struct {
	NAMe interface{} `json:"name"`
	NamE interface{} `json:"name"`
}

func main() {

	data := "{\"name\":\"张三\"}"
	str := []byte(data)

	stu1 := Stu1{}
	json.Unmarshal(str, &stu1)
	fmt.Println(stu1)

	stu2 := Stu2{}
	json.Unmarshal(str, &stu2)
	fmt.Println(stu2)

	stu3 := Stu3{}
	json.Unmarshal(str, &stu3)
	fmt.Println(stu3)

	stu4 := Stu4{}
	json.Unmarshal(str, &stu4)
	fmt.Println(stu4)
}
