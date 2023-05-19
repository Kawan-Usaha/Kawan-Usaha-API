package Model

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	Title     string    `gorm:"notNull;size:255" json:"title"`
	Image     string    `gorm:"null;size:255" json:"image"`
	Tags      []Tag     `gorm:"many2many:category_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;" json:"tags"`
	Articles  []Article `gorm:"many2many:article_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;" json:"articles"`
	CreatedAt time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
