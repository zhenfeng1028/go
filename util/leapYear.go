package main

import "fmt"

var yeartime int64

func main() {
	fmt.Println("请输入要判断的年份：")
	fmt.Scanln(&yeartime)
	if (yeartime%4 == 0 && yeartime%100 != 0) || yeartime%400 == 0 {
		fmt.Printf("输入的年份 %d 是闰年\n", yeartime)
	} else {
		fmt.Printf("输入的年份 %d 不是闰年\n", yeartime)
	}
}
