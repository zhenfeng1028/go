package main

import (
	"encoding/json"
	"fmt"
)

type Coordinate struct {
	Lat *float64 `json:"latitude,omitempty"`
	Lng *float64 `json:"longitude,omitempty"`
}

// 正确的写法也是将结构体内的定义改为指针
func main() {
	cData := `{
		"latitude": 0.0,
		"longitude": 0.0
	}`
	c := new(Coordinate)
	json.Unmarshal([]byte(cData), &c)

	// 具体处理逻辑...

	coordinateBytes, _ := json.MarshalIndent(c, "", "    ")
	fmt.Printf("%s\n", string(coordinateBytes))
}
