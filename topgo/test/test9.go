package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	layout := "2006-01-02+15:04:05"

	got := now.Format(layout)
	fmt.Println(got)
}
