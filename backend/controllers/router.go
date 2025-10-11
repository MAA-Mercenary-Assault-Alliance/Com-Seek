package controllers

import (
	"com-seek/backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	corsConfig := middlewares.SetupCors()
	router.Use(cors.New(corsConfig))

	authController := NewAuthController(db)
	studentController := NewStudentController(db)
	companyController := NewCompanyController(db)
	jobController := NewJobController(db)
	JobApplicationController := NewJobApplicationController(db)
	adminController := NewAdminController(db)

	authGroup := router.Group("/auth")
	authGroup.POST("/register", authController.CreateUser)
	authGroup.POST("/login", authController.Login)

	requiredLogin := router.Group("/", middlewares.CheckAuth)

	student := requiredLogin.Group("/student")
	student.GET("", studentController.GetStudentProfile)
	student.GET("/:id", studentController.GetStudentProfile)
	student.PATCH("", studentController.UpdateStudentProfile)

	company := requiredLogin.Group("/company")
	company.GET("", companyController.GetCompanyProfile)
	company.GET("/:id", companyController.GetCompanyProfile)
	company.PATCH("", companyController.UpdateCompanyProfile)

	job := requiredLogin.Group("/job")
	job.GET("", jobController.GetJobs)
	job.POST("", jobController.CreateJob)
	job.PATCH("/:id", jobController.UpdateJob)
	job.DELETE("/:id", jobController.DeleteJob)
	job.POST("/apply", JobApplicationController.CreateJobApplication)

	admin := requiredLogin.Group("/admin")
	admin.PATCH("review-company/:id", adminController.ReviewCompany)
	admin.PATCH("review-student/:id", adminController.ReviewStudent)
	admin.PATCH("review-job/:id", adminController.ReviewJob)
	admin.GET("/companies", adminController.GetPendingCompanies)
	admin.GET("/students", adminController.GetPendingStudents)
	admin.GET("/jobs", adminController.GetPendingJobs)

	return router
}
