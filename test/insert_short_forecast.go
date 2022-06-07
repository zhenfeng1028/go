package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Fc3SectionForecastShortTime struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('自增id') INT(11)"`
	Section     string    `json:"section" xorm:"comment('路段') VARCHAR(20)"`
	Direction   int       `json:"direction" xorm:"default 0 comment('0-上行，1-下行') SMALLINT"`
	AvgSpeed    float64   `json:"avg_speed" xorm:"default 0 comment('平均速度') FLOAT"`
	FlowCount   int       `json:"flow_count" xorm:"default 0 comment('流量') INTEGER"`
	SectionLen  float32   `json:"section_len" xorm:"default 0 comment('路段长度') FLOAT"`
	FcFTime     string    `json:"fc_f_time" xorm:"comment('预测生成时刻，例如202107051840，以10分钟为单位，所以最后一个字符应为0') index(short_idx_f_time) index(short_idx_f_time_d_time) VARCHAR(20)"`
	FcDTime     string    `json:"fc_d_time" xorm:"comment('被预测的时刻，例如202107051940，以10分钟为单位，所以最后一个字符应为0') index(short_idx_d_time) index(short_idx_f_time_d_time) index(short_idx_daily_d_time) VARCHAR(20)"`
	Day         string    `json:"day" xorm:"comment('被预测的日期，例如20210705，以天为单位') index(short_idx_daily) index(short_idx_daily_d_time) VARCHAR(20)"`
	TimeAdv     int       `json:"time_adv" xorm:"default 0 comment('预测提前量，单位分钟，缺省为60分钟') index(short_idx_daily) index(short_idx_daily_d_time) SMALLINT"`
	CreatedBy   int       `json:"created_by" xorm:"not null default 0 comment('创建人') INTEGER"`
	CreatedTime time.Time `json:"created_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdatedBy   int       `json:"updated_by" xorm:"not null default 0 comment('更新人') INTEGER"`
	UpdatedTime time.Time `json:"updated_time" xorm:"not null comment('更新时间') DATETIME"`
	ObjStatus   string    `json:"obj_status" xorm:"not null default '0' comment('数据有效性 0有效1无效') VARCHAR(1)"`
}

type ChargeUnit struct {
	ChargeUnitId   string  `json:"charge_unit_id" xorm:"not null comment('收费单元编号') VARCHAR(255)"`
	ChargeUnitName string  `json:"charge_unit_name" xorm:"not null comment('收费单元名称') VARCHAR(255)"`
	Distance       float32 `json:"distance" xorm:"default 0 comment('距离') FLOAT"`
	PrevChargeUnit string  `json:"prev_charge_unit" xorm:"comment('前一个收费单元') VARCHAR(50)"`
	NextChargeUnit string  `json:"next_charge_unit" xorm:"comment('后一个收费单元') VARCHAR(50)"`
	StartStake     string  `json:"start_stake" xorm:"comment('起始桩号') VARCHAR(50)"`
	EndStake       string  `json:"end_stake" xorm:"comment('结束桩号') VARCHAR(50)"`
	Direction      int     `json:"direction" xorm:"comment('方向') SMALLINT"`
	StakeDistance  float32 `json:"stake_distance" xorm:"default 0 comment('桩号距离') FLOAT"`
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.131:4000)/expressway_forecast?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	// if err := engine.Sync2(new(Fc3SectionForecastShortTime)); err != nil {
	// 	log.Fatal("数据表同步失败:", err)
	// }

	session := engine.NewSession()
	defer session.Close()

	res := make([]*ChargeUnit, 0)
	rows, err := session.Table(new(ChargeUnit)).Select("*").Rows(new(ChargeUnit))
	if err != nil {
		fmt.Println(session.LastSQL())
	}
	for rows.Next() {
		cu := new(ChargeUnit)
		err = rows.Scan(cu)
		if err != nil {
			fmt.Printf("rows.Scan ChargeUnit err : %v", err)
			continue
		}
		res = append(res, cu)
	}

	datas := make([]Fc3SectionForecastShortTime, 0)
	for _, v := range res {
		data := Fc3SectionForecastShortTime{
			Section:     v.ChargeUnitId,
			Direction:   v.Direction,
			AvgSpeed:    80,
			FlowCount:   200,
			Day:         "20210301",
			TimeAdv:     120,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		for fcDTime := 202103010000; fcDTime < 202103012359; fcDTime += 15 {
			data.FcDTime = strconv.Itoa(fcDTime)
			if fcDTime%100 == 45 {
				fcDTime += 40
			}
			datas = append(datas, data)
		}
	}

	_, err = session.Insert(datas)
	if err != nil {
		log.Println(err)
	}
}
