package product

import (
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
)

type Feature struct {
	gorm.Model `gorm:"embedded"`
	Name string
	Description string
}