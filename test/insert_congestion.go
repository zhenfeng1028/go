package main

import (
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type SectionCongestionStatistic struct {
	Id             int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Day            int    `json:"day" xorm:"comment('时间，日') index(IDX_day) INT(11)"`
	Hour           int    `json:"hour" xorm:"comment('时间，小时') index(IDX_hour) INT(11)"`
	Direction      int    `json:"direction" xorm:"comment('0上行1下行') INT(11)"`
	StartSectionId string `json:"start_section_id" xorm:"comment('起点桩号') VARCHAR(255)"`
	EndSectionId   string `json:"end_section_id" xorm:"comment('止点桩号') VARCHAR(255)"`
	StartMinute    int64  `json:"start_minute" xorm:"comment('拥堵起始分钟') BIGINT(20)"`
	EndMinute      int64  `json:"end_minute" xorm:"comment('拥堵结束分钟') BIGINT(20)"`
	Duration       int    `json:"duration" xorm:"comment('拥堵时长') INT(11)"`
	Mileage        int    `json:"mileage" xorm:"comment('拥堵里程') INT(11)"`
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.132:4000)/expressway_roadnet_dev_refactor?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := engine.Sync2(new(SectionCongestionStatistic)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}

	session := engine.NewSession()
	defer session.Close()

	stats := make([]SectionCongestionStatistic, 0)
	for i := 0; i <= 123; i++ {
		sectionId := "K" + strconv.Itoa(i) + "+" + "0" + "_" + "0"
		stat := SectionCongestionStatistic{Day: 20220222, Hour: 2022022215, Direction: 0, StartSectionId: sectionId, EndSectionId: sectionId, StartMinute: 202202221501, EndMinute: 202202221529, Duration: 29, Mileage: 1000}
		stats = append(stats, stat)
	}

	_, err = session.Insert(stats)
	if err != nil {
		log.Println(err)
	}
}
