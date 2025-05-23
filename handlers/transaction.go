package handlers

import (
	"encoding/json"
	"net/http"

	// "strconv"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
	// "github.com/gorilla/mux"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if transaction.UserID == 0 || transaction.PackageID == 0 || transaction.Amount <= 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Cek apakah user ada
	var user models.User
	if err := config.DB.First(&user, transaction.UserID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Cek apakah package ada
	var packageData models.Package
	if err := config.DB.First(&packageData, transaction.PackageID).Error; err != nil {
		http.Error(w, "Package not found", http.StatusNotFound)
		return
	}

	// Simpan transaksi
	if err := config.DB.Create(&transaction).Error; err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transaction created successfully"})
}
