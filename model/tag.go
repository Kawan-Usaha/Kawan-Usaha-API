package Model

type Tag struct {
	ID   uint   `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	Name string `gorm:"notNull;size:255" json:"name"`
}
