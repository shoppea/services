package product

import (
	_ "github.com/jinzhu/gorm"
	"common"
	"github.com/jinzhu/gorm"
)

type Feature struct {
	gorm.Model `gorm:"embedded"`
	Name string
	Description string
}