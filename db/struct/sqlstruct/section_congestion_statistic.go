package sqlstruct

type SectionCongestionStatistic struct {
	Id             int64  `json:"id" xorm:"pk autoincr unique BIGINT"`
	Day            int64  `json:"day" xorm:"index BIGINT"`
	StartSectionId string `json:"start_section_id" xorm:"VARCHAR(255)"`
	EndSectionId   string `json:"end_section_id" xorm:"VARCHAR(255)"`
	StartMinute    int64  `json:"start_minute" xorm:"BIGINT"`
	EndMinute      int64  `json:"end_minute" xorm:"BIGINT"`
	Duration       int64  `json:"duration" xorm:"BIGINT"`
	Mileage        int64  `json:"mileage" xorm:"BIGINT"`
	Hour           int64  `json:"hour" xorm:"index BIGINT"`
	Direction      int64  `json:"direction" xorm:"BIGINT"`
}
