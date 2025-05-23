package handlers

import (
    "encoding/json"
    "net/http"
    "tryoutcpns-project-backend/models"
    "tryoutcpns-project-backend/config"
)

func Register(w http.ResponseWriter, r *http.Request) {
    config.InitDB()
    defer config.DB.Close()

    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    query := "INSERT INTO users (name, email, phone, password, role) VALUES (?, ?, ?, ?, ?)"
    _, err := config.DB.Exec(query, user.Name, user.Email, user.Phone, user.Password, "user")
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}