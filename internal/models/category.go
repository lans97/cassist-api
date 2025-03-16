package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID int
	User   User
	Name   string
	Color  string `gorm:"type:VARCHAR(7);check:hex ~ '^#[0-9A-Fa-f]{6}$'"`
}
