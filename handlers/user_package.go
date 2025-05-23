package handlers

import (
	"encoding/json"
	"net/http"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
)

func CreateUserPackage(w http.ResponseWriter, r *http.Request) {
	var userPackage models.UserPackage
	if err := json.NewDecoder(r.Body).Decode(&userPackage); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if userPackage.UserID == 0 || userPackage.PackageID == 0 || userPackage.StartDate == "" || userPackage.EndDate == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Cek apakah user ada
	var user models.User
	if err := config.DB.First(&user, userPackage.UserID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Cek apakah package ada
	var packageData models.Package
	if err := config.DB.First(&packageData, userPackage.PackageID).Error; err != nil {
		http.Error(w, "Package not found", http.StatusNotFound)
		return
	}

	// Simpan user package
	if err := config.DB.Create(&userPackage).Error; err != nil {
		http.Error(w, "Failed to create user package", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User package created successfully"})
}
