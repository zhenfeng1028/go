package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User1 struct {
	Id   int64  `xorm:"'my_id' pk autoincr"` // 指定主键并自增
	Name string `xorm:"'my_name' unique"`    // 唯一的
	Time int64  `xorm:"'my_time' updated"`   // 修改后自动更新时间
}

type User2 struct {
	Id   int64  `xorm:"pk autoincr"` // 指定主键并自增
	Name string `xorm:"unique"`      // 唯一的
	Time int64  `xorm:"updated"`     // 修改后自动更新时间
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.15:4000)/lzf?charset=utf8") // lzf是数据库实例名称
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// engine.SetMapper(names.GonicMapper{})

	// preMapper := names.NewPrefixMapper(names.SnakeMapper{}, "prefix_")
	// engine.SetColumnMapper(preMapper)

	// sufMapper := names.NewSuffixMapper(names.SnakeMapper{}, "_suffix")
	// engine.SetColumnMapper(sufMapper)

	// err = engine.Ping()
	// if err != nil {
	// 	log.Fatal("数据库连接失败:", err)
	// }

	// if err := engine.Sync(new(User1)); err != nil {
	// 	log.Fatal("数据表同步失败:", err)
	// }

	user := new(User1)
	user.Name = "xxxxx"
	engine.Insert(user)

	// pUser2 := new(User2)
	// engine.CreateTables(pUser2)

	// var isExist bool
	// var isEmpty bool
	// isExist, err = engine.IsTableExist(pUser2)
	// fmt.Println(isExist)
	// isEmpty, err = engine.IsTableEmpty(pUser2)
	// fmt.Println(isEmpty)

	// err = engine.DropTables("MyUser")
	// if err != nil {
	// 	log.Fatal("数据表删除失败:", err)
	// }
}

// func (user *User1) TableName() string {
// 	return "MyUser"
// }
