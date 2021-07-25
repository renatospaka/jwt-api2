package models

type User struct {
	ID       uint   `json:"user_id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password []byte `json:"-"`
	Email    string `json:"email" gorm:"unique"`
}
