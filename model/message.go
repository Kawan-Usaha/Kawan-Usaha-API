package Model

import "time"

type Message struct {
	//ID: Autoincrement of ID for Message, primary key.
	//ChatId: Foreign key of Chat.
	//UserId: Foreign key of User.
	//Message: Message of the Chat.
	ID        uint      `gorm:"notNull;autoIncrement;primaryKey" json:"id"`
	ChatId    string    `gorm:"size:255;notNull" json:"chat_id"`
	UserId    string    `gorm:"notNull;size:255" json:"user_id"`
	Message   string    `gorm:"notNull;size:255" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
