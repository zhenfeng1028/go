package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	go startServer("127.0.0.1:8897")
	go startServer("127.0.0.1:8898")
	go startServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<-a
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func startServer(addr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 注册zk节点q
	// 连接zk
	conn, err := GetConnect()
	if err != nil {
		fmt.Println("connect zk error,", err)
	}
	defer conn.Close()
	// zk节点注册
	err = RegistServer(conn, addr)
	if err != nil {
		fmt.Println("regist node error,", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleCient(conn, addr)
	}
}

func handleCient(conn net.Conn, addr string) {
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(addr + ", " + daytime))
}

func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	checkError(err)
	return
}

func RegistServer(conn *zk.Conn, addr string) (err error) {
	_, err = conn.Create("/go_servers/"+addr, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}
