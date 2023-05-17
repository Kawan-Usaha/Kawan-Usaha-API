package Model

type Verification struct {
	VerificationCode string `gorm:"size:255;notNull" json:"verification_code"`
	UserId           string `gorm:"notNull;size:255,primaryKey" json:"user_id"`
	User             User   `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,OnDelete:SET NULL;"`
}
