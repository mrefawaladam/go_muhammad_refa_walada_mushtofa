package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model

	CategoryName string `json:"category_name" form:"category_name" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
}
