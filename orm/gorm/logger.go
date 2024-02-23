package main

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/xiaomi-tc/log15"
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

	db.SetLogger(&LogImpl{})
	db.LogMode(true)

	list := make([]*Product, 0)
	db.New().Exec("SELECT * FROM products WHERE name = ?", "水杯") // 该方式返回的行数为0
	db.New().Table("products").Select("*").Where("name = ?", "水杯").Scan(&list)
}

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

var LogFormatter = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			source          = fmt.Sprintf("\033[35m(%v)\033[0m", values[1])
		)

		messages = []interface{}{source}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf(" \033[36;1m[%.2fms]\033[0m", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); true {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders and ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, " \033[33m"+sql+"\033[0m")
			messages = append(messages, fmt.Sprintf(" \033[36;31m[%v]\033[0m ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned"))
		} else {
			messages = append(messages, "\033[31;1m")
			messages = append(messages, values[2:]...)
			messages = append(messages, "\033[0m")
		}
	}

	return
}

type LogImpl struct{}

func (l *LogImpl) Print(args ...interface{}) {
	messages := LogFormatter(args...)
	log.Info("gorm", "data", fmt.Sprint(messages))
}
