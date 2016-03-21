package category

import "common"

type Category struct {
	common.DBImplementer
	common.Model `gorm:"embedded"`
}

func (c *Category ) AddEntity() error{
	return nil
}

func (c *Category ) DeleteEntity() error{
	return nil
}