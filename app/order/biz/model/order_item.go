package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32  `gorm:"type:int(11)"`
	OrderIdRefer string  `gorm:"type:int(11)"`
	Quantity     uint32  `gorm:"type:int(11)"`
	Cost         float32 `gorm:"type:decimal(10,2)"`
}
