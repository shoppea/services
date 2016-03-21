package throw
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHealthz(c *gin.Context) {
	c.JSON(http.StatusInternalServerError,gin.H{
		"Status": "Server went down",
	})
}
