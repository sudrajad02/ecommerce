package models

type ProductCategory struct {
	CategoryId   int    `gorm:"primary_key;auto_increment" json:"id"`
	CategoryName string `gorm:"type:varchar(16);null" json:"name"`
	CreatedAt    string `gorm:"type:timestamp;default:null" json:"-"`
	UpdatedAt    string `gorm:"type:timestamp;default:null" json:"-"`
	DeletedAt    string `gorm:"type:timestamp;default:null" json:"-"`
}
