package models

type Checkout struct {
	CheckoutId               int    `gorm:"primary_key;auto_increment" json:"id"`
	CheckoutAccountId        int    `gorm:"type:int;null;" json:"account_id"`
	CheckoutAccountAddressId int    `gorm:"type:int;null;" json:"address_id"`
	CheckoutAccountProductId int    `gorm:"type:int;null;" json:"product_id"`
	CheckoutAmountProduct    int    `gorm:"type:int;default:0;" json:"amount"`
	CheckoutTotalPrice       int    `gorm:"type:float;default:0;" json:"total_price"`
	CheckoutStatusDelivery   string `gorm:"type:varchar(16);default:Belum Di Bayar;" json:"status_delivery"`
	CreatedAt                string `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt                string `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt                string `gorm:"type:timestamp;default:null" json:"deleted_at"`
}
