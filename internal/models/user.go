package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `json:"uuid"        gorm:"unique"`
	Username string `json:"username"    gorm:"unique"`
	Email    string `json:"email"       gorm:"unique"`
}
