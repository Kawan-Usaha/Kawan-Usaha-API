package Model

import "time"

type Message struct {
	ID        uint      `gorm:"notNull" json:"id"`
	ChatId    string    `gorm:"size:255;notNull;primaryKey" json:"chat_id"`
	UserId    string    `gorm:"notNull;size:255" json:"user_id"`
	Message   string    `gorm:"notNull;size:255" json:"message"`
	CreatedAt time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
}
