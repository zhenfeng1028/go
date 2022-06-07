package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type CityCode struct {
	CityCode int    `json:"city_code" xorm:"not null pk INT(11)"`
	ProvCode int    `json:"prov_code" xorm:"not null INT(11)"`
	CityName string `json:"city_name" xorm:"not null VARCHAR(255)"`
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

	res := make([]*CityCode, 0)
	rows, err := session.Table(new(CityCode)).Select("*").Rows(new(CityCode))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		c := new(CityCode)
		err = rows.Scan(c)
		if err != nil {
			fmt.Printf("rows.Scan CityCode err : %v", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("CityCode:", *v)
	}
}
