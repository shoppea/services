// ROUTE /categories
package api

import (
	"github.com/gin-gonic/gin"
	"product/category"
)

// HANDLER<POST> : SERVICE/PRODUCT
// Create Category
func CreateCategory(c *gin.Context) {
	var cat category.Category
	cat.CreateDBEntry(c)
}

// HANDLER<POST> : SERVICE/PRODUCT
// Create SubCategory
func CreateSubCategory(c *gin.Context) {
	var subcat category.SubCategory
	subcat.CreateDBEntry(c)
}