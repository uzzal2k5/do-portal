package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
    "dop-backend/models"
)

// DB is the global database connection object
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
    // Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
    // Retrieve the environment variables
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	// PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSLMODE)

	var db_err error
	DB, db_err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db_err != nil {
		log.Fatalf("Error connecting to the database: %v", db_err)
	}
	fmt.Println("Database connection established.")
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error closing the database connection: %v", err)
	}
	sqlDB.Close()
}

// SaveContactForm stores the submitted contact form in the database
func SaveContactForm(form models.ContactForm) error {
	err := DB.Exec(`
		INSERT INTO contact_queries (name, email, mobile, message, created_at)
		VALUES ($1, $2, $3, $4, $5)`, form.Name, form.Email, form.Mobile, form.Message, form.CreatedAt)
	if err != nil {
		return fmt.Errorf("unable to insert contact query: %v", err)
	}
	return nil
}

// GetPortfolioContent fetches the portfolio content from the database
func GetPortfolioContent() (models.Portfolio, error) {
	var portfolio models.Portfolio
	query := `
		SELECT id, title, description, skill_set, experience, created_at
		FROM portfolio
		ORDER BY created_at DESC
		LIMIT 1
	`

	err := DB.Raw(query).Scan(&portfolio).Error
	if err != nil {
		return portfolio, fmt.Errorf("error fetching portfolio content: %v", err)
	}
	return portfolio, nil
}