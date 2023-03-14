package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Camera struct {
	Id              int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	CameraId        string `json:"camera_id" xorm:"unique VARCHAR(255)"`
	Name            string `json:"name" xorm:"comment('摄像头名称') VARCHAR(255)"`
	ParentSectionId string `json:"parent_section_id" xorm:"comment('摄像头所属路段id，不分上下行') VARCHAR(255)"`
	Type            int    `json:"type" xorm:"comment('0道路摄像头 1收费站摄像头') INT(11)"`
	Direction       int    `json:"direction" xorm:"comment('收费站 0入口 1出口') INT(11)"`
	PileNo          int    `json:"pile_no" xorm:"comment('桩号') INT(11)"`
}

type Section struct {
	Id              int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	SectionId       string `json:"section_id" xorm:"comment('路段id，分上下行') unique VARCHAR(255)"`
	ParentSectionId string `json:"parent_section_id" xorm:"comment('所属父路段id，不分上下行') VARCHAR(255)"`
	Direction       int    `json:"direction" xorm:"comment('0上行 1下行') INT(11)"`
	Name            string `json:"name" xorm:"comment('路段名称') VARCHAR(255)"`
	Mileage         int    `json:"mileage" xorm:"comment('里程数') INT(11)"`
	PrevSectionId   string `json:"prev_section_id" xorm:"comment('相邻前一路段id') VARCHAR(255)"`
	NextSectionId   string `json:"next_section_id" xorm:"comment('相邻后一路段id') VARCHAR(255)"`
	Type            int    `json:"type" xorm:"comment('0道路摄像头 1收费站摄像头') INT(11)"`
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

	_, err = session.Where("1=1").Delete(new(Camera))
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}
	_, err = session.Where("1=1").Delete(new(Section))
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}
}
