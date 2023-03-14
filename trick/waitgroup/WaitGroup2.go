package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // 等待计数器-1
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) // 等待计数器+1
		go process(i, &wg)
	}
	wg.Wait() // 等待计数器减为零之前一直阻塞
	fmt.Println("All go routines finished executing")
}

// 上面激活了3个goroutine，每次激活goroutine之前，都先调用Add()方法增加一个需要等待的goroutine计数。
// 每个goroutine都运行process()函数，这个函数在执行完成时需要调用Done()方法来表示goroutine的结束。
// 激活3个goroutine后，main goroutine会执行到Wait()，由于每个激活的goroutine运行的process()都需要睡眠2秒，
// 所以main goroutine在Wait()这里会阻塞一段时间(大约2秒)，当所有goroutine都完成后，等待计数器减为0，
// Wait()将不再阻塞，于是main goroutine得以执行后面的Println()。

// 还有一点需要特别注意的是process()中使用指针类型的*sync.WaitGroup作为参数，这里不能使用值类型的sync.WaitGroup作为参数，
// 因为这意味着每个goroutine都拷贝一份wg，每个goroutine都使用自己的wg。这显然是不合理的，这3个goroutine应该共享一个wg，
// 才能知道这3个goroutine都完成了。实际上，如果使用值类型的参数，main goroutine将会永久阻塞而导致产生死锁。
