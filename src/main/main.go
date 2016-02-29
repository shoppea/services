package main
import (
	"db"
	"github.com/gin-gonic/gin"
	"api"
)

func main() {
	db.SharedConnection()
	gin := gin.Default()
	gin.POST("/user",api.CreateUser)
	gin.Run(":8080")
}
