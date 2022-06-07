// time.Time类型表示时间。我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
