package models

type Student struct {
	UserID         uint             `gorm:"primaryKey"`
	User           User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"-"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	Description    string           `json:"description"`
	IsAlum         bool             `json:"is_alum"`
	Approved       bool             `json:"approved"`
	GitHub         string           `json:"github"`
	LinkedIn       string           `json:"linkedin"`
	Facebook       string           `json:"facebook"`
	Instagram      string           `json:"instagram"`
	Twitter        string           `json:"twitter"`
	ProfileImageID *string          `gorm:"type:char(36);default:null" json:"profile_image_id"`
	ProfileImage   File             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProfileImageID" json:"-"`
	CoverImageID   *string          `gorm:"type:char(36);default:null" json:"cover_image_id"`
	CoverImage     File             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CoverImageID" json:"-"`
	TranscriptID   string           `gorm:"type:char(36)" json:"transcript_id"`
	Transcript     File             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TranscriptID" json:"-"`
	JobApplication []JobApplication `json:"-"`
}
