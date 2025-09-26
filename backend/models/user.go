package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email			string	`gorm:"type:varchar(255);unique;not null"`
	Password	string	`gorm:"not null"`
}

type Admin struct {
	gorm.Model
	UserID		uint
	User			User
}