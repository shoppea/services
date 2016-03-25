package category

import (
	"common"
	"github.com/jinzhu/gorm"
)

type SubCategory struct {
	CategoryID int
	gorm.Model
	Name string
	common.DBImplementer
}

func (s *SubCategory ) AddEntity() error {
	return nil
}

func (s *SubCategory ) DeleteEntity() error {
	return nil
}

