package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64 //创建一个新的类型

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

// 创建一个新的类型

// type ErrNegativeSqrt float64

// 并为其实现

// func (e ErrNegativeSqrt) Error() string

// 方法使其拥有 error 值，通过 ErrNegativeSqrt(-2).Error() 调用该方法应返回
// "cannot Sqrt negative number: -2"。

// 注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。
// 可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。
// 这是为什么呢？
// 第一次执行 ErrNegativeSqrt 的 Error()函数，是由 main函数中的 fmt.Println(Sqrt(-2)) 触发的。
// Sqrt(-2) 返回 error 为 ErrNegativeSqrt，在main中打印这个 error，就会执行 ErrNegativeSqrt 的 Error()函数。
// 然后在 Error()函数内部，执行 fmt.Sprintf("cannot Sqrt negative number: %v", e)，
// e为ErrNegativeSqrt ，所以，再一次打印 ErrNegativeSqrt，所以，还会调用 ErrNegativeSqrt 的 Error()函数。
// 这样就出现了死循环，直到调用栈溢出报错。
