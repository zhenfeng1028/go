// 设置当前程序并发时占用的CPU逻辑核心数
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 100; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 100; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // 两个任务只有一个逻辑核心
	// runtime.GOMAXPROCS(2) // 两个任务有两个逻辑核心
	go a()
	go b()
	time.Sleep(time.Second)
}
