package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	MoneyBucketID int
	Bucket        MoneyBucket `gorm:"foreignKey:MoneyBucketID"`
	Ammount       decimal.Decimal
	CategoryID    int
	Category      Category `gorm:"foreignKey:CategoryID"`
	Description   string
}
