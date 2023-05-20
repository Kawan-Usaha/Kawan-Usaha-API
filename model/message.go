package Model

import "time"

type Message struct {
	ID        uint      `gorm:"notNull" json:"id"`
	ChatID    string    `gorm:"size:255;notNull;primaryKey" json:"chat_id"`
	Message   string    `gorm:"notNull;size:255" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
