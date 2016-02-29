package api

import (
	"github.com/gin-gonic/gin"
	"os/user"
	"db"
	log "github.com/Sirupsen/logrus"
	"error"
)

func CreateUser(c *gin.Context)  {
	var user user.User
	if c.BindJSON(&user) == nil{
		db := db.DBContext.Create(user)
		if db.Error != nil {
			log.Error(db.Error)
			error.ErrorBadRequest(c)
		}
	}else{
		error.ErrorBadRequest(c)
	}
}
