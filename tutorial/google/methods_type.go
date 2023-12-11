package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// 你也可以为非结构体类型声明方法。

// 在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat。

// 接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。
