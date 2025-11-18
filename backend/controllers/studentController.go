package controllers

import (
	"com-seek/backend/helpers"
	"com-seek/backend/models"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB             *gorm.DB
	fileController *FileController
}

func NewStudentController(db *gorm.DB, fileController *FileController) *StudentController {
	return &StudentController{
		DB:             db,
		fileController: fileController,
	}
}

type StudentResponse struct {
	UserID         uint   `json:"user_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Description    string `json:"description"`
	IsAlum         bool   `json:"is_alum"`
	Approved       bool   `json:"approved"`
	GitHub         string `json:"github"`
	LinkedIn       string `json:"linkedin"`
	Facebook       string `json:"facebook"`
	Instagram      string `json:"instagram"`
	Twitter        string `json:"twitter"`
	ProfileImageID string `json:"profile_image_id"`
	CoverImageID   string `json:"cover_image_id"`
	TranscriptID   string `json:"transcript_id"`
	CVID           string `json:"cv_id"`
}

type JobApplicationResponse struct {
	ID               uint                    `json:"id"`
	JobID            uint                    `json:"job_id"`
	Title            string                  `json:"job_title"`
	Name             string                  `json:"company_name"`
	ProfileImageID   string                  `json:"company_profile_image_id"`
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
	var studentRes StudentResponse

	if idStr == "" {
		result := sc.DB.Table("students").
			Joins("LEFT JOIN users ON users.id = students.user_id").
			Where("user_id = ?", userID).Scan(&studentRes)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "not a student account"})
			return
		}

		id = userID
	} else {
		u, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		id = uint(u)

		result := sc.DB.Table("students").
			Joins("LEFT JOIN users ON users.id = students.user_id").
			Where("user_id = ?", id).Scan(&studentRes)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
			return
		}
	}

	if userID == id {
		var jobApplications []JobApplicationResponse
		if err := sc.DB.Table("job_applications").
			Joins("LEFT JOIN jobs ON jobs.id = job_applications.job_id").
			Joins("LEFT JOIN companies ON jobs.company_id = companies.user_id").
			Where("student_id = ?", studentRes.UserID).
			Select("job_applications.*, jobs.*, companies.name, companies.profile_image_id").
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
		FirstName    string                `form:"first_name" binding:"omitempty"`
		LastName     string                `form:"last_name" binding:"omitempty"`
		Description  *string               `form:"description" binding:"omitempty,max=4096"`
		IsAlum       *bool                 `form:"is_alum" binding:"omitempty"`
		GitHub       *string               `form:"github" binding:"omitempty,max=256"`
		LinkedIn     *string               `form:"linkedin" binding:"omitempty,max=256"`
		Facebook     *string               `form:"facebook" binding:"omitempty,max=256"`
		Instagram    *string               `form:"instagram" binding:"omitempty,max=256"`
		Twitter      *string               `form:"twitter" binding:"omitempty,max=256"`
		ProfileImage *multipart.FileHeader `form:"profile_image"`
		CoverImage   *multipart.FileHeader `form:"cover_image"`
		Transcript   *multipart.FileHeader `form:"transcript"`
		CV           *multipart.FileHeader `form:"cv"`
	}

	var input StudentProfileInput

	if err := c.ShouldBind(&input); err != nil {
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

	var oldProfileImageID, oldCoverImageID, oldCVID, oldTranscriptID string

	if student.ProfileImageID != nil {
		oldProfileImageID = *student.ProfileImageID
	}
	if student.CoverImageID != nil {
		oldCoverImageID = *student.CoverImageID
	}
	oldTranscriptID = student.TranscriptID
	if student.CVID != nil {
		oldCVID = *student.CVID
	}

	if input.FirstName != "" {
		student.FirstName = input.FirstName
	}

	if input.LastName != "" {
		student.LastName = input.LastName
	}

	if input.Description != nil {
		student.Description = *input.Description
	}

	if input.IsAlum != nil {
		student.IsAlum = *input.IsAlum
	}

	if input.GitHub != nil {
		student.GitHub = *input.GitHub
	}

	if input.LinkedIn != nil {
		student.LinkedIn = *input.LinkedIn
	}

	if input.Facebook != nil {
		student.Facebook = *input.Facebook
	}

	if input.Instagram != nil {
		student.Instagram = *input.Instagram
	}

	if input.Twitter != nil {
		student.Twitter = *input.Twitter
	}

	if input.ProfileImage != nil {
		profile, err := sc.fileController.SaveFile(c, userID, input.ProfileImage, models.FileCategoryProfile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if student.ProfileImageID == nil {
			student.ProfileImageID = new(string)
		}

		*student.ProfileImageID = profile.ID
	}

	if input.CoverImage != nil {
		cover, err := sc.fileController.SaveFile(c, userID, input.CoverImage, models.FileCategoryCover)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if student.CoverImageID == nil {
			student.CoverImageID = new(string)
		}

		*student.CoverImageID = cover.ID
	}

	if input.Transcript != nil {
		transcript, err := sc.fileController.SaveFile(c, userID, input.Transcript, models.FileCategoryTranscript)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		student.TranscriptID = transcript.ID
	}

	if input.CV != nil {
		cover, err := sc.fileController.SaveFile(c, userID, input.CV, models.FileCategoryCV)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if student.CVID == nil {
			student.CVID = new(string)
		}

		*student.CVID = cover.ID
	}

	res := sc.DB.Save(&student)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	if input.Transcript != nil {
		if oldTranscriptID != "" && oldTranscriptID != student.TranscriptID {
			go helpers.DeleteFileRecord(sc.DB, oldTranscriptID)
		}
	}

	if oldProfileImageID != "" && oldProfileImageID != *student.ProfileImageID {
		go helpers.DeleteFileRecord(sc.DB, oldProfileImageID)
	}

	if oldCoverImageID != "" && oldCoverImageID != *student.CoverImageID {
		go helpers.DeleteFileRecord(sc.DB, oldCoverImageID)
	}

	if oldCVID != "" && student.CVID != nil && oldCVID != *student.CVID {
		go helpers.DeleteFileRecord(sc.DB, oldCVID)
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated the profile"})
}
