package product
import (
	"common"
	"github.com/jinzhu/gorm"
)

type Manufacturer struct {
	gorm.Model
	Name string
}
