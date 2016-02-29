package seller


import (
	_ "github.com/jinzhu/gorm"
)

type Seller struct {
	Name string
	Contact string
}
