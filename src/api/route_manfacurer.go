package api
import (
	"github.com/gin-gonic/gin"
	"product"
)

func CreateManufacturer(c *gin.Context) {
	var m product.Manufacturer
	m.CreateDBEntry(c)
}
