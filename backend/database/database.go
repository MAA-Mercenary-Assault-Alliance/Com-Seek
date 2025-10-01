package database

import (
	"fmt"
	"os"

	"com-seek/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	Password string
	Name     string
}

func LoadDBConfig() (*DBConfig, error) {
	required := []string{"DB_HOST", "DB_PORT", "DB_ROOT_PASSWORD", "DB_NAME"}
	for _, key := range required {
		if os.Getenv(key) == "" {
			return nil, fmt.Errorf("environment variable %s is required", key)
		}
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_ROOT_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}, nil
}

func InitDB() (*gorm.DB, error) {
	config, err := LoadDBConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Password, config.Host, config.Port, config.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	allModels := []any{
		&models.User{},
		&models.Admin{},
		&models.Student{},
		&models.Company{},
		&models.Job{},
	}

	if err := db.AutoMigrate(allModels...); err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	return db, nil
}
