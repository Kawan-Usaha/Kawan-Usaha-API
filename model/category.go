package Model

import (
	"time"
)

type Category struct {
	//ID: Autoincrement of ID for Category, primary key.
	//Title: Title of the Category.
	//Image: Image of the Category (URL).
	//Tags: Tags of the Category.
	//Articles: Articles the Category has.
	ID        uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	Title     string    `gorm:"notNull;size:255" json:"title"`
	Image     string    `gorm:"null;size:255" json:"image"`
	Tags      []Tag     `gorm:"null;many2many:category_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags"`
	Articles  []Article `gorm:"null;many2many:article_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"articles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
