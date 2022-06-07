package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", GreetingJob{"dj"})
	c.Start()

	time.Sleep(5 * time.Second)
}

// 使用自定义的结构可以让任务携带状态（Name字段）。
