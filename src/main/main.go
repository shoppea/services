package main
import
(
	"db"
	"github.com/gin-gonic/gin"
	"api"
	//"fmt"
	"product/category"
	"product/gender"
	"product/electronic"
)

func init(){
  	conn :=	db.SharedConnection()
	conn.CreateTable(category.Category{})
	conn.CreateTable(category.SubCategory{})
	conn.CreateTable(gender.Gender{})
	conn.CreateTable(gender.GenderCategory{})
	conn.CreateTable(electronic.ElectronicProduct{})
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
	gin.POST("/sub_category", api.CreateSubCategory)

	// ----------------- Gender -----------------
	gin.POST("/gender",api.CreateGender)

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