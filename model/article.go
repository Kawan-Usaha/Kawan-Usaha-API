package Model

import (
	"time"
)

type Article struct {
	ID          uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	UserId      string    `gorm:"notNull;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:255" json:"user_id"`
	Title       string    `gorm:"notNull;size:255" json:"title"`
	Content     string    `gorm:"type:text;null" json:"content"`
	Image       string    `gorm:"null;size:255" json:"image"`
	IsPublished bool      `gorm:"default:true;notNull" json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
