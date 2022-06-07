package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type ProvinceCode struct {
	Provid   int    `json:"ProvId" xorm:"not null pk INT(11)"`
	Provname string `json:"ProvName" xorm:"not null VARCHAR(255)"`
}

func main() {
	var engine *xorm.Engine
	var err error
	// engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.131:4000)/expressway_taishi_dev?charset=utf8")
	// if err != nil {
	// 	log.Fatal("数据库连接失败:", err)
	// }
	engine, err = xorm.NewEngine("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "100.100.142.132", 25432, "postgres", "smai123", "expressway"))
	if err != nil {
		log.Fatalf("init engine err : %v", err)
	}

	session := engine.NewSession()
	defer session.Close()

	res := make([]*ProvinceCode, 0)
	rows, err := session.Table(new(ProvinceCode)).Select("*").Rows(new(ProvinceCode))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		c := new(ProvinceCode)
		err = rows.Scan(c)
		if err != nil {
			fmt.Printf("rows.Scan ProvinceCode err : %v", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("ProvinceCode:", *v)
	}
}
