package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Users []User `gorm:"many2many:user_teams;"`
}
