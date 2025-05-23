package main

import (
	"log"
	"net/http"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/handlers"
	"tryoutcpns-project-backend/migrations"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	migrations.CreateTables(config.DB)

	r := mux.NewRouter()

	// Route untuk registrasi
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")

	// Route untuk login
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Route untuk get user (dengan middleware)
	r.HandleFunc("/api/users/{id}", handlers.GetUser).Methods("GET").UseHandler(handlers.AuthMiddleware(handlers.GetUser))

	// Tambahkan route lain sesuai kebutuhan...

	log.Println("🚀 Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
