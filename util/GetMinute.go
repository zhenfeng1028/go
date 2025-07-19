package main

import (
	"fmt"
	"sort"
	"time"
)

const Minute_Format = "200601021504"

func main() {
	slice := []int{1, 15, 60, 120, 240}
	sort.Ints(slice)
	fmt.Println(slice)
	// time, _ := time.Parse(Minute_Format, "202205241136")
	time := time.Now()
	for _, step := range slice {
		minute := GetMinute(time, step)
		if minute == 0 {
			fmt.Printf("cannot get periodic time : %v\n", time)
			continue
		}
		fmt.Println(minute)
	}
}

func GetMinute(t time.Time, step int) int {
	if step <= 0 {
		return 0
	}
	if step < 60 {
		return t.Year()*100000000 + int(t.Month())*1000000 + t.Day()*10000 + t.Hour()*100 + (t.Minute()/step)*step
	}
	if step%60 == 0 {
		stepHour := step / 60
		return t.Year()*100000000 + int(t.Month()*1000000) + t.Day()*10000 + (t.Hour()/stepHour)*stepHour*100
	}
	return 0
}
