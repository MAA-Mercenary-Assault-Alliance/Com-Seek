package models

type Student struct {
	UserID         uint `gorm:"primaryKey"`
	User           User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
	FirstName      string
	LastName       string
	Description    string
	IsAlum         bool
	Approved       bool
	GitHub         string
	LinkedIn       string
	Facebook       string
	Instragram     string
	Twitter        string
	JobApplication []JobApplication `gorm:"foreignKey:ID"`
}
