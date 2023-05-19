package Model

import (
	"time"
)

type Usaha struct {
	ID        uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	User      string    `gorm:"notNull;size:255" json:"user"`
	UsahaName string    `gorm:"notNull;size:255" json:"usahaname"`
	Type      string    `gorm:"notNull;size:255" json:"type"`
	Tags      []Tag     `gorm:"many2many:usaha_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;" json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
