package api

import (
	"github.com/gin-gonic/gin"
 	"user"
)

// create use route handler
func CreateUser(c *gin.Context)  {
	var u user.User
	u.CreateDBEntry(c)
}