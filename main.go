package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		panic("cloud not connect to the database")
	}
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello sone")
	})
	app.Listen(":8000")
}
