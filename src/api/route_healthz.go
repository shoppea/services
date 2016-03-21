package api
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"throw"
	"db"
)

func Healthz(c *gin.Context) {
	db := db.DBConnection
	if db.Error != nil {
		throw.ErrorHealthz(c)
		return
	}
	if db.DB().Ping() != nil {
		throw.ErrorHealthz(c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status" : "ok",
	})
}
