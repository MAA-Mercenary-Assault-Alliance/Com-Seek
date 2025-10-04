package controllers

import (
	"com-seek/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	DB *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		DB: db,
	}
}

type ReviewPayload struct {
	Approved bool `json:"approved" binding:"required"`
}

func IsAdmin(db *gorm.DB, userID uint) bool {
	var admin models.Admin

	if err := db.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		return false
	}
	return true
}

func (sc *AdminController) GetPendingCompanies(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Check for Admin Status
	if !IsAdmin(sc.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var companies []models.Company
	if err := sc.DB.Where("approved = ?", false).Preload("User").Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"companies": companies})
}

func (sc *AdminController) ReviewCompany(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id") // Company UserID

	// Check for Admin Status
	if !IsAdmin(sc.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var payload ReviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	// Check for Company
	company := models.Company{
		UserID: userID,
	}

	if err := sc.DB.First(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if payload.Approved {
		company.Approved = true
		if err := sc.DB.Save(&company).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to approve company"})
			return
		}
		c.JSON(200, gin.H{"message": "Company approved"})
	} else {
		// Just delete the company
		if err := sc.DB.Where("user_id = ?", idStr).Delete(&company).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to reject and delete company"})
			return
		}
		c.JSON(200, gin.H{"message": "Company deleted"})
	}
}

func (sc *AdminController) GetPendingStudents(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Check for Admin Status
	if !IsAdmin(sc.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var students []models.Student
	if err := sc.DB.Where("approved = ?", false).Preload("User").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"students": students})
}

func (sc *AdminController) ReviewStudent(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id") // Student UserID

	// Check for Admin Status
	if !IsAdmin(sc.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var payload ReviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	student := models.Student{
		UserID: userID,
	}

	if err := sc.DB.First(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if payload.Approved {
		student.Approved = true
		if err := sc.DB.Save(&student).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to approve student"})
			return
		}
		c.JSON(200, gin.H{"message": "Student approved"})
	} else {
		// Just delete the student
		if err := sc.DB.Where("user_id = ?", idStr).Delete(&student).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to reject and delete student"})
			return
		}
		c.JSON(200, gin.H{"message": "Student deleted"})
	}

}

func (sc *AdminController) GetPendingJobs(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	if !IsAdmin(sc.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var jobs []models.Job
	if err := sc.DB.Where("check_needed = ?", true).Preload("Company.User").Find(&jobs).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch pending jobs"})
		return
	}
	c.JSON(200, gin.H{"jobs": jobs})
}

func (sc *AdminController) ReviewJob(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id") // Job ID

	// Check for Admin Status
	if !IsAdmin(sc.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var payload ReviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	// Check for Job
	var job models.Job
	if err := sc.DB.Where("id = ?", idStr).First(&job).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job not found"})
		return
	}

	if !job.CheckNeeded {
		c.JSON(304, gin.H{"message": "Job already reviewed"})
		return
	}

	if payload.Approved {
		job.CheckNeeded = false
		job.Approved = true
	} else {
		job.Approved = false
	}

	if err := sc.DB.Save(&job).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to review job"})
		return
	}
	c.JSON(200, gin.H{"message": "Job Reviewed"})

}
