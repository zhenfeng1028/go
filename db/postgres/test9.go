package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Gantry struct {
	Gantryid         string  `json:"GantryId" xorm:"not null pk comment('门架编号') VARCHAR(19)"`
	Gantryhex        string  `json:"GantryHex" xorm:"comment('ETC门架Hex值') index VARCHAR(6)"`
	Chargeunit       string  `json:"ChargeUnit" xorm:"comment('收费单元编号') index VARCHAR(16)"`
	Gantryname       string  `json:"GantryName" xorm:"comment('门架名称') VARCHAR(255)"`
	Gantrynameunique string  `json:"GantryNameUnique" xorm:"comment('门架名称，区别主副门架') VARCHAR(255)"`
	Label            int     `json:"Label" xorm:"default 0 comment('门架标签,0正常门架 ，虚拟门架 1 , 其他公司门架2') SMALLINT"`
	Direction        int     `json:"Direction" xorm:"default 0 comment('0-正向,1-反向') SMALLINT"`
	Istranshubaso    bool    `json:"IsTranshubAsO" xorm:"default false comment('作为起始点，是否可判断互通枢纽 ，0否，1是') BOOL"`
	Nameaso          string  `json:"NameAsO" xorm:"comment('OD名称，门架作为起始对应名称') VARCHAR(50)"`
	Istranshubasd    bool    `json:"IsTranshubAsD" xorm:"default false comment('作为终止点，是否可判断互通枢纽 ，0否，1是') BOOL"`
	Nameasd          string  `json:"NameAsD" xorm:"comment('OD名称，门架作为终点对应名称') VARCHAR(50)"`
	Tollstationaso   int64   `json:"TollStationAsO" xorm:"default 0 comment('作为入口时，所对应的收费站') BIGINT"`
	Tollstationasd   int64   `json:"TollStationAsD" xorm:"default 0 comment('作为出口时，所对应的收费站') BIGINT"`
	Longitude        float64 `json:"Longitude" xorm:"default 0 comment('经度') FLOAT"`
	Latitude         float64 `json:"Latitude" xorm:"default 0 comment('纬度') FLOAT"`
	Stake            string  `json:"Stake" xorm:"comment('桩号') VARCHAR(50)"`
	Lanenum          int64   `json:"LaneNum" xorm:"default 0 comment('门架车道数') BIGINT"`
	Directionname    string  `json:"DirectionName" xorm:"comment('方向名称') VARCHAR(50)"`
	Region           string  `json:"Region" xorm:"comment('所属管理处') VARCHAR(50)"`
	Section          string  `json:"Section" xorm:"comment('所属路段') VARCHAR(50)"`
	Unitname         string  `json:"UnitName" xorm:"comment('单元名称') VARCHAR(50)"`
	Road             string  `json:"Road" xorm:"comment('道路编码') VARCHAR(50)"`
	Management       string  `json:"Management" xorm:"comment('管理处') VARCHAR(50)"`
	Startplace       string  `json:"StartPlace" xorm:"comment('起始位置') VARCHAR(50)"`
	Endplace         string  `json:"EndPlace" xorm:"comment('结束位置') VARCHAR(50)"`
	Stakenumber      int64   `json:"StakeNumber" xorm:"comment('桩号，数字') BIGINT"`
}

const (
	host     = "100.100.142.132"
	port     = 25432
	user     = "postgres"
	password = "smai123"
	dbname   = "test"
)

func main() {
	var (
		engine *xorm.Engine
		err    error
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	engine, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal("init engine err: ", err)
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
			fmt.Printf("rows.Scan Gantry err: %v", err)
			continue
		}
		res = append(res, c)
	}
	for _, v := range res {
		fmt.Println("Gantry:", *v)
	}
}
