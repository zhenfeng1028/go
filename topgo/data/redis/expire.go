package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()

	// 10秒后过期
	_, err = c.Do("EXPIRE", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}
