package Model

import (
	"time"
)

type Usaha struct {
	ID        uint      `gorm:"notNull" json:"id"`
	UsahaName string    `gorm:"notNull;size:255" json:"usahaname"`
	UsahaId   string    `gorm:"size:255;notNull;primaryKey" json:"usaha_id"`
	Type      string    `gorm:"notNull;size:255" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
