package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type User struct {
	Id      int    `json:"id" xorm:"pk autoincr comment('自增id') INT(11)"`
	Name    string `json:"name" xorm:"not null unique comment('姓名') VARCHAR(255)"`
	Age     int    `json:"sex" xorm:"not null comment('年龄') INT(11)"`
	Address string `json:"class" xorm:"comment('住址') VARCHAR(255)"`
}

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("init engine err: ", err)
	}
	if err := engine.Sync2(new(User)); err != nil {
		log.Fatal("sync database err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// users := []User{
	// 	{Name: "lizhenfeng", Age: 26, Address: "Chenchun Road"},
	// 	{Name: "jiangjunqiao", Age: 24, Address: "Zhangjiang Road"},
	// }

	// _, err = session.Insert(users)
	// if err != nil {
	// 	log.Println(err)
	// }

	// 1，on duplicate key update 语句根据主键id来判断当前插入是否已存在。
	// 2，已存在时，只会更新on duplicate key update之后限定的字段。
	// var sql = "insert or replace into user(id,name,age,address) values(2,'baijiaao',24,'Yuqiao') on conflict(name) do update set age=excluded.age, address=excluded.address;"

	// 3，on duplicate key update 语句也可以根据唯一键来判断当前插入的记录是否已存在。
	// var sql = "insert or replace into user(name,age,address) values('lizhenfeng',18,'Chenchun Road 2') on conflict(name) do update set age=excluded.age, address=excluded.address;"

	// 没有主键或唯一键字段值相同，即判断当前记录不存在，新插入一条。
	// var sql = "insert or replace into user(name,age,address) values('lizhenfeng2',18,'Chenchun Road 2') on conflict(name) do update set age=excluded.age, address=excluded.address;"

	// 4，如果传递了主键，是可以修改唯一键字段内容的。
	sql := "insert or replace into user(id,name,age,address) values(1,'lizhenfeng1',18,'Chenchun Road 2') on conflict(name) do update set name=excluded.name, age=excluded.age, address=excluded.address;"

	_, err = session.Exec(sql)
	if err != nil {
		log.Println(err)
	}
}
