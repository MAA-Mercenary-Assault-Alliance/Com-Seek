package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID					uint	`gorm:"primaryKey"`
	Title       string
	CompanyID		uint	`gorm:"index`
  Company     Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CompanyID;references:UserID"` // belongs to Company
	Description string
	Location    string
	Salary      uint
	Description	string
	CreatedAt		time.Time
	Approved		bool
}
