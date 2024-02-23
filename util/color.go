package main

import (
	"fmt"
)

// 利用\033控制字体显示效果的格式
// "\033[参数1;参数2;参数3m要显示的内容\033[0m"
// 参数1表示字体背景颜色，参数2表示字体颜色，参数3表示字体格式

func main() {
	msg := "an old falcon"

	// 字体格式
	reset := "\033[0m"
	bold := "\033[1m"
	italic := "\033[3m"
	underline := "\033[4m"
	strike := "\033[9m"

	// 字体颜色
	cRed := "\033[31m"
	cGreen := "\033[32m"
	cYellow := "\033[33m"
	cBlue := "\033[34m"
	cPurple := "\033[35m"
	cCyan := "\033[36m"
	cWhite := "\033[37m"

	// 背景颜色
	bcRed := "\033[41m"
	bcGreen := "\033[42m"
	bcYellow := "\033[43m"
	bcBlue := "\033[44m"
	bcPurple := "\033[45m"
	bcCyan := "\033[46m"
	bcWhite := "\033[47m"

	fmt.Println(msg)

	fmt.Println(cRed + msg)
	fmt.Println(cGreen + msg)
	fmt.Println(cYellow + msg)
	fmt.Println(cBlue + msg)
	fmt.Println(cPurple + msg)
	fmt.Println(cWhite + msg)
	fmt.Println(cCyan + msg + reset)

	fmt.Println(bcRed + msg + reset)
	fmt.Println(bcGreen + msg + reset)
	fmt.Println(bcYellow + msg + reset)
	fmt.Println(bcBlue + msg + reset)
	fmt.Println(bcPurple + msg + reset)
	fmt.Println(bcWhite + msg + reset)
	fmt.Println(bcCyan + msg + reset)

	fmt.Println(bold + msg)
	fmt.Println(italic + msg + reset)
	fmt.Println(underline + msg + reset)
	fmt.Println(strike + msg + reset)
	fmt.Println(msg)
}
