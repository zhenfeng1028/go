package main

import (
	"fmt"
	"os"
)

const path = "/home/lizhenfeng/github/go/trick"

func main() {
	filePaths := []string{}
	findFile(path, &filePaths)
	fmt.Println(filePaths)
}

// 递归获取目录下所有文件
func findFile(dirpath string, filepaths *[]string) {
	files, _ := os.ReadDir(dirpath)
	for _, f := range files {
		path := dirpath + "/" + f.Name()
		if f.IsDir() {
			findFile(path, filepaths)
		}
		*filepaths = append(*filepaths, path)
	}
}
