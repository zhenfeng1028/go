package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	singleton *Singleton
	once      sync.Once
)

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleton = new(Singleton)
	})
	return singleton
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := GetSingletonObj()
			fmt.Printf("%p\n", obj)
		}()
	}
	wg.Wait()
}

/*----output----
Create Obj
0x116f130
0x116f130
0x116f130
0x116f130
0x116f130
***************/
