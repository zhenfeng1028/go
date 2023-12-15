package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(cancelCtx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("context value:", cancelCtx.Value("key1"))
	for {
		select {
		case val := <-ch:
			fmt.Println("read from ch value:", val)
		case <-cancelCtx.Done():
			fmt.Println("worker is cancelled")
			return
		}
	}
}

func main() {
	rootCtx := context.Background()
	childCtx := context.WithValue(rootCtx, "key1", "value1")
	cancelCtx, cancelFunc := context.WithCancel(childCtx)

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go worker(cancelCtx, ch, wg)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	cancelFunc()
	wg.Wait()
	close(ch)
}
