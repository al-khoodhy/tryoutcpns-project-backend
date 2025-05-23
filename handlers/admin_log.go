package handlers

import (
	"encoding/json"
	"net/http"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
)

func CreateAdminLog(w http.ResponseWriter, r *http.Request) {
	var adminLog models.AdminLog
	if err := json.NewDecoder(r.Body).Decode(&adminLog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if adminLog.AdminID == 0 || adminLog.Action == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Cek apakah admin ada
	var admin models.User
	if err := config.DB.First(&admin, adminLog.AdminID).Error; err != nil {
		http.Error(w, "Admin not found", http.StatusNotFound)
		return
	}

	// Simpan log admin
	if err := config.DB.Create(&adminLog).Error; err != nil {
		http.Error(w, "Failed to create admin log", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin log created successfully"})
}
