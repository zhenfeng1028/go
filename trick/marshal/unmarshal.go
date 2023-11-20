package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	Local interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}

func main() {
	// json字符中的"引号，需用\进行转义，否则编译出错
	// json字符串沿用上面的结果，但对key进行了大小的修改
	data := "{\"name\":\"张三\",\"Age\":18,\"local\":true,\"sex\":\"男\",\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str := []byte(data)

	// Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	// 第二个参数必须是指针，否则无法接收解析的数据，stu仍为空对象Stu{}
	stu := Stu{}
	err := json.Unmarshal(str, &stu)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stu)
}

// 总结
// 一、json字符串解析时，需要一个“接收体”接受解析后的数据，且Unmarshal时接收体必须传递指针。
// 否则解析虽不报错，但数据无法赋值到接受体中。如这里用的是Stu{}接收。

// 二、解析时，接收体可自行定义。json串中的key自动在接收体中寻找匹配的项进行赋值。匹配规则是：
// 1.先查找与key一样的json标签，找到则赋值给该标签对应的变量(如Name)。
// 2.没有json标签的，就从上往下依次查找变量名与key一样的变量，如Age。
// 或者变量名忽略大小写后与key一样的变量，如Class。
// 第一个匹配的就赋值，后面就算有匹配的也忽略。(前提是该变量必需是可导出的，即首字母大写)。

// 三、不可导出的变量无法被解析（如sex变量，虽然json串中有key为sex的k-v，解析后其值仍为nil,即空值）

// 四、当接收体中存在json串中匹配不了的项时，解析会自动忽略该项，该项仍保留原值。如变量Test，保留空值nil。

// 五、你一定会发现，变量Class貌似没有解析为我们期待样子。因为此时的Class是个interface{}类型的变量，
// 而json串中key为CLASS的value是个复合结构，不是可以直接解析的简单类型数据（如“张三”，18，true等）。
// 所以解析时，由于没有指定变量Class的具体类型，json自动将value为复合结构的数据解析为map[string]interface{}类型的项。
