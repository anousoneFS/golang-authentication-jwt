package database

import (
	"github.com/anousoneFS/golang-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		panic("cloud not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
