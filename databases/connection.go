package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		panic("cloud not connect to the database")
	}
}
