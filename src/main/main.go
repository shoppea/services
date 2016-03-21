package main
import
(
	"db"
	"github.com/gin-gonic/gin"
	"api"
	//"fmt"
)

func init(){
	db.SharedConnection()
}

type Name struct {
	Name string
}

func main() {
	//db.DBConnection.CreateTable(category.Category{})
	gin := gin.Default()

	// ----------------- PRODUCTS -----------------
	gin.GET("/products",api.GetProducts)
	gin.POST("/products",api.CreateProduct)

	// ----------------- CATEGORY -----------------
	gin.POST("/category",api.CreateCategory)

	// ---------------- HEALTHZ -------------------
	gin.GET("/healthz",api.Healthz)

	// ---------------- RUN ON PORT ---------------
	gin.Run(":8080")
}

//func main() {
//	var a int
//	var b int
//	var c int
//	c = a + b
//	fmt.Println("Addition is : ", c)
//}