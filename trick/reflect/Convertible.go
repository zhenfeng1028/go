package main

import (
	"fmt"
	"reflect"
)

func main() {
	BasicType()
	StrAndByteSlice()
	NotCompatible()
	CustomType()
}

func BasicType() {
	// 定义两个类型
	intType := reflect.TypeOf(0)
	floatType := reflect.TypeOf(0.0)

	// 检查 int 是否可以转换为 float64
	fmt.Println("int convertible to float64?", intType.ConvertibleTo(floatType)) // 输出: true

	// 检查 float64 是否可以转换为 int
	fmt.Println("float64 convertible to int?", floatType.ConvertibleTo(intType)) // 输出: true
}

func StrAndByteSlice() {
	// 定义两个类型
	stringType := reflect.TypeOf("")
	bytesType := reflect.TypeOf([]byte{})

	// 检查 string 是否可以转换为 []byte
	fmt.Println("string convertible to []byte?", stringType.ConvertibleTo(bytesType)) // 输出: true

	// 检查 []byte 是否可以转换为 string
	fmt.Println("[]byte convertible to string?", bytesType.ConvertibleTo(stringType)) // 输出: true
}

func NotCompatible() {
	// 定义两个类型
	intType := reflect.TypeOf(0)
	stringType := reflect.TypeOf("")

	// 检查 int 是否可以转换为 string
	fmt.Println("int convertible to string?", intType.ConvertibleTo(stringType)) // 输出: true

	// 检查 string 是否可以转换为 int
	fmt.Println("string convertible to int?", stringType.ConvertibleTo(intType)) // 输出: false
}

type (
	MyInt   int
	MyFloat float64
)

func CustomType() {
	// 定义两个自定义类型
	myIntType := reflect.TypeOf(MyInt(0))
	intType := reflect.TypeOf(0)

	// 检查 MyInt 是否可以转换为 int
	fmt.Println("MyInt convertible to int?", myIntType.ConvertibleTo(intType)) // 输出: true

	// 检查 int 是否可以转换为 MyInt
	fmt.Println("int convertible to MyInt?", intType.ConvertibleTo(myIntType)) // 输出: true

	// 定义另一个自定义类型
	myFloatType := reflect.TypeOf(MyFloat(0.0))

	// 检查 MyInt 是否可以转换为 MyFloat
	fmt.Println("MyInt convertible to MyFloat?", myIntType.ConvertibleTo(myFloatType)) // 输出: true

	// 检查 MyFloat 是否可以转换为 MyInt
	fmt.Println("MyFloat convertible to MyInt?", myFloatType.ConvertibleTo(myIntType)) // 输出: true
}
