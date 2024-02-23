package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street     string     `json:"street"`               // 街道
	Ste        string     `json:"suite,omitempty"`      // 单元（可以不存在）
	City       string     `json:"city"`                 // 城市
	State      string     `json:"state"`                // 州/省
	Zipcode    string     `json:"zipcode"`              // 邮编
	Coordinate Coordinate `json:"coordinate,omitempty"` // 经纬度
}

type Coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

// 读入原来的地址数据，处理后序列化输出，我们就会发现即使加上了 omitempty 关键字，输出的 json 还是带上了一个空的坐标信息
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
