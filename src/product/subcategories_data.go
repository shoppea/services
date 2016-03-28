package product

import (
	"github.com/jinzhu/gorm"
	"product/category"
)

type SubCategory struct {
	gorm.Model
	name string
	Category category.Category
	CategoryID int
}
