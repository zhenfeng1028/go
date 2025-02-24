package main

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置必要的头部信息
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 刷新缓冲区
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// 每秒发送一条消息
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: Message %d at %s\n\n", i, time.Now().Format(time.RFC3339))
		flusher.Flush() // 将缓冲区的内容发送给客户端
		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>SSE Example</title>
		</head>
		<body>
			<h1>Server-Sent Events Example</h1>
			<div id="messages"></div>
			<script>
	            const eventSource = new EventSource('/events');
	            eventSource.onmessage = function(event) {
	                document.getElementById('messages').innerHTML += '<p>' + event.data + '</p>';
	            };
	        </script>
		</body>
		</html>
		`)
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
