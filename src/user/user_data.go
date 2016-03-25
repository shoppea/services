package user

import (
	_ "github.com/jinzhu/gorm"
	"common"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName string `gorm:"column:last_name"`
	Contact string `gorm:"column:contact_no"`
}
