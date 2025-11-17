package models

type Company struct {
	UserID         uint `gorm:"primaryKey"`
	User           User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"-"`
	Name           string
	Website        string
	ContactEmail   string
	ContactNumber  string
	Location       string
	Description    string
	Approved       bool
	ProfileImageID *string `gorm:"type:char(36);default:null" json:"profile_image_id"`
	ProfileImage   File    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProfileImageID" json:"-"`
	CoverImageID   *string `gorm:"type:char(36);default:null" json:"cover_image_id"`
	CoverImage     File    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CoverImageID" json:"-"`
	Jobs           []Job   `gorm:"foreignKey:CompanyID" json:"-"`
}
