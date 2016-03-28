package api

import (
	"github.com/gin-gonic/gin"
      	"common"
	"product/gender"
	"db"
)

func CreateGender(c *gin.Context) {
	var g gender.Gender
	common.ParseRequestJSON(c,g)
	db.DBInsert(g)
	common.Success(c)
}
