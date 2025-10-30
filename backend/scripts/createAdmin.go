package main

import (
	"com-seek/backend/database"
	"com-seek/backend/models"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./create_admin <email>")
		return
	}

	email := os.Args[1]

	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	password := string(passwordBytes)
	fmt.Println()

	_ = godotenv.Load()

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	var count int64
	err = db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if count > 0 {
		log.Fatalf("User with email '%s' already exists.", email)
	} else if err != nil {
		log.Fatalf("Failed to check for existing user: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		user := models.User{
			Email:    email,
			Password: string(hashedPassword),
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		admin := models.Admin{
			UserID: user.ID,
		}

		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Printf("Admin user '%s' created successfully.\n", email)
}
