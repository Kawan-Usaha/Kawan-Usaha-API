package Model

type Tag struct {
	ID   uint   `gorm:"notNull" json:"id"`
	Name string `gorm:"notNull;size:255" json:"name"`
}
