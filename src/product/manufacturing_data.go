package product
import (
	"github.com/jinzhu/gorm"
)

type Manufacturer struct {
	gorm.Model
	Name string
}
