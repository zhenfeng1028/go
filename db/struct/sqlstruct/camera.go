package sqlstruct

type Camera struct {
	Id              int64  `json:"id" xorm:"pk autoincr unique BIGINT"`
	CameraId        string `json:"camera_id" xorm:"unique VARCHAR(255)"`
	Name            string `json:"name" xorm:"VARCHAR(255)"`
	Type            int64  `json:"type" xorm:"BIGINT"`
	ParentSectionId string `json:"parent_section_id" xorm:"VARCHAR(255)"`
	Direction       int64  `json:"direction" xorm:"BIGINT"`
	PileNo          int64  `json:"pile_no" xorm:"BIGINT"`
}
