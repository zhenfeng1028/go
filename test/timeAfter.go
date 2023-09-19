package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	go func() {
		// 开启pprof，监听请求
		if err := http.ListenAndServe(":6060", nil); err != nil { // 也可以写成 127.0.0.1:6060
			fmt.Printf("start pprof failed on %s,err%v \n", "6060", err)
		}
	}()

	test()

	fmt.Println("finished!!")
}

func test() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("before, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	var i int32
	ch := make(chan string)
	done := make(chan struct{}) // 设定的时间已到，通知结束循环，不要再往channel里面写数据

	go func() {
		for {
			select {
			default:
				atomic.AddInt32(&i, 1)
				ch <- fmt.Sprintf("%d", i)
			case <-done:
				fmt.Println("close channel")
				return
			}
		}
	}()

	go func() {
		time.Sleep(time.Second)
		done <- struct{}{}
	}()

	for {
		select {
		case res := <-ch:
			runtime.GC()
			runtime.ReadMemStats(&ms)
			fmt.Printf("receive: %s, now have %d goroutines, %d bytes allocated, %d heap object\n", res, runtime.NumGoroutine(), ms.Alloc, ms.HeapObjects)
		case <-time.After(2 * time.Second): // 计时器出发后，GC会回收这些Timer，并不会造成“孤儿内存”（正确的做法是在for循环外进行初始化，否则会造成内存泄露）
			runtime.GC()
			runtime.ReadMemStats(&ms)
			fmt.Printf("after, now have %d goroutines, %d bytes allocated, %d heap object\n", runtime.NumGoroutine(), ms.Alloc, ms.HeapObjects)
			return
		}
	}
}
