// bufio包实现了带缓冲区的读写，是对文件读写的封装
// 模式			 含义
// os.O_WRONLY	只写
// os.O_CREATE	创建文件
// os.O_RDONLY	只读
// os.O_RDWR	读写
// os.O_TRUNC	清空
// os.O_APPEND	追加

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr() {
	// 参数2：打开模式，所有模式都在上面
	// 参数3：权限控制
	// r读 w写 x执行  r 4   w 2   x 1
	file, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open("./xxx.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}

func main() {
	wr()
	re()
}
