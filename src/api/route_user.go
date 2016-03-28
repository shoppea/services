package api

import (
	"github.com/gin-gonic/gin"
 	"user"
	"common"
	"db"
)

// create use route handler
func CreateUser(c *gin.Context)  {
	var u user.User
	common.ParseRequestJSON(c,u)
	db.DBInsert(u)
	common.Success(c)
}