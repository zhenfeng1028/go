package sqlstruct

type LaneStat struct {
	Id               int64   `json:"id" xorm:"pk autoincr unique BIGINT"`
	CameraId         string  `json:"camera_id" xorm:"index(idx_20400_idx_id_direction_minute) VARCHAR(255)"`
	Direction        int64   `json:"direction" xorm:"index(idx_20400_idx_id_direction_minute) BIGINT"`
	Minute           int64   `json:"minute" xorm:"index(idx_20400_idx_id_direction_minute) BIGINT"`
	Lane             string  `json:"lane" xorm:"VARCHAR(255)"`
	TrafficStatus    int64   `json:"traffic_status" xorm:"BIGINT"`
	AvgSpeed         int64   `json:"avg_speed" xorm:"BIGINT"`
	AvgSpeedCount    int64   `json:"avg_speed_count" xorm:"BIGINT"`
	TimeInterval     float64 `json:"time_interval" xorm:"DOUBLE"`
	DistanceInterval float64 `json:"distance_interval" xorm:"DOUBLE"`
	TimeRatio        float64 `json:"time_ratio" xorm:"DOUBLE"`
	SpaceRatio       float64 `json:"space_ratio" xorm:"DOUBLE"`
	SmallCount       int64   `json:"small_count" xorm:"BIGINT"`
	MediumCount      int64   `json:"medium_count" xorm:"BIGINT"`
	LargeCount       int64   `json:"large_count" xorm:"BIGINT"`
	HeavyCount       int64   `json:"heavy_count" xorm:"BIGINT"`
	VehicleQueue     float64 `json:"vehicle_queue" xorm:"DOUBLE"`
}
