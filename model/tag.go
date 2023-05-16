package Model

type Tag struct {
	Id   uint   `gorm:"notNull" json:"id"`
	Name string `gorm:"notNull;size:255" json:"name"`
	//Category []Category `gorm:"many2many:tags_on_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"`
	//Usaha    []Usaha    `gorm:"many2many:tags_on_usaha;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"usaha"`
}
