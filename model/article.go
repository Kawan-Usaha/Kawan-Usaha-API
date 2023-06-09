package Model

import (
	"time"
)

type Article struct {
	//ID: Autoincrement of ID for Article, primary key.
	//UserId: Foreign key of User.
	//User: Foreign key of User (Many to One).
	//Title: Title of Article.
	//Content: Content of Article (Text).
	//Image: Image of Article (URL).
	//IsPublished: Is Article published or not.
	//Category: Category of Article (Many to One).
	ID          uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	UserID      string    `gorm:"notNull;size:255;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Title       string    `gorm:"notNull;size:255" json:"title"`
	Content     string    `gorm:"type:text;null" json:"content"`
	Image       string    `gorm:"null;size:255" json:"image"`
	IsPublished bool      `gorm:"notNull" json:"is_published"`
	CategoryID  uint      `gorm:"null;size:255;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
