package helpers

import (
	"bytes"
	"com-seek/backend/config"
	"com-seek/backend/models"
	"fmt"
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gabriel-vasile/mimetype"
	_ "golang.org/x/image/webp"
	"gorm.io/gorm"
)

var supportedImageExtension = map[models.FileExtension]bool{
	models.FileExtensionJPG:  true,
	models.FileExtensionPNG:  true,
	models.FileExtensionWEBP: true,
}

var supportedDocExtension = map[models.FileExtension]bool{
	models.FileExtensionPDF: true,
}

func DeleteFileRecord(db *gorm.DB, fileID string) {
	var fileRecordToDelete models.File
	if err := db.First(&fileRecordToDelete, "id = ?", fileID).Error; err == nil {

		if err := db.Delete(&fileRecordToDelete).Error; err != nil {
			log.Printf("Failed to delete file record %s: %v\n", fileID, err)
		}
	}
}

func ValidateAndGetExtension(
	fileBytes []byte,
	fileCategory models.FileCategory,
	fileConfig config.FileConfig) (
	bool, models.FileExtension, error) {
	fileSize := len(fileBytes)

	if fileSize > fileConfig.MaxFileSize {
		return false, "", fmt.Errorf("file size exceeds the limit of %d", fileConfig.MaxFileSize)
	}

	mimeType := mimetype.Detect(fileBytes).String()

	var extension models.FileExtension
	switch mimeType {
	case "image/jpeg":
		extension = models.FileExtensionJPG
	case "image/png":
		extension = models.FileExtensionPNG
	case "image/webp":
		extension = models.FileExtensionWEBP
	case "application/pdf", "application/x-pdf":
		extension = models.FileExtensionPDF
	default:
		return false, "", fmt.Errorf("unsupported file type: %s", mimeType)
	}

	switch fileCategory {
	case models.FileCategoryTranscript:
		if !supportedDocExtension[extension] {
			return false, "", fmt.Errorf("unsupported document file type: %s", extension)
		}
	case models.FileCategoryProfile, models.FileCategoryCover:
		if !supportedImageExtension[extension] {
			return false, "", fmt.Errorf("unsupported image file type: %s", extension)
		}

		reader := bytes.NewReader(fileBytes)
		config, _, err := image.DecodeConfig(reader)

		if err != nil {
			return false, "", fmt.Errorf("invalid image file: %w", err)
		}

		switch fileCategory {
		case models.FileCategoryProfile:
			if config.Height > fileConfig.MaxProfileHeight || config.Width > fileConfig.MaxProfileWidth {
				return false, "", fmt.Errorf(
					"image dimensions exceed the limit of %dx%d pixels",
					fileConfig.MaxProfileWidth,
					fileConfig.MaxProfileWidth)
			}
		case models.FileCategoryCover:
			if config.Height > fileConfig.MaxCoverHeight || config.Width > fileConfig.MaxCoverWidth {
				return false, "", fmt.Errorf(
					"image dimensions exceed the limit of %dx%d pixels",
					fileConfig.MaxCoverWidth,
					fileConfig.MaxCoverHeight)
			}
		}
	}

	return true, extension, nil
}
