package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 从wsl访问主机数据库
const connectStrWSL = "wsl_root:@tcp(172.24.192.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

type Product struct {
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price"`
	Name  string `gorm:"column:name"`
}

func main() {

	db, err := gorm.Open("mysql", connectStrWSL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create table
	// db.CreateTable(&Product{})

	// Create
	// db.Create(&Product{Code: "L1219", Price: 1006, Name: "fffff"})

	list := make([]*Product, 0)
	db.New().Table("products").Limit(3).Offset(2).Select("*").Scan(&list)
	bs, _ := json.Marshal(list)
	fmt.Println(string(bs))

	// Read
	// var product Product
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	// db.Delete(&product)

	db = db.Table("products").Where("price < 1011").Select("*")
	subQuery := db.SubQuery()
	fmt.Println(subQuery)

	var cnt int64
	db.New().Raw(`select count(1) from ? as a`, subQuery).Count(&cnt)
	fmt.Println(cnt)
}
