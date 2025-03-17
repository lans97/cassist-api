package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MoneyBucket struct {
	gorm.Model
	UserID      int
	User        User            `gorm:"foreignKey:UserID"`
	Name        string
	Color       string          `gorm:"type:VARCHAR(7)"`
	Balance     decimal.Decimal
}
