package handlers

import (
	"encoding/json"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	var leaderboard []models.Leaderboard
	config.DB.Find(&leaderboard)
	json.NewEncoder(w).Encode(leaderboard)
}
