package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Unit struct {
	UnitId   int     `json:"unit_id" xorm:"pk comment('单元编号') VARCHAR(255)"`
	UnitName string  `json:"unit_name" xorm:"not null comment('单元名称') VARCHAR(255)"`
	Distance float32 `json:"distance" xorm:"default 0 comment('距离') FLOAT"`
}

const (
	host     = "100.100.142.132"
	port     = 25432
	user     = "postgres"
	password = "smai123"
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

	err = engine.Sync2(
		new(Unit))
	if err != nil {
		log.Fatal("sync database err: ", err)
	}
}
