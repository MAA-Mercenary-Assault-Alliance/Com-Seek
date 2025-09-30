package controllers

import (
	"com-seek/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobController struct {
	DB *gorm.DB
}

func NewJobController(db *gorm.DB) *JobController {
	return &JobController{
		DB: db,
	}
}

func (jc *JobController) CreateJob(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	type CreateJobInput struct {
		Title            string `json:"title" binding:"required,max=256"`
		Location         string `json:"location" binding:"required,max=256"`
		JobType          string `json:"job_type" binding:"required,oneof='Software & Application Development' 'Data & AI' 'Cloud & Infrastructure' 'Cybersecurity' 'Product & Design' 'Testing & Quality' 'Hardware & Electronics' 'Management & Leadership' 'IT Support & Operations'"`
		EmploymentStatus string `json:"employment_status" binding:"required,oneof='part-time' 'full-time' 'contract' 'internship'"`
		Salary           uint   `json:"salary" binding:"required"`
		Description      string `json:"description" binding:"required,max=10240"`
	}

	company := models.Company{
		UserID: userID,
	}

	if err := jc.DB.First(&company).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not a company account"})
		return
	}

	if !company.Approved {
		c.JSON(http.StatusForbidden, gin.H{"error": "company account is not approved"})
		return
	}

	var input CreateJobInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job := models.Job{
		Title:            input.Title,
		CompanyID:        userID,
		Location:         input.Location,
		JobType:          models.JobType(input.JobType),
		EmploymentStatus: models.EmploymentStatus(input.EmploymentStatus),
		Salary:           input.Salary,
		Description:      input.Description,
		CheckNeeded:      true,
	}

	if err := jc.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (jc *JobController) GetJobs(c *gin.Context) {
	jobIDStr := c.Param("id")

	if jobIDStr != "" {
		u, err := strconv.ParseUint(jobIDStr, 10, strconv.IntSize)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		jobID := uint(u)

		job := models.Job{
			ID: jobID,
		}

		if err := jc.DB.Preload("Company.User").First(&job).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"job": job})
		return
	}

	var jobs []models.Job

	if err := jc.DB.Preload("Company.User").Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

func (jc *JobController) UpdateJob(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	jobIDStr := c.Param("id")

	var jobID uint

	u, err := strconv.ParseUint(jobIDStr, 10, strconv.IntSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jobID = uint(u)

	type UpdateJobInput struct {
		Title            string `json:"title" binding:"omitempty,max=256"`
		Location         string `json:"location" binding:"omitempty,max=256"`
		JobType          string `json:"job_type" binding:"omitempty,oneof='Software & Application Development' 'Data & AI' 'Cloud & Infrastructure' 'Cybersecurity' 'Product & Design' 'Testing & Quality' 'Hardware & Electronics' 'Management & Leadership' 'IT Support & Operations'"`
		EmploymentStatus string `json:"employment_status" binding:"omitempty,oneof='part-time' 'full-time' 'contract' 'internship'"`
		Salary           uint   `json:"salary" binding:"omitempty"`
		Description      string `json:"description" binding:"omitempty,max=10240"`
	}

	var input UpdateJobInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job := models.Job{
		ID: jobID,
	}

	if err := jc.DB.Preload("Company.User").First(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "job not found"})
		return
	}

	if job.CompanyID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	updates := map[string]interface{}{
		"CheckNeeded": true,
	}

	if input.Title != "" {
		updates["Title"] = input.Title
	}
	if input.Location != "" {
		updates["Location"] = input.Location
	}
	if input.JobType != "" {
		updates["JobType"] = input.JobType
	}
	if input.EmploymentStatus != "" {
		updates["EmploymentStatus"] = input.EmploymentStatus
	}
	if input.Salary != 0 {
		updates["Salary"] = input.Salary
	}
	if input.Description != "" {
		updates["Description"] = input.Description
	}

	if err := jc.DB.Model(&job).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated the job"})
}

func (jc *JobController) DeleteJob(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	jobIDStr := c.Param("id")

	var jobID uint

	u, err := strconv.ParseUint(jobIDStr, 10, strconv.IntSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jobID = uint(u)

	job := models.Job{
		ID: jobID,
	}

	if err := jc.DB.Preload("Company.User").First(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "job not found"})
		return
	}

	if job.CompanyID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	if err := jc.DB.Delete(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted the job"})
}
