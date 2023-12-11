package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type User struct {
	Id       int     `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	UserName *string `json:"user_name" xorm:"not null comment('用户名') VARCHAR(255)"`
	Age      *int    `json:"age" xorm:"comment('年龄') INT(11)"`
	Job      *string `json:"job" xorm:"comment('职业') VARCHAR(255)"`
	Hobby    *string `json:"hobby" xorm:"comment('兴趣') VARCHAR(255)"`
}

func main() {
	var (
		engine *xorm.Engine
		err    error
	)

	engine, err = xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("init engine err: ", err)
	}

	err = engine.Sync2(
		new(User),
	)
	if err != nil {
		log.Fatal("sync database err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// userName := "zhangsan"
	// // age := 28
	// job := "engineer"
	// hobby := "play football"
	// users := []User{
	// 	{UserName: &userName, Job: &job, Hobby: &hobby},
	// 	// {UserName: "lisi", Age: 27, Hobby: "watch video"},
	// 	// {UserName: "wangwu", Job: "engineer", Hobby: "play football"},
	// }

	var sql = "insert or replace into user(id,user_name,age,job,hobby) values(1,'lizhenfeng',18,'teacher','watch tv') on conflict do update set user_name=excluded.user_name, age=excluded.age, job=excluded.job, hobby=excluded.hobby;"

	_, err = session.Exec(sql)
	if err != nil {
		log.Println(err)
	}

	// _, err = session.Insert(users)
	// if err != nil {
	// 	log.Println(err)
	// }
}
