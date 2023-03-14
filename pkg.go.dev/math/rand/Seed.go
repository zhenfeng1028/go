package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 若不使用rand.Seed，那么每次生成的随机数都是一样的
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
