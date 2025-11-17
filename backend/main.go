package main

import (
	"log"
	"os"

	"com-seek/backend/config"
	"com-seek/backend/controllers"
	"com-seek/backend/database"

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

	router := controllers.NewRouter(db, *fileConfig)

	routerErr := router.Run(os.Getenv("SERVER_ADDRESS"))
	if routerErr != nil {
		log.Fatal("Could not run server")
	}
}
