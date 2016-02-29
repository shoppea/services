package db

import (
	_ "github.com/go-sql-driver/mysql"
	_ "sync"
	"github.com/jinzhu/gorm"
	_ "product"
	"fmt"
)

var DBContext *gorm.DB

func SharedConnection() {
	if DBContext == nil {
		DBContext,err := gorm.Open("mysql", "root:root@tcp(docker.me:3306)/shoppea?charset=utf8&parseTime=True&loc=Local")
		DBContext.LogMode(true)
		if err != nil {
			fmt.Println(err)
		}
	}
}

