package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	count   = 10000
	chanNum = 10
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func main() {
	ch := make(chan int, count)
	chans := make([]chan int, chanNum)

	for i := 0; i < chanNum; i++ {
		chans[i] = make(chan int)
	}

	go fanout(chans)
	fanin(chans, ch)

	fmt.Println(len(ch))
}

func fanin(chans []chan int, out chan int) {
	var wg sync.WaitGroup
	wg.Add(chanNum)
	for i := 0; i < chanNum; i++ {
		go func(ch chan int) {
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(chans[i])
	}
	wg.Wait()
}

func fanout(chans []chan int) {
	for i := 0; i < count; i++ {
		v := r.Int()
		chans[v%chanNum] <- v
	}
	for i := 0; i < chanNum; i++ {
		close(chans[i]) // 若不手动关闭channel，在40行会发生阻塞
	}
}
