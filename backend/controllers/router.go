package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authController := NewAuthController(db)

	authGroup := router.Group("/auth")
	authGroup.POST("/register", authController.CreateUser)

	return router
}