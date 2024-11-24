package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db.16 Error connecting", err)
	}
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("db.20 Failed to get database object from Gorm DB", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("db.23 Failed to ping database", err)
	}

	log.Println("db.26 Database connection established successfully")
	// database.AutoMigrate(&models.Category{}, &models.Organizer{}, &models.Industry{}, &models.EventState{}, &models.EventType{}, &models.Event_info{}, &models.User{}, &models.ReviewForm{}, &models.ReviewFormField{}, &models.UserRole{}, &models.UQEventContactInfo{}, &models.UqDataSource{}, models.ReviewSubmissionOption{})
	database.AutoMigrate()

	DB = database
}
