package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]bool)
	mu      sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 9999...")
	go acceptConnections(listener)

	// Broadcast messages from the server
	for {
		var message string
		fmt.Print("Enter message to broadcast: ")
		fmt.Scanln(&message)

		broadcastMessage(message)
	}
}

// Accept incoming connections
func acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		mu.Lock()
		clients[conn] = true
		mu.Unlock()

		fmt.Println("New client connected:", conn.RemoteAddr())
		go handleClient(conn)
	}
}

// Handle individual client connection
func handleClient(conn net.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr())
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received from %s: %s\n", conn.RemoteAddr(), message)
	}
}

// Broadcast a message to all connected clients
func broadcastMessage(message string) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range clients {
		_, err := fmt.Fprintln(conn, message)
		if err != nil {
			fmt.Println("Error sending message to", conn.RemoteAddr(), ":", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}
