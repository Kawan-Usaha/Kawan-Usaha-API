package Model

import "time"

type Chat struct {
	ID        uint      `gorm:"notNull" json:"id"`
	ChatId    string    `gorm:"size:255;notNull;primaryKey" json:"chat_id"`
	UserId    string    `gorm:"notNull;size:255" json:"user_id"`
	CreatedAt time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
