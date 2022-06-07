package sqlstruct

type CameraHourlyStat struct {
	Id                      int64   `json:"id" xorm:"pk autoincr unique BIGINT"`
	Hour                    int64   `json:"hour" xorm:"index(idx_20316_idx_n_hour_direction) unique(idx_20316_idx_n_hour_id_direction) BIGINT"`
	CameraId                string  `json:"camera_id" xorm:"unique(idx_20316_idx_n_hour_id_direction) index(idx_20316_idx_n_id_minute) VARCHAR(255)"`
	Direction               int64   `json:"direction" xorm:"index(idx_20316_idx_n_hour_direction) unique(idx_20316_idx_n_hour_id_direction) BIGINT"`
	Minute                  int64   `json:"minute" xorm:"index(idx_20316_idx_n_id_minute) BIGINT"`
	Status                  int64   `json:"status" xorm:"BIGINT"`
	AvgSpeed                int64   `json:"avg_speed" xorm:"BIGINT"`
	AvgSpeedCount           int64   `json:"avg_speed_count" xorm:"BIGINT"`
	CongestionMinute        int64   `json:"congestion_minute" xorm:"BIGINT"`
	SlowMinute              int64   `json:"slow_minute" xorm:"BIGINT"`
	ContinueMinute          int64   `json:"continue_minute" xorm:"BIGINT"`
	AvgSpeedOverall         int64   `json:"avg_speed_overall" xorm:"BIGINT"`
	AvgSpeedCountOverall    int64   `json:"avg_speed_count_overall" xorm:"BIGINT"`
	TimeIntervalOverall     float64 `json:"time_interval_overall" xorm:"DOUBLE"`
	DistanceIntervalOverall float64 `json:"distance_interval_overall" xorm:"DOUBLE"`
	TimeRatioOverall        float64 `json:"time_ratio_overall" xorm:"DOUBLE"`
	SpaceRatioOverall       float64 `json:"space_ratio_overall" xorm:"DOUBLE"`
	SmallCountOverall       int64   `json:"small_count_overall" xorm:"BIGINT"`
	MediumCountOverall      int64   `json:"medium_count_overall" xorm:"BIGINT"`
	LargeCountOverall       int64   `json:"large_count_overall" xorm:"BIGINT"`
	HeavyCountOverall       int64   `json:"heavy_count_overall" xorm:"BIGINT"`
	SlightCongestionMinute  int64   `json:"slight_congestion_minute" xorm:"BIGINT"`
	BasicSmoothMinute       int64   `json:"basic_smooth_minute" xorm:"BIGINT"`
	VehicleQueueOverall     float64 `json:"vehicle_queue_overall" xorm:"DOUBLE"`
}
