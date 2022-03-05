package main

import (
	database "github.com/anousoneFS/golang-jwt/databases"
	route "github.com/anousoneFS/golang-jwt/routes"
	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	// allow user credentail like: http authentication, cookie
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	route.Setup(app)
	app.Listen(":8000")
}
