package Model

import (
	"time"
)

type Usaha struct {
	ID        uint      `gorm:"notNull" json:"id"`
	UsahaName string    `gorm:"notNull;size:255" json:"usahaname"`
	Type      string    `gorm:"notNull;size:255" json:"type"`
	Tags      []Tag     `gorm:"foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;" json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
