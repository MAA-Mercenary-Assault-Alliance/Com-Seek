package controllers

import (
	"com-seek/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobApplicationController struct {
	DB *gorm.DB
}

func NewJobApplicationController(db *gorm.DB) *JobApplicationController {
	return &JobApplicationController{
		DB: db,
	}
}

func (jc *JobApplicationController) CreateJobApplication(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	type CreateJobApplicationInput struct {
		JobID uint `json:"job_id" binding:"required"`
	}

	student := models.Student{
		UserID: userID,
	}

	if err := jc.DB.First(&student).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not a student account"})
		return
	}

	if !student.Approved {
		c.JSON(http.StatusForbidden, gin.H{"error": "student account is not approved"})
		return
	}

	var input CreateJobApplicationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job := models.Job{
		ID: input.JobID,
	}

	if err := jc.DB.First(&job).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no job found with the given id"})
		return
	}

	if !job.Approved {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job is not approved"})
		return
	}

	jobApplication := models.JobApplication{
		StudentID: student.UserID,
		JobID:     job.ID,
	}

	if err := jc.DB.Create(&jobApplication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully applied"})
}
