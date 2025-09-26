package controllers

import (
	"com-seek/backend/models"
	"net/http"

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

	student := models.Student{
		UserID: userID,
	}

	if err := sc.DB.First(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": student})
}
