package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

const (
	Delete_Test_Sql = "delete from test where id = ?"
)

type Test struct {
	Id      int    `json:"id" xorm:"pk autoincr comment('自增id') INT(11)"`
	Name    string `json:"name" xorm:"not null unique comment('姓名') VARCHAR(255)"`
	Age     int    `json:"sex" xorm:"not null comment('年龄') INT(11)"`
	Address string `json:"class" xorm:"comment('住址') VARCHAR(255)"`
}

func main() {

	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("init engine err: ", err)
	}
	if err := engine.Sync2(new(Test)); err != nil {
		log.Fatal("sync database err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// test := []Test{
	// 	{Name: "lizhenfeng", Age: 26, Address: "Chenchun Road"},
	// 	{Name: "jiangjunqiao", Age: 24, Address: "Zhangjiang Road"},
	// }

	// _, err = session.Insert(test)
	// if err != nil {
	// 	log.Println(err)
	// }

	_, err = session.Exec(Delete_Test_Sql, 2)
	if err != nil {
		log.Println(err)
	}
}
