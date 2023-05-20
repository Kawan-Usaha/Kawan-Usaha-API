package Model

import (
	"time"
)

type User struct {
	ID           uint         `gorm:"notNull;autoIncrement" json:"id"`
	UserId       string       `gorm:"notNull;size:255;primaryKey" json:"user_id"`
	Name         string       `gorm:"notNull;size:255" json:"name"`
	Email        string       `gorm:"size:255;notNull;uniqueIndex" json:"email"`
	Password     string       `gorm:"notNull;size:255" json:"password"`
	Usaha        []Usaha      `gorm:"Null;foreignkey:User;association_foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Article      []Article    `gorm:"Null;foreignkey:UserId;association_foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"article"`
	Verified     bool         `gorm:"notNull;default:false" json:"verified"`
	Verification Verification `gorm:"Null;foreignKey:UserId;association_foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"verification"`
	RoleId       uint         `gorm:"notNull;default:0" json:"role_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type ChangePassword struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
