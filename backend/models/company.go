package models

type Company struct {
	UserID			uint	`gorm:"primaryKey"`
  User 				User	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
	Name				string
	Website			string
	Location		string
	Description string
	Jobs				[]Job	`gorm:"foreignKey:CompanyID"`
}
