package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Password []byte
	Email    string `gorm:"unique"`
}
