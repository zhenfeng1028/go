package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Camera struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	CameraId string `json:"camera_id" xorm:"unique VARCHAR(255)"`
	Name     string `json:"name" xorm:"comment('摄像头名称') VARCHAR(255)"`
	Location string `json:"location" xorm:"comment('位置') VARCHAR(255)"`
}

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

	err = engine.Sync2(new(Camera))
	if err != nil {
		log.Fatal("sync database err: ", err)
	}
}
