package Model

import "time"

type Chat struct {
	ID        uint      `gorm:"notNull" json:"id"`
	ChatId    string    `gorm:"size:255;notNull;primaryKey" json:"chat_id"`
	UserID    string    `gorm:"notNull;size:255" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Message   []Message `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
