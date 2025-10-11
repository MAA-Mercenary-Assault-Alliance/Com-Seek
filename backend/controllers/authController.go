package controllers

import (
	"net/http"
	"os"
	"time"

	"com-seek/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB: db,
	}
}

type RegisterInput struct {
	Email          string               `json:"email" binding:"required"`
	Password       string               `json:"password" binding:"required"`
	UserType       string               `json:"user_type" binding:"required,oneof=student company"`
	StudentProfile *StudentProfileInput `json:"student_profile,omitempty"`
	CompanyProfile *CompanyProfileInput `json:"company_profile,omitempty"`
}

type StudentProfileInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	IsAlum    *bool  `json:"is_alum" binding:"required"`
}

type CompanyProfileInput struct {
	Name          string `json:"name" binding:"required"`
	Location      string `json:"location" binding:"required"`
	ContactEmail  string `json:"contact_email" binding:"required"`
	ContactNumber string `json:"contact_number" binding:"required"`
}

func (ac *AuthController) CreateUser(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	ac.DB.Where("email=?", input.Email).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already registered"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(passwordHash),
	}

	tx := ac.DB.Begin()
	defer tx.Rollback()

	if err := tx.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	switch input.UserType {
	case "student":
		studentProfile := input.StudentProfile
		if studentProfile == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Student profile is required for student type"})
			return
		}

		profile := models.Student{
			UserID:    user.ID,
			FirstName: studentProfile.FirstName,
			LastName:  studentProfile.LastName,
			IsAlum:    *studentProfile.IsAlum,
		}

		if err := tx.Create(&profile).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	case "company":
		companyProfile := input.CompanyProfile
		if companyProfile == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Company profile is required for company type"})
			return
		}

		profile := models.Company{
			UserID:        user.ID,
			Name:          companyProfile.Name,
			Location:      companyProfile.Location,
			ContactEmail:  companyProfile.ContactEmail,
			ContactNumber: companyProfile.ContactNumber,
		}

		if err := tx.Create(&profile).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginInput LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	ac.DB.Where("email=?", loginInput.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	role := GetUserRole(ac.DB, userFound.ID)

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add((time.Minute * 30)).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	maxAge := int((time.Minute * 30) / time.Second)

	sameSite := http.SameSiteLaxMode
	secure := false
	if os.Getenv("ENV") != "dev" {
		sameSite = http.SameSiteNoneMode
		secure = true
	}

	c.SetSameSite(sameSite)
	c.SetCookie(
		"token",
		token,
		maxAge,
		"/",
		"",
		secure,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"email": userFound.Email,
		"role":  role,
	})
}

func GetUserRole(DB *gorm.DB, userID uint) string {
	var result struct{ UserID uint }

	if err := DB.Table("students").Where("user_id = ?", userID).First(&result).Error; err == nil {
		return "student"
	}
	if err := DB.Table("companies").Where("user_id = ?", userID).First(&result).Error; err == nil {
		return "company"
	}
	if err := DB.Table("admins").Where("user_id = ?", userID).First(&result).Error; err == nil {
		return "admin"
	}

	return ""
}
