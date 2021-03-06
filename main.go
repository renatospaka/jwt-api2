package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/renatospaka/go-jwt/routes"
)

func main() {
	//database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	route.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
