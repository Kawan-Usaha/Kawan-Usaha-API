package Model

import (
	"time"
)

type Article struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	UserId      string     `gorm:"notNull;size:255" json:"user_id"`
	Title       string     `gorm:"notNull;size:255" json:"title"`
	Content     string     `gorm:"type:text;null" json:"content"`
	Image       string     `gorm:"null;size:255" json:"image"`
	IsPublished bool       `gorm:"default:true;notNull" json:"is_published"`
	Category    []Category `gorm:"foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;" json:"category"`
	CreatedAt   time.Time  `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp" json:"updated_at"`
}
