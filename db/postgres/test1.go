package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

const (
	host     = "100.100.142.132"
	port     = 25432
	user     = "postgres"
	password = "smai123"
	dbname   = "expressway_roadnet_dev"
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
}
