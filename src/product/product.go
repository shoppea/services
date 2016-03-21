package product

import (
	"common"
	_"db"
)

type Product struct {
	common.Model `gorm:"embedded"`
	Features []Feature
	SubCategoryID SubCategory      // Product has one subcategory
}

func (p *Product ) AddProduct() error{
	return nil
}

func (p *Product ) AddEntity() error {
	return nil
}

