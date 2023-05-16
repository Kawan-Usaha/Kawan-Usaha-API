package Model

type TempCode struct {
	ID    uint   `gorm:"notNull;primaryKey" json:"id"`
	Email string `gorm:"notNull;size:255" json:"email"`
	Code  string `gorm:"notNull;size:255" json:"code"`
}
