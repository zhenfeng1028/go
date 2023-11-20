package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Sex   string `json:"sex"`
	Local bool   `json:"local"`
	Class *Class `json:"class"`
}

type Class struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func main() {
	stu := Stu{
		Name:  "张三",
		Age:   18,
		Sex:   "男",
		Local: true,
	}

	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3

	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	fmt.Println(string(jsonStu))
}

// 备注：
// 结构体中字段首字母需要大写才能转成JSON，否则不转
// 如果变量打上了json标签，如Name旁边的 `json:"name"` ，那么转化成的json key就用该标签“name”，否则取变量名作为key
// bool类型也是可以直接转换为json的value值
// 指针变量，编码时自动转换为它所指向的值，如cla变量
// 最后，强调一句：json编码成字符串后就是纯粹的字符串了
