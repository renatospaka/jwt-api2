package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/renatospaka/go-jwt/database"
	"github.com/renatospaka/go-jwt/route"
)

func main() {
	database.Connect()
	app := fiber.New()
	route.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
