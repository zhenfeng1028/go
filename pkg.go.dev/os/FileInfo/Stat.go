package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Stat(`/Users/lizhenfeng/test.sh`)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileinfo.Name())    // 获取文件名
	fmt.Println(fileinfo.IsDir())   // 判断是否是目录
	fmt.Println(fileinfo.ModTime()) // 获取文件修改时间
	fmt.Println(fileinfo.Mode())    // 获取文件权限
	fmt.Println(fileinfo.Size())    // 获取文件大小
	fmt.Println(fileinfo.Sys())     // 底层数据源
}
