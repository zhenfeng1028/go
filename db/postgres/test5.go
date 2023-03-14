package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

const (
	sql_test                      = "insert into tbl_test(id, name, age, address) values($1,$2,$3,$4) ON CONFLICT (id) DO UPDATE SET age=excluded.age, address=excluded.address;"
	Upsert_Camera_Hourly_Stat_Sql = "insert into camera_hourly_stat(camera_id, direction, hour, minute, status, avg_speed, avg_speed_count, congestion_minute, slow_minute, continue_minute, avg_speed_overall, avg_speed_count_overall, time_interval_overall, distance_interval_overall, time_ratio_overall, space_ratio_overall, small_count_overall, medium_count_overall, large_count_overall, heavy_count_overall) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20) ON CONFLICT (camera_id, direction, hour) DO UPDATE SET minute=excluded.minute, status=excluded.status, avg_speed=excluded.avg_speed, avg_speed_count=excluded.avg_speed_count, congestion_minute=excluded.congestion_minute, slow_minute=excluded.slow_minute, continue_minute=excluded.continue_minute, avg_speed_overall=excluded.avg_speed_overall, avg_speed_count_overall=excluded.avg_speed_count_overall, time_interval_overall=excluded.time_interval_overall, distance_interval_overall=excluded.distance_interval_overall, time_ratio_overall=excluded.time_ratio_overall, space_ratio_overall=excluded.space_ratio_overall, small_count_overall=excluded.small_count_overall, medium_count_overall=excluded.medium_count_overall, large_count_overall=excluded.large_count_overall, heavy_count_overall=excluded.heavy_count_overall"
)

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

	// result, err := session.Exec(sql_test, 1, "huahua1", 201, "京华市1")
	// fmt.Println(result)

	result, err := session.Exec(Upsert_Camera_Hourly_Stat_Sql, "aaa", 0, 2021083101, 202108310157, 0, 70, 80, 20, 20, 20, 80, 1000, 20, 20, 0.5, 0.5, 500, 250, 126, 124)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	} else {
		fmt.Println(result)
	}
}
