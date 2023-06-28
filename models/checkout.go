package models

type Checkout struct {
	CheckoutId               int            `gorm:"primary_key;auto_increment" json:"id"`
	CheckoutAccountId        int            `gorm:"type:int;null;" json:"account_id"`
	Account                  Account        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CheckoutAccountId;references:AccountId" json:"account"`
	CheckoutAccountAddressId int            `gorm:"type:int;null;" json:"address_id"`
	AccountAddress           AccountAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CheckoutAccountAddressId;references:AccountAddressId" json:"account_address"`
	CheckoutAccountProductId int            `gorm:"type:int;null;" json:"product_id"`
	Product                  Product        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CheckoutAccountProductId;references:ProductId" json:"product"`
	CheckoutAmountProduct    int            `gorm:"type:int;default:0;" json:"amount"`
	CheckoutTotalPrice       int            `gorm:"type:float;default:0;" json:"total_price"`
	CheckoutStatusDelivery   string         `gorm:"type:varchar(128);default:Belum Di Bayar;" json:"status_delivery"`
	CreatedAt                string         `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt                string         `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt                string         `gorm:"type:timestamp;default:null" json:"deleted_at"`
}
