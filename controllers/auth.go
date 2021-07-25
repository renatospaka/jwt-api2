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

// func init() {
// 	SecretKey = []byte("secret")
// }

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{}
	DB := database.Connect()
	DB.Where("email = ?", data["email"]).First((&user))
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "incorrect user or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect user or password",
		})
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.Itoa(int(user.ID)),
	}

	//mySigningKey := []byte("secret")
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claim.SignedString([]byte("secret"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			//"message": "login failed",
			"message": err.Error(),
		})
	}
	log.Printf("%v", token)

	// var data map[string]string

	// if err := c.BodyParser(&data); err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	// user := models.User{}
	// DB := database.Connect()
	// DB.Where("email = ?", data["email"]).First((&user))
	// if user.ID == 0 {
	// 	c.Status(fiber.StatusNotFound)
	// 	return c.JSON(fiber.Map{
	// 		//"message": "user not found",
	// 		"message": "incorrect user or password",
	// 	})
	// }

	// if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "incorrect user or password",
	// 	})
	// }

	// //issuer := strconv.Itoa(int(user.ID))
	// issuer := strconv.FormatUint(uint64(user.ID), 10)
	// claim := &jwt.StandardClaims{
	// 	Issuer:    issuer,
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// }
	// claims := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	// mySecret := []byte(SecretKey)
	// token, err := claims.SignedString(mySecret)
	// if err != nil {
	// 	c.Status(fiber.StatusInternalServerError)
	// 	return c.JSON(fiber.Map{
	// 		//"message": "login failed",
	// 		"message": err.Error(),
	// 	})
	// }

	// //return c.Status(fiber.StatusOK).JSON(token)
	// return c.JSON(token)
	return c.Status(fiber.StatusOK).JSON(token)
}
