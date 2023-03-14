package main

import (
	"fmt"
	"os"
)

func main() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err: ", err)
		return
	}
	name := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写入信息：%s\n", name)
}
