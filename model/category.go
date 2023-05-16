package Model

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"notNull;size:255" json:"title"`
	Tags      []Tag     `gorm:"foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;" json:"tags"`
	Image     string    `gorm:"null;size:255" json:"image"`
	CreatedAt time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
