// 上面的成员变量都是已知的类型，只能接收指定的类型，比如string类型的Name只能赋值string类型的数据。
// 但有时为了通用性，或使代码简洁，我们希望有一种类型可以接受各种类型的数据，并进行json编码。这就用到了interface{}类型。
// interface{}类型其实是个空接口，即没有方法的接口。go的每一种类型都实现了该接口。因此，任何其他类型的数据都可以赋值给interface{}类型。
package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name  interface{} `json:"name"`
	Age   interface{} `json:"age"`
	Sex   interface{} `json:"sex"`
	Local interface{} `json:"local"`
	Class interface{} `json:"class"`
}

type Class struct {
	Name  interface{} `json:"name"`
	Grade interface{} `json:"grade"`
}

func main() {
	// 实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name:  "张三",
		Age:   18,
		Sex:   "男",
		Local: true,
	}

	// 指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	
	stu.Class = cla

	// Marshal失败时err!=nil
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	// jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}
