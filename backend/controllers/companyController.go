package controllers

import (
	"com-seek/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyController struct {
	DB *gorm.DB
}

func NewCompanyController(db *gorm.DB) *CompanyController {
	return &CompanyController{
		DB: db,
	}
}

func (cc *CompanyController) GetCompanyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id")

	var id uint

	if idStr == "" {
		id = userID
	} else {
		bitSize := strconv.IntSize
		u, err := strconv.ParseUint(c.Param("id"), 10, bitSize)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		id = uint(u)
	}

	company := models.Company{
		UserID: id,
	}

	if err := cc.DB.Preload("User").First(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": company})
}

func (cc *CompanyController) UpdateCompanyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	type CompanyProfileInput struct {
		Name          string `json:"name" binding:"omitempty"`
		Website       string `json:"website" binding:"omitempty,max=256"`
		ContactEmail  string `json:"contact_email" binding:"omitempty,max=100"`
		ContactNumber string `json:"contact_number" binding:"omitempty,max=20"`
		Location      string `json:"location" binding:"omitempty,max=256"`
		Description   string `json:"description" binding:"omitempty,max=4096"`
	}

	var input CompanyProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
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

	res := cc.DB.Model(&company).Updates(&input)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated the profile"})
}
