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

	res := make([]*Camera, 0)
	rows, err := session.Table(new(Camera)).Select("*").Rows(new(Camera))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		c := new(Camera)
		err = rows.Scan(c)
		if err != nil {
			fmt.Printf("rows.Scan Camera err: %v\n", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("Camera:", *v)
	}
}
