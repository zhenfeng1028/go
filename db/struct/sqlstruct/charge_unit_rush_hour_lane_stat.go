package sqlstruct

type ChargeUnitRushHourLaneStat struct {
	Id              int64   `json:"id" xorm:"pk autoincr unique BIGINT"`
	Day             int64   `json:"day" xorm:"index(idx_20340_idx_day_charge_unit_lane_type) index(idx_20340_idx_day_lane_type) BIGINT"`
	ChargeUnit      string  `json:"charge_unit" xorm:"index(idx_20340_idx_day_charge_unit_lane_type) VARCHAR(255)"`
	Lane            int64   `json:"lane" xorm:"index(idx_20340_idx_day_charge_unit_lane_type) index(idx_20340_idx_day_lane_type) BIGINT"`
	Type            int64   `json:"type" xorm:"index(idx_20340_idx_day_charge_unit_lane_type) index(idx_20340_idx_day_lane_type) BIGINT"`
	HourCount       float64 `json:"hour_count" xorm:"DOUBLE"`
	CongestionIndex float64 `json:"congestion_index" xorm:"DOUBLE"`
	AvgSpeed        int64   `json:"avg_speed" xorm:"BIGINT"`
	FlowCount       int64   `json:"flow_count" xorm:"BIGINT"`
	Equivalent      float64 `json:"equivalent" xorm:"DOUBLE"`
	Ke1Count        int64   `json:"ke1_count" xorm:"BIGINT"`
	Ke2Count        int64   `json:"ke2_count" xorm:"BIGINT"`
	Ke3Count        int64   `json:"ke3_count" xorm:"BIGINT"`
	Ke4Count        int64   `json:"ke4_count" xorm:"BIGINT"`
	Huo1Count       int64   `json:"huo1_count" xorm:"BIGINT"`
	Huo2Count       int64   `json:"huo2_count" xorm:"BIGINT"`
	Huo3Count       int64   `json:"huo3_count" xorm:"BIGINT"`
	Huo4Count       int64   `json:"huo4_count" xorm:"BIGINT"`
	Huo5Count       int64   `json:"huo5_count" xorm:"BIGINT"`
	Huo6Count       int64   `json:"huo6_count" xorm:"BIGINT"`
	Zhuan1Count     int64   `json:"zhuan1_count" xorm:"BIGINT"`
	Zhuan2Count     int64   `json:"zhuan2_count" xorm:"BIGINT"`
	Zhuan3Count     int64   `json:"zhuan3_count" xorm:"BIGINT"`
	Zhuan4Count     int64   `json:"zhuan4_count" xorm:"BIGINT"`
	Zhuan5Count     int64   `json:"zhuan5_count" xorm:"BIGINT"`
	Zhuan6Count     int64   `json:"zhuan6_count" xorm:"BIGINT"`
	OtherCount      int64   `json:"other_count" xorm:"BIGINT"`
	Holiday         int64   `json:"holiday" xorm:"BIGINT"`
}
