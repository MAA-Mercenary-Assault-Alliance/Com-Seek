package controllers

import (
	"com-seek/backend/models"
	"com-seek/backend/services"
	"net/http"
	"strconv"

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

	tx := jc.DB.Begin()
	defer tx.Rollback()

	result := tx.FirstOrCreate(&jobApplication)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "job application already exists"})
		return
	}

	tx.Preload("Job.Company").Preload("Student").First(&jobApplication)

	subject := jobApplication.Job.Title + ": New applicant applied"
	body := (jobApplication.Student.FirstName +
		" " + jobApplication.Student.LastName +
		" has applied to your job posting: " +
		jobApplication.Job.Title)

	if err := services.SendEmail(
		jobApplication.Job.Company.ContactEmail,
		subject,
		body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit the application"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully applied"})
}

func (jc *JobApplicationController) DeleteJobApplication(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	jobApplicationIDStr := c.Param("id")

	var jobApplicationID uint

	u, err := strconv.ParseUint(jobApplicationIDStr, 10, strconv.IntSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jobApplicationID = uint(u)

	jobApplication := models.JobApplication{
		ID: jobApplicationID,
	}

	if err := jc.DB.Preload("Job").First(&jobApplication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "job application not found"})
		return
	}

	if jobApplication.Job.CompanyID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	if err := jc.DB.Delete(&jobApplication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted the job application"})
}
