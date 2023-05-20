package Model

import "time"

type Verification struct {
	Id               uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	UserId           string    `gorm:"notNull;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:255" json:"user_id"`
	VerificationCode string    `gorm:"size:255;notNull" json:"verification_code"`
	UsedCode         bool      `gorm:"default:false;notNull" json:"used_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
