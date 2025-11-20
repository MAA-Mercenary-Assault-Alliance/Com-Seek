package controllers

import (
	"com-seek/backend/models"
	"com-seek/backend/services"
	"fmt"
	"log"
	"strconv"

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
	Approved *bool `json:"approved" binding:"required"`
}

func IsAdmin(db *gorm.DB, userID uint) bool {
	var admin models.Admin

	if err := db.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		return false
	}
	return true
}

func ConvertIDtoUint(idStr string, c *gin.Context) (uint, error) {
	u, err := strconv.ParseUint(idStr, 10, strconv.IntSize)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return 0, err
	}
	var uintID = uint(u)

	return uintID, err
}

func SendDeletionEmail(recipient string) {
	if err := services.SendEmail(
		recipient,
		"Account Disapproval and Deletion Notification",
		"We're sorry to inform you that your account has been disapproved and subsequently deleted. If you believe this was done in error, please reach out to our support team. Otherwise, you are welcome to re-register if you'd like to use the site again in the future."); err != nil {
		log.Println("Error: failed to send account deletion notification to", recipient, ":", err)
	}
}

func (ac *AdminController) GetPendingCompanies(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Check for Admin Status
	if !IsAdmin(ac.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var companies []models.Company
	if err := ac.DB.Where("approved = ?", false).Preload("User").Find(&companies).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching pending companies"})
		return
	}

	var responses []models.CompanyResponse
	for _, c := range companies {
		responses = append(responses, models.CompanyResponse{
			UserID:        c.UserID,
			CreatedAt:     c.User.CreatedAt,
			Name:          c.Name,
			Website:       c.Website,
			ContactEmail:  c.ContactEmail,
			ContactNumber: c.ContactNumber,
			Location:      c.Location,
			Description:   c.Description,
			Approved:      c.Approved,
		})
	}

	c.JSON(200, gin.H{"companies": responses})
}

func (ac *AdminController) ReviewCompany(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id") // Company UserID

	// Check for Admin Status
	if !IsAdmin(ac.DB, userID) {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Company id: %s>: unauthorized", userID, idStr))
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var payload ReviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Company id: %s>: invalid payload", userID, idStr))
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	uintID, convertError := ConvertIDtoUint(idStr, c)
	if convertError != nil {
		return
	}

	// Check for Company
	company := models.Company{
		UserID: uintID,
	}

	if err := ac.DB.Preload("User").First(&company).Error; err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Company id: %d>: error fetching company", userID, uintID))
		c.JSON(500, gin.H{"error": "Error fetching company"})
		return
	}

	if *payload.Approved {
		company.Approved = true
		if err := ac.DB.Save(&company).Error; err != nil {
			logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Company id: %d>: failed to approve company", userID, uintID))
			c.JSON(500, gin.H{"error": "Failed to approve company"})
			return
		}
		logger.Info(fmt.Sprintf("<Admin id: %d> Successfully Reviewed <Company id: %d>: Approved", userID, uintID))
		c.JSON(200, gin.H{"message": "Company approved"})
	} else {
		notificationEmail := company.User.Email

		// Just delete the user
		if err := ac.DB.Unscoped().Delete(&company.User).Error; err != nil {
			logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Company id: %d>: failed to reject and delete company's user", userID, uintID))
			c.JSON(500, gin.H{"error": "Failed to delete company's user"})
			return
		}
		logger.Info(fmt.Sprintf("<Admin id: %d> Successfully Reviewed <Company id: %d>: Rejected", userID, uintID))
		c.JSON(200, gin.H{"message": "Company's user deleted"})

		go SendDeletionEmail(notificationEmail)
	}
}

func (ac *AdminController) GetPendingStudents(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Check for Admin Status
	if !IsAdmin(ac.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var students []models.Student
	if err := ac.DB.Where("approved = ?", false).Preload("User").Find(&students).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching pending students"})
		return
	}

	var responses []models.StudentResponse
	for _, c := range students {
		responses = append(responses, models.StudentResponse{
			UserID:         c.UserID,
			CreatedAt:      c.User.CreatedAt,
			FirstName:      c.FirstName,
			LastName:       c.LastName,
			Description:    c.Description,
			IsAlum:         c.IsAlum,
			Approved:       c.Approved,
			GitHub:         c.GitHub,
			LinkedIn:       c.LinkedIn,
			Facebook:       c.Facebook,
			Instagram:      c.Instagram,
			Twitter:        c.Twitter,
			ProfileImageID: c.ProfileImageID,
		})
	}

	c.JSON(200, gin.H{"students": responses})
}

