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
	// Load environment variables
	if err := utils.LoadEnv(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Initialize database
	config.InitDB()

	// Migrate tables
	// config.MigrateTables()

	// Inisialisasi rate limiter
	store := memory.NewStore()
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10,
	}
	rawLimiter := limiter.New(store, rate)
	rateLimiterMiddleware := middleware.RateLimitMiddleware(rawLimiter)

	// Buat router
	r := mux.NewRouter()

	// Daftarkan semua routes
	handlers.RegisterRoutes(r, rateLimiterMiddleware)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
