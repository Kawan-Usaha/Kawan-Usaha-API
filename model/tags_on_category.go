package Model

type TagsOnCategory struct {
	Id          uint       `gorm:"notNull;primaryKey" json:"id"`
	Tags        []Tag      `gorm:"foreignkey:Id;type:text;null" json:"tags"`
	Tags_id     uint       `gorm:"notNull;size:255" json:"tags_id"`
	Category    []Category `gorm:"foreignkey:Id;type:text;null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"`
	Category_id uint       `gorm:"notNull;size:255" json:"category_id"`
}
