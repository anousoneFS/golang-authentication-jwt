package main

import (
	database "github.com/anousoneFS/golang-jwt/databases"
	route "github.com/anousoneFS/golang-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	route.Setup(app)
	app.Listen(":8000")
}
