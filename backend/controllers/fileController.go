package controllers

import (
	"bytes"
	"com-seek/backend/config"
	"com-seek/backend/helpers"
	"com-seek/backend/models"
	"errors"
	"fmt"
	"image"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "golang.org/x/image/webp"
	"gorm.io/gorm"
)

type FileController struct {
	fileConfig config.FileConfig
	DB         *gorm.DB
}

func NewFileController(db *gorm.DB, fileConfig config.FileConfig) *FileController {
	return &FileController{
		fileConfig: fileConfig,
		DB:         db,
	}
}

func (fc *FileController) SaveFile(
	c *gin.Context,
	userID uint,
	fileHeader *multipart.FileHeader,
	fileCategory models.FileCategory) (*models.File, error) {
	img, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	defer img.Close()

	fileBytes, err := io.ReadAll(img)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	valid, extension, err := helpers.ValidateAndGetExtension(fileBytes, fileCategory, fc.fileConfig)

	if !valid {
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("file is not valid")
	}

	fileRecord := &models.File{
		ID:        uuid.NewString(),
		UserID:    userID,
		Extension: extension,
		Category:  fileCategory,
	}

	if err := fc.DB.Create(fileRecord).Error; err != nil {
		return nil, fmt.Errorf("failed to create file record: %w", err)
	}

	if err := os.MkdirAll(fc.fileConfig.SavePath, 0o755); err != nil {
		fc.DB.Delete(fileRecord)
		return nil, fmt.Errorf("failed to create save directory %s: %w", fc.fileConfig.SavePath, err)
	}

	filePath := filepath.Join(fc.fileConfig.SavePath, fileRecord.ID)

	if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
		fc.DB.Delete(fileRecord)
		return nil, fmt.Errorf("failed to save file to local: %w", err)
	}

	return fileRecord, nil
}

func (fc *FileController) ServeFile(c *gin.Context) {
	fileID := c.Param("uuid")

	if _, err := uuid.Parse(fileID); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "file not found",
		})
		return
	}

	var fileRecord models.File

	if err := fc.DB.Where("id = ?", fileID).First(&fileRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "file not found",
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to retrieve file record",
			})
		}
		return
	}

	filePath := filepath.Join(fc.fileConfig.SavePath, fileRecord.ID)

	securePath, err := filepath.Abs(filePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal path error",
		})
		return
	}

	savePathAbs, err := filepath.Abs(fc.fileConfig.SavePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal configuration error",
		})
		return
	}

	if !strings.HasPrefix(securePath, savePathAbs) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "invalid file path detected",
		})
		return
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "file not found on server storage",
		})
		return
	}

	contentType := mime.TypeByExtension(string(fileRecord.Extension))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	c.Header("Content-Type", contentType)

	c.File(filePath)
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
