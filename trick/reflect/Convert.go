package main

import (
	"fmt"
	"reflect"
)

func main() {
	convertInt2Float64()
	convertStr2ByteSlice()
	convertIntf2Int()
	convertNotCompatible()
}

func convertInt2Float64() {
	// 定义一个 int 类型的值
	var num int = 42
	value := reflect.ValueOf(num)

	// 将 int 转换为 float64
	convertedValue := value.Convert(reflect.TypeOf(float64(0)))
	fmt.Println("Converted value:", convertedValue.Float()) // 输出: 42.0
}

func convertStr2ByteSlice() {
	// 定义一个字符串
	str := "Hello, Go!"
	value := reflect.ValueOf(str)

	// 将字符串转换为 []byte
	convertedValue := value.Convert(reflect.TypeOf([]byte{}))
	fmt.Println("Converted value:", convertedValue.Bytes()) // 输出: [72 101 108 108 111 44 32 71 111 33]
}

func convertIntf2Int() {
	// 定义一个接口类型的值
	var i interface{} = 100
	value := reflect.ValueOf(i)

	// 将接口类型转换为 int
	convertedValue := value.Convert(reflect.TypeOf(0))
	fmt.Println("Converted value:", convertedValue.Int()) // 输出: 100
}

func convertNotCompatible() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()

	// 定义一个字符串
	str := "Hello"
	value := reflect.ValueOf(str)

	// 尝试将字符串转换为 int（不兼容类型）
	convertedValue := value.Convert(reflect.TypeOf(0))
	fmt.Println("Converted value:", convertedValue.Int())
}
