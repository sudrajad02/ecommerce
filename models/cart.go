package models

type Cart struct {
	CartId        int     `gorm:"primary_key;auto_increment" json:"id"`
	CartProductId int     `gorm:"type:int;null;index:idx_cart_product,unique" json:"product_id"`
	Product       Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CartProductId;references:ProductId" json:"product"`
	CartAccountId int     `gorm:"type:int;null;index:idx_cart_product" json:"account_id"`
	Account       Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CartAccountId;references:AccountId" json:"account"`
	AmountProduct int     `gorm:"type:int;null" json:"amount"`
	TotalPrice    int     `gorm:"type:int;null" json:"price"`
	CreatedAt     string  `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt     string  `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt     string  `gorm:"type:timestamp;default:null" json:"deleted_at"`
}
