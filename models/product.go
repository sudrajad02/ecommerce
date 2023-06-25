package models

type Product struct {
	ProductId    int    `gorm:"primary_key;auto_increment" json:"id"`
	ProductName  string `gorm:"type:varchar(64);null" json:"name"`
	ProductColor string `gorm:"type:varchar(8);null" json:"color"`
	ProductSize  string `gorm:"type:varchar(4);null" json:"size"`
	ProductPrice string `gorm:"type:float;DEFAULT 0" json:"price"`
	ProductStock string `gorm:"type:int(4);DEFAULT 0" json:"stock"`
	CategoryId   string `gorm:"type:int;foreignKey:CategoryId" json:"category_id"`
	CreatedAt    string `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt    string `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt    string `gorm:"type:timestamp;default:null" json:"deleted_at"`
}
