package models

type Account struct {
	AccountId       int    `gorm:"primary_key;auto_increment" json:"id"`
	AccountName     string `gorm:"type:varchar(64);null;" json:"name"`
	AccountEmail    string `gorm:"type:varchar(64);not null;unique" json:"email"`
	AccountUsername string `gorm:"type:varchar(16);not null;unique" json:"username"`
	AccountPassword string `gorm:"type:varchar(32);not null;" json:"password"`
	CreatedAt       string `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt       string `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt       string `gorm:"type:timestamp;default:null" json:"deleted_at"`
}
