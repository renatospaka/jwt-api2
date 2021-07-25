package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	//"github.com/renatospaka/go-jwt/database"
	"github.com/renatospaka/go-jwt/routes"
)

func main() {
	//database.Connect()
	app := fiber.New()
	route.Setup(app)
	//createToken()

	log.Fatal(app.Listen(":8000"))
}

func createToken() {
	log.Println("creating token")
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}
	log.Printf("ExpiresAt: %T %v\n", claims.ExpiresAt, claims.ExpiresAt)
	log.Printf("Issuer: %T %v\n", claims.Issuer, claims.Issuer)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Printf("%v %v", ss, err)
}