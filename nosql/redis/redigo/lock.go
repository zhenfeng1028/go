package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   100,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	success, randomKey := rdLock("abc")
	if !success {
		fmt.Println("rdLock failed")
		return
	}
	defer rdUnlock("abc", randomKey)
	// 业务逻辑
	time.Sleep(10 * time.Second)
}

func rdLock(key string) (bool, string) {
	randomKey := uuid.NewV4().String()
	cmds := []Transaction{
		{
			Cmd:  "SETNX",
			Args: []interface{}{key, randomKey},
		},
		{
			Cmd:  "EXPIRE",
			Args: []interface{}{key, 10},
		},
	}

	reply, err := redis.Ints(DoTransaction(cmds))
	if err != nil {
		fmt.Println("redis.Ints", "err", err)
		return false, ""
	}
	if len(reply) != 2 {
		fmt.Println("reply unexpect", "reply", reply)
		return false, ""
	} else if reply[0] == 0 {
		fmt.Println("lock not free", "reply", reply)
		return false, ""
	}
	return true, randomKey
}

func rdUnlock(key, randomKey string) {
	// 使用 Lua 脚本确保原子性删除锁
	script := redis.NewScript(1, `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`)
	con := pool.Get()
	defer con.Close()
	err := script.Load(con)
	if err != nil {
		fmt.Println("script.Load", "err", err)
		return
	}
	_, err = script.Do(con, key, randomKey)
	if err != nil {
		fmt.Println("script.Do", "err", err)
		return
	}
}

type Transaction struct {
	Cmd  string
	Args []interface{}
}

func DoTransaction(cmds []Transaction) (reply interface{}, err error) {
	rc := pool.Get()
	defer rc.Close()
	rc.Send("MULTI")
	for _, cmd := range cmds {
		sendErr := rc.Send(cmd.Cmd, cmd.Args...)
		if sendErr != nil {
			fmt.Println("Transacton failed:", sendErr)
			break
		}
	}
	return rc.Do("EXEC")
}
