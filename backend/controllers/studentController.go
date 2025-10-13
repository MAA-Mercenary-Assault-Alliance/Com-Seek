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

type StudentResponse struct {
	UserID      uint   `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	IsAlum      bool   `json:"is_alum"`
	Approved    bool   `json:"approved"`
	GitHub      string `json:"github"`
	LinkedIn    string `json:"linkedin"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Twitter     string `json:"twitter"`
}

type JobApplicationResponse struct {
	ID               uint                    `json:"id"`
	JobID            uint                    `json:"job_id"`
	Title            string                  `json:"job_title"`
	Name             string                  `json:"company_name"`
	Location         string                  `json:"job_location"`
	JobType          models.JobType          `json:"job_type"`
	EmploymentStatus models.EmploymentStatus `json:"job_employment_status"`
	MinSalary        uint                    `json:"job_min_salary"`
	MaxSalary        uint                    `json:"job_max_salary"`
	MinExperience    uint                    `json:"job_min_experience"`
	MaxExperience    uint                    `json:"job_max_experience"`
	CreatedAt        string                  `json:"created_at"`
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

	var studentRes StudentResponse

	if err := sc.DB.Table("students").Joins("LEFT JOIN users ON users.id = students.user_id").
		Where("user_id = ?", id).Scan(&studentRes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userID == id {
		var jobApplications []JobApplicationResponse
		if err := sc.DB.Table("job_applications").
			Joins("LEFT JOIN jobs ON jobs.id = job_applications.job_id").
			Joins("LEFT JOIN companies ON jobs.company_id = companies.user_id").
			Where("student_id = ?", id).
			Select("job_applications.*, jobs.*, companies.name").
			Scan(&jobApplications).
			Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"profile":      studentRes,
			"applications": jobApplications,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": studentRes})
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
		Instagram   *string `json:"instagram" binding:"omitempty,max=256"`
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
