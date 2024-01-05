package main

import (
	"runtime"
	"sync/atomic"
	"time"
)

type Golimit struct {
	l    int32
	size int32
}

func New(n int) *Golimit {
	return &Golimit{
		l:    0,
		size: int32(n),
	}
}

func (g *Golimit) Run(f func()) {
	cnt := 0
	for {
		l := atomic.LoadInt32(&g.l)
		if l < g.size {
			atomic.AddInt32(&g.l, 1)
			break
		}
		cnt++
		if cnt >= 2000 && cnt < 5000 {
			runtime.Gosched()
		} else if cnt >= 5000 && cnt < 10000 {
			time.Sleep(time.Millisecond * 2)
		} else if cnt >= 10000 {
			time.Sleep(time.Millisecond * 4)
		}
	}
	go func() {
		defer func() {
			atomic.AddInt32(&g.l, -1)
		}()
		f()
	}()
}

func (g *Golimit) Stop() {
	g.l = 0
}
