package models

import "time"

type StudentResponse struct {
	UserID         uint `gorm:"primaryKey"`
	User           User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"-"`
	CreatedAt      time.Time
	FirstName      string
	LastName       string
	Description    string
	IsAlum         bool
	Approved       bool
	GitHub         string
	LinkedIn       string
	Facebook       string
	Instagram      string
	Twitter        string
	JobApplication []JobApplication `json:"-"`
}
