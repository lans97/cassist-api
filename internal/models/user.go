package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID          string `json:"uuid"           gorm:"unique"`
	Email         string `json:"email"          gorm:"unique"`
	DisplayName   string `json:"display_name"   gorm:"null"`
	EmailVerified *bool  `json:"email_verified" gorm:"default:false"`
    RoleID        int    `json:"role_id"`
    Role          Role   `json:"role"           gorm:"foreignKey:RoleID"`
}
