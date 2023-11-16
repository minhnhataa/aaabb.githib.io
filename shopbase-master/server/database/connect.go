package database

import (
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/ecommerce?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Product{}, &models.CartItem{})
}
