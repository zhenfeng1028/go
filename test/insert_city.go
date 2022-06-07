package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type ChargeUnitVehicleSourceCityStat struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Day        int    `json:"day" xorm:"comment('时间，日') index(IDX_day) INT(11)"`
	ChargeUnit string `json:"charge_unit" xorm:"comment('收费单元') index(IDX_day_charge_unit) VARCHAR(255)"`
	ProvCode   string `json:"prov_code" xorm:"comment('流量来源地（省）') VARCHAR(255)"`
	CityCode   string `json:"city_code" xorm:"comment('流量来源地（地级市）') VARCHAR(255)"`
	FlowCount  int    `json:"flow_count" xorm:"comment('通行量') INT(11)"`
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.131:4000)/expressway_roadnet_dev_refactor?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := engine.Sync2(new(ChargeUnitVehicleSourceCityStat)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}

	session := engine.NewSession()
	defer session.Close()

	source := []ChargeUnitVehicleSourceCityStat{
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3301", FlowCount: 9000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3302", FlowCount: 8000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3303", FlowCount: 7000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3304", FlowCount: 6000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3305", FlowCount: 5000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3306", FlowCount: 4000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3307", FlowCount: 3000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "33", CityCode: "3308", FlowCount: 2000},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "65", CityCode: "6501", FlowCount: 100},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "65", CityCode: "6502", FlowCount: 100},
		{Day: 20210802, ChargeUnit: "G250133001000610", ProvCode: "65", CityCode: "6530", FlowCount: 100},
	}

	_, err = session.Insert(source)
	if err != nil {
		log.Println(err)
	}
}
