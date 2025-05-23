package handlers

import (
    "cpns-backend/models"
    "cpns-backend/config"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
    var leaderboard []models.Leaderboard
    config.DB.Find(&leaderboard)
    json.NewEncoder(w).Encode(leaderboard)
}