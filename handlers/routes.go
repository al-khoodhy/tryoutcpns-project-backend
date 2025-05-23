package handlers

import (
	"net/http"
	"tryoutcpns-project-backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterRoutes menambahkan semua route ke router
func RegisterRoutes(r *mux.Router, rateLimiter func(http.Handler) http.Handler) {
	// Auth routes
	r.HandleFunc("/api/register", Register).Methods("POST")
	r.HandleFunc("/api/login", Login).Methods("POST")

	// Protected routes (dibungkus pakai middleware AuthMiddleware)
	r.Handle("/api/users/{id}", middleware.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	r.Handle("/api/packages", middleware.AuthMiddleware(http.HandlerFunc(GetAllPackages))).Methods("GET")
	r.Handle("/api/questions", middleware.AuthMiddleware(http.HandlerFunc(GetAllQuestions))).Methods("GET")
	r.Handle("/api/results", middleware.AuthMiddleware(http.HandlerFunc(GetAllResults))).Methods("GET")
	r.Handle("/api/transactions", middleware.AuthMiddleware(http.HandlerFunc(CreateTransaction))).Methods("POST")

	// Apply global rate limiter (untuk semua route)
	r.Use(rateLimiter)
}
