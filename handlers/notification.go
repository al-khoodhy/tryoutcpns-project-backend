package handlers

import (
    "tryoutcpns-project-backend/models"
    "tryoutcpns-project-backend/config"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {
    var notification models.Notification
    if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validasi input
    if notification.UserID == 0 || notification.Message == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

    // Cek apakah user ada
    var user models.User
    if err := config.DB.First(&user, notification.UserID).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Simpan notifikasi
    if err := config.DB.Create(&notification).Error; err != nil {
        http.Error(w, "Failed to create notification", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Notification created successfully"})
}