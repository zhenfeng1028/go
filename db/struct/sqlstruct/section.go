package sqlstruct

type Section struct {
	Id              int64  `json:"id" xorm:"pk autoincr unique BIGINT"`
	SectionId       string `json:"section_id" xorm:"unique VARCHAR(255)"`
	ParentSectionId string `json:"parent_section_id" xorm:"VARCHAR(255)"`
	Direction       int64  `json:"direction" xorm:"BIGINT"`
	Name            string `json:"name" xorm:"VARCHAR(255)"`
	Mileage         int64  `json:"mileage" xorm:"BIGINT"`
	Type            int64  `json:"type" xorm:"BIGINT"`
	PrevSectionId   string `json:"prev_section_id" xorm:"VARCHAR(255)"`
	NextSectionId   string `json:"next_section_id" xorm:"VARCHAR(255)"`
}
