package main

import "fmt"

const (
	Hour_Last_Minute = 59

	Day_Last_Minute = 2359

	Jan_Last_Minute          = 1312359
	Non_Leap_Feb_Last_Minute = 2282359 // 平年
	Leap_Feb_Last_Minute     = 2292359 // 闰年
	Mar_Last_Minute          = 3312359
	Apr_Last_Minute          = 4302359
	May_Last_Minute          = 5312359
	Jun_Last_Minute          = 6302359
	Jul_Last_Minute          = 7312359
	Aug_Last_Minute          = 8312359
	Sep_Last_Minute          = 9302359
	Oct_Last_Minute          = 10312359
	Nov_Last_Minute          = 11302359
	Dec_Last_Minute          = 12312359
)

func main() {
	var (
		startMinute = 202112312357
		endMinute   = 202201010028
	)
	minuteMap := make(map[int]struct{})
	for m := startMinute; m <= endMinute; m++ {
		if _, ok := minuteMap[m]; !ok {
			minuteMap[m] = struct{}{}
		}

		// 闰年
		leapYear := false
		year := m / 100000000
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			leapYear = true
		}
		// 跨月
		monthLastMinute := m % 100000000
		switch monthLastMinute {
		case Jan_Last_Minute:
			m += 697640
		case Non_Leap_Feb_Last_Minute:
			if !leapYear {
				m += 727640
			}
		case Leap_Feb_Last_Minute:
			m += 717640
		case Mar_Last_Minute:
			m += 697640
		case Apr_Last_Minute:
			m += 707640
		case May_Last_Minute:
			m += 697640
		case Jun_Last_Minute:
			m += 707640
		case Jul_Last_Minute:
			m += 697640
		case Aug_Last_Minute:
			m += 697640
		case Sep_Last_Minute:
			m += 707640
		case Oct_Last_Minute:
			m += 697640
		case Nov_Last_Minute:
			m += 707640
		case Dec_Last_Minute:
			m += 88697640
		}

		// 跨天
		if m%10000 == 2359 {
			m += 7640
		}

		// 跨小时
		if m%100 == 59 {
			m += 40
		}
	}
	fmt.Println("duration:", len(minuteMap))
}
