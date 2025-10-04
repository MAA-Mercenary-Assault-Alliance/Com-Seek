package controllers

import (
	"com-seek/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{
		DB: db,
	}
}

func (sc *StudentController) GetStudentProfile(c *gin.Context) {
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

	student := models.Student{
		UserID: id,
	}

	if err := sc.DB.Preload("User").First(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": student})
}

func (sc *StudentController) UpdateStudentProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	type StudentProfileInput struct {
		FirstName   string  `json:"first_name" binding:"omitempty"`
		LastName    string  `json:"last_name" binding:"omitempty"`
		Description *string `json:"description" binding:"omitempty,max=4096"`
		IsAlum      *bool   `json:"is_alum" binding:"omitempty"`
		GitHub      *string `json:"github" binding:"omitempty,max=256"`
		LinkedIn    *string `json:"linkedin" binding:"omitempty,max=256"`
		Facebook    *string `json:"facebook" binding:"omitempty,max=256"`
		Instragram  *string `json:"instragram" binding:"omitempty,max=256"`
		Twitter     *string `json:"twitter" binding:"omitempty,max=256"`
	}

	var input StudentProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		UserID: userID,
	}

	if err := sc.DB.First(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := sc.DB.Model(&student).Updates(&input)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated the profile"})
}
