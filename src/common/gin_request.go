package common

import "github.com/gin-gonic/gin"

// Accepts gin context of receiving request
func ParseRequestJSON(c *gin.Context,obj interface{}) (err error) {
	err = c.BindJSON(obj)
	return
}
