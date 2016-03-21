package throw

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorBadRequest(c *gin.Context)  {
	log.Error("bad user request")
	c.JSON(http.StatusBadRequest,gin.H{
		"message": "Bad request",
	})
}

