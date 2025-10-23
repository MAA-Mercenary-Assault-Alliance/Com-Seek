package models

import "time"

type CompanyResponse struct {
	UserID        uint `gorm:"primaryKey"`
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"-"`
	CreatedAt     time.Time
	Name          string
	Website       string
	ContactEmail  string
	ContactNumber string
	Location      string
	Description   string
	Approved      bool
	Jobs          []Job `gorm:"foreignKey:CompanyID" json:"-"`
}
