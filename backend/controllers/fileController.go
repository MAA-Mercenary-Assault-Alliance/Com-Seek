package controllers

import (
	"com-seek/backend/models"
	"fmt"
	"image"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
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
	extension, err := ExtractExtension(fileHeader.Filename)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if fileHeader.Size > int64(fc.MaxFileSize) {
		return nil, fmt.Errorf("file size exceeds the limit of %d", fc.MaxFileSize)
	}

	img, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open image file: %w", err)
	}

	defer func() { _ = img.Close() }()

	config, _, err := image.DecodeConfig(img)
	if err != nil {
		return nil, fmt.Errorf("invalid image file")
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
		UserID:    userID,
		Extension: extension,
		Category:  fileCategory,
	}

	if err := db.Create(fileRecord).Error; err != nil {
		return nil, fmt.Errorf("failed to create file record: %w", err)
	}

	filePath := filepath.Join(fc.SavePath, fileRecord.ID)

	if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
		db.Delete(fileRecord)
		return nil, fmt.Errorf("failed to save file to local: %w", err)
	}

	return fileRecord, nil
}

func ExtractExtension(filename string) (models.FileExtension, error) {
	extension := strings.ToLower(filepath.Ext(filename))
	if extension == "" {
		return "", fmt.Errorf("no file extension found")
	}
	extension = extension[1:]

	switch extension {
	case "jpg":
		return models.FileExtensionJPG, nil
	case "jpeg":
		return models.FileExtensionJPEG, nil
	case "png":
		return models.FileExtensionPNG, nil
	case "webp":
		return models.FileExtensionWEBP, nil
	default:
		return "", fmt.Errorf("unsupported file: %s", extension)
	}
}
