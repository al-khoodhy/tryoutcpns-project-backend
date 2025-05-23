package handlers

import (
	"net/http"
	"strconv"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"

	"github.com/gorilla/mux"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var result models.Result
	if err := config.DB.First(&result, id).Error; err != nil {
		http.Error(w, "Result not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllResults(w http.ResponseWriter, r *http.Request) {
	var results []models.Result
	config.DB.Find(&results)
	json.NewEncoder(w).Encode(results)
}
