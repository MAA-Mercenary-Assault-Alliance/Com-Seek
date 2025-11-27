package controllers

import (
	"com-seek/backend/config"
	"com-seek/backend/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var logger = config.GetLogger()

func NewRouter(db *gorm.DB, fileConfig config.FileConfig) *gin.Engine {
	router := gin.Default()

	corsConfig := middlewares.SetupCors()
	router.Use(cors.New(corsConfig))
	router.Use(middlewares.SecurityHeaders())

	fileController := NewFileController(db, fileConfig)
	authController := NewAuthController(db, fileController)
	studentController := NewStudentController(db, fileController)
	companyController := NewCompanyController(db, fileController)
	jobController := NewJobController(db)
	JobApplicationController := NewJobApplicationController(db)
	adminController := NewAdminController(db)

	authGroup := router.Group("/auth")
	authGroup.POST("/register/student", middlewares.RateLimiterFor(http.MethodPost), authController.RegisterStudent)
	authGroup.POST("/register/company", middlewares.RateLimiterFor(http.MethodPost), authController.RegisterCompany)
	authGroup.POST("/login", middlewares.RateLimiterFor(http.MethodPost), authController.Login)
	authGroup.POST("/logout", middlewares.CheckAuth, authController.Logout)

	requiredLogin := router.Group("/", middlewares.CheckAuth)

	requiredLogin.GET("/file/:uuid", middlewares.RateLimiterFor(http.MethodGet), fileController.ServeFile)

	student := requiredLogin.Group("/student")
	student.GET("", middlewares.RateLimiterFor(http.MethodGet), studentController.GetStudentProfile)
	student.GET("/:id", middlewares.RateLimiterFor(http.MethodGet), studentController.GetStudentProfile)
	student.PATCH("", middlewares.RateLimiterFor(http.MethodPatch), studentController.UpdateStudentProfile)

	company := requiredLogin.Group("/company")
	company.GET("", middlewares.RateLimiterFor(http.MethodGet), companyController.GetCompanyProfile)
	company.GET("/:id", middlewares.RateLimiterFor(http.MethodGet), companyController.GetCompanyProfile)
	company.PATCH("", middlewares.RateLimiterFor(http.MethodPatch), companyController.UpdateCompanyProfile)
	company.GET("/jobs", middlewares.RateLimiterFor(http.MethodGet), companyController.GetCompanyJobs)

	job := requiredLogin.Group("/job")
	job.GET("", middlewares.RateLimiterFor(http.MethodGet), jobController.GetJobs)
	job.POST("", middlewares.RateLimiterFor(http.MethodPost), jobController.CreateJob)
	job.GET("/:id", middlewares.RateLimiterFor(http.MethodGet), jobController.GetJob)
	job.PATCH("/:id", middlewares.RateLimiterFor(http.MethodPatch), jobController.UpdateJob)
	job.DELETE("/:id", middlewares.RateLimiterFor(http.MethodDelete), jobController.DeleteJob)
	job.POST("/apply", middlewares.RateLimiterFor(http.MethodPost), JobApplicationController.CreateJobApplication)
	job.DELETE("/application/:id", middlewares.RateLimiterFor(http.MethodDelete), JobApplicationController.DeleteJobApplication)

	admin := requiredLogin.Group("/admin")
	admin.PATCH("review-company/:id", middlewares.RateLimiterFor(http.MethodPatch), adminController.ReviewCompany)
	admin.PATCH("review-student/:id", middlewares.RateLimiterFor(http.MethodPatch), adminController.ReviewStudent)
	admin.PATCH("review-job/:id", middlewares.RateLimiterFor(http.MethodPatch), adminController.ReviewJob)
	admin.GET("/companies", middlewares.RateLimiterFor(http.MethodGet), adminController.GetPendingCompanies)
	admin.GET("/students", middlewares.RateLimiterFor(http.MethodGet), adminController.GetPendingStudents)
	admin.GET("/jobs", middlewares.RateLimiterFor(http.MethodGet), adminController.GetPendingJobs)

	return router
}
