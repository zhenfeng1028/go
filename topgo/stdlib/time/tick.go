// 使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) // 每秒都会执行的任务
	}
}
