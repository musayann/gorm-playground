package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price uint64
}
