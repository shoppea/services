package common

import (
	"github.com/jinzhu/gorm"
//	log "github.com/Sirupsen/logrus"
	"time"
	"github.com/gin-gonic/gin"
	"throw"
	"db"
)

type DBImplementer interface {
	AddEntity() error
	DeleteEntity() error
}

// Base model for each individual unit
type Model struct {
	gorm.Model
	Name string `json:"name"`
}

// Getter method for getting name
func (d *Model) GetName() string{
	return d.Name
}

// gorm Callback to append create and update date
func (d *Model) BeforeSave() (err error) {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return
}

// gorm Callback to modify update date
func (d *Model) BeforeUpdate() (err error) {
	d.UpdatedAt = time.Now()
	return
}

// Accept any data model response from request and add it to database
func (d *Model) CreateDBEntry(c *gin.Context) (err error){
	err = c.BindJSON(&d)
	if err != nil {
		throw.ErrorBadRequest(c)
	}
	err = db.DBConnection.Create(d).Error
	return
}

