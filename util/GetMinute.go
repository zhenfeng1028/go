package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

const (
	Day_Format    = "20060102"
	Hour_Format   = "2006010215"
	Minute_Format = "200601021504"
	Second_Format = "20060102150405"
)

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
	hourLong, _ := strconv.Atoi(t.Format(Hour_Format))
	day, _ := strconv.Atoi(t.Format(Day_Format))
	hourShort := t.Hour()
	minute := t.Minute()
	if step < 60 {
		size := int(time.Hour / (time.Duration(step) * time.Minute))
		for i := 0; i < size; i++ {
			if minute < step*(i+1) {
				return hourLong*100 + step*i
			}
		}
	} else if step == 60 {
		return hourLong * 100
	} else {
		stepHour := step / 60
		size := int((24 * time.Hour) / (time.Duration(step) * time.Minute))
		for i := 0; i < size; i++ {
			if hourShort < stepHour*(i+1) {
				return (day*100 + stepHour*i) * 100
			}
		}
	}
	return 0
}