func (ac *AdminController) ReviewStudent(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id") // Student UserID

	// Check for Admin Status
	if !IsAdmin(ac.DB, userID) {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Student id: %s>: unauthorized", userID, idStr))
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var payload ReviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Student id: %s>: invalid payload", userID, idStr))
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	uintID, convertError := ConvertIDtoUint(idStr, c)
	if convertError != nil {
		return
	}

	student := models.Student{
		UserID: uintID,
	}

	if err := ac.DB.Preload("User").Where("user_id = ?", idStr).First(&student).Error; err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Student id: %d>: error fetching student", userID, uintID))
		c.JSON(500, gin.H{"error": "Error fetching pending students"})
		return
	}

	if *payload.Approved {
		student.Approved = true
		if err := ac.DB.Save(&student).Error; err != nil {
			logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Student id: %d>: failed to approve student", userID, uintID))
			c.JSON(500, gin.H{"error": "Failed to approve student"})
			return
		}
		logger.Info(fmt.Sprintf("<Admin id: %d> Successfully Reviewed <Student id: %d>: Approved", userID, uintID))
		c.JSON(200, gin.H{"message": "Student approved"})
	} else {
		notificationEmail := student.User.Email

		// Just delete the student
		if err := ac.DB.Unscoped().Delete(&student.User).Error; err != nil {
			logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Student id: %d>: failed to reject and delete student's user", userID, uintID))
			c.JSON(500, gin.H{"error": "Failed to reject and delete student"})
			return
		}
		logger.Info(fmt.Sprintf("<Admin id: %d> Successfully Reviewed <Student id: %d>: Rejected", userID, uintID))
		c.JSON(200, gin.H{"message": "Student's user deleted"})

		go SendDeletionEmail(notificationEmail)
	}
}

func (ac *AdminController) GetPendingJobs(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	if !IsAdmin(ac.DB, userID) {
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var jobs []models.Job
	if err := ac.DB.Where("check_needed = ?", true).Preload("Company.User").Find(&jobs).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch pending jobs"})
		return
	}
	c.JSON(200, gin.H{"jobs": jobs})
}

func (ac *AdminController) ReviewJob(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	idStr := c.Param("id") // Job ID

	// Check for Admin Status
	if !IsAdmin(ac.DB, userID) {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Job id: %s>: unauthorized", userID, idStr))
		c.JSON(403, gin.H{"error": "Unauthorized: For Admin only"})
		return
	}

	var payload ReviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Job id: %s>: invalid payload", userID, idStr))
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	uintID, convertError := ConvertIDtoUint(idStr, c)
	if convertError != nil {
		return
	}

	// Check for Job
	job := models.Job{
		ID: uintID,
	}
	if err := ac.DB.First(&job).Error; err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Job id: %s>: error fetching job", userID, idStr))
		c.JSON(404, gin.H{"error": "Job not found"})
		return
	}

	if !job.CheckNeeded {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Job id: %s>: job already reviewed", userID, idStr))
		c.JSON(304, gin.H{"message": "Job already reviewed"})
		return
	}

	if *payload.Approved {
		job.CheckNeeded = false
		job.Approved = true
		job.Visibility = true
	} else {
		job.CheckNeeded = false
		job.Approved = false
	}

	if err := ac.DB.Save(&job).Error; err != nil {
		logger.Error(fmt.Sprintf("<Admin id: %d> Attempt to Review <Job id: %s>: failed to save job review", userID, idStr))
		c.JSON(500, gin.H{"error": "Failed to review job"})
		return
	}

	logger.Info(fmt.Sprintf("<Admin id: %d> Successfully Reviewed <Job id: %s>: Approved=%t", userID, idStr, *payload.Approved))
	c.JSON(200, gin.H{"message": "Job Reviewed"})
}
