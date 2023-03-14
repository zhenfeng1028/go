package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type CameraHourlyStat struct {
	Id                      int     `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Hour                    int     `json:"hour" xorm:"comment('时间，小时') unique(IDX_n_hour_id_direction) index(IDX_n_hour_direction) INT(11)"`
	CameraId                string  `json:"camera_id" xorm:"unique(IDX_n_hour_id_direction) index(IDX_n_id_minute) VARCHAR(255)"`
	Direction               int     `json:"direction" xorm:"comment('0上行 1下行') unique(IDX_n_hour_id_direction)  index(IDX_n_hour_direction) INT(11)"`
	Minute                  int64   `json:"minute" xorm:"comment('更新时间，分钟') index(IDX_n_id_minute) BIGINT(20)"`
	Status                  int     `json:"status" xorm:"comment('当前状况 0畅通 1缓行 2拥堵') INT(11)"`
	AvgSpeed                int     `json:"avg_speed" xorm:"comment('最新平均车速') INT(11)"`
	AvgSpeedCount           int     `json:"avg_speed_count" xorm:"comment('最新车辆数') INT(11)"`
	CongestionMinute        int     `json:"congestion_minute" xorm:"comment('拥堵分钟数，包括预估的补漏时间') INT(11)"`
	SlowMinute              int     `json:"slow_minute" xorm:"comment('缓行分钟数，包括预估的补漏时间') INT(11)"`
	ContinueMinute          int     `json:"continue_minute" xorm:"comment('状态连续分钟数，包括预估的补漏时间') INT(11)"`
	AvgSpeedOverall         int     `json:"avg_speed_overall" xorm:"comment('平均车速') INT(11)"`
	AvgSpeedCountOverall    int     `json:"avg_speed_count_overall" xorm:"comment('车辆数') INT(11)"`
	TimeIntervalOverall     float32 `json:"time_interval_overall" xorm:"comment('平均车头时距') FLOAT"`
	DistanceIntervalOverall float32 `json:"distance_interval_overall" xorm:"comment('平均车头间距') FLOAT"`
	TimeRatioOverall        float32 `json:"time_ratio_overall" xorm:"comment('平均车道时间占有率') FLOAT"`
	SpaceRatioOverall       float32 `json:"space_ratio_overall" xorm:"comment('平均车道空间占有率') FLOAT"`
	// CountByType             map[string]int `json:"count_by_type" xorm:"comment('0小型车 1中型车 2大型车 3重型车') TEXT"`
	SmallCountOverall  int `json:"small_count_overall" xorm:"comment('小型车数量') INT(11)"`
	MediumCountOverall int `json:"medium_count_overall" xorm:"comment('中型车数量') INT(11)"`
	LargeCountOverall  int `json:"large_count_overall" xorm:"comment('大型车数量') INT(11)"`
	HeavyCountOverall  int `json:"heavy_count_overall" xorm:"comment('重型车数量') INT(11)"`
}

type Camera struct {
	Id              int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	CameraId        string `json:"camera_id" xorm:"unique VARCHAR(255)"`
	Name            string `json:"name" xorm:"comment('摄像头名称') VARCHAR(255)"`
	ParentSectionId string `json:"parent_section_id" xorm:"comment('摄像头所属路段id，不分上下行') VARCHAR(255)"`
	Type            int    `json:"type" xorm:"comment('0道路摄像头 1收费站摄像头') INT(11)"`
	Direction       int    `json:"direction" xorm:"comment('收费站 0入口 1出口') INT(11)"`
}

type PeriodDetailsQueryRes struct {
	Period           string  `json:"period"`
	PeriodValueCount int     `json:"period_value_count"`
	MaxDbPeriodValue int     `json:"max_db_period_value"`
	Count            int     `json:"count"`
	SpeedCount       int     `json:"speed_count"`
	AllSpeed         int     `json:"all_speed"`
	AvgSpeed         int     `json:"avg_speed"`
	Small            int     `json:"small"`
	Medium           int     `json:"medium"`
	Large            int     `json:"large"`
	Heavy            int     `json:"heavy"`
	TimeInterval     float32 `json:"time_interval"`
	DistanceInterval float32 `json:"distance_interval"`
	TimeRatio        float32 `json:"time_ratio"`
	SpaceRatio       float32 `json:"space_ratio"`
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.15:4000)/expressway_roadnet_dev?charset=utf8")
	if err != nil {
		log.Fatal("init engine err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	cmQuery := "hour%100 as period, count(distinct(hour)) as period_value_count, max(hour) as max_db_period_value, sum(avg_speed_count_overall) as speed_count, sum(avg_speed_count_overall*avg_speed_overall) as all_speed, count(1) as count"
	vtQuery := "sum(small_count_overall) as small, sum(medium_count_overall) as medium, sum(large_count_overall) as large, sum(heavy_count_overall) as heavy"
	irQuery := "sum(time_interval_overall) as time_interval, sum(distance_interval_overall) as distance_interval, sum(time_ratio_overall) as time_ratio, sum(space_ratio_overall) as space_ratio"

	session = session.Table(new(CameraHourlyStat)).Select(fmt.Sprintf("%s, %s, %s", cmQuery, vtQuery, irQuery))
	session = session.Join("LEFT", "camera", "camera_hourly_stat.camera_id = camera.camera_id")
	session = session.And("hour>=2021052500")
	session = session.And("hour<=2021052523")
	session = session.And("camera.type = 0")
	session = session.GroupBy("period")

	var (
		res  = make(map[string]*PeriodDetailsQueryRes)
		res2 = make([]*PeriodDetailsQueryRes, 0)
	)
	rows, err := session.Rows(new(PeriodDetailsQueryRes))
	if err != nil {
		fmt.Printf("GetPeriodDetails err: %v", err)
		fmt.Println(session.LastSQL())
		return
	}
	for rows.Next() {
		info := new(PeriodDetailsQueryRes)
		err = rows.Scan(info)
		if err != nil {
			_ = fmt.Errorf("rows.Scan PeriodDetailsQueryRes err: %v", err)
			continue
		}
		res[info.Period] = info
	}
	for _, v := range res {
		res2 = append(res2, v)
	}
	sort.Sort(PeriodDetailsQueryResSlice(res2))

	for _, v := range res2 {
		bs, _ := json.Marshal(v)
		fmt.Println(string(bs))
		fmt.Println()
	}
}

type PeriodDetailsQueryResSlice []*PeriodDetailsQueryRes

func (s PeriodDetailsQueryResSlice) Len() int      { return len(s) }
func (s PeriodDetailsQueryResSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s PeriodDetailsQueryResSlice) Less(i, j int) bool {
	ti, _ := strconv.Atoi(s[i].Period)
	tj, _ := strconv.Atoi(s[j].Period)
	return ti < tj
}
