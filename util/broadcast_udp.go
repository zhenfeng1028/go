package main

import (
	"fmt"
	"net"
	"time"
)

// Multicast configuration
const (
	multicastAddr = "224.0.0.1:9999" // Multicast address and port
)

func main() {
	go startReceiver() // Start the receiver
	time.Sleep(1 * time.Second)

	startSender() // Start the sender
}

// Receiver: listens for multicast messages
func startReceiver() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error starting receiver:", err)
		return
	}
	defer conn.Close()

	conn.SetReadBuffer(1024)
	fmt.Println("Receiver is listening for messages...")

	buf := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		fmt.Printf("Received message: %s from %s\n", string(buf[:n]), src)
	}
}

// Sender: broadcasts messages to the multicast group
func startSender() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error starting sender:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Sender is broadcasting messages...")
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Broadcast message %d", i)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
		time.Sleep(2 * time.Second)
	}
	fmt.Println("Sender finished broadcasting.")
}
