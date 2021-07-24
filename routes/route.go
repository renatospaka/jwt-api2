package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatospaka/go-jwt/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
