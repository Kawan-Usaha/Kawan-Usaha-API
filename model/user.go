package Model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"notNull" json:"id"`
	UserId    string    `gorm:"notNull;size:255,primaryKey" json:"user_id"`
	Name      string    `gorm:"notNull;size:255" json:"name"`
	Username  string    `gorm:"size:255;notNull;uniqueIndex" json:"username"`
	Email     string    `gorm:"size:255;notNull;uniqueIndex" json:"email"`
	Password  string    `gorm:"notNull;size:255" json:"password"`
	UsahaID   uint      `gorm:"notNull" json:"usaha_id"`
	Usaha     Usaha     `gorm:"foreignKey:UsahaId" json:"usaha" constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChangePassword struct {
	Password string `gorm:"notNull;size:255" json:"password"`
	Newpass1 string `gorm:"notNull;size:255" json:"newpass1"`
	Newpass2 string `gorm:"notNull;size:255" json:"newpass2"`
}
