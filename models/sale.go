package models

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ProductID uint    `gorm:"column:product_id"`
	Product   Product `gorm:"foreignKey:ID;references:ProductID"`
	Quantity  uint
	Total     uint
	
}

