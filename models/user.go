package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Fullname  string    `json:"fullname"`
	DOB       time.Time `json:"dob"`
	Bio       string    `json:"bio"`
	Passwords []Password
}
