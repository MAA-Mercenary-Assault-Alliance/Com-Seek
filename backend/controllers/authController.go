package controllers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"com-seek/backend/helpers"
	"com-seek/backend/models"
	"com-seek/backend/services"

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
	Email          string                `form:"email" binding:"required,email"`
	Password       string                `form:"password" binding:"required"`
	FirstName      string                `form:"first_name" binding:"required"`
	LastName       string                `form:"last_name" binding:"required"`
	IsAlumString   string                `form:"is_alum" binding:"required,oneof=true false 1 0"`
	Transcript     *multipart.FileHeader `form:"transcript" binding:"required"`
	ReCAPTCHAToken string                `form:"recaptcha_response" binding:"required"`
}

type RegisterCompanyInput struct {
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Location       string `json:"location" binding:"required"`
	ContactEmail   string `json:"contact_email" binding:"required,email"`
	ContactNumber  string `json:"contact_number" binding:"required"`
	ReCAPTCHAToken string `json:"recaptcha_response" binding:"required"`
}

func (ac *AuthController) RegisterStudent(c *gin.Context) {
	var input RegisterStudentInput

	if err := c.ShouldBind(&input); err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a student: %s", helpers.GetClientIP(c), err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.VerifyRecaptchaToken(input.ReCAPTCHAToken); err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a student: CAPTCHA verification failed: %s", helpers.GetClientIP(c), err.Error()))
		if errors.Is(err, services.ErrRecaptchaFailed) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			logger.Error("Attempt to Register" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		return
	}

	isAlum, err := strconv.ParseBool(input.IsAlumString)
	if err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a student: invalid is_alum value", helpers.GetClientIP(c)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for is_alum"})
		return
	}

	var userFound models.User
	if err := ac.DB.Where("email = ?", input.Email).First(&userFound).Error; err == nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a student: email already registered", helpers.GetClientIP(c)))
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a student: failed to hash password", helpers.GetClientIP(c)))
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
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a student: failed to create user: %s", helpers.GetClientIP(c), err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user: " + err.Error()})
		return
	}

	transcript, err := ac.fileController.SaveFile(c, user.ID, input.Transcript, models.FileCategoryTranscript)
	if err != nil {
		logger.Error(fmt.Sprintf("<User id: %d> attempt to register as a student: failed to save transcript file: %s", user.ID, err.Error()))
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
		logger.Error(fmt.Sprintf("<User id: %d> attempt to register as a student: failed to create student profile: %s", user.ID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student profile: " + err.Error()})
		helpers.DeleteFileRecord(ac.DB, transcript.ID)
		return
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error(fmt.Sprintf("<User id: %d> attempt to register as a student: transaction commit failed: %s", user.ID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction commit failed: " + err.Error()})
		helpers.DeleteFileRecord(ac.DB, transcript.ID)
		return
	}

	logger.Info(fmt.Sprintf("<Student id: %d> registered successfully", user.ID))
	c.JSON(http.StatusOK, gin.H{"message": "student user created successfully"})
}

func (ac *AuthController) RegisterCompany(c *gin.Context) {
	var input RegisterCompanyInput

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a company: %s", helpers.GetClientIP(c), err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.VerifyRecaptchaToken(input.ReCAPTCHAToken); err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a company: CAPTCHA verification failed: %s", helpers.GetClientIP(c), err.Error()))
		if errors.Is(err, services.ErrRecaptchaFailed) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		return
	}

	var userFound models.User
	if err := ac.DB.Where("email = ?", input.Email).First(&userFound).Error; err == nil {
		logger.Info(fmt.Sprintf("<User ip: %s> attempt to register as a company: email already registered", helpers.GetClientIP(c)))
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a company: failed to hash password", helpers.GetClientIP(c)))
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
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to register as a company: failed to create user: %s", helpers.GetClientIP(c), err.Error()))
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
		logger.Error(fmt.Sprintf("<User id: %d> attempt to register as a company: failed to create company profile: %s", user.ID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create company profile: " + err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error(fmt.Sprintf("<User id: %d> attempt to register as a company: transaction commit failed: %s", user.ID, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction commit failed: " + err.Error()})
		return
	}

	logger.Info(fmt.Sprintf("<Company id: %d> registered successfully", user.ID))
	c.JSON(http.StatusOK, gin.H{"message": "company user created successfully"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginInput LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to login: %s", helpers.GetClientIP(c), err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	ac.DB.Where("email=?", loginInput.Email).Find(&userFound)

	if userFound.ID == 0 {
		logger.Error(fmt.Sprintf("<User ip: %s> attempt to login: user not found", helpers.GetClientIP(c)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
		logger.Error(fmt.Sprintf("<User id: %d> attempt to login: invalid password", userFound.ID))
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
		logger.Error(fmt.Sprintf("<User id: %d> attempt to login: failed to generate token: %s", userFound.ID, err.Error()))
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

	logger.Info(fmt.Sprintf("<User id: %d> logged in successfully as a %s", userFound.ID, role))
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

	userID := c.MustGet("userID").(uint)
	role := GetUserRole(ac.DB, userID)

	logger.Info(fmt.Sprintf("<%s id: %d> logged out successfully", role, userID))
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
