package main

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Send events every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		// Write a simple event
		fmt.Fprintf(w, "data: Current time is %s\n\n", t.Format(time.RFC1123)) // Each message is delimited by a blank line (\n\n)
		// Flush to ensure it's sent immediately
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)
	fmt.Println("Listening on http://localhost:8080/events")
	http.ListenAndServe(":8080", nil)
}
