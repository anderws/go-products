package database

import (
	"fmt"
	"github.com/anderws/go-products/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error
	hostDb := os.Getenv("DB_HOST")
	userDb := os.Getenv("DB_USER")
	pwdDb := os.Getenv("DB_PASSWORD")
	nameDb := os.Getenv("DB_NAME")
	dsn := "host="+hostDb+" user="+userDb+" password="+pwdDb+" dbname="+nameDb+" port=5432 sslmode=disable TimeZone=Asia/Shanghai"

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.Product{})
}
