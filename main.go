package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/handlers"
	"tryoutcpns-project-backend/middleware"
	"tryoutcpns-project-backend/utils"

	"github.com/gorilla/mux"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
	// Load environment variables from .env
	if err := utils.LoadEnv(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Initialize database
	config.InitDB()

	// Migrate tables
	config.MigrateTables()

	// Initialize rate limiter
	rateLimiter := limiter.New(memory.NewStore(), limiter.Options{
		Max:     10, // Max 10 requests per minute
		Default: 1 * time.Minute,
	})

	// Create router
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Protected routes
	authMiddleware := middleware.AuthMiddleware
	r.HandleFunc("/api/users/{id}", handlers.GetUser).Methods("GET").Use(authMiddleware)
	r.HandleFunc("/api/packages", handlers.GetAllPackages).Methods("GET").Use(authMiddleware)
	r.HandleFunc("/api/questions", handlers.GetAllQuestions).Methods("GET").Use(authMiddleware)
	r.HandleFunc("/api/results", handlers.GetAllResults).Methods("GET").Use(authMiddleware)
	r.HandleFunc("/api/transactions", handlers.CreateTransaction).Methods("POST").Use(authMiddleware)

	// Apply rate limiter middleware
	r.Use(middleware.RateLimitMiddleware(rateLimiter))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
