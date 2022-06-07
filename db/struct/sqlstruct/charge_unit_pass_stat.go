package sqlstruct

type ChargeUnitPassStat struct {
	Id           int64   `json:"id" xorm:"pk autoincr unique BIGINT"`
	Day          int64   `json:"day" xorm:"index BIGINT"`
	PassCount    int64   `json:"pass_count" xorm:"BIGINT"`
	Equivalent   float64 `json:"equivalent" xorm:"DOUBLE"`
	Ke1Count     int64   `json:"ke1_count" xorm:"BIGINT"`
	Ke2Count     int64   `json:"ke2_count" xorm:"BIGINT"`
	Ke3Count     int64   `json:"ke3_count" xorm:"BIGINT"`
	Ke4Count     int64   `json:"ke4_count" xorm:"BIGINT"`
	Huo1Count    int64   `json:"huo1_count" xorm:"BIGINT"`
	Huo2Count    int64   `json:"huo2_count" xorm:"BIGINT"`
	Huo3Count    int64   `json:"huo3_count" xorm:"BIGINT"`
	Huo4Count    int64   `json:"huo4_count" xorm:"BIGINT"`
	Huo5Count    int64   `json:"huo5_count" xorm:"BIGINT"`
	Huo6Count    int64   `json:"huo6_count" xorm:"BIGINT"`
	Zhuan1Count  int64   `json:"zhuan1_count" xorm:"BIGINT"`
	Zhuan2Count  int64   `json:"zhuan2_count" xorm:"BIGINT"`
	Zhuan3Count  int64   `json:"zhuan3_count" xorm:"BIGINT"`
	Zhuan4Count  int64   `json:"zhuan4_count" xorm:"BIGINT"`
	Zhuan5Count  int64   `json:"zhuan5_count" xorm:"BIGINT"`
	Zhuan6Count  int64   `json:"zhuan6_count" xorm:"BIGINT"`
	OtherCount   int64   `json:"other_count" xorm:"BIGINT"`
	PassDistance int64   `json:"pass_distance" xorm:"BIGINT"`
	Holiday      int64   `json:"holiday" xorm:"BIGINT"`
}
