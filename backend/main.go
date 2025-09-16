package main

import (
	"log"
	"os"

	"com-seek/backend/controllers"
	"com-seek/backend/database"
	"com-seek/backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, dbErr := database.InitDB()
	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	router := controllers.NewRouter(db)

	corsConfig := middlewares.SetupCors()
	router.Use(cors.New(corsConfig))

	routerErr := router.Run(os.Getenv("SERVER_ADDRESS"))
	if routerErr != nil {
		log.Fatal("Could not run server")
	}
}
