package entity

import (
	"gorm.io/gorm"
)

type Items struct {
	*gorm.Model

	ItemName    string   `json:"" form:"item_name" validate:"required"`
	Description string   `json:"description" form:"description" validate:"required"`
	Stock       string   `json:"stock" form:"stock" validate:"required"`
	Price       string   `json:"price" form:"price" validate:"required"`
	CategoryId  string   `json:"category_id" form:"category_id" validate:"required"  gorm:"foreignKey:OrderId"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryId"`
}
