// ROUTE /categories
package api

import (
	"github.com/gin-gonic/gin"
	"product/category"
	"db"
	"common"
)

// HANDLER<POST> : SERVICE/PRODUCT
// Create Category
func CreateCategory(c *gin.Context) {
	var cat category.Category
	common.ParseRequestJSON(c,&cat)
	db.DBInsert(&cat)
}

// HANDLER<POST> : SERVICE/PRODUCT
// Create SubCategory
func CreateSubCategory(c *gin.Context) {
	var subcat category.SubCategory
	common.ParseRequestJSON(c,&subcat)
	db.DBInsert(&subcat)
}