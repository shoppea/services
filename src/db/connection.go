package db

import (
	_ "github.com/go-sql-driver/mysql"
	_ "sync"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Sirupsen/logrus"
	"sync"
	//"user"
	//"product"
)

var DBConnection gorm.DB
var once sync.Once

func SharedConnection() ( gorm.DB ){
	once.Do(func() {
		var err error
		DBConnection,err = gorm.Open("mysql", "root:root@tcp(docker.me:3306)/shoppea")
		DBConnection.LogMode(true)
		//DBConnection.CreateTable(product.Product{})
		//DBConnection.CreateTable()
		logrus.Info("Getting connection from MYSql")
		if err != nil {
			logrus.Info("Error encoutered while getting connection from mysql")
			fmt.Println(err)
		}
		return
	})
	return DBConnection
}