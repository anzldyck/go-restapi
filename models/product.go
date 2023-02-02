package models

type Product struct {
	ProductId int64 `gorm:"primaryKey" json:"product_id"`
	ProductName string `gorm:"type:varchar(50)" json:"product_name"`
	ProductDescription string `gorm:"type:text" json:"product_description"`
	CreatedAt dateTime `gorm:"type:dateTime" json:"created_at"`
}