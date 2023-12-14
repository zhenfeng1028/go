package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool // redis连接池

func init() {
	pool = &redis.Pool{ // 实例化一个连接池
		MaxIdle:     16,
		MaxActive:   100, // 最大连接数量
		IdleTimeout: 300, // 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { // 要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	c := pool.Get() // 从连接池，取一个连接
	defer c.Close() // 函数运行结束 ，将连接关闭

	_, err := c.Do("SET", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("GET", "abc"))
	if err != nil {
		fmt.Println("get abc faild,", err)
		return
	}

	fmt.Println(r)
}
