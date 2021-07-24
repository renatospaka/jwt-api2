package models

type User struct {
	ID       string
	Name     string
	Password []byte
	Email    string `gorm:"unique"`
}
