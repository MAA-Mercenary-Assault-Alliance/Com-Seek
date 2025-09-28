package controllers

import (
	"com-seek/backend/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authController := NewAuthController(db)
	studentController := NewStudentController(db)
	companyController := NewCompanyController(db)

	authGroup := router.Group("/auth")
	authGroup.POST("/register", authController.CreateUser)
	authGroup.POST("/login", authController.Login)

	requiredLogin := router.Group("/", middlewares.CheckAuth)

	student := requiredLogin.Group("/student")
	student.GET("/", studentController.GetStudentProfile)
	student.GET("/:id", studentController.GetStudentProfile)
	student.PATCH("/", studentController.UpdateStudentProfile)

	company := requiredLogin.Group("/company")
	company.GET("/", companyController.GetCompanyProfile)
	company.GET("/:id", companyController.GetCompanyProfile)
	company.PATCH("/", companyController.UpdateCompanyProfile)

	return router
}
