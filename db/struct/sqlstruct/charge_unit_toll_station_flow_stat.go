package sqlstruct

type ChargeUnitTollStationFlowStat struct {
	Id         int64  `json:"id" xorm:"pk autoincr unique BIGINT"`
	Day        int64  `json:"day" xorm:"index BIGINT"`
	ChargeUnit string `json:"charge_unit" xorm:"index VARCHAR(255)"`
	StationHex string `json:"station_hex" xorm:"VARCHAR(255)"`
	FlowCount  int64  `json:"flow_count" xorm:"BIGINT"`
}
