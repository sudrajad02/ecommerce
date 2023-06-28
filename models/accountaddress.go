package models

type AccountAddress struct {
	AccountAddressId        int    `gorm:"primary_key;auto_increment" json:"id"`
	AccountAddressAccountId int    `gorm:"type:int;null;" json:"account_id"`
	AccountAddressLocation  string `gorm:"type:text;null;" json:"address"`
	AccountAddressIsActive  string `gorm:"type:tinyint(1);default:0;" json:"is_active"`
	CreatedAt               string `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt               string `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt               string `gorm:"type:timestamp;default:null" json:"deleted_at"`
}
