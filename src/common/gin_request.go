package common

import (
	"github.com/gin-gonic/gin"
	"db"
	"net/http"
)

// Accepts gin context of receiving request
func ParseRequestJSON(c *gin.Context,obj interface{}) (err error) {
	err = c.BindJSON(obj)
	return
}

// Insert Struct into database
func DBInsert(obj interface{}) (err error){
	conn := db.SharedConnection()
	return conn.Create(obj).Error
}

func Success(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message": "OK",
	})
}