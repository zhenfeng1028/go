package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32          // 计数器
	wg      sync.WaitGroup // 信号量
)

func main() {
	threadNum := 1000
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go incCounter(i)
	}
	wg.Wait()
}

func incCounter(index int) {
	defer wg.Done()

	spinNum := 0
	for {
		// 原子操作
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread-%d, spinnum: %d\n", index, spinNum)
}

// 这里之所以使用无限循环是因为在高并发下每个线程执行 CAS 并不是每次都成功，失败了的线程需要重新获取变量当前的值，然后重新执行 CAS 操作。
// 读者可以把线程数改为 10000 或者更多就会发现输出 thread-5329, spinnum: 1 其中这个 1 就说明该线程尝试了两个 CAS 操作，第二次才成功。
