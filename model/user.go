package Model

import (
	"time"
)

type User struct {
	ID           uint         `gorm:"notNull;autoIncrement" json:"id"`
	UserId       string       `gorm:"notNull;size:255;primaryKey;column:user_id" json:"user_id"`
	Name         string       `gorm:"notNull;size:255" json:"name"`
	Email        string       `gorm:"size:255;notNull;uniqueIndex" json:"email"`
	Password     string       `gorm:"notNull;size:255" json:"password"`
	Usaha        []Usaha      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"usaha"`
	Article      []Article    `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"article"`
	Chat         []Chat       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"chat"`
	Verified     bool         `gorm:"notNull;default:false" json:"verified"`
	Verification Verification `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"verification"`
	RoleId       uint         `gorm:"notNull;default:0" json:"role_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type ChangePassword struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
