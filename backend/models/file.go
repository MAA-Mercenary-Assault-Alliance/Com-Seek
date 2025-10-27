package models

import "time"

type FileExtension string

const (
	FileExtensionJPG  FileExtension = "image"
	FileExtensionJPEG FileExtension = "jpeg"
	FileExtensionPNG  FileExtension = "png"
	FileExtensionWEBP FileExtension = "webp"
)

type FileCategory string

const (
	FileCategoryProfile FileCategory = "profile"
	FileCategoryCover   FileCategory = "cover"
)

type File struct {
	ID        string `gorm:"type:uuid;primarykey;" json:"id"`
	UserID    uint
	Extension FileExtension
	Category  FileCategory
	CreatedAt time.Time
	UpdatedAt time.Time
}
