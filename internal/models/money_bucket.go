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
	Color       string          `gorm:"type:VARCHAR(7);check:hex ~ '^#[0-9A-Fa-f]{6}$'"`
	Balance     decimal.Decimal
}
