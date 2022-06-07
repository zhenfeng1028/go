package main

import (
	"fmt"
	"os"
)

func main() {
	// 新建文件
	file, err := os.Create("./xxx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}
