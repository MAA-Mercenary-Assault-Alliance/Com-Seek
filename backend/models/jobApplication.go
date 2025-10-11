package models

import "time"

type JobApplication struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	Student   Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:StudentID;references:UserID"`
	JobID     uint
	Job       Job `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:JobID;references:ID"`
	CreatedAt time.Time
}
