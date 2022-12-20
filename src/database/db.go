package database

import (
	"github.com/anderws/go-products/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/products?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}	
}

func AutoMigrate() {
	DB.AutoMigrate(models.Product{})
}