// 时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()            // 获取当前时间
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	// 使用time.Unix()函数可以将时间戳转为时间格式
	timeObj := time.Unix(timestamp1, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	hour := timeObj.Hour()     // 小时
	minute := timeObj.Minute() // 分钟
	second := timeObj.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
