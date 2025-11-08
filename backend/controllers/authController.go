package controllers

import (
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"com-seek/backend/helpers"
	"com-seek/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB             *gorm.DB
	fileController *FileController
}

func NewAuthController(db *gorm.DB, fileController *FileController) *AuthController {
	return &AuthController{
		DB:             db,
		fileController: fileController,
	}
}

type RegisterStudentInput struct {
	Email        string                `form:"email" binding:"required,email"`
	Password     string                `form:"password" binding:"required"`
	FirstName    string                `form:"first_name" binding:"required"`
	LastName     string                `form:"last_name" binding:"required"`
	IsAlumString string                `form:"is_alum" binding:"required,oneof=true false 1 0"`
	Transcript   *multipart.FileHeader `form:"transcript" binding:"required"`
}

type RegisterCompanyInput struct {
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Location      string `json:"location" binding:"required"`
	ContactEmail  string `json:"contact_email" binding:"required,email"`
	ContactNumber string `json:"contact_number" binding:"required"`
}

func (ac *AuthController) RegisterStudent(c *gin.Context) {
	var input RegisterStudentInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAlum, err := strconv.ParseBool(input.IsAlumString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for is_alum"})
		return
	}

	var userFound models.User
	if err := ac.DB.Where("email = ?", input.Email).First(&userFound).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user: " + err.Error()})
		return
	}

	transcript, err := ac.fileController.SaveFile(c, user.ID, input.Transcript, models.FileCategoryTranscript)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profile := models.Student{
		UserID:       user.ID,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		IsAlum:       isAlum,
		TranscriptID: transcript.ID,
	}

	if err := tx.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student profile: " + err.Error()})
		helpers.DeleteFileRecord(ac.DB, transcript.ID)
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction commit failed: " + err.Error()})
		helpers.DeleteFileRecord(ac.DB, transcript.ID)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student user created successfully"})
}

func (ac *AuthController) RegisterCompany(c *gin.Context) {
	var input RegisterCompanyInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	if err := ac.DB.Where("email = ?", input.Email).First(&userFound).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user: " + err.Error()})
		return
	}

	profile := models.Company{
		UserID:        user.ID,
		Name:          input.Name,
		Location:      input.Location,
		ContactEmail:  input.ContactEmail,
		ContactNumber: input.ContactNumber,
	}

	if err := tx.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create company profile: " + err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction commit failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "company user created successfully"})
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

func (ac *AuthController) Logout(c *gin.Context) {
	sameSite := http.SameSiteLaxMode
	secure := false
	if os.Getenv("ENV") != "dev" {
		sameSite = http.SameSiteNoneMode
		secure = true
	}

	c.SetSameSite(sameSite)
	c.SetCookie(
		"token",
		"",
		-1,
		"/",
		"",
		secure,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
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
