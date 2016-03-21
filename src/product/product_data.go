package product

import (
	_ "github.com/jinzhu/gorm"
	"common"
)

type Feature struct {
	common.Model `gorm:"embedded"`
	Description string
}