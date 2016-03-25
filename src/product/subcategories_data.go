package product

import (
	"common"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"product/category"
)

type SubCategory struct {
	gorm.Model
	name string
	Category category.Category
	CategoryID int
}
