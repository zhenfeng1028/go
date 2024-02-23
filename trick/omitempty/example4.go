package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street     string      `json:"street"`               // 街道
	Ste        string      `json:"suite,omitempty"`      // 单元（可以不存在）
	City       string      `json:"city"`                 // 城市
	State      string      `json:"state"`                // 州/省
	Zipcode    string      `json:"zipcode"`              // 邮编
	Coordinate *Coordinate `json:"coordinate,omitempty"` // 经纬度
}

type Coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

// 为了达到我们想要的效果，可以把坐标定义为指针类型，这样 Golang 就能知道一个指针的“空值”是多少了，否则面对一个我们自定义的结构， Golang 是猜不出我们想要的空值的。
func main() {
	data := `{
			  "street": "200 Larkin St",
			  "city": "San Francisco",
			  "state": "CA",
			  "zipcode": "94102"
		  }`
	addr := new(Address)
	err := json.Unmarshal([]byte(data), &addr)
	if err != nil {
		log.Panicf("data unmarshal err: %s ", err)
	}

	// 处理了一番 addr 变量...

	addressBytes, _ := json.MarshalIndent(addr, "", "    ")
	fmt.Printf("%s\n", string(addressBytes))
}
