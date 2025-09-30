package controllers

import (
	"com-seek/backend/models"
	"net/http"

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
