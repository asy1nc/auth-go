package models

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	UserID   uint
	User     User
	Password string `json:"-"`
	Salt     string `json:"-"`
}
