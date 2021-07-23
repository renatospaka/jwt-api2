package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatospaka/go-jwt/controller"
)

func Setup(app *fiber.App) {
	app.Get("/", controller.Hello)
}
