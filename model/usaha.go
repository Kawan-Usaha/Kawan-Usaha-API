package Model

import (
	"time"
)

type Usaha struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UsahaName string    `gorm:"notNull;size:255" json:"usahaname"`
	Owner     string    `gorm:"size:255;notNull;foreignKey" json:"owner"`
	Type      string    `gorm:"notNull;size:255" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
