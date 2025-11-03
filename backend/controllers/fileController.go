package controllers

import (
	"bytes"
	"com-seek/backend/models"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "golang.org/x/image/webp"
	"gorm.io/gorm"
)

type FileController struct {
	SavePath         string
	MaxFileSize      int
	MaxProfileHeight int
	MaxProfileWidth  int
	MaxCoverHeight   int
	MaxCoverWidth    int
}

func NewFileController(
	savePath string,
	maxFileSize int,
	maxProfileHeight int,
	maxProfileWidth int,
	maxCoverHeight int,
	maxCoverWidth int) *FileController {
	return &FileController{
		SavePath:         savePath,
		MaxFileSize:      maxFileSize,
		MaxProfileHeight: maxProfileHeight,
		MaxProfileWidth:  maxProfileWidth,
		MaxCoverHeight:   maxCoverHeight,
		MaxCoverWidth:    maxCoverWidth,
	}
}

func (fc *FileController) SaveImage(
	c *gin.Context,
	db *gorm.DB,
	userID uint,
	fileHeader *multipart.FileHeader,
	fileCategory models.FileCategory) (*models.File, error) {

	if fileHeader.Size > int64(fc.MaxFileSize) {
		return nil, fmt.Errorf("file size exceeds the limit of %d", fc.MaxFileSize)
	}

	img, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open image file: %w", err)
	}

	defer img.Close()

	fileBytes, err := io.ReadAll(img)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	extension, config, err := CheckAndGetConfigFromBytes(fileBytes)
	if err != nil {
		return nil, err
	}

	if (fileCategory == models.FileCategoryProfile) &&
		(config.Height > fc.MaxProfileHeight || config.Width > fc.MaxProfileWidth) {
		return nil, fmt.Errorf("image dimensions exceed the limit of %dx%d pixels", fc.MaxProfileWidth, fc.MaxProfileHeight)
	}

	if (fileCategory == models.FileCategoryCover) &&
		(config.Height > fc.MaxCoverHeight || config.Width > fc.MaxCoverWidth) {
		return nil, fmt.Errorf("image dimensions exceed the limit of %dx%d pixels", fc.MaxCoverWidth, fc.MaxCoverHeight)
	}

	fileRecord := &models.File{
		ID:        uuid.NewString(),
		UserID:    userID,
		Extension: extension,
		Category:  fileCategory,
	}

	if err := db.Create(fileRecord).Error; err != nil {
		return nil, fmt.Errorf("failed to create file record: %w", err)
	}

	if err := os.MkdirAll(fc.SavePath, 0o755); err != nil {
		db.Delete(fileRecord)
		return nil, fmt.Errorf("failed to create save directory %s: %w", fc.SavePath, err)
	}

	filePath := filepath.Join(fc.SavePath, fileRecord.ID)

	if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
		db.Delete(fileRecord)
		return nil, fmt.Errorf("failed to save file to local: %w", err)
	}

	return fileRecord, nil
}

func CheckAndGetConfigFromBytes(fileBytes []byte) (models.FileExtension, image.Config, error) {
	mimeType := mimetype.Detect(fileBytes).String()

	var extension models.FileExtension
	switch mimeType {
	case "image/jpeg":
		extension = models.FileExtensionJPG
	case "image/png":
		extension = models.FileExtensionPNG
	case "image/webp":
		extension = models.FileExtensionWEBP
	default:
		return "", image.Config{}, fmt.Errorf("unsupported file type: %s", mimeType)
	}

	reader := bytes.NewReader(fileBytes)
	config, format, err := image.DecodeConfig(reader) // Capture the format

	if err != nil {
		// Log or print the format variable to see if it was recognized at all
		fmt.Printf("DecodeConfig failed. Detected format was: %s. Error: %v\n", format, err)
		return "", image.Config{}, fmt.Errorf("invalid image file: %w", err)
	}

	return extension, config, nil
}
