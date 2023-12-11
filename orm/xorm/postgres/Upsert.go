package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

const sql_test = "insert into user(id, name, age, address) values($1,$2,$3,$4) ON CONFLICT (id) DO UPDATE SET age=excluded.age, address=excluded.address;"

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "lzf123"
	dbname   = "test"
)

func main() {
	var (
		engine *xorm.Engine
		err    error
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	engine, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal("init engine err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	result, err := session.Exec(sql_test, 1, "huahua1", 201, "京华市1")
	fmt.Println(result)
}
