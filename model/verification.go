package Model

import "time"

type Verification struct {
	Id               uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	UserId           string    `gorm:"notNull;size:255" json:"user_id"`
	VerificationCode string    `gorm:"size:255;notNull" json:"verification_code"`
	CreatedAt        time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt        time.Time `gorm:"type:timestamp" json:"updated_at"`
}
