package main

import (
	"fmt"
	"os"
)

func main() {
	// 只读方式打开当前目录下的xxx.md文件
	file, err := os.Open("./xxx.md")
	if err != nil {
		fmt.Println("open file failed!, err: ", err)
		return
	}
	// 关闭文件
	file.Close()
}
