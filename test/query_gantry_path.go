package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type GantryPath struct {
	ID       int     `json:"id" xorm:"not null pk autoincr INT(11) 'id'"`
	FromID   string  `json:"from_id" xorm:"VARCHAR(255) index not null 'from_id'"`
	FromType int     `json:"from_type" xorm:"int not null 'from_type'"`
	ToID     string  `json:"to_id" xorm:"VARCHAR(255) index not null 'to_id'"`
	ToType   int     `json:"to_type" xorm:"int not null 'to_type'"`
	Distance float64 `json:"distance" xorm:"default 0 FLOAT"`
	Path     string  `json:"path" xorm:"VARCHAR(4000)"`
}

func main() {
	var engine *xorm.Engine
	var err error

	engine, err = xorm.NewEngine("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "100.100.142.132", 25432, "postgres", "smai123", "expressway"))
	if err != nil {
		log.Fatalf("init engine err : %v", err)
	}

	session := engine.NewSession()
	defer session.Close()

	res := make([]*GantryPath, 0)
	rows, err := session.Table(new(GantryPath)).Select("*").Rows(new(GantryPath))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		c := new(GantryPath)
		err = rows.Scan(c)
		if err != nil {
			fmt.Printf("rows.Scan GantryPath err : %v", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("GantryPath:", *v)
	}
}
