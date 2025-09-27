package controllers

import (
	"com-seek/backend/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authController := NewAuthController(db)
	StudentController := NewStudentController(db)

	authGroup := router.Group("/auth")
	authGroup.POST("/register", authController.CreateUser)
	authGroup.POST("/login", authController.Login)

	requiredLogin := router.Group("/", middlewares.CheckAuth)

	student := requiredLogin.Group("/students")
	student.GET("/", StudentController.GetStudentProfile)
	student.PATCH("/", StudentController.UpdateStudentProfile)

	return router
}
