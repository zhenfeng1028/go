package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)
	var durationAsInt64 = int64(duration)
	fmt.Println(duration)
	fmt.Println(durationAsInt64)

	m1 := 202105191458
	m2 := 202105191448
	minute1, _ := time.Parse("200601021504", strconv.Itoa((m1)))
	fmt.Println(minute1)
	minute2, _ := time.Parse("200601021504", strconv.Itoa((m2)))
	fmt.Println(minute2)
	duration = minute1.Sub(minute2)
	duraMinute := int64(duration)/int64(time.Minute) + 1
	fmt.Println(duraMinute)

	sub := time.Duration(math.Abs(float64(minute1.Sub(minute2)))) / time.Minute // 除法得到的结果是duration的最小单位ns
	fmt.Println(sub)                                                            // 10ns
}
