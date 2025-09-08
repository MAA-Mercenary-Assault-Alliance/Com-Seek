package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email			string	`json:"email" gorm:"type:varchar(255);uniqueIndex:idx_email,priority:1,length:191;not null"`
	Password	string	`json:"password";gorm:"not null"`
}

type Admin struct {
	UserID		uint
	User			User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}