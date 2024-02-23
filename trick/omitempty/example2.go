package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street  string `json:"street"`          // 街道
	Ste     string `json:"suite,omitempty"` // 单元（可以不存在）
	City    string `json:"city"`            // 城市
	State   string `json:"state"`           // 州/省
	Zipcode string `json:"zipcode"`         // 邮编
}

// 可以在 Golang 的结构体定义中添加 omitempty 关键字，来表示这条信息如果没有提供，在序列化成 json 的时候就不要包含其默认值。
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
