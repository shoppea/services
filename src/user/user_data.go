package user

import (
	_ "github.com/jinzhu/gorm"
	"common"
)

type User struct {
	common.Model
	FirstName string `gorm:"column:first_name"`
	LastName string `gorm:"column:last_name"`
	Contact string `gorm:"column:contact_no"`
}
