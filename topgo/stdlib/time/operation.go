// func (t Time) Add(d Duration) Time
// func (t Time) Sub(u Time) Duration
// func (t Time) Equal(u Time) bool
// func (t Time) Before(u Time) bool
// func (t Time) After(u Time) bool

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}
