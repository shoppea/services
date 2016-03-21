package category

import "common"

type SubCategory struct {
	CategoryID int
	common.Model
	common.DBImplementer
}

func (s *SubCategory ) AddEntity() error {
	return nil
}

func (s *SubCategory ) DeleteEntity() error {
	return nil
}

