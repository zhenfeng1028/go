package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type FloatTest struct {
	Id       int     `json:"id" xorm:"pk autoincr comment('自增id') INT(11)"`
	Distance float32 `json:"distance" xorm:"comment('距离')  decimal"`
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "100.100.142.132", 25432, "postgres", "smai123", "test"))
	if err != nil {
		log.Fatal("init engine err: ", err)
	}
	if err := engine.Sync2(new(FloatTest)); err != nil {
		log.Fatal("sync database err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	data := []FloatTest{
		{Distance: 1234567.8},
	}

	_, err = session.Insert(data)
	if err != nil {
		log.Println(err)
	}
}
