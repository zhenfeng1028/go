package main

import (
	"fmt"
	"sync"
	"time"
)

// CPU缓存行：缓存是由缓存行组成的，通常是64字节（常用处理器的缓存行是64字节的），并且它有效地引用主内存中的一块地址

// 当多线程修改互相独立的变量时，如果这些变量共享同一个缓存行，就会无意中影响彼此的性能，这就是伪共享。

type Point struct {
	x int64
	y int64
}

func main() {
	now := time.Now()
	p := Point{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000000; i++ {
			p.x++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000000; i++ {
			p.y++
		}
	}()
	wg.Wait()
	fmt.Println(time.Since(now))
}

// x和y完全没有任何关系，但是更新x的时候会把其它包含x的缓存行失效，同时y也就失效了，运行这段程序的输出时间为236.561168ms

/*
	如何避免伪共享：
	方法一：
		在两个int64类型变量之间再加7个int64类型
		type Point struct {
			x                          int64
			p1, p2, p3, p4, p5, p6, p7 int64
			y                          int64
		}
		(使用该方法运行这段程序的输出时间为143.56029ms)
	方法二：
		自定义int64类型
		type MyInt64 struct {
			v                          int64
			p1, p2, p3, p4, p5, p6, p7 int64
		}
		type Point struct {
			x MyInt64
			y MyInt64
		}
		(使用该方法运行这段程序的输出时间为142.372801ms)
*/
