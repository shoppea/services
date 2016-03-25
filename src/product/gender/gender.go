package gender

import "github.com/jinzhu/gorm"

type Gender struct {
	gorm.Model
	Name string
}
