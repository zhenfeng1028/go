package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Camera struct {
	Id              int    `json:"id" xorm:"pk autoincr comment('自增id') INT(11)"`
	CameraId        string `json:"camera_id" xorm:"unique VARCHAR(255)"`
	Name            string `json:"name" xorm:"comment('摄像头名称') VARCHAR(255)"`
	ParentSectionId string `json:"parent_section_id" xorm:"comment('摄像头所属路段id，不分上下行') VARCHAR(255)"`
	Type            int    `json:"type" xorm:"comment('0道路摄像头 1收费站摄像头') INT(11)"`
	Direction       int    `json:"direction" xorm:"comment('收费站 0入口 1出口') INT(11)"`
	PileNo          int    `json:"pile_no" xorm:"comment('桩号') INT(11)"`
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
		new(Camera))
	if err != nil {
		log.Fatal("sync database err: ", err)
	}
}
