package main

import (
    "fmt"
	"log"
	"os"
    "os/signal"
    "syscall"
	"net/http"
	"dop-backend/database"
	"dop-backend/models"
	"github.com/joho/godotenv"
	"dop-backend/routes"
)

func main() {
    //  Load environment variables from .env file
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the PORT value from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
    // Initialize the database connection
	database.InitDB()
	// Automatically migrate the schema, this will create/update tables based on the models
    db_err := database.DB.AutoMigrate(&models.Service{})
    if db_err != nil {
    	log.Fatalf("Error migrating database schema: %v", db_err)
    }
    fmt.Println("Database migrated successfully!")

    // Gracefully shut down the application and close the DB connection
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		fmt.Println("Shutting down...")
		database.CloseDB()
	}()
    // Initialize router
	router := routes.DopRouter()

    // Apply CORS middleware
// 	corsHandler := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"http://localhost:3000"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Content-Type"},
// 		AllowCredentials: true,
// 	})
// 	handler := corsHandler.Handler(router)
    // Apply CORS middleware
// 	corsOptions := handlers.AllowedOrigins([]string{"*"})
// 	http.Handle("/", handlers.CORS(corsOptions)(router)) // Apply CORS middleware

    // Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		fmt.Println("Shutting down...")
		database.CloseDB()
		os.Exit(0)
	}()
    // Start the server
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}


