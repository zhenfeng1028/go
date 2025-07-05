package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Test both cases
	fmt.Println("=== With DisableCompression: false (default behavior) ===")
	requestWithCompression(false)

	fmt.Println("\n=== With DisableCompression: true (manual gzip handling) ===")
	requestWithCompression(true)
}

func requestWithCompression(disable bool) {
	transport := &http.Transport{
		DisableCompression: disable,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, _ := http.NewRequest("GET", "https://httpbin.org/gzip", nil)

	// Uncomment the next line to simulate user explicitly requesting gzip
	// req.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))

	var bodyReader io.ReadCloser

	// Manually handle gzip if compression is disabled and content is gzipped
	if disable && resp.Header.Get("Content-Encoding") == "gzip" {
		gzReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("Failed to create gzip reader:", err)
			return
		}
		defer gzReader.Close()
		bodyReader = gzReader
	} else {
		// Either compression was enabled, or server returned plain text
		bodyReader = resp.Body
	}

	bodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println("Body:\n", strings.TrimSpace(string(bodyBytes)))
}
