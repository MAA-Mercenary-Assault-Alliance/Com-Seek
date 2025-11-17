package controllers

import (
	"com-seek/backend/models"
	"fmt"
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
		MinSalary        *uint  `json:"min_salary" binding:"required"`
		MaxSalary        *uint  `json:"max_salary" binding:"required"`
		MinExperience    *uint  `json:"min_experience" binding:"required"`
		MaxExperience    *uint  `json:"max_experience" binding:"required"`
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

	if *input.MaxSalary < *input.MinSalary {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "max salary must be greater than or equal to min salary"})
		return
	}

	if *input.MaxExperience < *input.MinExperience {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "max experience must be greater than or equal to min experience"})
		return
	}

	job := models.Job{
		Title:            input.Title,
		CompanyID:        userID,
		Location:         input.Location,
		JobType:          models.JobType(input.JobType),
		EmploymentStatus: models.EmploymentStatus(input.EmploymentStatus),
		MinSalary:        *input.MinSalary,
		MaxSalary:        *input.MaxSalary,
		MinExperience:    *input.MinExperience,
		MaxExperience:    *input.MaxExperience,
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
	var jobs []models.Job
	query := jc.DB.Preload("Company").Where("approved = 1 AND visibility = 1")

	if location := c.Query("location"); location != "" {
		locationPattern := fmt.Sprintf("%%%s%%", location)
		query = query.Where("location LIKE ?", locationPattern)
	}
	if jobType := c.Query("job_type"); jobType != "" {
		query = query.Where("job_type = ?", jobType)
	}
	if employmentStatus := c.Query("employment_status"); employmentStatus != "" {
		query = query.Where("employment_status = ?", employmentStatus)
	}
	if minSalary := c.Query("min_salary"); minSalary != "" {
		if minSal, err := strconv.Atoi(minSalary); err == nil {
			query = query.Where("min_salary >= ?", minSal)
		}
	}
	if maxSalary := c.Query("max_salary"); maxSalary != "" {
		if maxSal, err := strconv.Atoi(maxSalary); err == nil {
			query = query.Where("max_salary <= ?", maxSal)
		}
	}
	if minExperience := c.Query("min_experience"); minExperience != "" {
		if minExp, err := strconv.Atoi(minExperience); err == nil {
			query = query.Where("min_experience >= ?", minExp)
		}
	}
	if maxExperience := c.Query("max_experience"); maxExperience != "" {
		if maxExp, err := strconv.Atoi(maxExperience); err == nil {
			query = query.Where("max_experience <= ?", maxExp)
		}
	}
	if keyword := c.Query("keyword"); keyword != "" {
		keywordPattern := fmt.Sprintf("%%%s%%", keyword)
		query = query.Where("title LIKE ? OR description LIKE ?", keywordPattern, keywordPattern)
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "30"))
	if err != nil || limit <= 0 {
		limit = 30
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

func (jc *JobController) GetJob(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id")

	bitSize := strconv.IntSize
	u, err := strconv.ParseUint(idStr, 10, bitSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	id := uint(u)

	job := models.Job{
		ID: id,
	}

	if err := jc.DB.Preload("Company").First(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type JobApplicantResponse struct {
		ID             uint   `json:"id"`
		StudentID      uint   `json:"student_id"`
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		IsAlum         bool   `json:"is_alum"`
		CreatedAt      string `json:"created_at"`
		ProfileImageID string `json:"profile_image_id"`
	}

	if job.CompanyID == userID {
		var applicants []JobApplicantResponse
		if err := jc.DB.Table("job_applications").
			Joins("LEFT JOIN students ON students.user_id = job_applications.student_id").
			Where("job_id = ?", job.ID).
			Scan(&applicants).
			Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"job":        job,
			"applicants": applicants,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
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
		Visibility       *bool  `json:"visibility" binding:"omitempty"`
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
	if input.Visibility != nil {
		job.Visibility = *input.Visibility
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
