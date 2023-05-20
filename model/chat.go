package Model

import "time"

type Chat struct {
	//ID: Autoincrement of ID for Chat, primary key.
	//ChatId: ID of the Chat.
	//UserId: Foreign key of User.
	//Messages: Messages of the Chat.
	ID        uint      `gorm:"notNull;autoIncrement" json:"id"`
	ChatId    string    `gorm:"size:255;notNull;primaryKey" json:"chat_id"`
	UserID    string    `gorm:"notNull;size:255;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Messages  []Message `gorm:"foreignKey:ChatId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
