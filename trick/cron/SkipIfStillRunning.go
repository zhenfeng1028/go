package main

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/robfig/cron/v3"
)

type skipJob struct {
	count int32
}

func (s *skipJob) Run() {
	atomic.AddInt32(&s.count, 1)
	log.Printf("%d: hello world\n", s.count)
	if atomic.LoadInt32(&s.count) == 1 {
		time.Sleep(2 * time.Second)
	}
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&skipJob{}))
	c.Start()

	time.Sleep(10 * time.Second)
}
