package controllers

import (
	"com-seek/backend/helpers"
	"com-seek/backend/models"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyController struct {
	DB             *gorm.DB
	fileController *FileController
}

func NewCompanyController(db *gorm.DB, fileController *FileController) *CompanyController {
	return &CompanyController{
		DB:             db,
		fileController: fileController,
	}
}

func (cc *CompanyController) GetCompanyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id")

	var company models.Company

	if idStr == "" {
		if err := cc.DB.First(&company, "user_id = ?", userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusForbidden, gin.H{"error": "not a company account"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := cc.DB.First(&company, "user_id = ?", id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var jobs []models.Job

	if err := cc.DB.Table("jobs").Preload("Company").
		Where("jobs.company_id = ? AND jobs.approved = 1", company.UserID).Find(&jobs).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profile": company,
		"jobs":    jobs,
	})
}

func (cc *CompanyController) UpdateCompanyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	type CompanyProfileInput struct {
		Name          string                `form:"name" binding:"omitempty"`
		Website       *string               `form:"website" binding:"omitempty,max=256"`
		ContactEmail  *string               `form:"contact_email" binding:"omitempty,max=100"`
		ContactNumber *string               `form:"contact_number" binding:"omitempty,max=20"`
		Location      *string               `form:"location" binding:"omitempty,max=256"`
		Description   *string               `form:"description" binding:"omitempty,max=4096"`
		ProfileImage  *multipart.FileHeader `form:"profile_image"`
		CoverImage    *multipart.FileHeader `form:"cover_image"`
	}

	var input CompanyProfileInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := models.Company{
		UserID: userID,
	}

	if err := cc.DB.First(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var oldProfileImageID, oldCoverImageID string

	if company.ProfileImageID != nil && *company.ProfileImageID != "" {
		oldProfileImageID = *company.ProfileImageID
	}
	if company.CoverImageID != nil && *company.CoverImageID != "" {
		oldCoverImageID = *company.CoverImageID
	}

	if input.Name != "" {
		company.Name = input.Name
	}

	if input.Website != nil {
		company.Website = *input.Website
	}

	if input.ContactEmail != nil {
		company.ContactEmail = *input.ContactEmail
	}

	if input.ContactNumber != nil {
		company.ContactNumber = *input.ContactNumber
	}

	if input.Location != nil {
		company.Location = *input.Location
	}

	if input.Description != nil {
		company.Description = *input.Description
	}

	if input.ProfileImage != nil {
		profile, err := cc.fileController.SaveFile(c, userID, input.ProfileImage, models.FileCategoryProfile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if company.ProfileImageID == nil {
			company.ProfileImageID = new(string)
		}

		*company.ProfileImageID = profile.ID
	}

	if input.CoverImage != nil {
		cover, err := cc.fileController.SaveFile(c, userID, input.CoverImage, models.FileCategoryCover)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if company.CoverImageID == nil {
			company.CoverImageID = new(string)
		}

		*company.CoverImageID = cover.ID
	}

	res := cc.DB.Save(&company)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	if oldProfileImageID != "" && oldProfileImageID != *company.ProfileImageID {
		go helpers.DeleteFileRecord(cc.DB, oldProfileImageID)
	}

	if oldCoverImageID != "" && oldCoverImageID != *company.CoverImageID {
		go helpers.DeleteFileRecord(cc.DB, oldCoverImageID)
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated the profile"})
}

func (cc *CompanyController) GetCompanyJobs(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	company := models.Company{
		UserID: userID,
	}

	if err := cc.DB.First(&company).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not a company account"})
		return
	}

	var jobs []models.Job

	if err := cc.DB.Table("jobs").Preload("Company").
		Where("jobs.company_id = ?", company.UserID).Find(&jobs).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}
