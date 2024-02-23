package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street  string `json:"street"`  // 街道
	Ste     string `json:"suite"`   // 单元（可以不存在）
	City    string `json:"city"`    // 城市
	State   string `json:"state"`   // 州/省
	Zipcode string `json:"zipcode"` // 邮编
}

// 多了一行 "suite": "", ，而这则信息在原本的 json 数据中是没有的
// 但我们更希望的是，在一个地址有 suite 号码的时候输出，不存在 suite 的时候就不输出
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
