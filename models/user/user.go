package models

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Password []byte
	Email    string `gorm:"unique"`
}
