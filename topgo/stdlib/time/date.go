package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 获取当前时间
	var (
		start time.Time
		end   time.Time
	)
	start = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
	fmt.Println(start)
	end = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(time.Minute * -10)
	fmt.Println(end)
}
