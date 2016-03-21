package api

import (
	"github.com/gin-gonic/gin"
	"product"
	"github.com/Sirupsen/logrus"
	"db"
)

// HANDLER<POST> : SERVICE/PRODUCT
// Create Product
func CreateProduct(c *gin.Context)  {
	var product product.Product
	product.AddEntity()
}

// HANDLER<GET> : SERVICE/PRODUCT
// Getting list of products
func GetProducts(c *gin.Context)  {
	var Products []product.Product
	database := db.DBConnection
	database.Find(&Products)
	logrus.Info("product: ", len(Products))
	c.JSON(200,Products)
}

// HANDLER<DELETE> : SERVICE/PRODUCT
// Delete product by id from DB
func DeleteProduct(c *gin.Context){
	db.DBConnection.Delete(c.BindJSON(&product.Product{}))
}

