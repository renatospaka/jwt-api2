package database

import (
	"log"
	"os"
	"strings"

	models "github.com/renatospaka/go-jwt/models/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn      string
	user     string
	pwd      string
	database string
	DB       *gorm.DB
)

func init() {
	godotenv.Load()
	user = strings.Trim(os.Getenv("MYSQL_USR"), " ")
	pwd = strings.Trim(os.Getenv("MYSQL_PWD"), " ")
	database = strings.Trim(os.Getenv("MYSQL_DATABASE"), " ")
	dsn = user + ":" + pwd + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func Connect() *gorm.DB {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Println("Cannot connect to Database")
		return nil
	}

	connection.AutoMigrate(&models.User{})
	return connection
}
