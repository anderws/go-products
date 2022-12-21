package database

import (
	"fmt"
	"github.com/anderws/go-products/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error
	servicoDb := os.Getenv("DB_HOST")
	userDb := os.Getenv("DB_USER")
	pwdDb := os.Getenv("DB_PASSWORD")
	portDd := os.Getenv("DB_PORT")

	stringUrl := userDb + ":" + pwdDb + "@tcp(" + servicoDb + ":" + portDd + ")/products?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(stringUrl)

	DB, err = gorm.Open(mysql.Open(stringUrl), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.Product{})
}
