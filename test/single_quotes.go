package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type A struct {
	EventTime time.Time `json:"event_time"`
}

func main() {
	s := `{"event_time": "2021-06-22T16:13:14.583+08:00"}`
	var a A
	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
