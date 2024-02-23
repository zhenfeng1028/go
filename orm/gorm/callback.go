package main

import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/xiaomi-tc/log15"
)

const (
	injectNation = "X-manual-nation"
	connectStr   = "root:@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
)

type Person struct {
	Id     int64  `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Gender string `gorm:"column:gender"`
	Age    int64  `gorm:"column:age"`
	Nation string `gorm:"column:nation"`
}

func main() {
	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.LogMode(true)

	RegisterCallback(db)

	// Migrate the schema
	db.AutoMigrate(&Person{})

	// Create
	db.InstantSet(injectNation, "汉族")
	db.Create(&Person{Name: "lzf", Gender: "男", Age: 28})

	// Read
	db.InstantSet(injectNation, "回族")
	var person Person
	db.Table(person.TableName()).Where("age = ?", 26).Scan(&person)
}

func (Person) TableName() string {
	return "person"
}

func RegisterCallback(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", injectNationCreate)
	db.Callback().Query().Before("gorm:query").Register("my_plugin:before_query", injectNationQuery)
}

var injectNationCreate = func(scope *gorm.Scope) {
	if scope == nil || scope.HasError() {
		return
	}
	field, ok := scope.FieldByName("nation")
	if !ok {
		return
	}
	if field.IsBlank && field.Struct.Type.Kind() == reflect.String {
		i, ok := scope.Get(injectNation)
		if !ok {
			return
		}
		nation, ok := i.(string)
		if ok {
			if err := field.Set(nation); err != nil {
				log.Error("injectNationCreate", "field.Set", err)
			}
		}
	}
}

var injectNationQuery = func(scope *gorm.Scope) {
	if scope == nil || scope.HasError() {
		return
	}
	i, ok := scope.Get(injectNation)
	if !ok {
		return
	}
	nation, ok := i.(string)
	if ok {
		scope.Search.Where("nation = ?", nation)
	}
}
