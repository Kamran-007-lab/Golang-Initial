package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string
	IsAdmin  bool
	Teams    []Team `gorm:"many2many:user_teams;"`
}
