package Model

type TagsOnUsaha struct {
	Id       uint    `gorm:"notNull;primaryKey" json:"id"`
	Tags     []Tag   `gorm:"foreignkey:Id;type:text;null" json:"tags"`
	Tags_id  uint    `gorm:"notNull;size:255" json:"tags_id"`
	Usaha    []Usaha `gorm:"foreignkey:Id;type:text;null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"usaha"`
	Usaha_id uint    `gorm:"notNull;size:255" json:"usaha_id"`
}
