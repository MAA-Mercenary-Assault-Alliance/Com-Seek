package models

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

type FileExtension string

const (
	FileExtensionJPG  FileExtension = "jpg"
	FileExtensionPNG  FileExtension = "png"
	FileExtensionWEBP FileExtension = "webp"
)

type FileCategory string

const (
	FileCategoryProfile FileCategory = "profile"
	FileCategoryCover   FileCategory = "cover"
)

type File struct {
	ID        string        `gorm:"type:char(36);primarykey" json:"id"`
	UserID    uint          `gorm:"not null" json:"user_id"`
	Extension FileExtension `gorm:"type:varchar(10);not null" json:"extension"`
	Category  FileCategory  `gorm:"type:varchar(10);not null" json:"category"`
	CreatedAt time.Time
}

func (f *File) AfterDelete(tx *gorm.DB) (err error) {
	savePath := os.Getenv("FILE_SAVE_PATH")
	if savePath == "" {
		return fmt.Errorf("file save path not configured")
	}

	filePath := filepath.Join(savePath, f.ID)

	if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
		return nil
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("could not delete the file from storage: %w", err)
	}

	return nil
}
