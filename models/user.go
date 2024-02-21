package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
}
