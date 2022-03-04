package route

import (
	"github.com/anousoneFS/golang-jwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
