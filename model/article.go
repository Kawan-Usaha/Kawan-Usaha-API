package Model

import (
	"time"
)

type Article struct {
	ID          uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	UserID      string    `gorm:"notNull;size:255;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Title       string    `gorm:"notNull;size:255" json:"title"`
	Content     string    `gorm:"type:text;null" json:"content"`
	Image       string    `gorm:"null;size:255" json:"image"`
	IsPublished bool      `gorm:"default:true;notNull" json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
