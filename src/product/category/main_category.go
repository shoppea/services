package category

import (
	"common"
	"github.com/jinzhu/gorm"
)

type Category struct {
	common.DBImplementer
	Name string
	gorm.Model `gorm:"embedded"`
}

func (c *Category ) AddEntity() error{
	return nil
}

func (c *Category ) DeleteEntity() error{
	return nil
}