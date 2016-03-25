package gender

import (
	"github.com/jinzhu/gorm"
	"product/category"
)

type GenderCategory struct {
	gorm.Model
	Gender Gender
	GenderId int
	SubCategory category.SubCategory
	SubCategoryId int
}
