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
		MinSalary        uint   `json:"min_salary" binding:"required"`
		MaxSalary        uint   `json:"max_salary" binding:"required"`
		MinExperience    uint   `json:"min_experience" binding:"required"`
		MaxExperience    uint   `json:"max_experience" binding:"required"`
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

	if input.MaxSalary < input.MinSalary {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "max salary must be greater than or equal to min salary"})
		return
	}

	if input.MaxExperience < input.MinExperience {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "max experience must be greater than or equal to min experience"})
		return
	}

	job := models.Job{
		Title:            input.Title,
		CompanyID:        userID,
		Location:         input.Location,
		JobType:          models.JobType(input.JobType),
		EmploymentStatus: models.EmploymentStatus(input.EmploymentStatus),
		MinSalary:        input.MinSalary,
		MaxSalary:        input.MaxSalary,
		MinExperience:    input.MinExperience,
		MaxExperience:    input.MaxExperience,
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
		MinSalary        *uint  `json:"min_salary" binding:"omitempty"`
		MaxSalary        *uint  `json:"max_salary" binding:"omitempty"`
		MinExperience    *uint  `json:"min_experience" binding:"omitempty"`
		MaxExperience    *uint  `json:"max_experience" binding:"omitempty"`
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

	if input.Title != "" {
		job.Title = input.Title
	}
	if input.Location != "" {
		job.Location = input.Location
	}
	if input.JobType != "" {
		job.JobType = models.JobType(input.JobType)
	}
	if input.EmploymentStatus != "" {
		job.EmploymentStatus = models.EmploymentStatus(input.EmploymentStatus)
	}
	if input.MinSalary != nil {
		job.MinSalary = *input.MinSalary
	}
	if input.MaxSalary != nil {
		job.MaxSalary = *input.MaxSalary
	}
	if input.MinExperience != nil {
		job.MinExperience = *input.MinExperience
	}
	if input.MaxExperience != nil {
		job.MaxExperience = *input.MaxExperience
	}
	if input.Description != "" {
		job.Description = input.Description
	}
	if job.MaxSalary < job.MinSalary {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "max salary must be greater than or equal to min salary"})
		return
	}
	if job.MaxExperience < job.MinExperience {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "max experience must be greater than or equal to min experience"})
		return
	}

	job.CheckNeeded = true

	if err := jc.DB.Save(&job).Error; err != nil {
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
