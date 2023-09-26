package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	var buffer bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		// 对行做一些处理
		if err != nil {
			if err == io.EOF {
				fmt.Println("read finished!")
				break
			} else {
				panic(err)
			}
		}
		buffer.WriteString(line)
	}
	err = os.WriteFile("output.txt", buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
