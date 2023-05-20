package Model

import "time"

type Tag struct {
	ID        uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	Name      string    `gorm:"notNull;size:255" json:"name"`
	Usaha     []Usaha   `gorm:"Null;many2many:usaha_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"usaha"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
