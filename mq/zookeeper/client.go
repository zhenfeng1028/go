package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		startClient()
		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	// service := "127.0.0.1:8899"
	// 获取地址
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Println("get server host failed,", err)
		return
	}

	fmt.Println("connect host:", serverHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)

	result, err := io.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
}

func getServerHost() (host string, err error) {
	conn, err := GetConnect()
	if err != nil {
		fmt.Println("connect zk error,", err)
		return
	}
	defer conn.Close()
	serverList, err := GetServerList(conn)
	if err != nil {
		fmt.Println("get server list error,", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty")
		return
	}

	// 随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	checkError(err)
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}
