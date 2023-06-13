package Model

import (
	"time"
)

type User struct {
	//ID: Autoincrement ID user.
	//UserId: Primary key user, uses uuid.
	//Name: Name of user.
	//Email: Email of user.
	//Password: Password of user. Saved as hash.
	//Usaha: Usaha/UMKM of user.
	//Article: Articles made by user.
	//Chat: Chats made by user.
	//Verified: Is user verified or not.
	//Verification: Verification code of user.
	//Role: Role of user. 0: User, 1: Admin (WIP).
	ID               uint         `gorm:"notNull;autoIncrement" json:"id"`
	UserId           string       `gorm:"notNull;size:255;primaryKey;column:user_id" json:"user_id"`
	Name             string       `gorm:"notNull;size:255" json:"name"`
	Email            string       `gorm:"size:255;notNull;uniqueIndex" json:"email"`
	Password         string       `gorm:"notNull;size:255" json:"password"`
	Usaha            []Usaha      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"usaha"`
	Article          []Article    `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"article"`
	Verified         bool         `gorm:"notNull;default:false" json:"verified"`
	Verification     Verification `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"verification"`
	FavoriteArticles []Article    `gorm:"many2many:userfavoritearticles;" json:"favorite_articles"`
	RoleId           uint         `gorm:"notNull;default:0" json:"role_id"`
	Image            string       `gorm:"null;size:255" json:"image"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
}

type UserFavoriteArticle struct {
	UserID    string
	ArticleID uint
}

type ChangePassword struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
