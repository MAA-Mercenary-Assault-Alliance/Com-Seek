package models

import (
	"time"
)

type JobType string

const (
	SoftwareDev         JobType = "Software & Application Development"
	DataAI              JobType = "Data & AI"
	CloudInfrastructure JobType = "Cloud & Infrastructure"
	Cybersecurity       JobType = "Cybersecurity"
	ProductDesign       JobType = "Product & Design"
	TestingQA           JobType = "Testing & Quality"
	Hardware            JobType = "Hardware & Electronics"
	Management          JobType = "Management & Leadership"
	Support             JobType = "IT Support & Operations"
)

type EmploymentStatus string

const (
	PartTime   EmploymentStatus = "part-time"
	FullTime   EmploymentStatus = "full-time"
	Contract   EmploymentStatus = "contract"
	Internship EmploymentStatus = "internship"
)

type Job struct {
	ID               uint `gorm:"primaryKey"`
	Title            string
	CompanyID        uint    `gorm:"index"`
	Company          Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CompanyID;references:UserID"`
	Location         string
	JobType          JobType
	EmploymentStatus EmploymentStatus
	MinSalary        uint
	MaxSalary        uint
	MinExperience    uint
	MaxExperience    uint
	Description      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Approved         bool
	Visibility       bool
	CheckNeeded      bool             `json:"-"`
	JobApplication   []JobApplication `json:"-"`
}
