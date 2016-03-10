package product


import (
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name string
	CategoryID uint
}

