package main

import (
	"log"
	"os"

	"com-seek/backend/config"
	"com-seek/backend/controllers"
	"com-seek/backend/database"
	"fmt"

	"io"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, dbErr := database.InitDB()
	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	fileConfig, err := config.LoadFileConfig()
	if err != nil {
		log.Fatalf("Error loading file configuration: %v", err)
	}

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	os.Mkdir("logs", 0755)
	// Logging to a file.
	os.Mkdir("logs/requests", 0755)
	logFileName := config.GenerateLogFileName("req")
	f, _ := os.Create(fmt.Sprintf("logs/requests/%s", logFileName))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := controllers.NewRouter(db, *fileConfig)

	routerErr := router.Run(os.Getenv("SERVER_ADDRESS"))
	if routerErr != nil {
		log.Fatal("Could not run server")
	}
}
