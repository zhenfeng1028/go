package sqlstruct

type SectionHourlyStat struct {
	Id                      int64   `json:"id" xorm:"pk autoincr unique BIGINT"`
	SectionId               string  `json:"section_id" xorm:"unique(idx_20427_idx_id_hour) VARCHAR(255)"`
	Hour                    int64   `json:"hour" xorm:"unique(idx_20427_idx_id_hour) index BIGINT"`
	Minute                  int64   `json:"minute" xorm:"index BIGINT"`
	Status                  int64   `json:"status" xorm:"BIGINT"`
	AvgSpeed                int64   `json:"avg_speed" xorm:"BIGINT"`
	AvgSpeedCount           int64   `json:"avg_speed_count" xorm:"BIGINT"`
	CongestionMinute        int64   `json:"congestion_minute" xorm:"BIGINT"`
	SlowMinute              int64   `json:"slow_minute" xorm:"BIGINT"`
	ContinueMinute          int64   `json:"continue_minute" xorm:"BIGINT"`
	MaxConCongestionMinute  int64   `json:"max_con_congestion_minute" xorm:"BIGINT"`
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
}
