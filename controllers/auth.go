package controller

import (
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/renatospaka/go-jwt/database"
	models "github.com/renatospaka/go-jwt/models/user"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	DB := database.Connect()
	DB.Create(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{}
	DB := database.Connect()
	DB.Where("email = ?", data["email"]).First((&user))
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			//"message": "user not found",
			"message": "incorrect user or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect user or password",
		})
	}
	//issuer := strconv.Itoa(int(user.ID))
	issuer := strconv.FormatUint(uint64(user.ID), 10)
	log.Printf("issuer: %s\n", issuer)
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Issuer:    "5",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	log.Printf("claims: typeof(%T) %v\n", claims, claims)

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			//"message": "login failed",
			"message": err.Error(),
		})
	}

	//return c.Status(fiber.StatusOK).JSON(token)
	return c.JSON(token)
}
