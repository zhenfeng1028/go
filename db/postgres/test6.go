package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type CameraHourlyStat struct {
	Id                      int     `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Hour                    int     `json:"hour" xorm:"comment('时间，小时') unique(IDX_n_hour_id_direction) index(IDX_n_hour_direction) INT(11)"`
	CameraId                string  `json:"camera_id" xorm:"unique(IDX_n_hour_id_direction) index(IDX_n_id_minute) VARCHAR(255)"`
	Direction               int     `json:"direction" xorm:"comment('0上行 1下行') unique(IDX_n_hour_id_direction)  index(IDX_n_hour_direction) INT(11)"`
	Minute                  int64   `json:"minute" xorm:"comment('更新时间，分钟') index(IDX_n_id_minute) BIGINT(20)"`
	Status                  int     `json:"status" xorm:"comment('当前状况 0畅通 1缓行 2拥堵') INT(11)"`
	CongestionIndex         float32 `json:"congestion_index" xorm:"comment('拥堵指数') FLOAT"`
	AvgSpeed                int     `json:"avg_speed" xorm:"comment('最新平均车速') INT(11)"`
	AvgSpeedCount           int     `json:"avg_speed_count" xorm:"comment('最新车辆数') INT(11)"`
	CongestionMinute        int     `json:"congestion_minute" xorm:"comment('拥堵分钟数，包括预估的补漏时间') INT(11)"`
	SlowMinute              int     `json:"slow_minute" xorm:"comment('缓行分钟数，包括预估的补漏时间') INT(11)"`
	SlightCongestionMinute  int     `json:"slight_congestion_minute" xorm:"comment('轻度拥堵分钟数，包括预估的补漏时间') INT(11)"`
	BasicSmoothMinute       int     `json:"basic_smooth_minute" xorm:"comment('基本畅通分钟数，包括预估的补漏时间') INT(11)"`
	ContinueMinute          int     `json:"continue_minute" xorm:"comment('状态连续分钟数，包括预估的补漏时间') INT(11)"`
	AvgSpeedOverall         int     `json:"avg_speed_overall" xorm:"comment('平均车速') INT(11)"`
	AvgSpeedCountOverall    int     `json:"avg_speed_count_overall" xorm:"comment('车辆数') INT(11)"`
	TimeIntervalOverall     float32 `json:"time_interval_overall" xorm:"comment('平均车头时距') FLOAT"`
	DistanceIntervalOverall float32 `json:"distance_interval_overall" xorm:"comment('平均车头间距') FLOAT"`
	TimeRatioOverall        float32 `json:"time_ratio_overall" xorm:"comment('平均车道时间占有率') FLOAT"`
	SpaceRatioOverall       float32 `json:"space_ratio_overall" xorm:"comment('平均车道空间占有率') FLOAT"`
	VehicleQueueOverall     float32 `json:"vehicle_queue_overall" xorm:"comment('平均车辆排队长度') FLOAT"`
	// CountByType             map[string]int `json:"count_by_type" xorm:"comment('0小型车 1中型车 2大型车 3重型车') TEXT"`
	SmallCountOverall  int `json:"small_count_overall" xorm:"comment('小型车数量') INT(11)"`
	MediumCountOverall int `json:"medium_count_overall" xorm:"comment('中型车数量') INT(11)"`
	LargeCountOverall  int `json:"large_count_overall" xorm:"comment('大型车数量') INT(11)"`
	HeavyCountOverall  int `json:"heavy_count_overall" xorm:"comment('重型车数量') INT(11)"`
}

type SectionHourlyStat struct {
	Id                      int     `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	SectionId               string  `json:"section_id" xorm:"unique(IDX_id_hour) VARCHAR(255)"`
	Hour                    int     `json:"hour" xorm:"comment('时间，小时') unique(IDX_id_hour) index INT(11)"`
	Minute                  int64   `json:"minute" xorm:"comment('更新时间，分钟') index BIGINT(20)"`
	Status                  int     `json:"status" xorm:"comment('当前状况 0畅通 1缓行 2拥堵') INT(11)"`
	CongestionIndex         float32 `json:"congestion_index" xorm:"comment('拥堵指数') FLOAT"`
	AvgSpeed                int     `json:"avg_speed" xorm:"comment('最新平均车速') INT(11)"`
	AvgSpeedCount           int     `json:"avg_speed_count" xorm:"comment('最新车辆数') INT(11)"`
	CongestionMinute        int     `json:"congestion_minute" xorm:"comment('拥堵分钟数，包括预估的补漏时间') INT(11)"`
	SlowMinute              int     `json:"slow_minute" xorm:"comment('缓行分钟数，包括预估的补漏时间') INT(11)"`
	SlightCongestionMinute  int     `json:"slight_congestion_minute" xorm:"comment('轻度拥堵分钟数，包括预估的补漏时间') INT(11)"`
	BasicSmoothMinute       int     `json:"basic_smooth_minute" xorm:"comment('基本畅通分钟数，包括预估的补漏时间') INT(11)"`
	ContinueMinute          int     `json:"continue_minute" xorm:"comment('状态连续分钟数，包括预估的补漏时间') INT(11)"`
	MaxConCongestionMinute  int     `json:"max_con_congestion_minute" xorm:"comment('最长连续拥堵分钟数') INT(11)"`
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

	err = engine.Sync2(
		new(CameraHourlyStat),
		new(SectionHourlyStat),
	)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}
}
