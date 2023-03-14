package sqlstruct

type ChargeUnitVehicleSourceCityStat struct {
	Id         int64  `json:"id" xorm:"pk autoincr unique BIGINT"`
	Day        int64  `json:"day" xorm:"index BIGINT"`
	ChargeUnit string `json:"charge_unit" xorm:"index VARCHAR(255)"`
	ProvCode   string `json:"prov_code" xorm:"VARCHAR(255)"`
	CityCode   string `json:"city_code" xorm:"VARCHAR(255)"`
	PassCount  int64  `json:"pass_count" xorm:"BIGINT"`
	FlowCount  int64  `json:"flow_count" xorm:"BIGINT"`
}
