package api
import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
	"db"
)

func CreateManufacturer(c *gin.Context) {
	var m product.Manufacturer
	common.ParseRequestJSON(c,m)
	db.DBInsert(m)
	common.Success(c)
}
