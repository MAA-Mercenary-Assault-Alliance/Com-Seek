package controllers

import (
	"com-seek/backend/models"
	"com-seek/backend/services"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
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
		JobID          uint                  `form:"job_id" binding:"required"`
		File           *multipart.FileHeader `form:"file" binding:"required"`
		ReCAPTCHAToken string                `form:"recaptcha_response" binding:"required"`
	}

	student := models.Student{
		UserID: userID,
	}

	if err := jc.DB.First(&student).Error; err != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application: not a student account", userID))
		c.JSON(http.StatusForbidden, gin.H{"error": "not a student account"})
		return
	}

	if !student.Approved {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application: student account is not approved", userID))
		c.JSON(http.StatusForbidden, gin.H{"error": "student account is not approved"})
		return
	}

	var input CreateJobApplicationInput

	if err := c.ShouldBind(&input); err != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: %s", userID, input.JobID, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.VerifyRecaptchaToken(input.ReCAPTCHAToken); err != nil {
		if errors.Is(err, services.ErrRecaptchaFailed) {
			logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: reCAPTCHA verification failed", userID, input.JobID))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: reCAPTCHA verification error: %s", userID, input.JobID, err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		return
	}

	job := models.Job{
		ID: input.JobID,
	}

	if err := jc.DB.First(&job).Error; err != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: no job found with the given id", userID, input.JobID))
		c.JSON(http.StatusBadRequest, gin.H{"error": "no job found with the given id"})
		return
	}

	if !job.Approved {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: job is not approved", userID, input.JobID))
		c.JSON(http.StatusBadRequest, gin.H{"error": "job is not approved"})
		return
	}

	jobApplication := models.JobApplication{
		StudentID: student.UserID,
		JobID:     job.ID,
	}

	file, err := input.File.Open()
	if err != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: failed to open attachment file", userID, input.JobID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open attachment file"})
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: failed to read file", userID, input.JobID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	mime := mimetype.Detect(buffer)

	if mime.String() != "application/pdf" && mime.String() != "application/x-pdf" {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: invalid file type", userID, input.JobID))
		c.JSON(http.StatusBadRequest, gin.H{"error": "only PDF files are allowed"})
		return
	}

	tx := jc.DB.Begin()
	defer tx.Rollback()

	result := tx.Where("student_id = ? AND job_id = ?", student.UserID, job.ID).FirstOrCreate(&jobApplication)

	if result.Error != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: %s", userID, input.JobID, result.Error.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: job application already exists", userID, input.JobID))
		c.JSON(http.StatusConflict, gin.H{"error": "job application already exists"})
		return
	}

	tx.Preload("Job.Company.User").Preload("Student").First(&jobApplication)

	subject := jobApplication.Job.Title + ": New applicant applied"
	body := (jobApplication.Student.FirstName +
		" " + jobApplication.Student.LastName +
		" has applied to your job posting: " +
		jobApplication.Job.Title)

	if err := services.SendEmailWithAttachment(
		jobApplication.Job.Company.User.Email,
		subject,
		body,
		input.File); err != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: failed to send email to company: %s", userID, input.JobID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error(fmt.Sprintf("<Student id: %d> Attempt to Create Job Application for <Job id: %d>: failed to submit the application: %s", userID, input.JobID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit the application"})
		return
	}

	logger.Info(fmt.Sprintf("<Student id: %d> Successfully Created Job Application for <Job id: %d>", userID, input.JobID))
	c.JSON(http.StatusOK, gin.H{"message": "successfully applied"})
}

func (jc *JobApplicationController) DeleteJobApplication(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	jobApplicationIDStr := c.Param("id")

	var jobApplicationID uint

	u, err := strconv.ParseUint(jobApplicationIDStr, 10, strconv.IntSize)

	if err != nil {
		logger.Error(fmt.Sprintf("<Company id: %d> Attempt to Reject Job Application: invalid Job Application ID", userID))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jobApplicationID = uint(u)

	jobApplication := models.JobApplication{
		ID: jobApplicationID,
	}

	if err := jc.DB.Preload("Job").First(&jobApplication).Error; err != nil {
		logger.Error(fmt.Sprintf("<Company id: %d> Attempt to Reject <Job Application id: %d> of <Student id: %d> in <Job id: %d>: job application not found", userID, jobApplicationID, jobApplication.StudentID, jobApplication.JobID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "job application not found"})
		return
	}

	if jobApplication.Job.CompanyID != userID {
		logger.Error(fmt.Sprintf("<Company id: %d> Attempt to Reject <Job Application id: %d> of <Student id: %d> in <Job id: %d>: unauthorized", userID, jobApplicationID, jobApplication.StudentID, jobApplication.JobID))
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	if err := jc.DB.Delete(&jobApplication).Error; err != nil {
		logger.Error(fmt.Sprintf("<Company id: %d> Attempt to Reject <Job Application id: %d> of <Student id: %d> in <Job id: %d>: %s", userID, jobApplicationID, jobApplication.StudentID, jobApplication.JobID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info(fmt.Sprintf("<Company id: %d> Successfully Rejected <Job Application id: %d> of <Student id: %d> in <Job id: %d>", userID, jobApplicationID, jobApplication.StudentID, jobApplication.JobID))
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted the job application"})
}
