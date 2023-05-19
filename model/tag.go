package Model

import "time"

type Tag struct {
	ID        uint      `gorm:"notNull;autoIncrement;primarykey" json:"id"`
	Name      string    `gorm:"notNull;size:255" json:"name"`
	CreatedAt time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
