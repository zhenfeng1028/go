package sqlstruct

type TollStationCongestionStatistic struct {
	Id            int64  `json:"id" xorm:"pk autoincr unique BIGINT"`
	Day           int64  `json:"day" xorm:"index BIGINT"`
	TollStationId string `json:"toll_station_id" xorm:"VARCHAR(255)"`
	Duration      int64  `json:"duration" xorm:"BIGINT"`
}
