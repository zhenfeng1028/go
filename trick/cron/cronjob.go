package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	c.Start()
	time.Sleep(time.Second * 5)
}

/*
	使用非常简单，创建cron对象，这个对象用于管理定时任务。

	调用cron对象的AddFunc()方法向管理器中添加定时任务。
	AddFunc()接受两个参数，参数 1 以字符串形式指定触发时间规则，参数 2 是一个无参的函数，每次触发时调用。
	@every 1s表示每秒触发一次，@every后加一个时间间隔，表示每隔多长时间触发一次。
	例如@every 1h表示每小时触发一次，@every 1m2s表示每隔 1 分 2 秒触发一次。
	time.ParseDuration()支持的格式都可以用在这里。

	调用c.Start()启动定时循环。

	注意一点，因为c.Start()启动一个新的 goroutine 做循环检测，我们在代码最后加了一行time.Sleep(time.Second * 5)防止主 goroutine 退出。
*/
