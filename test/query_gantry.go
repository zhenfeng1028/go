package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Gantry struct {
	Gantryid         string  `json:"GantryId" xorm:"not null pk comment('门架编号') VARCHAR(19)"`
	Gantryhex        string  `json:"GantryHex" xorm:"comment('ETC门架Hex值') index VARCHAR(6)"`
	Chargeunit       string  `json:"ChargeUnit" xorm:"comment('收费单元编号') index VARCHAR(16)"`
	Gantryname       string  `json:"GantryName" xorm:"comment('收费单元名称') VARCHAR(255)"`
	Gantrynameunique string  `json:"GantryNameUnique" xorm:"comment('门架名称，区别主副门架') VARCHAR(255)"`
	Label            int     `json:"Label" xorm:"default 0 comment('门架标签,0正常门架 ，虚拟门架 1 , 其他公司门架2') TINYINT(4)"`
	Direction        int     `json:"Direction" xorm:"default 0 comment('0-正向,1-反向') TINYINT(4)"`
	Istranshubaso    int     `json:"IsTranshubAsO" xorm:"default 0 comment('作为起始点，是否可判断互通枢纽 ，0否，1是') TINYINT(1)"`
	Nameaso          string  `json:"NameAsO" xorm:"comment('OD名称，门架作为起始对应名称') VARCHAR(255)"`
	Istranshubasd    int     `json:"IsTranshubAsD" xorm:"default 0 comment('作为终止点，是否可判断互通枢纽 ，0否，1是') TINYINT(1)"`
	Nameasd          string  `json:"NameAsD" xorm:"comment('OD名称，门架作为终点对应名称') VARCHAR(255)"`
	Tollstationaso   int     `json:"TollStationAsO" xorm:"default 0 comment('作为入口时，所对应的收费站') INT(10)"`
	Tollstationasd   int     `json:"TollStationAsD" xorm:"default 0 comment('作为出口时，所对应的收费站') INT(10)"`
	Longitude        float32 `json:"Longitude" xorm:"default 0 comment('经度') FLOAT"`
	Latitude         float32 `json:"Latitude" xorm:"default 0 comment('纬度') FLOAT"`
	Stake            string  `json:"Stake" xorm:"comment('桩号') VARCHAR(255)"`
	Lanenum          int     `json:"LaneNum" xorm:"default 0 comment('门架车道数') INT(11)"`
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

	res := make([]*Gantry, 0)
	rows, err := session.Table(new(Gantry)).Select("*").Rows(new(Gantry))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		c := new(Gantry)
		err = rows.Scan(c)
		if err != nil {
			fmt.Printf("rows.Scan Gantry err : %v", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("Gantry:", *v)
	}
}
