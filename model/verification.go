package Model

import "time"

type Verification struct {
	//Id: Autoincrement of ID for Verification, primary key.
	//UserId: Foreign key of User.
	//VerificationCode: Verification Code of the User.
	//UsedCode: Is Verification Code used or not.
	Id               uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	UserID           string    `gorm:"notNull;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:255" json:"user_id"`
	VerificationCode string    `gorm:"size:255;notNull" json:"verification_code"`
	UsedCode         bool      `gorm:"default:false;notNull" json:"used_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
