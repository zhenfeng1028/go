package main

import (
	"fmt"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second)
	done = true
	fmt.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

// 1. done 即多个 Goroutine 阻塞等待的条件。
// 2. read() 调用 Wait() 等待通知，直到 done 为 true。
// 3. write() 发送数据，发送完成后，将 done 置为 true，调用 Broadcast() 通知所有等待的协程。
