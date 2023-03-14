// tcp/server/main.go
package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端
// 处理流程：
// 1 监听端口
// 2 接受客户端请求建立连接
// 3 创建goroutine处理连接

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		sendStr := "I have recieved " + recvStr + "'s message"
		conn.Write([]byte(sendStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept() // 接收连接
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
