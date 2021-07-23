package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	err      error
	DSN      string
	user     string
	pwd      string
	database string
)

func init() {
	godotenv.Load()
	user = strings.Trim(os.Getenv("MYSQL_USR"), " ")
	pwd = strings.Trim(os.Getenv("MYSQL_PWD"), " ")
	database = strings.Trim(os.Getenv("MYSQL_DATABASE"), " ")
	DSN = user + ":" + pwd + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func main() {
	_, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Cannot connect to Database")
	}
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx) error {
			return c.SendString("Hello, JWT over Fiber in Golang!")
	})

	log.Fatal(app.Listen(":8000"))
}