package Model

import "time"

type Tag struct {
	//ID: Autoincrement of ID for Tag, primary key.
	//Name: Name of the Tag.
	ID        uint       `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	Name      string     `gorm:"notNull;size:255" json:"name"`
	Usaha     []Usaha    `gorm:"Null;many2many:usaha_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"usaha"`
	Category  []Category `gorm:"Null;many2many:category_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
