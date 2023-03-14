package main

import (
	"encoding/json"
	"fmt"
)

type Coordinate struct {
	Lat float64 `json:"latitude,omitempty"`
	Lng float64 `json:"longitude,omitempty"`
}

// 另一个“陷阱”是，对于用 omitempty 定义的 field ，如果给它赋的值恰好等于默认空值的话，在转为 json 之后也不会输出这个 field 。
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
