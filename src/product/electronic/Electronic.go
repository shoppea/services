package electronic

import (
	"product"
	"github.com/jinzhu/gorm"
	"product/category"
)

type ElectronicProduct struct {
	product.Product `gorm:"-"`
	gorm.Model
	Category category.SubCategory
	CategoryId int
	Price float64
	Description string
	Status string
}