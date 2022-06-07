package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type ChargeUnit struct {
	ChargeUnitId   string  `json:"charge_unit_id" xorm:"not null comment('收费单元编号') VARCHAR(255)"`
	ChargeUnitName string  `json:"charge_unit_name" xorm:"not null comment('收费单元名称') VARCHAR(255)"`
	Distance       float32 `json:"distance" xorm:"default 0 comment('距离') FLOAT"`
	PrevChargeUnit string  `json:"prev_charge_unit" xorm:"comment('前一个收费单元') VARCHAR(50)"`
	NextChargeUnit string  `json:"next_charge_unit" xorm:"comment('后一个收费单元') VARCHAR(50)"`
	StartStake     string  `json:"start_stake" xorm:"comment('起始桩号') VARCHAR(50)"`
	EndStake       string  `json:"end_stake" xorm:"comment('结束桩号') VARCHAR(50)"`
	Direction      int     `json:"direction" xorm:"comment('方向') SMALLINT"`
}

func main() {
	var engine *xorm.Engine
	var err error
	// engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.131:4000)/expressway_taishi_dev?charset=utf8")
	// if err != nil {
	// 	log.Fatal("数据库连接失败:", err)
	// }
	engine, err = xorm.NewEngine("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "100.100.142.132", 25432, "postgres", "smai123", "zhuyong"))
	if err != nil {
		log.Fatalf("init engine err : %v", err)
	}

	session := engine.NewSession()
	defer session.Close()

	res := make([]*ChargeUnit, 0)
	rows, err := session.Table(new(ChargeUnit)).Select("*").Rows(new(ChargeUnit))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		c := new(ChargeUnit)
		err = rows.Scan(c)
		if err != nil {
			fmt.Printf("rows.Scan ChargeUnit err : %v", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("ChargeUnit:", *v)
	}
}
